package match

import (
	"github.com/szokodiakos/r8m8/sql"
	"github.com/szokodiakos/r8m8/transaction"
)

type matchRepositorySQL struct {
}

func (mrs *matchRepositorySQL) Create(transaction transaction.Transaction) (int64, error) {
	var createdID int64
	query := `
		INSERT INTO matches
			(created_at)
		VALUES
			(current_timestamp)
		RETURNING id;
	`

	sqlTransaction := transaction.ConcreteTransaction.(sql.Transaction)
	res := sqlTransaction.QueryRow(query)
	err := res.Scan(&createdID)
	if err != nil {
		return createdID, err
	}

	return createdID, nil
}

// NewRepositorySQL factory
func NewRepositorySQL() Repository {
	return &matchRepositorySQL{}
}
