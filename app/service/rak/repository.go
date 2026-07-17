package rak

import (
	"database/sql"
	"library/app/model"
)

type Repository interface {
	Create(rak *model.Rak) error
	GetAll() (*sql.Rows, error)
	Update(rak *model.Rak) error
	Delete(id int) error
	GetById(id int) *sql.Row
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(rak *model.Rak) error {
	_, err := r.db.Exec("insert into rak(no_rak) values(?)", rak.NoRak)
	return err
}

func (r *repository) GetAll() (*sql.Rows, error) {
	return r.db.Query("select * from rak")
}

func (r *repository) Update(rak *model.Rak) error {
	_, err := r.db.Exec("update rak set no_rak=? where id=?", rak.NoRak, rak.Id)
	return err
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("delete from rak where id=?", id)
	return err
}

func (r *repository) GetById(id int) *sql.Row {
	return r.db.QueryRow("select * from rak where id=?", id)
}
