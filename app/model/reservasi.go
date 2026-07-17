package model

import "time"

type Reservasi struct{
	Id int
	AnggotaId int
	BukuId int
	BatasPengambilan time.Time
}