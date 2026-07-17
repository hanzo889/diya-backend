package bukuhub

import (
	"database/sql"
	"library/app/model"
)

type Repository interface {
	Create(bukuhub *model.BukuHub) error
	GetAll() (*sql.Rows, error)
	Update(bukuhub *model.BukuHub) error
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

func (r *repository) Create(bukuhub *model.BukuHub) error {
	_, err := r.db.Exec("insert into buku_hub (barcode,buku_id,list_kondisi_id,anggota_id,rak_id) values(?,?,?,?,?)", bukuhub.Barcode, bukuhub.BukuId, bukuhub.ListKondisiId, bukuhub.AnggotaId, bukuhub.RakId)
	return err
}

func (r *repository) GetAll() (*sql.Rows, error) {
	return r.db.Query("select * from buku")
}

func (r *repository) Update(bukuhub *model.BukuHub) error {
	_, err := r.db.Exec("update buku_hub set barcode=?,buku_id=?,list_kondisi_id=?,anggota_id=?,rak_id=? where id=?", bukuhub.Barcode, bukuhub.BukuId, bukuhub.ListKondisiId, bukuhub.AnggotaId, bukuhub.RakId, bukuhub.ID)
	return err
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("delete from buku_hub where id=?", id)
	return err
}

func (r *repository) GetById(id int) *sql.Row {
	return r.db.QueryRow("select * from buku_hub where id=?", id)
}
