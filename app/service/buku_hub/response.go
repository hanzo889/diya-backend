package bukuhub

type responseBukuHub struct{
	Id int  `json:"id"`
	Barcode string `json:"barcode"`
	BukuId int	`json:"bukuId"`
	ListKondisiId int `json:"list_kondisi_id"`
	AnggotaId int `json:"anggota"`
	RakId int `json:"rakId"`
}