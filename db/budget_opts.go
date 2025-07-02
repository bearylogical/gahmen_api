package storage

import (
	"fmt"

	"gahmen-api/types"

	"github.com/jmoiron/sqlx"
)

type BudgetOptsStore struct {
	db *sqlx.DB
}

func NewBudgetOptsStore(db *sqlx.DB) *BudgetOptsStore {
	return &BudgetOptsStore{db: db}
}

func (s *BudgetOptsStore) GetBudgetOpts() ([]*types.BudgetOpts, error) {
	sqlStmt := `select
					distinct("ValueType"),
					"ValueYear"
				from
					budgetexpenditure b`
	opts := []*types.BudgetOpts{}
	err := s.db.Select(&opts, sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("error getting budget options: %w", err)
	}
	return opts, nil
}
