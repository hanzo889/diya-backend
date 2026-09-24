package pinjaman

import "time"

type ResponsePinjaman struct {
	Id              int       `json:"id"`
	AnggotaId       int       `json:"anggota_id"`
	BukuId          int       `json:"buku_id"`
	TglPinjam       time.Time `json:"tgl_pinjam"`
	TglBalik        time.Time `json:"tgl_balik"`
	PetugasPinjamId int       `json:"petugas_pinjam_id"`
	PetugasBalikId  int       `json:"petugas_balik_id"`
	KondisiAwalId   int       `json:"kondisi_awal_id"`
	KondisiAkhirId  int       `json:"kondisi_akhir_id"`
	Status          string    `json:"status"`
}
type TampungPinjamanAnggota struct {
	Id        int    `json:"id"`
	NoAnggota string `json:"no_anggota"`
	Nama      string `json:"nama"`
	MaksBuku  int    `json:"maks_buku"`
	MaksHari  int    `json:"maks_hari"`
}
type ResponseBuku struct {
	Id        int       `json:"id"`
	Judul     string    `json:"judul"`
	TglPinjam time.Time `json:"tgl_pinjam"`
}

type ResponsePinjamanAnggota struct {
	Id          int            `json:"id"`
	NoAnggota   string         `json:"no_anggota"`
	Nama        string         `json:"nama"`
	BolehPinjam *bool          `json:"boleh_pinjam"`
	Buku        []ResponseBuku `json:"buku"`
}

type responseGetPinjamanByAnggotaId struct{
	AnggotaId int
	BukuId int
	MaksBuku int
}