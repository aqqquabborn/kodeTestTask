package usecases

import (
	"context"
	"kodeTestTask/internal/api/models"
	. "kodeTestTask/internal/api/repositories"
)

type NotesUsecase interface {
	CreateNote(ctx context.Context, note *models.Note) error
	GetNoteByID(ctx context.Context, noteID int) (*models.Note, error)
	GetAllByUserID(ctx context.Context, userID int) ([]*models.Note, error)
	UpdateNote(ctx context.Context, note *models.Note) error
	DeleteNote(ctx context.Context, noteID int) error
}

type notesUsecase struct {
	notesRepository NotesRepository
}

func NewNotesUsecase(repository NotesRepository) NotesUsecase {
	return &notesUsecase{repository}
}

func (uc *notesUsecase) CreateNote(ctx context.Context, note *models.Note) error {
	return uc.notesRepository.Create(ctx, note)
}

func (uc *notesUsecase) GetNoteByID(ctx context.Context, noteID int) (*models.Note, error) {
	return uc.notesRepository.GetByID(ctx, noteID)
}

func (uc *notesUsecase) UpdateNote(ctx context.Context, note *models.Note) error {
	return uc.notesRepository.Update(ctx, note)
}

func (uc *notesUsecase) GetAllByUserID(ctx context.Context, userID int) ([]*models.Note, error) {
	return uc.notesRepository.GetAllByUserID(ctx, userID)
}

func (uc *notesUsecase) DeleteNote(ctx context.Context, noteID int) error {
	return uc.notesRepository.Delete(ctx, noteID)
}
