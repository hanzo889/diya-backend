package listkategori

import (
	"library/app/model"
	"log"
)

type ListKategori interface {
	Get() []model.ListKategori
	CreateListKategori(kategori string)
	Update(list_kodisi model.ListKategori)
	Delete(id int)
	GetById(id int) *model.ListKategori
}

type ListKategoriService struct {
	repo Repository
}

var data []model.ListKategori

func NewService(repo Repository) ListKategori {
	return &ListKategoriService{repo}
}

func (b *ListKategoriService) CreateListKategori(kategori string) {
	ListKategori := &model.ListKategori{
		Kategori: kategori,
	}
	err := b.repo.Create(ListKategori)
	if err != nil {
		log.Println(err)
	}

}
func (b *ListKategoriService) Get() []model.ListKategori {
	var kategorikategori []model.ListKategori
	kategorikategoriRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return kategorikategori
	}
	// log.Println(kategorikategoriRepo)
	for kategorikategoriRepo.Next() {

		var kategori model.ListKategori
		if err := kategorikategoriRepo.Scan(&kategori.Id, &kategori.Kategori); err != nil {
			return kategorikategori
		}
		kategorikategori = append(kategorikategori, kategori)
	}
	return kategorikategori
}

func (b *ListKategoriService) Update(list_kategori model.ListKategori) {
	err := b.repo.Update(&list_kategori)
	if err != nil {
		log.Println(err)
	}
}

func (b *ListKategoriService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b ListKategoriService) GetById(id int) *model.ListKategori {
	bariskategori := b.repo.GetById(id)

	var kategori model.ListKategori
	if err := bariskategori.Scan(&kategori.Id, kategori.Kategori); err != nil {
		log.Println(err)
	}
	return &kategori
}
