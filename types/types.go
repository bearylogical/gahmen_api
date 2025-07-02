package types

import (
	"time"
)

type Ministry struct {
	ID        int       `json:"id"`
	Name      string    `db:"Name" json:"name"`
	CreatedAt time.Time `db:"CreatedAt" json:"createdAt"`
}

type Document struct {
	DocumentID   int       `db:"documentid" json:"document_id"`
	Ministry     string    `json:"ministry"`
	DocumentName string    `json:"document_name"`
	Year         int       `json:"year"`
	DocumentPath string    `json:"document_path"`
	MD5Hash      string    `json:"md5_hash"`
	CreatedAt    time.Time `json:"createdAt"`
}

type SGDILink struct {
	ChildName  string  `db:"child" json:"child_name"`
	ParentName *string `db:"parent" json:"parent_name"`
	ChildURL   string  `db:"child_url" json:"child_url"`
	ParentURL  *string `db:"parent_url" json:"parent_url"`
}

type ProjectExpenditure struct {
	ProjectID     int     `json:"project_id" db:"ProjectID"`
	ProjectTitle  string  `json:"project_title" db:"ProjectTitle"`
	Ministry      string  `json:"ministry" db:"Ministry"`
	ValueType     string  `json:"value_type" db:"ValueType"`
	ValueAmount   float32 `json:"value_amount" db:"ValueAmount"`
	ValueYear     int     `json:"value_year" db:"ValueYear"`
	Category      string  `json:"parent_header" db:"Category"`
	DocumentYear  int     `json:"document_year" db:"DocumentYear"`
	MinistryID    int     `json:"ministry_id" db:"MinistryID"`
	DocumentID    int     `json:"document_id" db:"DocumentID"`
	ExpenditureID int     `json:"expenditure_id" db:"ExpenditureID"`
}

type ProgrammeExpenditure struct {
	ProgrammeID            int     `json:"programme_id" db:"ProgrammeID"`
	ProgrammeTitle         string  `json:"programme_title" db:"ProgrammeTitle"`
	Ministry               string  `json:"ministry" db:"Ministry"`
	ValueCode              string  `json:"value_code" db:"ValueCode"`
	ValueAmount            float32 `json:"value_amount" db:"ValueAmount"`
	ValueYear              int     `json:"value_year" db:"ValueYear"`
	ValueName              string  `json:"value_name" db:"ValueName"`
	DocumentYear           int     `json:"document_year" db:"DocumentYear"`
	MinistryID             int     `json:"ministry_id" db:"MinistryID"`
	DocumentID             int     `json:"document_id" db:"DocumentID"`
	ProgrammeExpenditureID int     `json:"expenditure_id" db:"ProgrammeExpenditureID"`
}

type MinistryProject struct {
	ProjectTitle string `db:"ProjectTitle" json:"project_title"`
	MinistryID   int    `db:"MinistryID" json:"ministry_id"`
	ProjectId    int    `db:"ProjectID" json:"project_id"`
}

type Expenditures struct {
	MinistryName    string  `db:"MinistryName" json:"ministry_id"`
	ObjectPath      string  `db:"ObjectPath" json:"object_path"`
	ObjectClass     string  `db:"ObjectClass" json:"object_class"`
	ObjectCode      string  `db:"ObjectCode" json:"object_code"`
	ExpenditureType string  `db:"ExpenditureType" json:"expenditure_type"`
	ValueType       string  `db:"ValueType" json:"value_type"`
	ValueAmount     float32 `db:"ValueAmount" json:"value_amount"`
	ValueYear       int     `db:"ValueYear" json:"value_year"`
}

type Expenditure struct {
	MinistryName string  `db:"MinistryName" json:"ministry_id"`
	ValueType    string  `db:"ValueType" json:"value_type"`
	ValueAmount  float32 `db:"ValueAmount" json:"value_amount"`
	ValueYear    int     `db:"ValueYear" json:"value_year"`
}

type BudgetOpts struct {
	ValueType string `db:"ValueType" json:"value_type"`
	ValueYear int    `db:"ValueYear" json:"value_year"`
}

type Personnel struct {
	Category       string  `db:"category" json:"personnel_type"`
	ParentCategory string  `db:"ParentCategory" json:"category"`
	ValueAmount    *int    `db:"ValueAmount" json:"value_amount"`
	ValueYear      *int    `db:"ValueYear" json:"value_year"`
	ValueType      *string `db:"ValueType" json:"value_type"`
}

type MinistryData struct {
	MinistryName          string                  `db:"MinistryName" json:"ministry_name"`
	MinistryID            int                     `db:"MinistryID" json:"ministry_id"`
	ProgrammeExpenditures []*ProgrammeExpenditure `json:"programme_expenditures"`
	ProjectExpenditures   []*ProjectExpenditure   `json:"project_expenditures"`
	MinistryExpenditures  []*Expenditures `json:"ministry_expenditures"`
	MinistryPersonnel     []*Personnel    `json:"ministry_personnel"`
}

func NewMinistry(name string) *Ministry {
	return &Ministry{
		Name:      name,
		CreatedAt: time.Now().UTC(),
	}
}


