package listkondisi

import (
	"library/app/model"
	"log"
)

type ListKondisi interface {
	Get() []reponseListKondisi
	CreateListKondisi(kondisiR CreateRequest)
	Update(kondisiR *CreateRequest,id int)
	Delete(id int)
	GetById(id int) *reponseListKondisi
}

type ListKondisiService struct {
	repo Repository
}



func NewService(repo Repository) ListKondisi {
	return &ListKondisiService{repo}
}

func (b *ListKondisiService) CreateListKondisi(kondisR CreateRequest) {
	listkondisi := &model.ListKondisi{
		Kondisi: kondisR.Kondisi,
	}
	err := b.repo.Create(listkondisi)
	if err != nil {
		log.Println(err)
	}

}
func (b *ListKondisiService) Get() []reponseListKondisi {
	var kondisikondisi []reponseListKondisi
	kondisikondisiRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return kondisikondisi
	}
	// log.Println(kondisikondisiRepo)
	for kondisikondisiRepo.Next() {

		var kondisi reponseListKondisi
		if err := kondisikondisiRepo.Scan(&kondisi.Id, &kondisi.Kondisi); err != nil {
			return kondisikondisi
		}
		kondisikondisi = append(kondisikondisi, kondisi)
	}
	return kondisikondisi
}

func (b *ListKondisiService) Update(kondisiR *CreateRequest,id int) {
	kondisi:=&model.ListKondisi{
		Id: id,
		Kondisi: kondisiR.Kondisi,
	}
	err := b.repo.Update(kondisi)
	if err != nil {
		log.Println(err)
	}
}

func (b *ListKondisiService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b ListKondisiService) GetById(id int) *reponseListKondisi {
	bariskondisi := b.repo.GetById(id)

	var kondisi reponseListKondisi
	if err := bariskondisi.Scan(&kondisi.Id, kondisi.Kondisi); err != nil {
		log.Println(err)
	}
	return &kondisi
}
