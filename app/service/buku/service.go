package buku

import (
	"fmt"
	"library/app/model"
	"library/app/service/anggota"
	bukuhub "library/app/service/buku_hub"
	listkategori "library/app/service/list_kategori"
	"log"
	"strconv"
)

type Buku interface {
	Get() []ResponseBuku
	CreateBuku(bukuRequest CreateRequest) (int, string)
	Update(bukuRequest *CreateRequest, id int)
	Delete(id int)
	GetById(id int) *ResponseBuku
	GetBukuByBarcode(barcode string) (*ResponseBuku, error)
	CreateBarcode(bukuHubRequest CreateRequestBarcode, bukuId int) (int, error)
	GetMaxBarcode() (string, error)
}

type bukuService struct {
	repo             Repository
	repoListKategori listkategori.Repository
	repoAnggota      anggota.Repository
	repoBukuHub      bukuhub.Repository
}

func NewService(repo Repository, repoListKategori listkategori.Repository, repoAnggota anggota.Repository, repoBukuHub bukuhub.Repository) Buku {
	return &bukuService{repo, repoListKategori, repoAnggota, repoBukuHub}
}

func (b *bukuService) CreateBuku(bukuRequest CreateRequest) (int, string) {
	buku := model.Buku{
		Judul:          bukuRequest.Judul,
		ListKategoriId: bukuRequest.ListKategoriId,
		Penulis:        bukuRequest.Penulis,
	}
	err := b.repo.Create(&buku)
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

func (b *bukuService) GetBukuByBarcode(barcode string) (*ResponseBuku, error) {
	bukuRepo := b.repo.GetBukuByBarcode(barcode)
	var buku ResponseBuku
	if err := bukuRepo.Scan(&buku.Id, &buku.Judul); err != nil {
		return nil, err
	}
	return &buku, nil
}
func (b *bukuService) CreateBarcode(bukuHubRequest CreateRequestBarcode, bukuId int) (int, error) {
	var noAnggota_ *int
	if bukuHubRequest.NoAnggota != nil {
		dataAnggota := b.repoAnggota.GetAnggotaByNoAnggota(*bukuHubRequest.NoAnggota)
		var tampungDataAnggota model.Anggota
		if err := dataAnggota.Scan(&tampungDataAnggota.ID, &tampungDataAnggota.NoAnggota, &tampungDataAnggota.Nama, &tampungDataAnggota.Alumni); err != nil {
			fmt.Println("tampungggggg",tampungDataAnggota)
			return 405, err
		}
		noAnggota_ = &tampungDataAnggota.ID
	}
	dataBuku := b.GetById(bukuId)
	getMaxBarcode, err := b.GetMaxBarcode()

	if err != nil {
		return 405, err
		
	}

	bukuhub := model.BukuHub{
		Barcode:       getMaxBarcode,
		BukuId:        dataBuku.Id,
		ListKondisiId: bukuHubRequest.ListKondisiId,
		AnggotaId:     noAnggota_,
		RakId:         bukuHubRequest.RakId,
	}
	fmt.Println("bukuHubbb",bukuhub)

	err = b.repoBukuHub.Create(&bukuhub)

	if err != nil {
		log.Println(err)
		return 405, err
	}
	dataBuku.Stock++
	b.Update(&CreateRequest{
		Judul:          dataBuku.Judul,
		ListKategoriId: dataBuku.ListKategoriId,
		Penulis:        dataBuku.Penulis,
		Stock:          dataBuku.Stock,
	}, bukuId)

	return 200, err

}
func (b *bukuService) GetMaxBarcode() (string, error) {
	barcode := b.repoBukuHub.GetMaxBarcode()
	var maxBarcode *string

	if err := barcode.Scan(&maxBarcode); err != nil {
		return "", err
	}
	if maxBarcode == nil {
		return "B001", nil
	} else if maxBarcode != nil {
		maxBarcodeTerakhirTampung := *maxBarcode
		barcodeTerakhir := maxBarcodeTerakhirTampung[len(maxBarcodeTerakhirTampung)-3:]
		barcodeTerakhirInt, err := strconv.Atoi(barcodeTerakhir)
		if err != nil {
			log.Println(err)
		}
		return fmt.Sprintf("B%03d", barcodeTerakhirInt+1), nil
	}
	return "", fmt.Errorf("error")
}
