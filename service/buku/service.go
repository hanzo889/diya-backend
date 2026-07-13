package buku

import (
	"library/model"
	"log"
)

type Buku interface {
	Get() []model.Buku
	CreateBuku(judul string, listKategoriId int, stock int)
	Update(buku model.Buku)
	Delete(id int)
	GetById(id int) *model.Buku
}

type bukuService struct {
	repo Repository
}

var data []model.Buku

func NewService(repo Repository) Buku {
	return &bukuService{repo}
}

func (b *bukuService) CreateBuku(judul string, listKategoriId int, stock int) {
	buku := &model.Buku{
		Judul:          judul,
		ListKategoriId: listKategoriId,
		Stock:          stock,
	}
	err := b.repo.Create(buku)
	if err != nil {
		log.Println(err)
	}

}
func (b *bukuService) Get() []model.Buku {
	var bukuBuku []model.Buku
	bukuBukuRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return bukuBuku
	}
	// log.Println(bukuBukuRepo)
	for bukuBukuRepo.Next() {

		var buku model.Buku
		if err := bukuBukuRepo.Scan(&buku.Id, &buku.Judul, &buku.ListKategoriId, &buku.Stock); err != nil {
			return bukuBuku
		}
		bukuBuku = append(bukuBuku, buku)
	}
	return bukuBuku
}

func (b *bukuService) Update(buku model.Buku) {
	err := b.repo.Update(&buku)
	if err != nil {
		log.Println(err)
	}
}

func (b *bukuService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b *bukuService) GetById(id int) *model.Buku {
	barisBuku := b.repo.GetById(id)

	var buku model.Buku
	if err := barisBuku.Scan(&buku.Id, &buku.Judul, &buku.ListKategoriId, &buku.Stock); err != nil {
		log.Println(err)
	}
	return &buku
}
