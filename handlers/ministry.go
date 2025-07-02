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
// @Success 200 {array} types.Ministry
// @Router /ministries [get]
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
// @Success 200 {object} types.Ministry
// @Router /ministries/{ministry_id} [get]
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
// @Success 200 {object} types.Ministry
// @Router /ministries [post]
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
// @Param ministryID query int true "Ministry ID"
// @Param topN query int true "Top N"
// @Param startYear query int true "Start Year"
// @Success 200 {object} types.Ministry
// @Router /v2/budget [get]
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