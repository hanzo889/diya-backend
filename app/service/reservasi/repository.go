package reservasi

import (
	"database/sql"
	"library/app/model"
)

type Repository interface {
	Create(reservasi *model.Reservasi) error
	GetAll() (*sql.Rows, error)
	Update(reservasi *model.Reservasi) error
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

func (r *repository) Create(reservasi *model.Reservasi) error {
	_, err := r.db.Exec("insert into reservasi (anggota_id,buku_id,batas_pengembaliaan) values(?,?,?)",reservasi.AnggotaId,reservasi.BukuId,reservasi.BatasPengambilan)
	return err
}

func (r *repository) GetAll() (*sql.Rows, error) {
	return r.db.Query("select * from reservasi")
}

func (r *repository) Update(reservasi *model.Reservasi) error {
	_, err := r.db.Exec("update reservasi set anggota_id=?,buku_id=?,batas_pengembaliaan=? where id=?",reservasi.AnggotaId,reservasi.BukuId,reservasi.BatasPengambilan,reservasi.Id)
	return err
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("delete from reservasi where id=?", id)
	return err
}

func (r *repository) GetById(id int) *sql.Row {
	return r.db.QueryRow("select * from reservasi where id=?", id)
}
