package anggota

import (
	"library/model"
)

type Anggota interface {
	Get() []model.Anggota
	CreateAnggota(anggota model.Anggota)
}

type AnggotaService struct{
	data []model.Anggota
}

func New() Anggota {
	return &AnggotaService{}
}

func (a *AnggotaService) CreateAnggota(anggota model.Anggota) {
	a.data = append(a.data, anggota)
}

func (a *AnggotaService) Get() []model.Anggota {
	return a.data
}
