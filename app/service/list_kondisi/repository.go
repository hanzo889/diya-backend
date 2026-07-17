package listkondisi

import (
	"database/sql"
	"library/app/model"
)

type Repository interface {
	Create(listKondisi *model.ListKondisi) error
	GetAll() (*sql.Rows, error)
	Update(ListKondisi *model.ListKondisi) error
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

func (r *repository) Create(listKondisi *model.ListKondisi) error {
	_, err := r.db.Exec("insert into list_kondisi(kondisi) values(?)", listKondisi.Kondisi)
	return err
}

func (r *repository) GetAll() (*sql.Rows, error) {
	return r.db.Query("select * from list_kondisi")
}

func (r *repository) Update(list_kondisi *model.ListKondisi) error {
	_, err := r.db.Exec("update list_kondisi set kondisi=? where id=?", list_kondisi.Kondisi, list_kondisi.Id)
	return err
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("delete from list_kondisi where id=?", id)
	return err
}

func (r *repository) GetById(id int) *sql.Row {
	return r.db.QueryRow("select * from list_kondisi where id=?", id)
}
