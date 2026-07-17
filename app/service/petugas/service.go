package petugas

import (
	"library/app/model"
	"log"
)

type Petugas interface {
	Get() []model.Petugas
	CreatePetugas(anggotaid int)
	Update(petugas model.Petugas)
	Delete(id int)
	GetById(id int) *model.Petugas
}

type petugasService struct {
	repo Repository
}

var data []model.Petugas

func NewService(repo Repository) Petugas {
	return &petugasService{repo}
}

func (b *petugasService) CreatePetugas(anggotaid int) {
	petugas := model.Petugas{}
	err := b.repo.CreatePetugas(&petugas)
	if err != nil {
		log.Println(err)
	}

}
func (b *petugasService) Get() []model.Petugas {
	var petugaspetugas []model.Petugas
	petugaspetugasRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return petugaspetugas
	}
	// log.Println(petugaspetugasRepo)
	for petugaspetugasRepo.Next() {

		var petugas model.Petugas
		if err := petugaspetugasRepo.Scan(&petugas.Id, &petugas.AnggotaId); err != nil {
			return petugaspetugas
		}
		petugaspetugas = append(petugaspetugas, petugas)
	}
	return petugaspetugas
}

func (b *petugasService) Update(petugas model.Petugas) {
	err := b.repo.Update(&petugas)
	if err != nil {
		log.Println(err)
	}
}

func (b *petugasService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b *petugasService) GetById(id int) *model.Petugas {
	barispetugas := b.repo.GetById(id)

	var petugas model.Petugas
	if err := barispetugas.Scan(&petugas.Id, &petugas.AnggotaId); err != nil {
		log.Println(err)
	}
	return &petugas
}
