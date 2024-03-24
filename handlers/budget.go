package handlers

import (
	"gahmen-api/helpers"
	"net/http"
)

//	func (s *APIServer) handleBudgetProjects(w http.ResponseWriter, r *http.Request) error {
//		switch r.Method {
//		case "GET":
//			return s.handleListProjectByMinistryID(w, r)
//		}
//		return fmt.Errorf("method not allowed %s", r.Method)
//	}
func (h *Handler) ListProjectByMinistryID(w http.ResponseWriter, r *http.Request) error {
	ministry_id, err := helpers.GetIDByResponseField(r, "ministry_id")
	if err != nil {
		return err
	}
	documents, err := h.store.ListProjectExpenditureByMinistryID(ministry_id)
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, documents)
}
