package database

import (
	"github.com/google/uuid"
	"grpcbank/src/application/domain"
	"log"
	"time"
)

func (a *DatabaseAdapter) GetBankAccountByAccountNumber(accountNumber string) (BankAccountOrm, error) {
	var bankAccountOrm BankAccountOrm

	if err := a.db.First(&bankAccountOrm, "account_number = ?", accountNumber).Error; err != nil {
		log.Printf("Can't find bank account number %v : %v\n", accountNumber, err)
		return bankAccountOrm, err
	}

	return bankAccountOrm, nil
}

func (a *DatabaseAdapter) CreateExchangeRate(exchangeRate BankExchangeRateOrm) (uuid.UUID, error) {
	if err := a.db.Create(exchangeRate).Error; err != nil {
		return uuid.Nil, err
	}

	return exchangeRate.ExchangeRateUuid, nil
}

func (a *DatabaseAdapter) GetExchangeRateAtTimestamp(fromCur string, toCur string,
	timeStamp time.Time) (BankExchangeRateOrm, error) {
	var exchangeRateOrm BankExchangeRateOrm

	err := a.db.First(&exchangeRateOrm, "from_currency = ? "+" AND to_currency = ? "+
		" AND (? BETWEEN valid_from_timestamp and valid_to_timestamp)", fromCur, toCur, timeStamp).Error

	return exchangeRateOrm, err
}

func (a *DatabaseAdapter) CreateTransaction(acct BankAccountOrm, bankTrx BankTransactionOrm) (uuid.UUID, error) {
	tx := a.db.Begin()

	if err := tx.Create(bankTrx).Error; err != nil {
		tx.Rollback()
		return uuid.Nil, err
	}

	newAmount := bankTrx.Amount

	if bankTrx.TransactionType == domain.TransactionTypeOut {
		newAmount = -1 * bankTrx.Amount
	}

	newAccountBalance := acct.CurrentBalance + newAmount

	if err := tx.Model(&acct).Updates(
		map[string]interface{}{
			"current_balance": newAccountBalance,
			"updated_at":      time.Now(),
		},
	).Error; err != nil {
		tx.Rollback()
		return uuid.Nil, err
	}

	tx.Commit()

	return bankTrx.TransactionUuid, nil
}

func (a *DatabaseAdapter) CreateTransfer(transfer BankTransferOrm) (uuid.UUID, error) {
	if err := a.db.Create(transfer).Error; err != nil {
		return uuid.Nil, err
	}

	return transfer.TransferUuid, nil
}

func (a *DatabaseAdapter) CreateTransferTransactionPair(fromAccountOrm BankAccountOrm,
	toAccountOrm BankAccountOrm, fromTransactionOrm BankTransactionOrm,
	toTransactionOrm BankTransactionOrm) (bool, error) {
	tx := a.db.Begin()

	if err := tx.Create(fromTransactionOrm).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	if err := tx.Create(toTransactionOrm).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	fromAccountBalanceNew := fromAccountOrm.CurrentBalance - fromTransactionOrm.Amount

	if err := tx.Model(&fromAccountOrm).Updates(
		map[string]interface{}{
			"current_balance": fromAccountBalanceNew,
			"updated_at":      time.Now(),
		},
	).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	toAccountBalanceNew := toAccountOrm.CurrentBalance + toTransactionOrm.Amount

	if err := tx.Model(&toAccountOrm).Updates(
		map[string]interface{}{
			"current_balance": toAccountBalanceNew,
			"updated_at":      time.Now(),
		},
	).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()

	return true, nil
}

func (a *DatabaseAdapter) UpdateTransferStatus(transfer BankTransferOrm, status bool) error {
	if err := a.db.Model(&transfer).Updates(
		map[string]interface{}{
			"transfer_success": status,
			"updated_at":       time.Now(),
		},
	).Error; err != nil {
		return err
	}

	return nil
}
