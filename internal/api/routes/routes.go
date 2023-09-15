package routes

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
	"kodeTestTask/internal/api/handlers"
	"kodeTestTask/internal/api/repositories"
	"kodeTestTask/internal/api/usecases"
)

func CreateRouter(db *sql.DB) *chi.Mux {

	notesRepository := repositories.NewNotesRepository(db)
	notesUsecase := usecases.NewNotesUsecase(notesRepository)

	usersRepository := repositories.NewUsersRepository(db)
	usersUsecase := usecases.NewUsersUsecase(usersRepository)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	notesHandler := handlers.NewNotesHandler(notesUsecase)

	r.Post("/api/notes", notesHandler.CreateNote)
	r.Get("/api/notes/{noteID}", notesHandler.GetNoteByID)
	r.Put("/api/notes/{noteID}", notesHandler.UpdateNote)
	r.Delete("/api/notes/{noteID}", notesHandler.DeleteNote)
	r.Get("/api/notes", notesHandler.GetAllByUserID)

	usersHandler := handlers.NewUsersHandler(usersUsecase)

	r.Post("/api/users", usersHandler.CreateUser)
	r.Get("/api/users/{userID}", usersHandler.GetUserByID)
	r.Get("/api/users/{username}", usersHandler.GetUserByUsername)
	r.Put("/api/users/{userID}", usersHandler.UpdateUser)
	r.Delete("/api/users/{userID}", usersHandler.DeleteUser)
	r.Post("/api/authenticate", usersHandler.AuthenticateUser)

	return r
}
