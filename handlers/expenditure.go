package handlers

import (
	"net/http"

	"gahmen-api/helpers"
)

// @Summary Get programme expenditure by ministry ID
// @Description Get programme expenditure by ministry ID
// @Tags expenditures
// @Produce  json
// @Param ministry_id path int true "Ministry ID"
// @Success 200 {array} types.ProgrammeExpenditure "OK" "[{"programme_id": 1, "programme_title": "Education Programme", "ministry": "Ministry of Education", "value_code": "VC001", "value_amount": 1000000.0, "value_year": 2023, "value_name": "Salaries", "document_year": 2023, "ministry_id": 1, "document_id": 1, "expenditure_id": 1}]"
// @Router /api/v1/expenditures/programmes/{ministry_id} [get]
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

// @Summary List all expenditures by ministry ID
// @Description Get a list of all expenditures for a specific ministry
// @Tags expenditures
// @Produce  json
// @Param ministry_id path int true "Ministry ID"
// @Success 200 {array} types.Expenditures "OK" "[{"ministry_id": "Ministry of Finance", "object_path": "Ministry of Finance/OPERATING/5000", "object_class": "5000", "object_code": "5000", "expenditure_type": "OPERATING", "value_type": "Actual", "value_amount": 1000000.0, "value_year": 2023}]"
// @Router /api/v1/expenditures/ministry/{ministry_id} [get]
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

// @Summary List all expenditures
// @Description Get a list of all expenditures filtered by value year and type
// @Tags expenditures
// @Produce  json
// @Param valueYear query int true "Value Year"
// @Param valueType query string true "Value Type"
// @Success 200 {array} types.Expenditures "OK" "[{"ministry_id": "Ministry of Finance", "object_path": "Ministry of Finance/OPERATING/5000", "object_class": "5000", "object_code": "5000", "expenditure_type": "OPERATING", "value_type": "Actual", "value_amount": 1000000.0, "value_year": 2023}]"
// @Router /api/v1/expenditures [get]
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
