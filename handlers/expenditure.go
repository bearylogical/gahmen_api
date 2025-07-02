package handlers

import (
	"net/http"

	"gahmen-api/helpers"
)

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
