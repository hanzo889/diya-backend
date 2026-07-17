package pinjaman

import "time"

type CreateRequest struct {
	AnggotaId       int       `json:"anggota_id"`
	BukuId          int       `json:"buku_id"`
	TglPinjam       time.Time `json:"tgl_pinjam"`
	TglBalik        time.Time `json:"tgl_balik"`
	PetugasPinjamId int       `json:"petugas_pinjam_id"`
	ListKondisiId   int       `json:"list_kondisi_id"`
	PetugasBalikId  int       `json:"petugas_balik_id"`
}
