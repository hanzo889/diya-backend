package buku

type ResponseBuku struct {
	Id             int `json:"id"`
	Judul          string `json:"judul"`
	ListKategoriId int `json:"list_kategori_id"`
	Stock          int	`json:"stock"`
}

