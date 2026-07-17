package rak

import (
	"library/app/model"
	"log"
)

type Rak interface {
	Get() []model.Rak
	CreateRak(noRak int)
	Update(rak model.Rak)
	Delete(id int)
	GetById(id int) *model.Rak
}

type rakService struct {
	repo Repository
}

var data []model.Rak

func NewService(repo Repository) Rak {
	return &rakService{repo}
}

func (b *rakService) CreateRak(noRak int) {
	rak := model.Rak{}
	err := b.repo.Create(&rak)
	if err != nil {
		log.Println(err)
	}

}
func (b *rakService) Get() []model.Rak {
	var rakrak []model.Rak
	rakrakRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return rakrak
	}
	// log.Println(rakrakRepo)
	for rakrakRepo.Next() {

		var rak model.Rak
		if err := rakrakRepo.Scan(&rak.Id, &rak.NoRak); err != nil {
			return rakrak
		}
		rakrak = append(rakrak, rak)
	}
	return rakrak
}

func (b *rakService) Update(rak model.Rak) {
	err := b.repo.Update(&rak)
	if err != nil {
		log.Println(err)
	}
}

func (b *rakService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b *rakService) GetById(id int) *model.Rak {
	barisrak := b.repo.GetById(id)

	var rak model.Rak
	if err := barisrak.Scan(&rak.Id, &rak.NoRak); err != nil {
		log.Println(err)
	}
	return &rak
}
