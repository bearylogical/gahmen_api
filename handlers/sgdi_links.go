package handlers

import (
	"net/http"

	"gahmen-api/helpers"
)

func (h *Handler) ListSGDILinksByMinistryID(w http.ResponseWriter, r *http.Request) error {
	ministry_id, err := helpers.GetIDByResponseField(r, "ministry_id")
	if err != nil {
		return err
	}
	links, err := h.store.GetSGDILinkByMinistryID(ministry_id)
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, links)
}
