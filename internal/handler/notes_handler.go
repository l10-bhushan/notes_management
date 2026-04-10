package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
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

func (handler *NotesHandler) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := handler.service.GetAllNotes()
	if err != nil {
		error := model.Error{
			Status:  false,
			Message: "Failed to fetch data",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error)
		return
	}

	result := model.Success{
		Status:  true,
		Message: "Fetched all notes",
		Data:    notes,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func (handler *NotesHandler) GetNotesById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Println("Id is : ", id)
	note, err := handler.service.GetNotesById(id)
	if err != nil {
		err := model.Error{
			Status:  false,
			Message: "Error fetching note",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	success := model.Success{
		Status:  true,
		Message: "Note successfully found",
		Data:    note,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success)
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		errMsg := model.Error{
			Status:  false,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errMsg)
		return
	}
	result := model.Success{
		Status:  true,
		Message: "Note created successfully",
		Data:    note,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (handler *NotesHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := handler.service.DeleteNote(id)

	if err != nil {
		error := model.Error{
			Status:  false,
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error)
		return
	}

	success := model.Success{
		Status:  true,
		Message: "Note deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success)
}

func (handler *NotesHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var note model.Notes
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		error := model.Error{
			Status:  false,
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error)
		return
	}
	err = handler.service.UpdateNote(id, note.Title, note.Content)
	if err != nil {
		error := model.Error{
			Status:  false,
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error)
		return
	}

	success := model.Success{
		Status:  true,
		Message: "Note updated successfully",
		Data: model.Notes{
			Id:         note.Id,
			Title:      note.Content,
			Updated_At: time.Now().Local().String(),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success)

}
