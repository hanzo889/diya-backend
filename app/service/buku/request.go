package buku

type CreateRequest struct {
	Judul          string `json:"judul"`
	ListKategoriId int    `json:"list_kategori_id"`
	Stock          int    `json:"stock"`
	Penulis        string `json:"penulis"`
}

type CreateRequestBarcode struct{
	ListKondisiId int `json:"list_kondisi_id"`
	RakId int `json:"rak_id"`
	NoAnggota *string `json:"no_anggota"`
}


