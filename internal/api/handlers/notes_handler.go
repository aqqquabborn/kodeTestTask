package handlers

import (
	"encoding/json"
	"kodeTestTask/internal/api/auth"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"kodeTestTask/internal/api/models"
	"kodeTestTask/internal/api/usecases"
)

type NotesHandler struct {
	notesUsecase usecases.NotesUsecase
}

func NewNotesHandler(usecase usecases.NotesUsecase) *NotesHandler {
	return &NotesHandler{usecase}
}

func (h *NotesHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenString := r.Header.Get("Authorization")

	currentUserID, err := auth.GetCurrentUserIDFromToken(tokenString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if note.UserID != currentUserID {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	if err := h.notesUsecase.CreateNote(r.Context(), &note); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func (h *NotesHandler) GetNoteByID(w http.ResponseWriter, r *http.Request) {
	noteID, err := strconv.Atoi(chi.URLParam(r, "noteID"))
	if err != nil {
		http.Error(w, "Invalid noteID", http.StatusBadRequest)
		return
	}

	tokenString := r.Header.Get("Authorization")

	currentUserID, err := auth.GetCurrentUserIDFromToken(tokenString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	note, err := h.notesUsecase.GetNoteByID(r.Context(), noteID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if note.UserID != currentUserID {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(note)
}

func (h *NotesHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	noteID, err := strconv.Atoi(chi.URLParam(r, "noteID"))
	if err != nil {
		http.Error(w, "Invalid noteID", http.StatusBadRequest)
		return
	}

	var updatedNote models.Note
	if err := json.NewDecoder(r.Body).Decode(&updatedNote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenString := r.Header.Get("Authorization")

	currentUserID, err := auth.GetCurrentUserIDFromToken(tokenString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if updatedNote.UserID != currentUserID {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	updatedNote.ID = noteID
	if err := h.notesUsecase.UpdateNote(r.Context(), &updatedNote); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedNote)
}

func (h *NotesHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	noteID, err := strconv.Atoi(chi.URLParam(r, "noteID"))
	if err != nil {
		http.Error(w, "Invalid noteID", http.StatusBadRequest)
		return
	}

	// Получите строку токена из заголовка Authorization
	tokenString := r.Header.Get("Authorization")

	// Получите идентификатор текущего пользователя из токена
	currentUserID, err := auth.GetCurrentUserIDFromToken(tokenString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	note, err := h.notesUsecase.GetNoteByID(r.Context(), noteID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Проверьте, что идентификатор текущего пользователя совпадает с UserID заметки
	if note.UserID != currentUserID {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	if err := h.notesUsecase.DeleteNote(r.Context(), noteID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *NotesHandler) GetAllByUserID(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	currentUserID, err := auth.GetCurrentUserIDFromToken(tokenString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	notes, err := h.notesUsecase.GetAllByUserID(r.Context(), currentUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}
