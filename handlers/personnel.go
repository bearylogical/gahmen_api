package handlers

import (
	"net/http"

	"gahmen-api/helpers"
)

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
