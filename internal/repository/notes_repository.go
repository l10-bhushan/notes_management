package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/l10-bhushan/notes_management/internal/model"
)

// Here we created the repository interface that would have functions that need to be
// Taken care of
type NotesRepository interface {
	// GetAllNotes(ctx context.Context) any
	// GetNotesById(ctx context.Context) any
	CreateNote(data model.Notes) (model.Notes, error)
	// DeleteNote(ctx context.Context)
	// UpdateNote(ctx context.Context, data any) any
	// FilterNote(ctx context.Context, category string) any
}

type PostgresNotesRepository struct {
	db *pgxpool.Pool
}

func NewPostGresNotesRepository(db *pgxpool.Pool) *PostgresNotesRepository {
	return &PostgresNotesRepository{
		db: db,
	}
}

func (repo *PostgresNotesRepository) CreateNote(data model.Notes) (model.Notes, error) {
	query := `INSERT INTO notes (id , title , content , archived , created_at , updated_at) VALUES ($1 , $2, $3 , $4 , $5, $6)`
	_, err := repo.db.Exec(context.Background(), query, data.Id, data.Title, data.Content, data.Archived, data.Created_At, data.Updated_At)

	if err != nil {
		log.Printf("Error while creating note: %s", err)
		return model.Notes{}, err
	}

	return data, nil
}
