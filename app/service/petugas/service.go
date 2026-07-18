package petugas

import (
	"library/app/model"
	"log"
)

type Petugas interface {
	Get() []responsePetugas
	CreatePetugas(petugas CreateRequest)
	Update(petugas *CreateRequest,id int)
	Delete(id int)
	GetById(id int) *responsePetugas
}

type petugasService struct {
	repo Repository
}

var data []model.Petugas

func NewService(repo Repository) Petugas {
	return &petugasService{repo}
}

func (b *petugasService) CreatePetugas(ptg CreateRequest) {
	petugas := &model.Petugas{
		AnggotaId: ptg.AnggotaId,
	}
	err := b.repo.CreatePetugas(petugas)
	if err != nil {
		log.Println(err)
	}

}
func (b *petugasService) Get() []responsePetugas {
	var petugaspetugas []responsePetugas
	petugaspetugasRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return petugaspetugas
	}
	// log.Println(petugaspetugasRepo)
	for petugaspetugasRepo.Next() {

		var petugas responsePetugas
		if err := petugaspetugasRepo.Scan(&petugas.Id, &petugas.AnggotaId); err != nil {
			return petugaspetugas
		}
		petugaspetugas = append(petugaspetugas, petugas)
	}
	return petugaspetugas
}

func (b *petugasService) Update(ptg *CreateRequest,id int) {
	petugas := &model.Petugas{
		AnggotaId: ptg.AnggotaId,
	}

	err := b.repo.Update(petugas)
	if err != nil {
		log.Println(err)
	}
}

func (b *petugasService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b *petugasService) GetById(id int) *responsePetugas {
	barispetugas := b.repo.GetById(id)

	var petugas responsePetugas
	if err := barispetugas.Scan(&petugas.Id, &petugas.AnggotaId); err != nil {
		log.Println(err)
	}
	return &petugas
}
