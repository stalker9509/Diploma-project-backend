package repository

import (
	"Diploma-project-backend/internal/model"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type GroupPostgres struct {
	db *sqlx.DB
}

func NewGroupPostgres(db *sqlx.DB) *GroupPostgres {
	return &GroupPostgres{db: db}
}

func (r *GroupPostgres) CreateGroup(group model.Group) (string, error) {
	group.ID = uuid.NewString()
	var id string

	query := fmt.Sprintf("INSERT INTO %s (id, name, curator, timetable) values ($1, $2, $3, $4) RETURNING id", groupTable)

	row := r.db.QueryRow(query, group.ID, group.Name, group.Curator, group.Timetable)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil

}

func (r *GroupPostgres) GetGroup(id string) (model.Group, error) {
	var group model.Group
	query := fmt.Sprintf(" SELECT * FROM %s WHERE id=$1", groupTable)

	err := r.db.Get(&group, query, id)

	return group, err
}

func (r *GroupPostgres) GetAllGroups() ([]model.Group, error) {
	var group []model.Group
	query := fmt.Sprintf(" SELECT * FROM %s", groupTable)

	err := r.db.Get(&group, query)

	return group, err
}

func (r *GroupPostgres) UpdateGroup(group model.Group) (string, error) {
	var id string

	query := fmt.Sprintf(`
		UPDATE %s 
		SET name = $1, curator = $2 , timtable = $3
		WHERE id = $4
		RETURNING id`, groupTable)

	row := r.db.QueryRow(query, group.Name, group.Curator, group.Timetable, group.ID)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil

}

func (r *GroupPostgres) DeleteGroup(id string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, groupTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
