package pinjaman

type CreateRequest struct {
	AnggotaId       int    `json:"anggota_id"`
	BukuId          int    `json:"buku_id"`
	PetugasPinjamId *int   `json:"petugas_pinjam_id"`
	PetugasBalikId  *int   `json:"petugas_balik_id"`
	KondisiAwalId   int    `json:"kondisi_awal_id"`
	KondisiAkhirId  int    `json:"kondisi_akhir_id"`
	Status          string `json:"status"`
}
