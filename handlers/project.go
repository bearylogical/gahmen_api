package handlers

import (
	"net/http"
	"strings"

	"gahmen-api/helpers"
)

// @Summary List projects by ministry ID
// @Description Get a list of projects for a specific ministry
// @Tags projects
// @Produce  json
// @Param ministry_id path int true "Ministry ID"
// @Success 200 {array} types.MinistryProject "OK" "[{"project_title": "Project Alpha", "ministry_id": 1, "project_id": 1}]"
// @Router /api/v1/projects/ministry/{ministry_id} [get]
// @Security BearerAuth
func (h *Handler) ListProjectByMinistryID(w http.ResponseWriter, r *http.Request) error {
	ministry_id, err := helpers.GetIntByResponseField(r, "ministry_id")
	if err != nil {
		return err
	}
	projects, err := h.store.ListProjectsByMinistryID(ministry_id)
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, projects)
}

// @Summary Get project expenditure by ID
// @Description Get project expenditure by ID
// @Tags projects
// @Produce  json
// @Param project_id path int true "Project ID"
// @Success 200 {object} types.ProjectExpenditure "OK" "{"project_id": 1, "project_title": "Project Alpha", "ministry": "Ministry of Finance", "value_type": "Actual", "value_amount": 500000.0, "value_year": 2023, "parent_header": "Infrastructure", "document_year": 2023, "ministry_id": 1, "document_id": 1, "expenditure_id": 1}"
// @Router /api/v1/projects/{project_id} [get]
// @Security BearerAuth
func (h *Handler) GetProjectExpenditureByID(w http.ResponseWriter, r *http.Request) error {
	project_id, err := helpers.GetIntByResponseField(r, "project_id")
	if err != nil {
		return err
	}
	documents, err := h.store.GetProjectExpenditureByID(project_id)
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, documents)
}

// @Summary Search project expenditure
// @Description Search project expenditure by query with pagination
// @Tags projects
// @Produce  json
// @Param query query string true "Search query (supports AND, OR, NOT, and quoted phrases)"
// @Param limit query int false "Number of results to return (default: 10)"
// @Param offset query int false "Number of results to skip (default: 0)"
// @Success 200 {array} types.ProjectExpenditure
// @Router /api/v1/projects [post]
// @Security BearerAuth
func (h *Handler) GetProjectExpenditureByQuery(w http.ResponseWriter, r *http.Request) error {
	query, err := helpers.GetStringByResponseQuery(r, "query")
	if err != nil {
		return err
	}
	// replace all spaces with ' & ' to allow for multiple search terms
	query = strings.Replace(query, " ", " & ", -1)
	documents, err := h.store.GetProjectExpenditureByQuery(query)
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, documents)
}

// @Summary List project expenditure by ministry ID
// @Description List project expenditure by ministry ID
// @Tags projects
// @Produce  json
// @Param ministry_id path int true "Ministry ID"
// @Success 200 {array} types.ProjectExpenditure "OK" "[{"project_id": 1, "project_title": "Project Alpha", "ministry": "Ministry of Finance", "value_type": "Actual", "value_amount": 500000.0, "value_year": 2023, "parent_header": "Infrastructure", "document_year": 2023, "ministry_id": 1, "document_id": 1, "expenditure_id": 1}]"
// @Security BearerAuth
func (h *Handler) ListProjectExpenditureByMinistryID(w http.ResponseWriter, r *http.Request) error {
	ministry_id, err := helpers.GetIntByResponseField(r, "ministry_id")
	if err != nil {
		return err
	}
	projects, err := h.store.ListProjectExpenditureByMinistryID(ministry_id)
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, projects)
}
