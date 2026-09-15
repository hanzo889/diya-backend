package anggota

import (
	"database/sql"
	"library/app/model"
)

type Repository interface {
	Create(anggota *model.Anggota) (sql.Result, error)
	GetAll() (*sql.Rows, error)
	Update(anggota *model.Anggota) error
	Delete(id int) error
	GetById(id int) *sql.Row
	GetMaxNoAnggota() *sql.Row
	GetAnggotaByNoAnggota(noAnggota string) *sql.Row
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(anggota *model.Anggota) (sql.Result, error) {
	data, err := r.db.Exec("insert into anggota(no_anggota,nama,alumni) values(?,?,?)", anggota.NoAnggota, anggota.Nama, anggota.Alumni)
	return data, err
}

func (r *repository) GetAll() (*sql.Rows, error) {
	return r.db.Query("select a.id, a.no_anggota,a.nama,a.alumni,k.klasifikasi_anggota_id as klasifikasi_anggota_id from anggota as a inner join klasifikasi_anggota_hub as k on a.id=k.anggota_id;")
}

func (r *repository) Update(anggota *model.Anggota) error {
	_, err := r.db.Exec("update anggota set  nama=? ,alumni=? where id=?", anggota.Nama, anggota.Alumni, anggota.ID)
	return err
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("delete from anggota where id=?", id)
	return err
}

func (r *repository) GetById(id int) *sql.Row {
	return r.db.QueryRow("select * from anggota where id=?", id)
}
func (r *repository) GetMaxNoAnggota() *sql.Row {
	return r.db.QueryRow("select max(no_anggota) from anggota")
}
func (r *repository) GetAnggotaByNoAnggota(noAnggota string) *sql.Row {
	return r.db.QueryRow("select * from anggota where no_anggota=?", noAnggota)
}
