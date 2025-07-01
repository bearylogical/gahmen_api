package handlers

import (
	"gahmen-api/helpers"
	"net/http"
	"strings"
)

//	func (s *APIServer) handleBudgetProjects(w http.ResponseWriter, r *http.Request) error {
//		switch r.Method {
//		case "GET":
//			return s.handleListProjectByMinistryID(w, r)
//		}
//		return fmt.Errorf("method not allowed %s", r.Method)
//	}

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

// @Summary Get programme expenditure by ministry ID
// @Description Get programme expenditure by ministry ID
// @Tags budget
// @Produce  json
// @Param ministry_id path int true "Ministry ID"
// @Success 200 {array} types.ProgrammeExpenditure
// @Router /budget/{ministry_id}/programmes [get]
// @Security BearerAuth
func (h *Handler) GetProgrammeExpenditureByMinistryID(w http.ResponseWriter, r *http.Request) error {
	ministry_id, err := helpers.GetIntByResponseField(r, "ministry_id")
	if err != nil {
		return err
	}
	documents, err := h.store.GetProgrammeExpenditureByMinistryID(ministry_id)
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

// @Summary List expenditure by ministry
// @Description List expenditure by ministry
// @Tags budget
// @Produce  json
// @Param ministry_id path int true "Ministry ID"
// @Success 200 {array} types.Expenditure
// @Router /budget/{ministry_id} [get]
// @Security BearerAuth
func (h *Handler) ListExpenditureByMinistry(w http.ResponseWriter, r *http.Request) error {
	ministry_id, err := helpers.GetIntByResponseField(r, "ministry_id")
	if err != nil {
		return err
	}
	documents, err := h.store.ListExpenditureByMinistryID(ministry_id)
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, documents)
}

// @Summary List expenditure
// @Description List expenditure
// @Tags budget
// @Produce  json
// @Param valueYear query int true "Value Year"
// @Param valueType query string true "Value Type"
// @Success 200 {array} types.Expenditure
// @Router /budget [get]
// @Security BearerAuth
func (h *Handler) ListExpenditure(w http.ResponseWriter, r *http.Request) error {
	value_year, err := helpers.GetIntByResponseQuery(r, "valueYear")
	if err != nil {
		return err
	}
	value_type, err := helpers.GetStringByResponseQuery(r, "valueType")
	if err != nil {
		return err
	}
	documents, err := h.store.ListExpenditure(value_type, value_year)
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, documents)
}

// @Summary Get budget options
// @Description Get budget options
// @Tags budget
// @Produce  json
// @Success 200 {object} types.BudgetOpts
// @Router /budget/opts [get]
// @Security BearerAuth
func (h *Handler) GetBudgetOpts(w http.ResponseWriter, r *http.Request) error {
	opts, err := h.store.GetBudgetOpts()
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, opts)
}

// @Summary List top N personnel by ministry ID
// @Description List top N personnel by ministry ID
// @Tags personnel
// @Produce  json
// @Param ministryID query int true "Ministry ID"
// @Param topN query int true "Top N"
// @Param startYear query int true "Start Year"
// @Success 200 {array} types.Personnel
// @Router /personnel [get]
// @Security BearerAuth
func (h *Handler) ListTopNPersonnelByMinistryID(w http.ResponseWriter, r *http.Request) error {
	ministry_id, err := helpers.GetIntByResponseQuery(r, "ministryID")
	if err != nil {
		return err
	}
	top_n, err := helpers.GetIntByResponseQuery(r, "topN")
	if err != nil {
		return err
	}

	start_year, err := helpers.GetIntByResponseQuery(r, "startYear")
	if err != nil {
		return err
	}
	personnel, err := h.store.ListTopNPersonnelByMinistryID(ministry_id, top_n, start_year)
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, personnel)
}
