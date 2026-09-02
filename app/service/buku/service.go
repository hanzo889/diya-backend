package buku

import (
	"fmt"
	"library/app/model"
	"log"
)

type Buku interface {
	Get() []ResponseBuku
	CreateBuku(bukuRequest CreateRequest)error
	Update(bukuRequest *CreateRequest, id int)
	Delete(id int)
	GetById(id int) *ResponseBuku
}

type bukuService struct {
	repo Repository
}

func NewService(repo Repository) Buku {
	return &bukuService{repo}
}

func (b *bukuService) CreateBuku(bukuRequest CreateRequest) error{
	buku := &model.Buku{
		Judul:          bukuRequest.Judul,
		ListKategoriId: bukuRequest.ListKategoriId,
		Stock:          bukuRequest.Stock,
		Penulis:        bukuRequest.Penulis,
	}
	err := b.repo.Create(buku)
	if err != nil {
		log.Println("***",err)
	}
return err
}
func (b *bukuService) Get() []ResponseBuku {
	var bukuBuku []ResponseBuku
	bukuBukuRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return bukuBuku
	}
	// log.Println(bukuBukuRepo)
	for bukuBukuRepo.Next() {

		var buku ResponseBuku
		fmt.Println(bukuBukuRepo)
		if err := bukuBukuRepo.Scan(&buku.Id, &buku.Judul, &buku.ListKategoriId, &buku.Stock, &buku.Penulis); err != nil {
			
			return bukuBuku
		}
		bukuBuku = append(bukuBuku, buku)
	}
	return bukuBuku
}

func (b *bukuService) Update(bukuRequest *CreateRequest, id int) {
	buku := &model.Buku{
		Id:             id,
		Judul:          bukuRequest.Judul,
		ListKategoriId: bukuRequest.ListKategoriId,
		Stock:          bukuRequest.Stock,
		Penulis:        bukuRequest.Penulis,
	}
	err := b.repo.Update(buku)
	if err != nil {
		log.Println(err)
	}
}

func (b *bukuService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b *bukuService) GetById(id int) *ResponseBuku {
	barisBuku := b.repo.GetById(id)

	var buku ResponseBuku
	if err := barisBuku.Scan(&buku.Id, &buku.Judul, &buku.ListKategoriId, &buku.Stock, &buku.Penulis); err != nil {
		log.Println(err)
	}
	return &buku
}
