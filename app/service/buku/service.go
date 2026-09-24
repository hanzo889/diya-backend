package buku

import (

	"library/app/model"
	listkategori "library/app/service/list_kategori"
	"log"
)

type Buku interface {
	Get() []ResponseBuku
	CreateBuku(bukuRequest CreateRequest) (int, string)
	Update(bukuRequest *CreateRequest, id int)
	Delete(id int)
	GetById(id int) *ResponseBuku
	GetBukuByBarcode(barcode string) (*ResponseBuku,error)
}

type bukuService struct {
	repo             Repository
	repoListKategori listkategori.Repository
}

func NewService(repo Repository, repoListKategori listkategori.Repository) Buku {
	return &bukuService{repo, repoListKategori}
}

func (b *bukuService) CreateBuku(bukuRequest CreateRequest) (int, string) {
	// getListKategori, err := b.repoListKategori.GetAll()
	// if err != nil {
	// 	return 405, "not allowed"
	// }
	// var listKategori ResponseListKategori
	// if err := getListKategori.Scan(&listKategori.Id, &listKategori.Kategori); err != nil {
	// 	return 405, "not allowed"
	// }
	buku := model.Buku{
		Judul:          bukuRequest.Judul,
		ListKategoriId: bukuRequest.ListKategoriId,
		Stock:          bukuRequest.Stock,
		Penulis:        bukuRequest.Penulis,
	}
	err := b.repo.Create(&buku)
	// err = b.repoListKategori.Create(&model.ListKategori{Kategori: listKategori.Kategori})
	if err != nil {
		return 405, "not allowed"
	}

	return 200, "created"
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

func (b *bukuService)GetBukuByBarcode(barcode string) (*ResponseBuku,error){
	bukuRepo:=b.repo.GetBukuByBarcode(barcode)
	var buku ResponseBuku
	if err:=bukuRepo.Scan(&buku.Id,&buku.Judul);err!=nil{
		return nil,err
	}
	return &buku,nil
}