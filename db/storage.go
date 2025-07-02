package storage

import (
	"gahmen-api/config"
	"gahmen-api/types"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage interface {
	GetMinistryDataByID(int, int, int) (*types.MinistryData, error)
	// ProjectStore methods
	ListProjectsByMinistryID(ministryID int) ([]*types.ProjectExpenditure, error)
	// ExpenditureStore methods
	GetProgrammeExpenditureByMinistryID(ministryID int) ([]*types.ProgrammeExpenditure, error)
	GetProjectExpenditureByQuery(query string, limit, offset int) ([]*types.ProjectExpenditure, error)
	ListProjectExpenditureByMinistryID(ministryID int) ([]*types.ProjectExpenditure, error)
	GetProjectExpenditureByID(projectID int) ([]*types.ProjectExpenditure, error)
	ListExpenditureByMinistryID(ministryID int) ([]*types.Expenditures, error)
	ListExpenditure(valueType string, valueYear int) ([]*types.Expenditures, error)
	// PersonnelStore methods
	ListTopNPersonnelByMinistryID(ministryID int, n int, startYear int) ([]*types.Personnel, error)
	// BudgetOptsStore methods
	GetBudgetOpts() ([]*types.BudgetOpts, error)
	// DocumentStore methods
	ListDocumentByMinistryID(ministryID int) ([]*types.Document, error)
	GetDocumentByID(id int) (*types.Document, error)
	// MinistryStore methods
	CreateMinistry(m *types.Ministry) error
	DeleteMinistry(id int) error
	GetMinistryByID(id int) (*types.Ministry, error)
	ListMinistries(isMinistry bool) ([]*types.Ministry, error)
	// SGDILinkStore methods
	GetSGDILinkByMinistryID(ministryID int) ([]*types.SGDILink, error)
}

type PostgresStore struct {
	*MinistryStore
	*DocumentStore
	*ExpenditureStore
	*PersonnelStore
	*SGDILinkStore
	*BudgetOptsStore
	*ProjectStore
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

	return &PostgresStore{
		MinistryStore:    NewMinistryStore(db),
		DocumentStore:    NewDocumentStore(db),
		ExpenditureStore: NewExpenditureStore(db),
		PersonnelStore:   NewPersonnelStore(db),
		SGDILinkStore:    NewSGDILinkStore(db),
		BudgetOptsStore:  NewBudgetOptsStore(db),
		ProjectStore:     NewProjectStore(db),
		db:               db,
	}, nil
}

func (s *PostgresStore) GetMinistryDataByID(ministryID int, n int, startYear int) (*types.MinistryData, error) {

	ministry := &types.MinistryData{}

	programmes, err := s.ExpenditureStore.GetProgrammeExpenditureByMinistryID(ministryID)
	if err != nil {
		return nil, err
	}
	projects, err := s.ExpenditureStore.ListProjectExpenditureByMinistryID(ministryID)
	if err != nil {
		return nil, err
	}
	expenditures, err := s.ExpenditureStore.ListExpenditureByMinistryID(ministryID)
	if err != nil {
		return nil, err
	}
	personnel, err := s.PersonnelStore.ListTopNPersonnelByMinistryID(ministryID, n, startYear)
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
