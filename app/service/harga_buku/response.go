package hargabuku

type responseHargaBuku struct {
	Id     int `json:"id"`
	BukuId int `json:"buku_id"`
	Harga  int `json:"harga"`
}
