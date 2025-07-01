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
	GetProjectExpenditureByID(int) ([]*types.ProjectExpenditure, error)
	GetProjectExpenditureByQuery(string) ([]*types.ProjectExpenditure, error)
	ListProjectsByMinistryID(int) ([]*types.MinistryProject, error)
	GetSGDILinkByMinistryID(int) ([]*types.SGDILink, error)
	GetProgrammeExpenditureByMinistryID(int) ([]*types.ProgrammeExpenditure, error)
	ListExpenditureByMinistryID(int) ([]*types.Expenditures, error)
	ListExpenditure(string, int) ([]*types.Expenditures, error)
	GetBudgetOpts() ([]*types.BudgetOpts, error)
	ListTopNPersonnelByMinistryID(int, int, int) ([]*types.Personnel, error)
	GetMinistryDataByID(int, int, int) (*types.MinistryData, error)
}

type PostgresStore struct {
	db *sqlx.DB
}

// create a postgres store from a config struc
func NewPostgresStore(c *config.Config) (*PostgresStore, error) {
	connStr := "host=" + c.PostgresHost + " user=" + c.PostgresUser + " dbname=" + c.PostgresDBName + " password=" + c.PostgresPasword + " sslmode=disable" + " port=" + c.PostgresPort
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

func (s *PostgresStore) GetProgrammeExpenditureByMinistryID(ministryID int) ([]*types.ProgrammeExpenditure, error) {
	sqlStmt := `select
				a2."MinistryID" as "MinistryID",
				a2."ProgrammeID" as "ProgrammeID",
				a2."ProgrammeTitle" as "ProgrammeTitle" ,
				a2."ValueCode" as "ValueCode",
				a2."ValueName" as "ValueName",
				a2."ValueAmount" as "ValueAmount",
				a2."ValueYear" as "ValueYear",
				a2."ProgrammeExpenditureID",
				a2."DocumentID" as "DocumentID",
				a2."Name" as "Ministry",
				b3."Year" as "DocumentYear"
			from
				(
				select
					*
				from
					(
					select
						b2."MinistryID" as "MinistryID",
						b2.id as "ProgrammeID",
						b2."ProgrammeTitle" as "ProgrammeTitle" ,
						b."ValueCode" as "ValueCode",
						b."ValueName" as "ValueName",
						b."ValueYear" as "ValueYear",
						b."ValueAmount" as "ValueAmount",
						b.id as "ProgrammeExpenditureID",
						b."DocumentID" as "DocumentID"
					from
						budgetprogrammeexpenditure
					b
					inner join budgetprogramme b2 on
						b."ProgrammeID" = b2.id
					where
						b2."MinistryID" = $1
			) a1
				inner join ministry m on
					a1."MinistryID" = m.id) a2
			inner join budgetdocuments b3 on
				a2."DocumentID" = b3.id
				where a2."ProgrammeTitle" != ''
			`
	programmes := []*types.ProgrammeExpenditure{}
	err := s.db.Select(&programmes, sqlStmt, ministryID)
	return programmes, err
}

func (s *PostgresStore) GetProjectExpenditureByQuery(query string) ([]*types.ProjectExpenditure, error) {
	sqlStmt := `select
				g."ProjectID",
				g."ProjectTitle",
				g."ValueType",
				g."ValueAmount",
				g."ValueYear",
				g."DocumentID",
				g."MinistryID",
				g."Ministry",
				g."Category",
				g."ExpenditureID",
				b5."Year" as "DocumentYear"
			from
				(
				select
					b4."ProjectID",
					b4."ProjectTitle",
					b4."ValueType",
					b4."ValueAmount",
					b4."ValueYear",
					b4."DocumentID",
					b4."MinistryID",
					b4."Ministry",
					b4."ExpenditureID",
					d."Category"
				from
					(
					select
						*,
						m."Name" as "Ministry"
					from
						(
						select
							b.id as "ProjectID",
							b."ProjectTitle",
							b2."ValueType" as "ValueType",
								b2."ValueAmount" as "ValueAmount" ,
								b2."ValueYear" as "ValueYear",
								b2."DocumentID" as "DocumentID",
								b."MinistryID" as "MinistryID",
								b2.id as "ExpenditureID"
						from
							budgetprojects b
						inner join budgetdevelopmentprojectsexpenditure b2 on
							b.id = b2."ProjectID"
							and b2."ValueType" not in ('', 'OF')
						where
							b.search_vector @@ to_tsquery('english',
							$1)
			) b3
					inner join ministry m on
						m.id = b3."MinistryID") b4
				inner join (
					select
						distinct a."ProjectID",
						c."ParentHeader" as "Category"
					from
						(
						select
							max("ValueYear") as "ValueYear",
							"ProjectID"
						from
							budgetdevelopmentprojectsexpenditure u
						group by
							"ProjectID") a
					inner join budgetdevelopmentprojectsexpenditure c on
						a."ProjectID" = c."ProjectID"
						and a."ValueYear" = c."ValueYear"
			) d on
					d."ProjectID" = b4."ProjectID") g
			inner join budgetdocuments b5 on
				b5.id = g."DocumentID"`

	projects := []*types.ProjectExpenditure{}
	err := s.db.Select(&projects, sqlStmt, query)
	return projects, err
}

func (s *PostgresStore) ListProjectExpenditureByMinistryID(ministryID int) ([]*types.ProjectExpenditure, error) {
	sqlStmt := `select g."ProjectID" as "ProjectID",
			g."ProjectTitle" as "ProjectTitle",
			g."ValueType" as "ValueType",
			g."ValueAmount" as "ValueAmount",
			g."ValueYear" as "ValueYear",
			g."DocumentID" as "DocumentID",
			g."MinistryID" as "MinistryID",
			g."Ministry" as "Ministry",
			g."Category" as "Category",
			g."ExpenditureID" as "ExpenditureID",
			b5."Year" as "DocumentYear" from (
		select
			b4."ProjectID",
			b4."ProjectTitle",
			b4."ValueType",
			b4."ValueAmount",
			b4."ValueYear",
			b4."DocumentID",
			b4."MinistryID",
			b4."Ministry",
			b4."ExpenditureID",
			d."Category"
		from
			(
			select
				*,
				m."Name" as "Ministry"
			from
				(
				select
					b.id as "ProjectID",
					b."ProjectTitle",
					b2."ValueType" as "ValueType",
							b2."ValueAmount" as "ValueAmount" ,
							b2."ValueYear" as "ValueYear",
							b2."DocumentID" as "DocumentID",
							b."MinistryID" as "MinistryID",
							b2.id as "ExpenditureID"
				from
					budgetprojects b
				inner join budgetdevelopmentprojectsexpenditure b2 on
					b.id = b2."ProjectID"
					and b2."ValueType" not in ('', 'OF')
				where
					b."MinistryID" = $1) b3
			inner join ministry m on
				m.id = b3."MinistryID") b4
		inner join (
			select
				distinct a."ProjectID",
				c."ParentHeader" as "Category"
			from
				(
				select
					max("ValueYear") as "ValueYear",
					"ProjectID"
				from
					budgetdevelopmentprojectsexpenditure u
				group by
					"ProjectID") a
			inner join budgetdevelopmentprojectsexpenditure c on
				a."ProjectID" = c."ProjectID"
				and a."ValueYear" = c."ValueYear"
		) d on d."ProjectID" = b4."ProjectID") g inner join budgetdocuments b5 on b5.id  = g."DocumentID"`

	projects := []*types.ProjectExpenditure{}
	err := s.db.Select(&projects, sqlStmt, ministryID)
	return projects, err

}

func (s *PostgresStore) GetProjectExpenditureByID(projectID int) ([]*types.ProjectExpenditure, error) {
	sqlStmt := `select g."ProjectID",
				g."ProjectTitle",
				g."ValueType",
				g."ValueAmount",
				g."ValueYear",
				g."DocumentID",
				g."MinistryID",
				g."Ministry",
				g."Category",
				g."ExpenditureID",
				b5."Year" as "DocumentYear" from (
			select
				b4."ProjectID",
				b4."ProjectTitle",
				b4."ValueType",
				b4."ValueAmount",
				b4."ValueYear",
				b4."DocumentID",
				b4."MinistryID",
				b4."Ministry",
				b4."ExpenditureID",
				d."Category"
			from
				(
				select
					*,
					m."Name" as "Ministry"
				from
					(
					select
						b.id as "ProjectID",
						b."ProjectTitle",
						b2."ValueType" as "ValueType",
								b2."ValueAmount" as "ValueAmount" ,
								b2."ValueYear" as "ValueYear",
								b2."DocumentID" as "DocumentID",
								b."MinistryID" as "MinistryID",
								b2.id as "ExpenditureID"
					from
						budgetprojects b
					inner join budgetdevelopmentprojectsexpenditure b2 on
						b.id = b2."ProjectID"
						and b2."ValueType" not in ('', 'OF')
					where
						b.id = $1) b3
				inner join ministry m on
					m.id = b3."MinistryID") b4
			inner join (
				select
					distinct a."ProjectID",
					c."ParentHeader" as "Category"
				from
					(
					select
						max("ValueYear") as "ValueYear",
						"ProjectID"
					from
						budgetdevelopmentprojectsexpenditure u
					group by
						"ProjectID") a
				inner join budgetdevelopmentprojectsexpenditure c on
					a."ProjectID" = c."ProjectID"
					and a."ValueYear" = c."ValueYear"
			) d on d."ProjectID" = b4."ProjectID") g inner join budgetdocuments b5 on b5.id  = g."DocumentID"`
	projects := []*types.ProjectExpenditure{}
	err := s.db.Select(&projects, sqlStmt, projectID)
	return projects, err
}

func (s *PostgresStore) ListExpenditureByMinistryID(ministryID int) ([]*types.Expenditures, error) {
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
								budgetexpenditure b
			where
			b."ObjectCode" <> '9999') a
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
	expenditures := []*types.Expenditures{}
	err := s.db.Select(&expenditures, sqlStmt, ministryID)
	return expenditures, err
}

func (s *PostgresStore) GetSGDILinkByMinistryID(ministryID int) ([]*types.SGDILink, error) {
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
	return links, err
}

func (s *PostgresStore) ListProjectsByMinistryID(ministryID int) ([]*types.MinistryProject, error) {
	sqlStmt := `select
					id as "ProjectID",
				"ProjectTitle" ,
				"MinistryID"
				from
				budgetprojects b
				where
				"MinistryID" = $1
	`

	projects := []*types.MinistryProject{}
	err := s.db.Select(&projects, sqlStmt, ministryID)
	return projects, err
}

func (s *PostgresStore) ListExpenditure(valueType string, valueYear int) ([]*types.Expenditures, error) {
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
	expenditures := []*types.Expenditures{}
	err := s.db.Select(&expenditures, sqlStmt, valueYear, valueType)
	return expenditures, err
}

func (s *PostgresStore) GetBudgetOpts() ([]*types.BudgetOpts, error) {
	sqlStmt := `select
					distinct("ValueType"),
					"ValueYear"
				from
					budgetexpenditure b`
	opts := []*types.BudgetOpts{}
	err := s.db.Select(&opts, sqlStmt)
	return opts, err
}

func (s *PostgresStore) ListTopNPersonnelByMinistryID(ministryID int, n int, startYear int) ([]*types.Personnel, error) {
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
	return opts, err
}

func (s *PostgresStore) GetMinistryDataByID(ministryID int, n int, startYear int) (*types.MinistryData, error) {

	ministry := &types.MinistryData{}

	programmes, err := s.GetProgrammeExpenditureByMinistryID(ministryID)
	if err != nil {
		return nil, err
	}
	projects, err := s.ListProjectExpenditureByMinistryID(ministryID)
	if err != nil {
		return nil, err
	}
	expenditures, err := s.ListExpenditureByMinistryID(ministryID)
	if err != nil {
		return nil, err
	}
	personnel, err := s.ListTopNPersonnelByMinistryID(ministryID, n, startYear)
	if err != nil {
		return nil, err
	}

	ministry.MinistryID = ministryID
	ministry.MinistryName = programmes[0].Ministry
	ministry.MinistryExpenditures = expenditures
	ministry.ProgrammeExpenditures = programmes
	ministry.ProjectExpenditures = projects
	ministry.MinistryPersonnel = personnel

	return ministry, nil
}
