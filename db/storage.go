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
	ListMinistries(bool) ([]*types.Ministry, error)
	ListDocumentByMinistryID(int) ([]*types.Document, error)
	GetDocumentByID(int) (*types.Document, error)
	ListProjectExpenditureByMinistryID(int) ([]*types.ProjectExpenditure, error)
	GetSGDILinkByMinistryID(int) ([]*types.SGDILINK, error)
	ListExpenditureByMinistryID(int) ([]*types.MinistryExpenditureType, error)
	ListExpenditure(string, int) ([]*types.MinistryExpenditureType, error)
	GetBudgetOpts() ([]*types.MinistryExpenditureOptions, error)
	ListTopNPersonnelByMinistryID(int, int, int) ([]*types.MinistryPersonnel, error)
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
	sqlStmt := `SELECT * FROM ministry WHERE id = $1`
	ministry := types.Ministry{}
	err := s.db.Get(&ministry, sqlStmt, id)

	return &ministry, err
}

func (s *PostgresStore) ListMinistries(isMinistry bool) ([]*types.Ministry, error) {

	ministries := []*types.Ministry{}
	if isMinistry {
		sqlStmt := `
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
		err := s.db.Select(&ministries, sqlStmt)
		return ministries, err // Return the slice of ministries directly
	} else {
		sqlStmt := `select
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
		err := s.db.Select(&ministries, sqlStmt)
		return ministries, err // Return the slice of ministries directly
	}

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
				ParentHeader as Category,
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

func (s *PostgresStore) ListExpenditureByMinistryID(ministryID int) ([]*types.MinistryExpenditureType, error) {
	sqlStmt := `select
				"ValueAmount",
							"ExpenditureType" ,
							"ValueType",
							"ValueYear",
							"Name" as "MinistryName"
			from
				(
				select
							sum("ValueAmount") as "ValueAmount",
							"ExpenditureType" ,
							"ValueType",
							"ValueYear",
							"MinistryID"
				from
							(
					select
								distinct "ObjectClass",
								"ObjectCode",
								"ValueType",
								"ValueYear",
								"ValueAmount",
								"MinistryID",
								case
									when cast("ObjectCode" as int) > 5200 then 'OTHER'
							when cast("ObjectCode" as int) > 5000 then 'DEVELOPMENT'
							else 'OPERATING'
						end as "ExpenditureType"
					from
								budgetexpenditure b ) a
				where
							a."MinistryID" = $1
				group by
							a."MinistryID",
							a."ExpenditureType",
							a."ValueType",
							a."ValueYear") c
			left join ministry m on
				m.id = c."MinistryID"
			where
				m."Name" != ''
			`
	expenditures := []*types.MinistryExpenditureType{}
	err := s.db.Select(&expenditures, sqlStmt, ministryID)
	return expenditures, err
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

func (s *PostgresStore) ListExpenditure(valueType string, valueYear int) ([]*types.MinistryExpenditureType, error) {
	sqlStmt := `select
					concat(c."MinistryName",
					'/',
					c."ExpenditureType",
					'/',
					c."ObjectClass" ) as "ObjectPath",
					"ObjectClass",
						"ObjectCode",
						"ValueType",
						"ValueYear",
						"ValueAmount",
						"MinistryName",
						"ExpenditureType"
				from
					(
					select
						distinct "ObjectClass",
						"ObjectCode",
						"ValueType",
						"ValueYear",
						"ValueAmount",
						"MinistryID",
						"Name" as "MinistryName",
						case
							when cast("ObjectCode" as int) > 5200 then 'Other Expenditure'
							when cast("ObjectCode" as int) > 5000 then 'Development Expenditure'
							else 'Operating Expenditure'
						end as "ExpenditureType"
					from
						budgetexpenditure b
					left join ministry m on
						m.id = b."MinistryID"
					where
						"ValueYear" = $1 and
						"ValueType" = $2 and "Name" != '' ) c
			`
	expenditures := []*types.MinistryExpenditureType{}
	err := s.db.Select(&expenditures, sqlStmt, valueYear, valueType)
	return expenditures, err
}

func (s *PostgresStore) GetBudgetOpts() ([]*types.MinistryExpenditureOptions, error) {
	sqlStmt := `select
					distinct("ValueType"),
					"ValueYear"
				from
					budgetexpenditure b`
	opts := []*types.MinistryExpenditureOptions{}
	err := s.db.Select(&opts, sqlStmt)
	return opts, err
}

func (s *PostgresStore) ListTopNPersonnelByMinistryID(ministryID int, n int, startYear int) ([]*types.MinistryPersonnel, error) {
	sqlStmt := `select
					category,
					"ParentHeader" as "ParentCategory" ,
					"ValueAmount",
					"ValueYear",
					"ValueType"
				from
					budgetpersonnel b
				where
					"MinistryID" = $1`
	if startYear > 0 {
		sqlStmt += fmt.Sprintf(" and \"ValueYear\" >= %d ", startYear)
	}
	if n > 0 {
		sqlStmt += fmt.Sprintf(" limit %d", n)
	}

	sqlStmt += `order by
	"ValueAmount" desc`

	println(sqlStmt)
	opts := []*types.MinistryPersonnel{}
	err := s.db.Select(&opts, sqlStmt, ministryID)
	return opts, err
}
