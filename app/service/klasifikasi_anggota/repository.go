package klasifikasianggota

import (
	"database/sql"
	"library/app/model"
)

type Repository interface {
	Create(klA *model.KlasifikasiAnggota) error
	GetAll() (*sql.Rows, error)
	Update(klA *model.KlasifikasiAnggota) error
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

func (r *repository) Create(klA *model.KlasifikasiAnggota) error {
	_, err := r.db.Exec("insert into  klasifikasi_anggota(klasifikasi,maks_buku,maks_hari) values(?,?,?)", klA.Klasifikasi, klA.MaksBuku, klA.MaksHari)
	return err
}

func (r *repository) GetAll() (*sql.Rows, error) {
	return r.db.Query("select * from klasifikasi_anggota")
}

func (r *repository) Update(klA *model.KlasifikasiAnggota) error {
	_, err := r.db.Exec("update klasifikasi_anggota set klasifikasi=?, maks_buku=?, maks_hari=? where id=?", klA.Klasifikasi, klA.MaksBuku, klA.MaksHari, klA.Id)
	return err
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("delete from klasifikasi_anggota where id=?", id)
	return err
}

func (r *repository) GetById(id int) *sql.Row {
	return r.db.QueryRow("select * from klasifikasi_anggota where id=?", id)
}
