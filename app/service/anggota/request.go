package anggota

type CreateRequest struct {
	Nama                 string `json:"nama"`
	Alumni               bool   `json:"alumni"`
	KlasifikasiAnggotaId int    `json:"klasifikasi_anggota_id"`
}
