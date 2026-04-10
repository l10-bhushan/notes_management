package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/l10-bhushan/notes_management/internal/model"
	"github.com/l10-bhushan/notes_management/internal/service"
)

type NotesHandler struct {
	service *service.NotesService
}

func NewNotesHandler(service *service.NotesService) *NotesHandler {
	return &NotesHandler{
		service: service,
	}
}

func (handler *NotesHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var req model.NotesCreationRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		log.Fatal("Failed to parse: Bad request")
		return
	}

	note, err := handler.service.CreateNote(req)
	if err != nil {
		log.Fatal("Failed to create user...")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}
