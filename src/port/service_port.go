package port

import (
	"github.com/google/uuid"
	"grpcbank/src/application/domain"
	"time"
)

type BankServicePort interface {
	FindCurrentBalance(accountNumber string) (float64, error)
	CreateExchangeRate(exchangeRate domain.ExchangeRate) (uuid.UUID, error)
	FindExchangeRate(fromCur string, toCur string, ts time.Time) (float64, error)
	CreateTransaction(acct string, bankTrx domain.Transaction) (uuid.UUID, error)
	CalculateTransactionSummary(trxSummary *domain.TransactionSummary, bankTrx domain.Transaction) error
	Transfer(transferTrx domain.TransferTransaction) (uuid.UUID, bool, error)
}
