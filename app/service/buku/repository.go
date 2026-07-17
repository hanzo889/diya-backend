package buku

import (
	"database/sql"
	"library/app/model"
)

type Repository interface {
	Create(buku *model.Buku) error
	GetAll() (*sql.Rows, error)
	Update(buku *model.Buku) error
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

func (r *repository) Create(buku *model.Buku) error {
	_, err := r.db.Exec("insert into buku (judul,list_katagori_id,stock) values(?,?,?)", buku.Judul, buku.ListKategoriId, buku.Stock)
	return err
}

func (r *repository) GetAll() (*sql.Rows, error) {
	return r.db.Query("select * from buku")
}

func (r *repository) Update(buku *model.Buku) error {
	_, err := r.db.Exec("update buku set judul=?, list_katagori_id=?, stock=? where id=?", buku.Judul, buku.ListKategoriId, buku.Stock, buku.Id)
	return err
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("delete from buku where id=?", id)
	return err
}

func (r *repository) GetById(id int) *sql.Row {
	return r.db.QueryRow("select * from buku where id=?", id)
}
