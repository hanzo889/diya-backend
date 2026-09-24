package bukuhub

type responseBukuHub struct {
	Id            int    `json:"id"`
	Barcode       string `json:"barcode"`
	BukuId        int    `json:"buku_id"`
	ListKondisiId int    `json:"list_kondisi_id"`
	AnggotaId     int    `json:"anggota_id"`
	RakId         int    `json:"rak_id"`
}
