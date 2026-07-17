package model

import "time"

type Pinjaman struct{
	Id int
	AnggotaId int 
	BukuId int
	TglPinjam time.Time
	TglBalik time.Time
	PetugasPinjamId int
	ListKondisiId int
	PetuagasBalikId int
}