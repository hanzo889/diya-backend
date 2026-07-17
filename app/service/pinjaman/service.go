package pinjaman

import (
	"library/app/model"
	"log"
)

type Pinjaman interface {
	Get() []ResponsePinjaman
	CreatePinjaman(pinjamanRequest CreateRequest)
	Update(pinjamanRequest *CreateRequest, id int)
	Delete(id int)
	GetById(id int) *ResponsePinjaman
}

type pinjamanService struct {
	repo Repository
}

var data []model.Pinjaman

func NewService(repo Repository) Pinjaman {
	return &pinjamanService{repo}
}

func (b *pinjamanService) CreatePinjaman(pinjamanRequest CreateRequest) {
	pinjaman := &model.Pinjaman{
		AnggotaId:       pinjamanRequest.AnggotaId,
		BukuId:          pinjamanRequest.BukuId,
		TglPinjam:       pinjamanRequest.TglPinjam,
		TglBalik:        pinjamanRequest.TglBalik,
		PetugasPinjamId: pinjamanRequest.PetugasPinjamId,
		ListKondisiId:   pinjamanRequest.ListKondisiId,
		PetugasBalikId:  pinjamanRequest.PetugasBalikId,
	}
	err := b.repo.Create(pinjaman)
	if err != nil {
		log.Println(err)
	}

}
func (b *pinjamanService) Get() []ResponsePinjaman {
	var pinjamanpinjaman []ResponsePinjaman
	pinjamanpinjamanRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return pinjamanpinjaman
	}
	// log.Println(pinjamanpinjamanRepo)
	for pinjamanpinjamanRepo.Next() {

		var pinjaman ResponsePinjaman
		if err := pinjamanpinjamanRepo.Scan(&pinjaman.Id, &pinjaman.AnggotaId, &pinjaman.BukuId, &pinjaman.TglPinjam, &pinjaman.TglBalik, &pinjaman.ListKondisiId, &pinjaman.ListKondisiId, &pinjaman.PetugasBalikId); err != nil {
			return pinjamanpinjaman
		}
		pinjamanpinjaman = append(pinjamanpinjaman, pinjaman)
	}
	return pinjamanpinjaman
}

func (b *pinjamanService) Update(pinjamanRequest *CreateRequest, id int) {
	pinjaman := &model.Pinjaman{
		Id:              id,
		AnggotaId:       pinjamanRequest.AnggotaId,
		BukuId:          pinjamanRequest.BukuId,
		TglPinjam:       pinjamanRequest.TglPinjam,
		TglBalik:        pinjamanRequest.TglBalik,
		PetugasPinjamId: pinjamanRequest.PetugasPinjamId,
		ListKondisiId:   pinjamanRequest.ListKondisiId,
		PetugasBalikId:  pinjamanRequest.PetugasBalikId,
	}
	err := b.repo.Update(pinjaman)
	if err != nil {
		log.Println(err)
	}
}

func (b *pinjamanService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b *pinjamanService) GetById(id int) *ResponsePinjaman {
	barispinjaman := b.repo.GetById(id)

	var pinjaman ResponsePinjaman
	if err := barispinjaman.Scan(&pinjaman.Id, &pinjaman.AnggotaId, &pinjaman.BukuId, &pinjaman.TglPinjam, &pinjaman.TglBalik, &pinjaman.ListKondisiId, &pinjaman.ListKondisiId, &pinjaman.PetugasBalikId); err != nil {
		log.Println(err)
	}
	return &pinjaman
}
