package handlers

import (
	"fmt"
	"net/http"

	"gahmen-api/helpers"
)

func (h *Handler) ListMinistry(w http.ResponseWriter, r *http.Request) error {
	ministries, err := h.store.ListMinistry()
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, ministries)
}

func (h *Handler) GetMinistryByID(w http.ResponseWriter, r *http.Request) error {
	id, err := helpers.GetIDByResponseField(r, "ministry_id")
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
	id, err := helpers.GetIDByResponseField(r, "ministry_id")
	if err != nil {
		return err
	}
	if err := h.store.DeleteMinistry(id); err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, fmt.Sprintf("ministry with id %d deleted", id))
}
