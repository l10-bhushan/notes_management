package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/l10-bhushan/notes_management/internal/model"
	"github.com/l10-bhushan/notes_management/internal/repository"
)

type NotesService struct {
	repo *repository.PostgresNotesRepository
}

func NewService(repo *repository.PostgresNotesRepository) *NotesService {
	return &NotesService{
		repo: repo,
	}
}

func (service *NotesService) GetAllNotes() ([]model.Notes, error) {
	notes, err := service.repo.GetAllNotes()
	fmt.Println(err)
	if err != nil {
		return []model.Notes{}, err
	}
	return notes, nil
}

func (service *NotesService) GetNotesById(id string) (model.Notes, error) {
	note, err := service.repo.GetNotesById(id)
	fmt.Println("Error from service: ", err)
	if err != nil {
		return model.Notes{}, err
	}

	return note, nil
}

func (service *NotesService) CreateNote(data model.NotesCreationRequest) (model.Notes, error) {

	id := uuid.New().String()
	title := data.Title
	content := data.Content
	note := model.Notes{
		Id:         id,
		Title:      title,
		Content:    content,
		Archived:   false,
		Created_At: time.Now().Local().String(),
		Updated_At: time.Now().Local().String(),
	}

	createdNote, err := service.repo.CreateNote(note)

	if err != nil {
		return model.Notes{}, err
	}

	return createdNote, err
}

func (service *NotesService) DeleteNote(id string) error {
	err := service.repo.DeleteNote(id)
	return err
}
