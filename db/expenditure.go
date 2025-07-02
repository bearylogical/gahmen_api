package storage

import (
	"fmt"
	"gahmen-api/types"

	"github.com/jmoiron/sqlx"
)

type ExpenditureStore struct {
	db *sqlx.DB
}

func NewExpenditureStore(db *sqlx.DB) *ExpenditureStore {
	return &ExpenditureStore{db: db}
}

func (s *ExpenditureStore) GetProgrammeExpenditureByMinistryID(ministryID int) ([]*types.ProgrammeExpenditure, error) {
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
	if err != nil {
		return nil, fmt.Errorf("error getting programme expenditure by ministry ID: %w", err)
	}
	return programmes, nil
}

func (s *ExpenditureStore) GetProjectExpenditureByQuery(query string, limit, offset int) ([]*types.ProjectExpenditure, error) {
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
			) d on					d."ProjectID" = b4."ProjectID") g
			inner join budgetdocuments b5 on
				b5.id = g."DocumentID" LIMIT $2 OFFSET $3`

	projects := []*types.ProjectExpenditure{}
	err := s.db.Select(&projects, sqlStmt, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error getting project expenditure by query: %w", err)
	}
	return projects, nil
}

func (s *ExpenditureStore) ListProjectExpenditureByMinistryID(ministryID int) ([]*types.ProjectExpenditure, error) {
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
	if err != nil {
		return nil, fmt.Errorf("error listing project expenditure by ministry ID: %w", err)
	}
	return projects, nil

}

func (s *ExpenditureStore) GetProjectExpenditureByID(projectID int) ([]*types.ProjectExpenditure, error) {
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
	if err != nil {
		return nil, fmt.Errorf("error getting project expenditure by ID: %w", err)
	}
	return projects, nil
}

func (s *ExpenditureStore) ListExpenditureByMinistryID(ministryID int) ([]*types.Expenditures, error) {
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
	if err != nil {
		return nil, fmt.Errorf("error listing expenditure by ministry ID: %w", err)
	}
	return expenditures, nil
}

func (s *ExpenditureStore) ListExpenditure(valueType string, valueYear int) ([]*types.Expenditures, error) {
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
	if err != nil {
		return nil, fmt.Errorf("error listing expenditure: %w", err)
	}
	return expenditures, nil
}
