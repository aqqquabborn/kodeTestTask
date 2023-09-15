package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"kodeTestTask/internal/api/models"
	"log"
)

type UsersRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, userID int) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, userID int) error
}

type usersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) UsersRepository {
	return &usersRepository{db}
}

func (r *usersRepository) Create(ctx context.Context, user *models.User) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password).Scan(&user.ID)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}
	return nil
}

func (r *usersRepository) GetByID(ctx context.Context, userID int) (*models.User, error) {
	query := "SELECT id, username, password FROM users WHERE id = $1"
	var user models.User
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		log.Printf("Error getting user by ID: %v", err)
		return nil, err
	}
	return &user, nil
}

func (r *usersRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	query := "SELECT id, username, password FROM users WHERE username = $1"
	var user models.User
	err := r.db.QueryRowContext(ctx, query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		log.Printf("Error getting user by username: %v", err)
		return nil, err
	}
	return &user, nil
}

func (r *usersRepository) Update(ctx context.Context, user *models.User) error {
	query := "UPDATE users SET username = $1, password = $2 WHERE id = $3"
	_, err := r.db.ExecContext(ctx, query, user.Username, user.Password, user.ID)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}
	return nil
}

func (r *usersRepository) Delete(ctx context.Context, userID int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, userID)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}
	return nil
}
