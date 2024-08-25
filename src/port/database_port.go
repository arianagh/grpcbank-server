package port

import (
	"github.com/google/uuid"
	"grpcbank/src/adapter/database"
	"time"
)

type BankDatabasePort interface {
	GetBankAccountByAccountNumber(accountNumber string) (database.BankAccountOrm, error)
	CreateExchangeRate(exchangeRate database.BankExchangeRateOrm) (uuid.UUID, error)
	GetExchangeRateAtTimestamp(fromCur string, toCur string, timeStamp time.Time) (database.BankExchangeRateOrm, error)
	CreateTransaction(acct database.BankAccountOrm, bankTrx database.BankTransactionOrm) (uuid.UUID, error)
	CreateTransfer(transfer database.BankTransferOrm) (uuid.UUID, error)
	CreateTransferTransactionPair(fromAccountOrm database.BankAccountOrm, toAccountOrm database.BankAccountOrm,
		fromTransactionOrm database.BankTransactionOrm, toTransactionOrm database.BankTransactionOrm) (bool, error)
	UpdateTransferStatus(transfer database.BankTransferOrm, status bool) error
}
