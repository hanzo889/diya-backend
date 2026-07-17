package buku

type CreateRequest struct {
	Judul          string `json:"judul"`
	ListKategoriId int    `json:"list_kategori_id"`
	Stock          int    `json:"stock"`
}


