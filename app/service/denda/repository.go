package denda

import (
	"database/sql"
	"library/app/model"
)

type Repository interface {
	Create(denda *model.Denda) error
	GetAll() (*sql.Rows, error)
	Update(denda *model.Denda) error
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

func (r *repository) Create(denda *model.Denda) error {
	_, err := r.db.Exec("insert into denda (petugas_id,buku_id,list_kondisi_id,total_hari,denda,harga_buku_id,total_denda) values(?,?,?,?,?,?,?)", denda.PetugasId, denda.BukuId, denda.ListKondisiId, denda.TotalHari, denda.HargaBukuId, denda.TotalDenda)
	return err
}

func (r *repository) GetAll() (*sql.Rows, error) {
	return r.db.Query("select * from denda")
}

func (r *repository) Update(denda *model.Denda) error {
	_, err := r.db.Exec("update denda set petugas_id=?,buku_id=?,list_kondisi_id=?,total_hari=?,denda=?,harga_buku_id=?,total_denda where id=?", denda.PetugasId, denda.BukuId, denda.ListKondisiId, denda.TotalHari, denda.HargaBukuId, denda.TotalDenda)
	return err
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("delete from denda where id=?", id)
	return err
}

func (r *repository) GetById(id int) *sql.Row {
	return r.db.QueryRow("select * from denda where id=?", id)
}
