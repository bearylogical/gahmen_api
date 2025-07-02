package handlers

import (
	"net/http"

	"gahmen-api/helpers"
)

// @Summary Get budget options
// @Description Get budget options
// @Tags budget_options
// @Produce  json
// @Success 200 {array} types.BudgetOpts "OK" "[{"value_type": "Actual", "value_year": 2023}, {"value_type": "Estimated", "value_year": 2023}]"
// @Router /api/v1/budget/options [get]
// @Security BearerAuth
func (h *Handler) GetBudgetOpts(w http.ResponseWriter, r *http.Request) error {
	opts, err := h.store.GetBudgetOpts()
	if err != nil {
		return err
	}
	return helpers.WriteJSON(w, http.StatusOK, opts)
}
