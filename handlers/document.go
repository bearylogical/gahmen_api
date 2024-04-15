package handlers

import (
	"net/http"

	"gahmen-api/helpers"
)

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
