package hargabuku

import (
	"database/sql"
	"library/app/model"
)

type Repository interface {
	CreateHargaBuku(hargaBuku *model.HargaBuku) error
	GetAll() (*sql.Rows, error)
	Update(hargaBuku *model.HargaBuku) error
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

func (r *repository) CreateHargaBuku(hargaBuku *model.HargaBuku) error {
	_, err := r.db.Exec("insert into harga_buku(buku_id,harga) values(?,?)", hargaBuku.BukuId, hargaBuku.Harga)
	return err
}

func (r *repository) GetAll() (*sql.Rows, error) {
	return r.db.Query("select * from harga_buku")
}

func (r *repository) Update(hargabuku *model.HargaBuku) error {
	_, err := r.db.Exec("update harga_buku set buku_id=?,harga=? where id=?", hargabuku.BukuId, hargabuku.Harga, hargabuku.Id)
	return err
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("delete from harga_buku  where id=?", id)
	return err
}

func (r *repository) GetById(id int) *sql.Row {
	return r.db.QueryRow("select * from harga_buku  where id=?", id)
}
