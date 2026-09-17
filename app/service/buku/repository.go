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
	GetBukuByBarcode(barcode string) *sql.Row
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
	_, err := r.db.Exec("insert into buku (judul,list_kategori_id,stock,penulis) values(?,?,?,?)", buku.Judul, buku.ListKategoriId, buku.Stock, buku.Penulis)
	return err
}

func (r *repository) GetAll() (*sql.Rows, error) {
	return r.db.Query("select * from buku")
}

func (r *repository) Update(buku *model.Buku) error {
	_, err := r.db.Exec("update buku set judul=?, list_kategori_id=?, stock=?, penulis=? where id=?", buku.Judul, buku.ListKategoriId, buku.Stock, buku.Penulis, buku.Id)
	return err
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("delete from buku where id=?", id)
	return err
}

func (r *repository) GetById(id int) *sql.Row {
	return r.db.QueryRow("select * from buku where id=?", id)
}

func (r *repository) GetBukuByBarcode(barcode string) *sql.Row {
	return r.db.QueryRow("select b.id, b.judul from buku as b inner join buku_hub as bh on b.id=bh.buku_id where barcode=?", barcode)
}
