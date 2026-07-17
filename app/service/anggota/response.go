package anggota

type ResponseAnggota struct {
	Id                   int    `json:"id"`
	NoAnggota            string `json:"no_anggota"`
	Nama                 string `json:"nama"`
	KlasifikasiAnggotaId int    `json:"klasifikasi_anggota"`
	Alumni               bool   `json:"alumni"`
}
