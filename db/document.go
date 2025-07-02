package storage

import (
	"fmt"

	"gahmen-api/types"

	"github.com/jmoiron/sqlx"
)

type DocumentStore struct {
	db *sqlx.DB
}

func NewDocumentStore(db *sqlx.DB) *DocumentStore {
	return &DocumentStore{db: db}
}

func (s *DocumentStore) ListDocumentByMinistryID(ministryID int) ([]*types.Document, error) {
	sqlStmt := `select
				b.id as DocumentID,
				m."Name" as Ministry,
				b."Title" as DocumentName,
				b."Year"  as Year,
				b."DocumentPath" as DocumentPath ,
				b."MD5Hash" as MD5Hash ,
				b."CreatedAt" as CreatedAt
			from
				budgetdocuments b
			inner join ministry m on
				m.id = b."MinistryID"
			where
			b."MinistryID" = $1`
	documents := []*types.Document{}
	err := s.db.Select(&documents, sqlStmt, ministryID)
	if err != nil {
		return nil, fmt.Errorf("error listing documents by ministry ID: %w", err)
	}
	return documents, nil
}

func (s *DocumentStore) GetDocumentByID(id int) (*types.Document, error) {
	sqlStmt := "SELECT * FROM budgetdocuments WHERE id = $1"
	document := types.Document{}
	err := s.db.Get(&document, sqlStmt, id)
	if err != nil {
		return nil, fmt.Errorf("error getting document by ID: %w", err)
	}
	return &document, nil
}
