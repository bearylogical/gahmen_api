package storage

import (
	"fmt"

	"gahmen-api/types"

	"github.com/jmoiron/sqlx"
)

type PersonnelStore struct {
	db *sqlx.DB
}

func NewPersonnelStore(db *sqlx.DB) *PersonnelStore {
	return &PersonnelStore{db: db}
}

func (s *PersonnelStore) ListTopNPersonnelByMinistryID(ministryID int, n int, startYear int) ([]*types.Personnel, error) {
	sqlStmt := `select
					category,
					"ParentHeader" as "ParentCategory" ,
					"ValueAmount",
					"ValueYear",
					"ValueType"
				from
					budgetpersonnel b
				where
					"MinistryID" = $1
				and
					"category" != 'TOTAL'	
					`

	if startYear > 0 {
		sqlStmt += fmt.Sprintf(" and \"ValueYear\" >= %d ", startYear)
	}
	if n > 0 {
		sqlStmt += fmt.Sprintf(" limit %d", n)
	}

	sqlStmt += `order by
	"ValueAmount" desc`

	opts := []*types.Personnel{}
	err := s.db.Select(&opts, sqlStmt, ministryID)
	if err != nil {
		return nil, fmt.Errorf("error listing top N personnel by ministry ID: %w", err)
	}
	return opts, nil
}

