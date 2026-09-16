package pinjaman

import (
	"database/sql"
	"fmt"
	"library/app/model"
)

type Repository interface {
	Create(pinjaman *model.Pinjaman) error
	GetAll() (*sql.Rows, error)
	Update(pinjaman *model.Pinjaman) error
	Delete(id int) error
	GetById(id int) *sql.Row
	GetBukuPinjamanByAnggotaId(id int) *sql.Rows
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(pinjaman *model.Pinjaman) error {
	_, err := r.db.Exec("insert into pinjaman (anggota_id,buku_id,tgl_pinjam,tgl_balik,petugas_pinjam_id,list_kondisi_id,petugas_balik_id,kondisi_awal_id,kondisi_akhir_id,status) values(?,?,?,?,?,?,?,?,?,?)", pinjaman.AnggotaId, pinjaman.BukuId, pinjaman.TglPinjam, pinjaman.TglBalik, pinjaman.PetugasPinjamId, pinjaman.PetugasBalikId, pinjaman.Status)
	return err
}

func (r *repository) GetAll() (*sql.Rows, error) {
	return r.db.Query("select * from pinjaman")
}

func (r *repository) Update(pinjaman *model.Pinjaman) error {
	_, err := r.db.Exec("update pinjaman set anggota_id=?,buku_id=?=?,tgl_pinjam=?,tgl_balik=?,petugas_pinjam_id=?,petugas_balik_id=?,kondisi_awal_id=?,kondisi_akhir_id=?,status=?  where id=?", pinjaman.AnggotaId, pinjaman.BukuId, pinjaman.TglPinjam, pinjaman.TglBalik, pinjaman.PetugasPinjamId, pinjaman.PetugasBalikId, pinjaman.KondisiAwalId, pinjaman.KondisiAkhirId, pinjaman.Status, pinjaman.Id)
	return err
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("delete from pinjaman where id=?", id)
	return err
}

func (r *repository) GetById(id int) *sql.Row {
	return r.db.QueryRow("select * from pinjaman where id=?", id)
}
func (r *repository) GetBukuPinjamanByAnggotaId(id int) *sql.Rows {
	data, err := r.db.Query("select p.id,b.judul,p.tgl_pinjam from buku as b inner join pinjaman as p on p.buku_id=b.id where p.anggota_id=?", id)
	fmt.Println(err)
	return data
}
