package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gahmen-api/helpers"
	"gahmen-api/types"
)

// @Summary List all ministries
// @Description Get all ministries
// @Tags ministries
// @Produce  json
// @Success 200 {array} types.Ministry "OK" "[{"id": 1, "name": "Ministry of Finance", "createdAt": "2023-01-01T00:00:00Z"}]"
// @Router /api/v1/ministries [get]
// @Security BearerAuth
func (h *Handler) ListMinistries(w http.ResponseWriter, r *http.Request) error {
	ministryFlag, err := helpers.GetBoolByResponseQuery(r, "isMinistry")
	if err != nil {
		return err
	}
	ministries, err := h.store.ListMinistries(ministryFlag)
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, ministries)
}

// @Summary Get a ministry by ID
// @Description Get a ministry by ID
// @Tags ministries
// @Produce  json
// @Param ministry_id path int true "Ministry ID"
// @Success 200 {object} types.Ministry "OK" "{"id": 1, "name": "Ministry of Finance", "createdAt": "2023-01-01T00:00:00Z"}"
// @Router /api/v1/ministries/{ministry_id} [get]
// @Security BearerAuth
func (h *Handler) GetMinistryByID(w http.ResponseWriter, r *http.Request) error {
	id, err := helpers.GetIntByResponseField(r, "ministry_id")
	if err != nil {
		return err
	}
	ministry, err := h.store.GetMinistryByID(id)
	if err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, ministry)
}

// @Summary Create a new ministry
// @Description Create a new ministry
// @Tags ministries
// @Accept  json
// @Produce  json
// @Param ministry body types.Ministry true "Ministry object to be created"
// @Success 200 {object} types.Ministry "OK" "{"id": 1, "name": "Ministry of Finance", "createdAt": "2023-01-01T00:00:00Z"}"
// @Router /api/v1/ministries [post]
// @Security BearerAuth
func (h *Handler) CreateMinistry(w http.ResponseWriter, r *http.Request) error {
	ministry := new(types.Ministry)
	if err := json.NewDecoder(r.Body).Decode(ministry); err != nil {
		return err
	}

	if err := h.store.CreateMinistry(ministry); err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, ministry)
}

// @Summary Delete a ministry by ID
// @Description Delete a ministry by ID
// @Tags ministries
// @Produce  json
// @Param ministry_id path int true "Ministry ID"
// @Success 200 {string} string "ministry with id {id} deleted"
// @Router /api/v1/ministries/{ministry_id} [delete]
// @Security BearerAuth
func (h *Handler) DeleteMinistryByID(w http.ResponseWriter, r *http.Request) error {
	id, err := helpers.GetIntByResponseField(r, "ministry_id")
	if err != nil {
		return err
	}
	if err := h.store.DeleteMinistry(id); err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, fmt.Sprintf("ministry with id %d deleted", id))
}

// @Summary Get ministry data by ID
// @Description Get ministry data by ID
// @Tags ministries
// @Produce  json
// @Param ministryID query int true "Ministry ID" ["1", "2", "3"]
// @Param topN query int false "Top N" ["1", "5", "10"]
// @Param startYear query int true "Start Year" ["2020", "2021", "2022", "2023"]
// @Success 200 {object} types.MinistryData "OK" "{"ministry_name": "Ministry of Finance", "ministry_id": 1, "programme_expenditures": [{"programme_id": 1, "programme_title": "Education Programme", "ministry": "Ministry of Education", "value_code": "VC001", "value_amount": 1000000.0, "value_year": 2023, "value_name": "Salaries", "document_year": 2023, "ministry_id": 1, "document_id": 1, "expenditure_id": 1}], "project_expenditures": [{"project_id": 1, "project_title": "Project Alpha", "ministry": "Ministry of Finance", "value_type": "Actual", "value_amount": 500000.0, "value_year": 2023, "parent_header": "Infrastructure", "document_year": 2023, "ministry_id": 1, "document_id": 1, "expenditure_id": 1}], "ministry_expenditures": [{"ministry_id": "Ministry of Finance", "object_path": "Ministry of Finance/OPERATING/5000", "object_class": "5000", "object_code": "5000", "expenditure_type": "OPERATING", "value_type": "Actual", "value_amount": 1000000.0, "value_year": 2023}], "ministry_personnel": [{"personnel_type": "Admin", "category": "Support", "value_amount": 100, "value_year": 2023, "value_type": "Headcount"}]}"
// @Router /api/v2/budget [get]
// @Security BearerAuth
func (h *Handler) GetMinistryDataV2(w http.ResponseWriter, r *http.Request) error {
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
	ministry, err := h.store.GetMinistryDataByID(ministry_id, top_n, start_year)
	if err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, ministry)
}
