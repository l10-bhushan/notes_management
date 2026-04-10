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
	GetAllNotes() ([]model.Notes, error)
	GetNotesById(id string) (model.Notes, error)
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

func (repo *PostgresNotesRepository) GetAllNotes() ([]model.Notes, error) {
	rows, err := repo.db.Query(context.Background(), `SELECT * FROM notes`)
	if err != nil {

		log.Printf("Error while fetching data: %s", err)
		return []model.Notes{}, err
	}
	defer rows.Close()

	var notes []model.Notes

	for rows.Next() {
		var n model.Notes
		err := rows.Scan(&n.Id, &n.Title, &n.Content, &n.Archived, &n.Created_At, &n.Updated_At)
		if err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating rows: %s", err)
		return nil, err
	}

	return notes, nil
}

func (repo *PostgresNotesRepository) GetNotesById(id string) (model.Notes, error) {

	query := `SELECT * FROM notes WHERE ID = $1`
	row := repo.db.QueryRow(context.Background(), query, id)

	var note model.Notes
	err := row.Scan(&note.Id, &note.Title, &note.Content, &note.Archived, &note.Created_At, &note.Updated_At)
	if err != nil {
		return model.Notes{}, err
	}

	return note, nil
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
