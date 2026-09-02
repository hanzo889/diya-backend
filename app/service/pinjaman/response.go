package pinjaman

import "time"

type ResponsePinjaman struct {
	Id              int       `json:"id"`
	AnggotaId       int       `json:"anggota_id"`
	BukuId          int       `json:"buku_id"`
	TglPinjam       time.Time `json:"tgl_pinjam"`
	TglBalik        time.Time `json:"tgl_balik"`
	PetugasPinjamId int       `json:"petugas_pinjam_id"`

	PetugasBalikId int    `json:"petugas_balik_id"`
	KondisiAwalId  int    `json:"kondisi_awal_id"`
	KondisiAkhirId int    `json:"kondisi_akhir_id"`
	Status         string `json:"status"`
}
