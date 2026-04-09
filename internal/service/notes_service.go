package service

import (
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

func (service *NotesService) CreateNote(data model.Notes) (model.Notes, error) {
	data, err := service.repo.CreateNote(data)

	if err != nil {
		return model.Notes{}, err
	}

	return data, err
}
