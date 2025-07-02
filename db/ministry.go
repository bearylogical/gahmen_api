package storage

import (
	"fmt"

	"gahmen-api/types"

	"github.com/jmoiron/sqlx"
)

type MinistryStore struct {
	db *sqlx.DB
}

func NewMinistryStore(db *sqlx.DB) *MinistryStore {
	return &MinistryStore{db: db}
}

func (s *MinistryStore) CreateMinistry(m *types.Ministry) error {
	sqlStmt := "INSERT INTO ministry (name, created_at) VALUES ($1, $2)"
	_, err := s.db.Exec(sqlStmt, m.Name, m.CreatedAt)
	if err != nil {
		return fmt.Errorf("error creating ministry: %w", err)
	}
	return nil
}

func (s *MinistryStore) DeleteMinistry(id int) error {
	sqlStmt := "DELETE FROM ministry WHERE id = $1"
	_, err := s.db.Exec(sqlStmt, id)
	if err != nil {
		return fmt.Errorf("error deleting ministry: %w", err)
	}
	return nil
}

func (s *MinistryStore) GetMinistryByID(id int) (*types.Ministry, error) {
	sqlStmt := `SELECT * FROM ministry WHERE id = $1`
	ministry := types.Ministry{}
	err := s.db.Get(&ministry, sqlStmt, id)
	if err != nil {
		return nil, fmt.Errorf("error getting ministry by ID: %w", err)
	}
	return &ministry, nil
}

func (s *MinistryStore) ListMinistries(isMinistry bool) ([]*types.Ministry, error) {
	ministries := []*types.Ministry{}
	var sqlStmt string
	if isMinistry {
		sqlStmt = `
		select
			distinct("MinistryID") as id,
			m."Name" ,
			m."CreatedAt" 
		from
			organisation o
		left join ministry m on
			o."MinistryID" = m.id
		order by
			id asc`
	} else {
		sqlStmt = `select
					m.id as id,
					m."Name",
					m."CreatedAt"
				from
					(
					select
						"MinistryID"
					from
						budgetdocuments c
					where
						c."Year" = (
						select
							"Year"
						from
							budgetdocuments b
						order by
							"Year" desc
						limit 1)) d
				inner join ministry m on
					d."MinistryID" = m.ID
				where
					m."Name" != ''`
	}
	err := s.db.Select(&ministries, sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("error listing ministries: %w", err)
	}
	return ministries, nil
}
