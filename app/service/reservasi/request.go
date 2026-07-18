package reservasi

import "time"

 type CreateRequest struct{
	AnggotaId int `json:"anggota_id"`
	BukuId int	`json:"buku_id"`
	BatasPengambilan time.Time `json:"batas_pengambilan"`
}