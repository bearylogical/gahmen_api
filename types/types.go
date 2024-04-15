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

type SGDILINK struct {
	ChildName  string  `db:"child" json:"child_name"`
	ParentName *string `db:"parent" json:"parent_name"`
	ChildURL   string  `db:"child_url" json:"child_url"`
	ParentURL  *string `db:"parent_url" json:"parent_url"`
}

type ProjectExpenditure struct {
	ProjectTitle string  `json:"project_title"`
	Ministry     string  `json:"ministry"`
	ValueType    string  `json:"value_type"`
	ValueAmount  float32 `json:"value_amount"`
	ValueYear    int     `json:"value_year"`
	Category     string  `json:"parent_header"`
	DocumentYear int     `json:"document_year"`
	MinistryID   int     `json:"ministry_id"`
	DocumentID   int     `json:"document_id"`
	BudgetID     int     `json:"budget_id"`
}

type MinistryExpenditureType struct {
	MinistryName    string  `db:"MinistryName" json:"ministry_id"`
	ObjectPath      string  `db:"ObjectPath" json:"object_path"`
	ObjectClass     string  `db:"ObjectClass" json:"object_class"`
	ObjectCode      string  `db:"ObjectCode" json:"object_code"`
	ExpenditureType string  `db:"ExpenditureType" json:"expenditure_type"`
	ValueType       string  `db:"ValueType" json:"value_type"`
	ValueAmount     float32 `db:"ValueAmount" json:"value_amount"`
	ValueYear       int     `db:"ValueYear" json:"value_year"`
}

type MinistryExpenditure struct {
	MinistryName string  `db:"MinistryName" json:"ministry_id"`
	ValueType    string  `db:"ValueType" json:"value_type"`
	ValueAmount  float32 `db:"ValueAmount" json:"value_amount"`
	ValueYear    int     `db:"ValueYear" json:"value_year"`
}

type MinistryExpenditureOptions struct {
	ValueType string `db:"ValueType" json:"value_type"`
	ValueYear int    `db:"ValueYear" json:"value_year"`
}

type MinistryPersonnel struct {
	Category       string  `db:"category" json:"personnel_type"`
	ParentCategory string  `db:"ParentCategory" json:"category"`
	ValueAmount    *int    `db:"ValueAmount" json:"value_amount"`
	ValueYear      *int    `db:"ValueYear" json:"value_year"`
	ValueType      *string `db:"ValueType" json:"value_type"`
}

func NewMinistry(name string) *Ministry {
	return &Ministry{
		Name:      name,
		CreatedAt: time.Now().UTC(),
	}
}

func NewDocument(ministry string, documentName string, year int) *Document {
	return &Document{
		Ministry:     ministry,
		DocumentName: documentName,
		Year:         year,
	}
}
