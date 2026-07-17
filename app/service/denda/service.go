package denda

import (
	"library/app/model"
	"log"
)

type Denda interface {
	Get() []model.Denda
	CreateDenda(petugas_id int, bukuID int, listKondsiID int, total_hari int, denda int, hargaBukuID int, totalDenda int)
	Update(denda model.Denda)
	Delete(id int)
	GetById(id int) *model.Denda
}

type dendaService struct {
	repo Repository
}

var data []model.Denda

func NewService(repo Repository) Denda {
	return &dendaService{repo}
}

func (b *dendaService) CreateDenda(petugas_id int, bukuID int, listKondsiID int, total_hari int, denda int, hargaBukuID int, totalDenda int) {
	d := &model.Denda{}
	err := b.repo.Create(d)
	if err != nil {
		log.Println(err)
	}

}
func (b *dendaService) Get() []model.Denda {
	var dendadenda []model.Denda
	dendadendaRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return dendadenda
	}
	// log.Println(dendadendaRepo)
	for dendadendaRepo.Next() {

		var denda model.Denda
		if err := dendadendaRepo.Scan(&denda.Id, &denda.PetugasId, &denda.BukuId, &denda.ListKondisiId, &denda.TotalHari, &denda.Denda, &denda.HargaBukuId, &denda.TotalDenda); err != nil {
			return dendadenda
		}
		dendadenda = append(dendadenda, denda)
	}
	return dendadenda
}

func (b *dendaService) Update(denda model.Denda) {
	err := b.repo.Update(&denda)
	if err != nil {
		log.Println(err)
	}
}

func (b *dendaService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b *dendaService) GetById(id int) *model.Denda {
	barisDenda := b.repo.GetById(id)

	var denda model.Denda
	if err := barisDenda.Scan(&denda.Id, &denda.PetugasId, &denda.BukuId, &denda.ListKondisiId, &denda.TotalHari, &denda.Denda, &denda.HargaBukuId, &denda.TotalDenda); err != nil {
		log.Println(err)
	}
	return &denda
}
