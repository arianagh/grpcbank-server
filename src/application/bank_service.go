package application

import (
	"fmt"
	"github.com/google/uuid"
	"grpcbank/src/adapter/database"
	"grpcbank/src/application/domain"
	"grpcbank/src/port"
	"log"
	"time"
)

type BankService struct {
	db port.BankDatabasePort
}

func NewBankService(dbPort port.BankDatabasePort) *BankService {
	return &BankService{
		db: dbPort,
	}
}

func (s *BankService) FindCurrentBalance(accountNumber string) (float64, error) {
	bankAccount, err := s.db.GetBankAccountByAccountNumber(accountNumber)

	if err != nil {
		log.Println("Error on FindCurrentBalance :", err)
		return 0, err
	}

	return bankAccount.CurrentBalance, nil
}

func (s *BankService) CreateExchangeRate(exchangeRate domain.ExchangeRate) (uuid.UUID, error) {
	newUuid := uuid.New()
	now := time.Now()

	exchangeRateOrm := database.BankExchangeRateOrm{
		ExchangeRateUuid:   newUuid,
		FromCurrency:       exchangeRate.FromCurrency,
		ToCurrency:         exchangeRate.ToCurrency,
		Rate:               exchangeRate.Rate,
		ValidFromTimestamp: exchangeRate.ValidFromTimestamp,
		ValidToTimestamp:   exchangeRate.ValidToTimestamp,
		CreatedAt:          now,
		UpdatedAt:          now,
	}

	return s.db.CreateExchangeRate(exchangeRateOrm)
}

func (s *BankService) FindExchangeRate(fromCur string, toCur string, ts time.Time) (float64, error) {
	exchangeRate, err := s.db.GetExchangeRateAtTimestamp(fromCur, toCur, ts)

	if err != nil {
		return 0, err
	}

	return exchangeRate.Rate, nil
}

func (s *BankService) CreateTransaction(acct string, bankTrx domain.Transaction) (uuid.UUID, error) {
	newUuid := uuid.New()
	now := time.Now()

	bankAccountOrm, err := s.db.GetBankAccountByAccountNumber(acct)

	if err != nil {
		log.Printf("Can't create transaction for %v : %v\n", acct, err)
		return uuid.Nil, fmt.Errorf("can't find account number %v : %v", acct, err.Error())
	}

	if bankTrx.TransactionType == domain.TransactionTypeOut && bankAccountOrm.CurrentBalance < bankTrx.Amount {
		return bankAccountOrm.AccountUuid, fmt.Errorf(
			"insufficient account balance %v for [out] transaction amount %v",
			bankAccountOrm.CurrentBalance, bankTrx.Amount,
		)
	}

	transactionOrm := database.BankTransactionOrm{
		TransactionUuid:      newUuid,
		AccountUuid:          bankAccountOrm.AccountUuid,
		TransactionTimestamp: now,
		Amount:               bankTrx.Amount,
		TransactionType:      bankTrx.TransactionType,
		Notes:                bankTrx.Notes,
		CreatedAt:            now,
		UpdatedAt:            now,
	}

	savedUuid, err := s.db.CreateTransaction(bankAccountOrm, transactionOrm)

	return savedUuid, err
}

func (s *BankService) CalculateTransactionSummary(trxSummary *domain.TransactionSummary,
	bankTrx domain.Transaction) error {
	switch bankTrx.TransactionType {
	case domain.TransactionTypeIn:
		trxSummary.SumIn += bankTrx.Amount
	case domain.TransactionTypeOut:
		trxSummary.SumOut += bankTrx.Amount
	default:
		return fmt.Errorf("unknown transaction type %v", bankTrx.TransactionType)
	}

	trxSummary.SumTotal = trxSummary.SumIn - trxSummary.SumOut

	return nil
}

func (s *BankService) Transfer(transferTrx domain.TransferTransaction) (uuid.UUID, bool, error) {
	now := time.Now()

	fromAccountOrm, err := s.db.GetBankAccountByAccountNumber(transferTrx.FromAccountNumber)

	if err != nil {
		log.Printf("Can't find transfer from account %v : %v\n", transferTrx.FromAccountNumber, err)
		return uuid.Nil, false, domain.ErrTransferSourceAccountNotFound
	}

	if fromAccountOrm.CurrentBalance < transferTrx.Amount {
		return uuid.Nil, false, domain.ErrTransferTransactionPair
	}

	toAccountOrm, err := s.db.GetBankAccountByAccountNumber(transferTrx.ToAccountNumber)

	if err != nil {
		log.Printf("Can't find transfer to account %v : %v\n", transferTrx.ToAccountNumber, err)
		return uuid.Nil, false, domain.ErrTransferDestinationAccountNotFound
	}

	fromTransactionOrm := database.BankTransactionOrm{
		TransactionUuid:      uuid.New(),
		TransactionTimestamp: now,
		TransactionType:      domain.TransactionTypeOut,
		AccountUuid:          fromAccountOrm.AccountUuid,
		Amount:               transferTrx.Amount,
		Notes:                "Transfer out to " + transferTrx.ToAccountNumber,
		CreatedAt:            now,
		UpdatedAt:            now,
	}

	toTransactionOrm := database.BankTransactionOrm{
		TransactionUuid:      uuid.New(),
		TransactionTimestamp: now,
		TransactionType:      domain.TransactionTypeIn,
		AccountUuid:          toAccountOrm.AccountUuid,
		Amount:               transferTrx.Amount,
		Notes:                "Transfer in from " + transferTrx.FromAccountNumber,
		CreatedAt:            now,
		UpdatedAt:            now,
	}

	newTransferUuid := uuid.New()

	transferOrm := database.BankTransferOrm{
		TransferUuid:      newTransferUuid,
		FromAccountUuid:   fromAccountOrm.AccountUuid,
		ToAccountUuid:     toAccountOrm.AccountUuid,
		Currency:          transferTrx.Currency,
		Amount:            transferTrx.Amount,
		TransferTimestamp: now,
		TransferSuccess:   false,
		CreatedAt:         now,
		UpdatedAt:         now,
	}

	if _, err := s.db.CreateTransfer(transferOrm); err != nil {
		log.Printf("Can't create transfer from %v to %v : %v\n", transferTrx.FromAccountNumber, transferTrx.ToAccountNumber, err)
		return uuid.Nil, false, domain.ErrTransferRecordFailed
	}

	if transferPairSuccess, _ := s.db.CreateTransferTransactionPair(fromAccountOrm,
		toAccountOrm, fromTransactionOrm, toTransactionOrm); transferPairSuccess {
		err := s.db.UpdateTransferStatus(transferOrm, true)
		if err != nil {
			return uuid.Nil, false, err
		}
		return newTransferUuid, true, nil
	} else {
		return newTransferUuid, false, domain.ErrTransferTransactionPair
	}
}
