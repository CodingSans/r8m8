package transaction

import (
	"github.com/szokodiakos/r8m8/sql"
)

type transactionService struct {
	db sql.DB
}

func (ts *transactionService) Start() (Transaction, error) {
	var transaction Transaction
	sqlTransaction, err := ts.db.Begin()
	if err != nil {
		return transaction, err
	}
	transaction = Transaction{
		concreteTransaction: sqlTransaction,
	}
	return transaction, nil
}

func (ts *transactionService) Commit(transaction Transaction) error {
	sqlTransaction := GetSQLTransaction(transaction)
	return sqlTransaction.Commit()
}

func (ts *transactionService) Rollback(transaction Transaction) error {
	sqlTransaction := GetSQLTransaction(transaction)
	return sqlTransaction.Rollback()
}

// NewServiceSQL factory
func NewServiceSQL(db sql.DB) Service {
	return &transactionService{
		db: db,
	}
}
