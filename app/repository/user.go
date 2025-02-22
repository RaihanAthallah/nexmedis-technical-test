package repository

import (
	"database/sql"
	"log"
	"nexmedis-technical-test/app/model/entity"
)

type UserRepository interface {
	GetAllUsers() ([]entity.User, error)
	GetUserByID(id int) (*entity.User, error)
	CreateUser(user entity.User) error
	UpdateUser(user entity.User) error
	DeleteUser(id int) error
	GetUserByUsername(username string) (*entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAllUsers() ([]entity.User, error) {
	rows, err := r.db.Query("SELECT id, username, email FROM users")
	if err != nil {
		log.Println("Error fetching users:", err)
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			log.Println("Error scanning user:", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) GetUserByID(id int) (*entity.User, error) {
	var user entity.User
	err := r.db.QueryRow("SELECT id, username, email FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		log.Println("Error fetching user by ID:", err)
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) CreateUser(user entity.User) error {
	_, err := r.db.Exec("INSERT INTO users (username, email) VALUES ($1, $2)", user.Username, user.Email)
	if err != nil {
		log.Println("Error inserting user:", err)
	}
	return err
}

func (r *userRepository) UpdateUser(user entity.User) error {
	_, err := r.db.Exec("UPDATE users SET username = $1, email = $2 WHERE id = $3", user.Username, user.Email, user.ID)
	if err != nil {
		log.Println("Error updating user:", err)
	}
	return err
}

func (r *userRepository) DeleteUser(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		log.Println("Error deleting user:", err)
	}
	return err
}

func (r *userRepository) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := r.db.QueryRow("SELECT id, username, email FROM users WHERE username = $1", username).
		Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		log.Println("Error fetching user by username:", err)
		return nil, err
	}

	return &user, nil
}
