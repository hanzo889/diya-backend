package klasifikasianggota

type responseKlasifikasiAnggota struct {
	Id          int    `json:"id"`
	Klasifikasi string `json:"klasifikasi"`
	MaksBuku    int    `json:"maks_buku"`
	MaksHari    int    `json:"maks_hari"`
}
