package bukuhub

import (
	"library/app/model"
	"log"
)

type BukuHub interface {
	Get() []model.BukuHub
	CreateBukuHub(barcode int, buku_id int, list_kondisi_id int, anggota_id int, rak_id int)
	Update(bukuhub model.BukuHub)
	Delete(id int)
	GetById(id int) *model.BukuHub
}

type bukuHubService struct {
	repo Repository
}

var data []model.BukuHub

func NewService(repo Repository) BukuHub {
	return &bukuHubService{repo}
}

func (b *bukuHubService) CreateBukuHub(barcode int, buku_id int, list_kondisi_id int, anggota_id int, rak_id int) {
	bukuhub := &model.BukuHub{}
	err := b.repo.Create(bukuhub)
	if err != nil {
		log.Println(err)
	}

}
func (b *bukuHubService) Get() []model.BukuHub {
	var bukuBukuHub []model.BukuHub
	bukuBukuHubRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return bukuBukuHub
	}
	// log.Println(bukuBukuHubRepo)
	for bukuBukuHubRepo.Next() {

		var bukuhub model.BukuHub
		if err := bukuBukuHubRepo.Scan(&bukuhub.ID, &bukuhub.Barcode, &bukuhub.BukuId, &bukuhub.ListKondisiId, &bukuhub.AnggotaId, &bukuhub.RakId); err != nil {
			return bukuBukuHub
		}
		bukuBukuHub = append(bukuBukuHub, bukuhub)
	}
	return bukuBukuHub
}

func (b *bukuHubService) Update(buku model.BukuHub) {
	err := b.repo.Update(&buku)
	if err != nil {
		log.Println(err)
	}
}

func (b *bukuHubService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b *bukuHubService) GetById(id int) *model.BukuHub {
	barisBukuHub := b.repo.GetById(id)

	var bukuhub model.BukuHub
	if err := barisBukuHub.Scan(&bukuhub.ID, &bukuhub.Barcode, &bukuhub.BukuId, &bukuhub.ListKondisiId, &bukuhub.AnggotaId, &bukuhub.RakId); err != nil {
		log.Println(err)
	}
	return &bukuhub
}
