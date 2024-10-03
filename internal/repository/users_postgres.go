package repository

import (
	"Diploma-project-backend/internal/model"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UsersPostgres struct {
	db *sqlx.DB
}

func NewUsersPostgres(db *sqlx.DB) *UsersPostgres {
	return &UsersPostgres{db: db}
}

func (r *UsersPostgres) CreateUser(user model.Users) (string, error) {
	user.ID = uuid.NewString()
	var id string
	var isProfessor = user.IsProfessor

	if isProfessor {
		query := fmt.Sprintf("INSERT INTO %s (id, firstName, lastName, patronymic, email, password, jobTitle) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id", professorTable)

		row := r.db.QueryRow(query, user.ID, user.FirstName, user.LastName, user.Patronymic, user.Email, user.Password, user.JobTitle)
		if err := row.Scan(&id); err != nil {
			return "", err
		}

		return id, nil
	} else {
		query := fmt.Sprintf("INSERT INTO %s (id, firstName, lastName, patronymic, email, password, enrollment, graduation, groupid, imageavatar) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id", studentsTable)

		row := r.db.QueryRow(query, user.ID, user.FirstName, user.LastName, user.Patronymic, user.Email, user.Password, user.Enrollment, user.Graduation, user.GroupId, user.ImageAvatar)
		if err := row.Scan(&id); err != nil {
			return "", err
		}

		return id, nil
	}
}

func (r *UsersPostgres) GetUser(email, password string) (model.Users, error) {
	var user model.Users
	query := fmt.Sprintf(`
    SELECT * FROM %s WHERE email=$1 AND password=$2
    UNION
    SELECT * FROM %s WHERE email=$1 AND password=$2
	`, professorTable, studentsTable)

	err := r.db.Get(&user, query, email, password)

	return user, err
}

func (r *UsersPostgres) UpdateUser(user model.Users) (string, error) {
	var id string
	var isProfessor = user.IsProfessor

	if isProfessor {
		query := fmt.Sprintf(`
		UPDATE %s 
		SET firstName = $1, lastName = $2 , patronymic = $3, email = $4, password = $5, jobTitle = $6
		WHERE id = $7
		RETURNING id`, professorTable)

		row := r.db.QueryRow(query, user.FirstName, user.LastName, user.Patronymic, user.Email, user.Password, user.JobTitle, user.ID)
		if err := row.Scan(&id); err != nil {
			return "", err
		}

		return id, nil
	} else {
		query := fmt.Sprintf(`
		UPDATE %s
		SET firstName = $1, lastName = $2 , patronymic = $3, email = $4, password = $5
		WHERE id = $6
		RETURNING id`, studentsTable)

		row := r.db.QueryRow(query, user.FirstName, user.LastName, user.Patronymic, user.Email, user.Password, user.ID)
		if err := row.Scan(&id); err != nil {
			return "", err
		}

		return id, nil
	}
}

func (r *UsersPostgres) DeleteUser(id string, isProfessor bool) error {
	if isProfessor {
		query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, professorTable)
		_, err := r.db.Exec(query, id)
		if err != nil {
			return err
		}
	} else {
		query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, studentsTable)
		_, err := r.db.Exec(query, id)
		if err != nil {
			return err
		}
	}
	return nil
}
