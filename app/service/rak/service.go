package rak

import (
	"library/app/model"
	"log"
)

type Rak interface {
	Get() []ResponseRak
	CreateRak(rakRequest CreateRequest )
	Update(rakRequest *CreateRequest,id int)
	Delete(id int)
	GetById(id int) *ResponseRak
}

type rakService struct {
	repo Repository
}



func NewService(repo Repository) Rak {
	return &rakService{repo}
}

func (b *rakService)CreateRak(rakRequest CreateRequest ) {
	rak := model.Rak{
		NoRak: rakRequest.NoRak,
	}
	err := b.repo.Create(&rak)
	if err != nil {
		log.Println(err)
	}

}
func (b *rakService) Get() []ResponseRak {
	var rakrak []ResponseRak
	rakrakRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return rakrak
	}
	// log.Println(rakrakRepo)
	for rakrakRepo.Next() {

		var rak ResponseRak
		if err := rakrakRepo.Scan(&rak.Id, &rak.NoRak); err != nil {
			return rakrak
		}
		rakrak = append(rakrak, rak)
	}
	return rakrak
}

func (b *rakService) Update(rakRequest *CreateRequest,id int) {
	rak:=&model.Rak{
		Id: id,
		NoRak: rakRequest.NoRak,
	}
	err := b.repo.Update(rak)
	if err != nil {
		log.Println(err)
	}
}

func (b *rakService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b *rakService) GetById(id int) *ResponseRak {
	barisrak := b.repo.GetById(id)

	var rak ResponseRak
	if err := barisrak.Scan(&rak.Id, &rak.NoRak); err != nil {
		log.Println(err)
	}
	return &rak
}
