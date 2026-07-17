package denda

import (
	"library/app/model"
	"log"
)

type Denda interface {
	Get() []responseDenda
	CreateDenda(dendaR CreateRequest)
	Update(dendaR *CreateRequest, id int)
	Delete(id int)
	GetById(id int) *responseDenda
}

type dendaService struct {
	repo Repository
}

func NewService(repo Repository) Denda {
	return &dendaService{repo}
}

func (b *dendaService) CreateDenda(dendaR CreateRequest) {
	d := &model.Denda{
		PetugasId:     dendaR.PetugasId,
		BukuId:        dendaR.BukuId,
		ListKondisiId: dendaR.ListKondisiId,
		TotalHari:     dendaR.TotalHari,
		Denda:         dendaR.Denda,
		HargaBukuId:   dendaR.HargaBukuId,
		TotalDenda:    dendaR.TotalDenda,
	}
	err := b.repo.Create(d)
	if err != nil {
		log.Println(err)
	}

}
func (b *dendaService) Get() []responseDenda {
	var dendadenda []responseDenda
	dendadendaRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return dendadenda
	}
	// log.Println(dendadendaRepo)
	for dendadendaRepo.Next() {

		var denda responseDenda
		if err := dendadendaRepo.Scan(&denda.Id, &denda.PetugasId, &denda.BukuId, &denda.ListKondisiId, &denda.TotalHari, &denda.Denda, &denda.HargaBukuId, &denda.TotalDenda); err != nil {
			return dendadenda
		}
		dendadenda = append(dendadenda, denda)
	}
	return dendadenda
}

func (b *dendaService) Update(dendaR *CreateRequest, id int) {
	denda := &model.Denda{
		Id: id,
		PetugasId:     dendaR.PetugasId,
		BukuId:        dendaR.BukuId,
		ListKondisiId: dendaR.ListKondisiId,
		TotalHari:     dendaR.TotalHari,
		Denda:         dendaR.Denda,
		HargaBukuId:   dendaR.HargaBukuId,
		TotalDenda:    dendaR.TotalDenda,
	}
	err := b.repo.Update(denda)
	if err != nil {
		log.Println(err)
	}
}

func (b *dendaService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b *dendaService) GetById(id int) *responseDenda {
	barisDenda := b.repo.GetById(id)

	var denda responseDenda
	if err := barisDenda.Scan(&denda.Id, &denda.PetugasId, &denda.BukuId, &denda.ListKondisiId, &denda.TotalHari, &denda.Denda, &denda.HargaBukuId, &denda.TotalDenda); err != nil {
		log.Println(err)
	}
	return &denda
}
