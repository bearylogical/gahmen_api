package handlers

import (
	"gahmen-api/db"
)

type Handler struct {
	store      storage.Storage
}

func NewHandler(store storage.Storage) *Handler {
	return &Handler{store: store}
}