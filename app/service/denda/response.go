package denda

type responseDenda struct{
	Id int `json:"id"`
	PetugasId int `json:"petugas_id"`
	BukuId int `json:"buku_id"`
	ListKondisiId int `json:"list_kondisi_id"`
	TotalHari int `json:"total_hari"`
	Denda int `json:"denda"`
	HargaBukuId int `json:"harga_buku_id"`
	TotalDenda int `json:"total_denda"`
}