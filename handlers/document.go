package handlers

import (
	"net/http"

	"gahmen-api/helpers"
)

// @Summary List documents by ministry ID
// @Description List documents by ministry ID
// @Tags documents
// @Produce  json
// @Param ministry_id path int true "Ministry ID"
// @Success 200 {array} types.Document
// @Router /budget/{ministry_id}/documents [get]
// @Security BearerAuth
func (h *Handler) ListDocumentByMinistryID(w http.ResponseWriter, r *http.Request) error {
	ministry_id, err := helpers.GetIntByResponseField(r, "ministry_id")
	if err != nil {
		return err
	}
	documents, err := h.store.ListDocumentByMinistryID(ministry_id)
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, documents)
}

func (h *Handler) GetDocumentByID(w http.ResponseWriter, r *http.Request) error {
	id, err := helpers.GetIntByResponseField(r, "document_id")
	if err != nil {
		return err
	}

	document, err := h.store.GetDocumentByID(id)
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, document)
}
