package klasifikasianggotahub

import (
	"database/sql"
	"library/app/model"
)

type Repository interface {
	Create(kah *model.KlasifikasiAnggotaHub) error
	Delete(anggotaId int) error
	Update(klasifikasiAnggotaId int, anggotaId int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(kah *model.KlasifikasiAnggotaHub) error {
	_, err := r.db.Exec("insert into klasifikasi_anggota_hub (anggota_id,klasifikasi_anggota_id) values(?,?)", kah.AnggotaId, kah.KlasifikasiAnggotaId)
	return err
}

func (r *repository) Delete(anggotaId int) error {
	_, err := r.db.Exec("delete from klasifikasi_anggota_hub where anggota_id=?", anggotaId)
	return err
}

func (r *repository) Update(klasifikasiAnggotaId int, anggotaId int) error {
	_, err := r.db.Exec("update klasifikasi_anggota_hub set klasifikasi_anggota_id=? where anggota_id=?", klasifikasiAnggotaId, anggotaId)
	return err
}
