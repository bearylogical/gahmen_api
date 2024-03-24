package types

import (
	"time"
)

type CreateMinistryRequest struct {
	Name string `json:"name"`
}

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
	ChildName  string `db:"child" json:"child_name"`
	ParentName string `db:"parent" json:"parent_name"`
	URL        string `db:"url" json:"url"`
}

type ProjectExpenditure struct {
	ProjectTitle string  `json:"project_title"`
	Ministry     string  `json:"ministry"`
	ValueType    string  `json:"value_type"`
	ValueAmount  float32 `json:"value_amount"`
	ValueYear    int     `json:"value_year"`
	DocumentYear int     `json:"document_year"`
	MinistryID   int     `json:"ministry_id"`
	DocumentID   int     `json:"document_id"`
	BudgetID     int     `json:"budget_id"`
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
