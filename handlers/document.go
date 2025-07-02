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
// @Success 200 {array} types.Document "OK" "[{"document_id": 1, "ministry": "Ministry of Finance", "document_name": "Budget 2023", "year": 2023, "document_path": "/docs/budget2023.pdf", "md5_hash": "abcdef12345", "createdAt": "2023-01-01T00:00:00Z"}]"
// @Router /api/v1/documents/ministry/{ministry_id} [get]
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

// @Summary Get document by ID
// @Description Get document by ID
// @Tags documents
// @Produce  json
// @Param document_id path int true "Document ID"
// @Success 200 {object} types.Document "OK" "{"document_id": 1, "ministry": "Ministry of Finance", "document_name": "Budget 2023", "year": 2023, "document_path": "/docs/budget2023.pdf", "md5_hash": "abcdef12345", "createdAt": "2023-01-01T00:00:00Z"}"
// @Router /api/v1/documents/{document_id} [get]
// @Security BearerAuth
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
