package petugas

import (
	"database/sql"
	"library/app/model"
)

type Repository interface {
	CreatePetugas(petugas *model.Petugas) error
	GetAll() (*sql.Rows, error)
	Update(petugas *model.Petugas) error
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

func (r *repository) CreatePetugas(petugas *model.Petugas) error {
	_, err := r.db.Exec("insert into petugas(anggota_id) values(?)", petugas.AnggotaId)
	return err
}

func (r *repository) GetAll() (*sql.Rows, error) {
	return r.db.Query("select * from petugas")
}

func (r *repository) Update(petugas *model.Petugas) error {
	_, err := r.db.Exec("update petugas set anggota_id=? where id=?", petugas.AnggotaId, petugas.Id)
	return err
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("delete from petugas where id=?", id)
	return err
}

func (r *repository) GetById(id int) *sql.Row {
	return r.db.QueryRow("select * from petugas where id=?", id)
}
