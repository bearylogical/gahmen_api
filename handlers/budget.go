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

func (h *Handler) GetProjectExpenditureByQuery(w http.ResponseWriter, r *http.Request) error {
	query, err := helpers.GetStringByResponseQuery(r, "query")
	if err != nil {
		return err
	}
	// replace all spaces with ' & ' to allow for multiple search terms
	query = strings.Replace(query, " ", " & ", -1)
	println(query)
	documents, err := h.store.GetProjectExpenditureByQuery(query)
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, documents)
}

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

func (h *Handler) GetBudgetOpts(w http.ResponseWriter, r *http.Request) error {
	opts, err := h.store.GetBudgetOpts()
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, opts)
}

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
