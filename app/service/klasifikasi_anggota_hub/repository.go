package klasifikasianggotahub

import (
	"database/sql"
	"library/app/model"
)

type Repository interface {
	Create(kah *model.KlasifikasiAnggotaHub) error
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

