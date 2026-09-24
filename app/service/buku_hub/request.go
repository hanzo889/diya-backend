package bukuhub

type CreateRequest struct {
	Barcode       string `json:"barcode"`
	BukuId        int    `json:"buku_id"`
	ListKondisiId int    `json:"list_kondisi_id"`
	AnggotaId     int    `json:"anggota_id"`
	RakId         int    `json:"rak_id"`
}
