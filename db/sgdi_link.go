package storage

import (
	"fmt"

	"gahmen-api/types"

	"github.com/jmoiron/sqlx"
)

type SGDILinkStore struct {
	db *sqlx.DB
}

func NewSGDILinkStore(db *sqlx.DB) *SGDILinkStore {
	return &SGDILinkStore{db: db}
}

func (s *SGDILinkStore) GetSGDILinkByMinistryID(ministryID int) ([]*types.SGDILink, error) {
	sqlStmt := `select
				o2."Name" as child,
				o."Name" as parent,
				o2."URL_LINK" as child_url,
				o."URL_LINK" as parent_url
			from
				organisation o2
			left join (
				select
					id,
					"Name",
					"URL_LINK"
				from
					organisation) o
			on
				o2."ParentID" = o.id
			where
				o2."MinistryID" = $1`
			links := []*types.SGDILink{}
	err := s.db.Select(&links, sqlStmt, ministryID)
	if err != nil {
		return nil, fmt.Errorf("error getting SGDI link by ministry ID: %w", err)
	}
	return links, nil
}
