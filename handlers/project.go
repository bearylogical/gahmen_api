package handlers

import (
	"net/http"
	"strings"

	"gahmen-api/helpers"
)

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
// @Tags budget
// @Produce  json
// @Param project_id path int true "Project ID"
// @Success 200 {object} types.ProjectExpenditure
// @Router /projects/{project_id} [get]
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

// @Summary Get project expenditure by query
// @Description Get project expenditure by query
// @Tags budget
// @Produce  json
// @Param query query string true "Query"
// @Success 200 {array} types.ProjectExpenditure
// @Router /projects [post]
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
// @Tags budget
// @Produce  json
// @Param ministry_id path int true "Ministry ID"
// @Success 200 {array} types.ProjectExpenditure
// @Router /budget/{ministry_id}/projects [get]
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
