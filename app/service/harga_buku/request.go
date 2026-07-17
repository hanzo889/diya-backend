package  hargabuku

type CreateRequest struct{
	BukuId int `json:"buku_id"`
	Harga int `json:"harga"`
}