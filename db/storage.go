package storage

import (
	"fmt"

	"gahmen-api/config"
	"gahmen-api/types"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateMinistry(*types.Ministry) error
	DeleteMinistry(int) error
	UpdateMinistry(*types.Ministry) error
	GetMinistryByID(int) (*types.Ministry, error)
	ListMinistry() ([]*types.Ministry, error)
	ListDocumentByMinistryID(int) ([]*types.Document, error)
	GetDocumentByID(int) (*types.Document, error)
	ListProjectExpenditureByMinistryID(int) ([]*types.ProjectExpenditure, error)
	GetSGDILinkByMinistryID(int) ([]*types.SGDILINK, error)
}

type PostgresStore struct {
	db *sqlx.DB
}

// create a postgres store from a config struc
func NewPostgresStore(c *config.Config) (*PostgresStore, error) {
	connStr := "user=" + c.PostgresUser + " dbname=" + c.PostgresDBName + " password=" + c.PostgresPasword + " sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) CreateMinistry(m *types.Ministry) error {
	sqlStmt := "INSERT INTO ministry (name, created_at) VALUES ($1, $2)"
	resp, err := s.db.Query(sqlStmt, m.Name, m.CreatedAt)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)
	return nil
}

func (s *PostgresStore) DeleteMinistry(id int) error {
	sqlStmt := "DELETE FROM ministry WHERE id = $1"
	_, err := s.db.Query(sqlStmt, id)
	return err
}

func (s *PostgresStore) UpdateMinistry(m *types.Ministry) error {
	return nil
}

func (s *PostgresStore) GetMinistryByID(id int) (*types.Ministry, error) {
	sqlStmt := "SELECT * FROM ministry WHERE id = $1"
	ministry := types.Ministry{}
	err := s.db.Select(&ministry, sqlStmt)

	return &ministry, err
}

func (s *PostgresStore) ListMinistry() ([]*types.Ministry, error) {
	sqlStmt := `SELECT * FROM ministry m where  m."Name" != ''`

	ministries := []*types.Ministry{}
	err := s.db.Select(&ministries, sqlStmt)
	return ministries, err // Return the slice of ministries directly

}

func (s *PostgresStore) ListDocumentByMinistryID(ministryID int) ([]*types.Document, error) {
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
	return documents, err

}

func (s *PostgresStore) GetDocumentByID(id int) (*types.Document, error) {
	sqlStmt := "SELECT * FROM budgetdocuments WHERE id = $1"
	document := types.Document{}
	err := s.db.Get(&document, sqlStmt, id)
	return &document, err
}

func (s *PostgresStore) ListProjectExpenditureByMinistryID(ministryID int) ([]*types.ProjectExpenditure, error) {
	sqlStmt := `select
				ProjectTitle,
				Ministry,
				ValueType,
				ValueAmount,
				ValueYear,
				"Year" as DocumentYear,
				b2."MinistryID" as MinistryID,
				BudgetID,
				DocumentID
			from
				(
				select
					b."ProjectTitle" as ProjectTitle ,
					m."Name" as Ministry,
					b."ValueType" as ValueType,
					b."ValueAmount" as ValueAmount ,
					b."ValueYear" as ValueYear,
					b."DocumentID" as DocumentID,
					b.id as BudgetID,
					b."ParentHeader" as ParentHeader
				from
					budgetdevelopmentprojectsexpenditure b
				inner join ministry m on
					b."MinistryID" = m.id
				where
					b."MinistryID" = $1
					and "ValueType" not in ('', 'OF')
			) s
			inner join budgetdocuments b2 on
				s.DocumentID = b2.id`

	projects := []*types.ProjectExpenditure{}
	err := s.db.Select(&projects, sqlStmt, ministryID)
	return projects, err

}

func (s *PostgresStore) GetSGDILinkByMinistryID(ministryID int) ([]*types.SGDILINK, error) {
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
	links := []*types.SGDILINK{}
	err := s.db.Select(&links, sqlStmt, ministryID)
	return links, err
}
