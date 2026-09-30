package pinjaman

import (
	"database/sql"
	"fmt"
	"library/app/model"
	"time"
)

type Repository interface {
	Create(pinjaman *model.Pinjaman) error
	GetAll() (*sql.Rows, error)
	Update(pinjaman *model.Pinjaman) error
	Delete(id int) error
	GetById(id int) *sql.Row
	GetBukuPinjamanByAnggotaId(id int) *sql.Rows
	GetPinjamanByAnggotaId(anggotaId int) *sql.Rows
	UpdateTglBalik(time_ time.Time, id int) error
	BukuDipinjam(bukuId int) (int, error)
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
	_, err := r.db.Exec("insert into pinjaman (anggota_id,buku_id,tgl_pinjam,tgl_balik,petugas_pinjam_id,petugas_balik_id,kondisi_awal_id,kondisi_akhir_id,status) values(?,?,?,?,?,?,?,?,?)", pinjaman.AnggotaId, pinjaman.BukuId, pinjaman.TglPinjam, pinjaman.TglBalik, pinjaman.PetugasPinjamId, pinjaman.PetugasBalikId, pinjaman.KondisiAwalId, pinjaman.KondisiAkhirId, pinjaman.Status)
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
	data, err := r.db.Query("select p.id,b.judul,p.tgl_pinjam from buku as b inner join pinjaman as p on p.buku_id=b.id where p.anggota_id=? and p.tgl_balik is null", id)
	fmt.Println(err)
	return data
}
func (r *repository) GetPinjamanByAnggotaId(anggotaId int) *sql.Rows {
	data, err := r.db.Query("select p.anggota_id, p.buku_id, k.maks_buku from pinjaman as p inner join klasifikasi_anggota_hub as kah on kah.anggota_id=p.anggota_id inner join klasifikasi_anggota as k on k.id=kah.klasifikasi_anggota_id where p.anggota_id=?", anggotaId)
	fmt.Println(err)
	return data
}

func (r *repository) UpdateTglBalik(time_ time.Time, id int) error {
	_, err := r.db.Exec("update pinjaman set tgl_balik=?, status=? where id=?", time_, "dikembalikam", id)
	return err
}
func (r *repository) BukuDipinjam(bukuId int) (int, error) {
	jumlah := 0
	data, err := r.db.Query("select * from pinjaman where buku_id=? and tgl_balik is null", bukuId)
	if err != nil {
		return jumlah, err
	}

	for data.Next() {
		jumlah += 1
	}
	return jumlah, nil
}
