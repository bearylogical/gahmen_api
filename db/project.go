package storage

import (
	"fmt"

	"gahmen-api/types"

	"github.com/jmoiron/sqlx"
)

type ProjectStore struct {
	db *sqlx.DB
}

func NewProjectStore(db *sqlx.DB) *ProjectStore {
	return &ProjectStore{db: db}
}

func (s *ProjectStore) ListProjectsByMinistryID(ministryID int) ([]*types.ProjectExpenditure, error) {
	sqlStmt := `select
				id as "ProjectID",
				"ProjectTitle" ,
				"Ministry" ,
				"ValueType" ,
				"ValueAmount" ,
				"ValueYear" ,
				"Category" ,
				"DocumentYear" ,
				"MinistryID" ,
				"DocumentID" ,
				"ExpenditureID"
				from
				budgetprojects b
				where
				"MinistryID" = $1
	`

	projects := []*types.ProjectExpenditure{}
	err := s.db.Select(&projects, sqlStmt, ministryID)
	if err != nil {
		return nil, fmt.Errorf("error listing projects by ministry ID: %w", err)
	}
	return projects, nil
}
