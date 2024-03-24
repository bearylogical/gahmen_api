package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetIDByResponseField(r *http.Request, fieldName string) (int, error) {
	idStr := r.PathValue(fieldName)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id given %s", idStr)
	}

	return id, nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
