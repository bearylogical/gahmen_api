package handlers

import (
	"net/http"

	"gahmen-api/helpers"
)

// @Summary Get budget options
// @Description Get budget options
// @Tags budget
// @Produce  json
// @Success 200 {object} types.BudgetOpts
// @Router /budget/opts [get]
// @Security BearerAuth
func (h *Handler) GetBudgetOpts(w http.ResponseWriter, r *http.Request) error {
	opts, err := h.store.GetBudgetOpts()
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, opts)
}
