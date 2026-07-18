package listkategori

import (
	"library/app/model"
	"log"
)

type ListKategori interface {
	Get() []responseListKategori
	CreateListKategori(lkr CreateRequest)
	Update(list_kodisi *CreateRequest,id int)
	Delete(id int)
	GetById(id int) *responseListKategori
}

type ListKategoriService struct {
	repo Repository
}

var data []model.ListKategori

func NewService(repo Repository) ListKategori {
	return &ListKategoriService{repo}
}

func (b *ListKategoriService) CreateListKategori(lkr CreateRequest) {
	ListKategori := &model.ListKategori{
		Kategori: lkr.Kategori,
	}
	err := b.repo.Create(ListKategori)
	if err != nil {
		log.Println(err)
	}

}
func (b *ListKategoriService) Get() []responseListKategori{
	var kategorikategori []responseListKategori
	kategorikategoriRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return kategorikategori
	}
	// log.Println(kategorikategoriRepo)
	for kategorikategoriRepo.Next() {

		var kategori responseListKategori
		if err := kategorikategoriRepo.Scan(&kategori.Id, &kategori.Kategori); err != nil {
			return kategorikategori
		}
		kategorikategori = append(kategorikategori, kategori)
	}
	return kategorikategori
}

func (b *ListKategoriService)Update(lkr *CreateRequest, id int) {
	ListKategori := &model.ListKategori{
		Kategori: lkr.Kategori,
	}
	err := b.repo.Update(ListKategori)
	if err != nil {
		log.Println(err)
	}
}

func (b *ListKategoriService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b ListKategoriService) GetById(id int) *responseListKategori {
	bariskategori := b.repo.GetById(id)

	var kategori responseListKategori
	if err := bariskategori.Scan(&kategori.Id, kategori.Kategori); err != nil {
		log.Println(err)
	}
	return &kategori
}
