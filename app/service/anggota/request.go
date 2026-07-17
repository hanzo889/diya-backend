package anggota

type CreateRequest struct {
	NoAnggota            string `json:"no_anggota"`
	Nama                 string `json:"nama"`
	KlasifikasiAnggotaId int    `json:"klasifikasi_anggota"`
	Alumni               bool   `json:"alumni"`
}
