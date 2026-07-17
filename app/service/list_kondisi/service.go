package listkondisi

import (
	"library/app/model"
	"log"
)

type ListKondisi interface {
	Get() []model.ListKondisi
	CreateListKondisi(kondisi string)
	Update(list_kodisi model.ListKondisi)
	Delete(id int)
	GetById(id int) *model.ListKondisi
}

type ListKondisiService struct {
	repo Repository
}

var data []model.ListKondisi

func NewService(repo Repository) ListKondisi {
	return &ListKondisiService{repo}
}

func (b *ListKondisiService) CreateListKondisi(kondisi string) {
	listkondisi := &model.ListKondisi{
		Kondisi: kondisi,
	}
	err := b.repo.Create(listkondisi)
	if err != nil {
		log.Println(err)
	}

}
func (b *ListKondisiService) Get() []model.ListKondisi {
	var kondisikondisi []model.ListKondisi
	kondisikondisiRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return kondisikondisi
	}
	// log.Println(kondisikondisiRepo)
	for kondisikondisiRepo.Next() {

		var kondisi model.ListKondisi
		if err := kondisikondisiRepo.Scan(&kondisi.Id, &kondisi.Kondisi); err != nil {
			return kondisikondisi
		}
		kondisikondisi = append(kondisikondisi, kondisi)
	}
	return kondisikondisi
}

func (b *ListKondisiService) Update(list_kodisi model.ListKondisi) {
	err := b.repo.Update(&list_kodisi)
	if err != nil {
		log.Println(err)
	}
}

func (b *ListKondisiService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b ListKondisiService) GetById(id int) *model.ListKondisi {
	bariskondisi := b.repo.GetById(id)

	var kondisi model.ListKondisi
	if err := bariskondisi.Scan(&kondisi.Id, kondisi.Kondisi); err != nil {
		log.Println(err)
	}
	return &kondisi
}
