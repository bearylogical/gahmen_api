package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetIntByResponseField(r *http.Request, fieldName string) (int, error) {
	idStr := r.PathValue(fieldName)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid field value given %s", idStr)
	}

	return id, nil
}
func GetIntByResponseQuery(r *http.Request, fieldName string) (int, error) {
	idStr := r.URL.Query().Get(fieldName)
	if idStr == "" {
		return -1, nil
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid field value given %s", idStr)
	}

	return id, nil
}

func GetStringByResponseQuery(r *http.Request, fieldName string) (string, error) {
	idStr := r.URL.Query().Get(fieldName)
	if idStr == "" {
		return idStr, fmt.Errorf("missing field value %s", idStr)
	}
	return idStr, nil
}

func GetBoolByResponseQuery(r *http.Request, fieldName string) (bool, error) {
	idStr := r.URL.Query().Get(fieldName)
	if idStr == "" {
		return false, nil
	}
	// idStr := r.PathValue(fieldName)
	boolVal, err := strconv.ParseBool(idStr)
	if err != nil {
		return boolVal, fmt.Errorf("invalid field given %s", idStr)
	}

	return boolVal, nil
}

func GetQueryParams(r *http.Request, fieldName string) (string, error) {
	query := r.URL.Query()
	value := query.Get(fieldName)
	if value == "" {
		return value, fmt.Errorf("missing query param %s", fieldName)
	}
	return value, nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
