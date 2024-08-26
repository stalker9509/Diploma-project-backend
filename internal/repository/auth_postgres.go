package repository

import (
	"Diploma-project-backend/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.Users) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (firstname, lastname, patronymic, email, password, imageavatar) values ($1, $2, $3, $4, $5, $6) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.FirstName, user.LastName, user.Patronymic, user.Email, user.Password, user.ImageAvatar)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (model.Users, error) {
	var user model.Users
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}
