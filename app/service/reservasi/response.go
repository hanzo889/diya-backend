package reservasi

import "time"

 type responReservasi struct{
	Id int `json:"id"`
	AnggotaId int `json:"anggota_id"`
	BukuId int	`json:"buku_id"`
	BatasPengambilan time.Time `json:"batas_pengambilan"`
}