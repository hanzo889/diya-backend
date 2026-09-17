package listkategori

import (
	"database/sql"
	"library/app/model"
)

type Repository interface {
	Create(listKategori *model.ListKategori) error
	GetAll() (*sql.Rows, error)
	Update(ListKategori *model.ListKategori) error
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

func (r *repository) Create(listKategori *model.ListKategori) error {
	_, err := r.db.Exec("insert into list_kategori(kategori) values(?)", listKategori.Kategori)
	return err
}

func (r *repository) GetAll() (*sql.Rows, error) {
	return r.db.Query("select * from list_kategori")
}

func (r *repository) Update(list_kategori *model.ListKategori) error {
	_, err := r.db.Exec("update list_kategori set kategori=? where id=?", list_kategori.Kategori, list_kategori.Id)
	return err
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("delete from list_kategori where id=?", id)
	return err
}

func (r *repository) GetById(id int) *sql.Row {
	return r.db.QueryRow("select * from list_kategori where id=?", id)
}
