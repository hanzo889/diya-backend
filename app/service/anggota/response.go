package anggota

type ResponseAnggota struct {
	Id        int    `json:"id"`
	NoAnggota string `json:"no_anggota"`
	Nama      string `json:"nama"`
	Alumni bool `json:"alumni"`
}
