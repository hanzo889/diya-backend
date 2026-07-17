package klasifikasianggota

type CreateRequest struct {
	Klasifikasi string `json:"klasifikasi"`
	MaksBuku    int `json:"maks_buku"`
	MaksHari    int `json:"maks_hari"`
}
