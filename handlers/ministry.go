package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gahmen-api/helpers"
	"gahmen-api/types"
)

func (h *Handler) ListMinistry(w http.ResponseWriter, r *http.Request) error {
	ministries, err := h.store.ListMinistry()
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, ministries)
}

func (h *Handler) CreateMinistry(w http.ResponseWriter, r *http.Request) error {
	createAccountRequest := new(types.CreateMinistryRequest)
	if err := json.NewDecoder(r.Body).Decode(createAccountRequest); err != nil {
		return err
	}

	ministry := types.NewMinistry(createAccountRequest.Name)

	if err := h.store.CreateMinistry(ministry); err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusCreated, ministry)
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
