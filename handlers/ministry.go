package handlers

import (
	"fmt"
	"net/http"

	"gahmen-api/helpers"
)

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

func (h *Handler) UpdateMinistryByID(w http.ResponseWriter, r *http.Request) error {
	return nil
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
