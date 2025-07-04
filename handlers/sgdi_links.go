package handlers

import (
	"net/http"

	"gahmen-api/helpers"
)

// ListSGDILinksByMinistryID retrieves a list of SGDILinks based on the provided ministry ID.
// It takes an HTTP response writer and request as parameters and returns an error if any.
// @Summary List SGDI links by ministry ID
// @Description List SGDI links by ministry ID
// @Tags sgdi_links
// @Produce  json
// @Param ministry_id path int true "Ministry ID"
// @Success 200 {array} types.SGDILink "OK" "[{"child_name": "Child Org", "parent_name": "Parent Org", "child_url": "http://child.example.com", "parent_url": "http://parent.example.com"}]"
// @Router /api/v1/sgdi-links/ministry/{ministry_id} [get]
// @Security BearerAuth
func (h *Handler) ListSGDILinksByMinistryID(w http.ResponseWriter, r *http.Request) error {
	ministry_id, err := helpers.GetIntByResponseField(r, "ministry_id")
	if err != nil {
		return err
	}
	links, err := h.store.GetSGDILinkByMinistryID(ministry_id)
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, links)
}
