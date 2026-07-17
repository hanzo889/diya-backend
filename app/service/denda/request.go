package  denda

type CreateRequest struct{
	PetugasId int `json:"petugasId"`
	BukuId int `json:"bukuId"`
	ListKondisiId int `json:"list_kondisi_id"`
	TotalHari int `json:"total_hari"`
	Denda int `json:"denda"`
	HargaBukuId int `json:"harga_buku_id"`
	TotalDenda int `json:"total_denda"`
}
