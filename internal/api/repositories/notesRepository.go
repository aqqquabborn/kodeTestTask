package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"kodeTestTask/internal/api/models"
	"log"
)

type NotesRepository interface {
	Create(ctx context.Context, note *models.Note) error
	GetByID(ctx context.Context, noteID int) (*models.Note, error)
	GetAllByUserID(ctx context.Context, userID int) ([]*models.Note, error)
	Update(ctx context.Context, note *models.Note) error
	Delete(ctx context.Context, noteID int) error
}

type notesRepository struct {
	db *sql.DB
}

func NewNotesRepository(db *sql.DB) NotesRepository {
	return &notesRepository{db}
}

func (r *notesRepository) Create(ctx context.Context, note *models.Note) error {
	query := "INSERT INTO notes (title, content, user_id) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, note.Title, note.Content, note.UserID).Scan(&note.ID)
	if err != nil {
		log.Printf("Error creating note: %v", err)
		return err
	}
	return nil
}

func (r *notesRepository) GetByID(ctx context.Context, noteID int) (*models.Note, error) {
	query := "SELECT id, title, content, user_id, created_at, updated_at FROM notes WHERE id = $1"
	var note models.Note
	err := r.db.QueryRowContext(ctx, query, noteID).Scan(&note.ID, &note.Title, &note.Content, &note.UserID, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("note not found")
		}
		log.Printf("Error getting note by ID: %v", err)
		return nil, err
	}
	return &note, nil
}

func (r *notesRepository) Update(ctx context.Context, note *models.Note) error {
	query := "UPDATE notes SET title = $1, content = $2, updated_at = NOW() WHERE id = $3"
	_, err := r.db.ExecContext(ctx, query, note.Title, note.Content, note.ID)
	if err != nil {
		log.Printf("Error updating note: %v", err)
		return err
	}
	return nil
}

func (r *notesRepository) Delete(ctx context.Context, noteID int) error {
	query := "DELETE FROM notes WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, noteID)
	if err != nil {
		log.Printf("Error deleting note: %v", err)
		return err
	}
	return nil
}

func (r *notesRepository) GetAllByUserID(ctx context.Context, userID int) ([]*models.Note, error) {
	query := "SELECT id, title, content, user_id, created_at, updated_at FROM notes WHERE user_id = $1"
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		log.Printf("Error getting all notes by user ID: %v", err)
		return nil, err
	}
	defer rows.Close()

	var notes []*models.Note
	for rows.Next() {
		var note models.Note
		err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.UserID, &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			log.Printf("Error scanning note row: %v", err)
			return nil, err
		}
		notes = append(notes, &note)
	}
	return notes, nil
}
