package bukuhub

import (
	"library/app/model"
	"log"
)

type BukuHub interface {
	Get() []responseBukuHub
	CreateBukuHub(bukuHubRequest CreateRequest)
	Update(bukuRequest *CreateRequest, id int)
	Delete(id int)
	GetById(id int) *responseBukuHub
}

type bukuHubService struct {
	repo Repository
}

var data []model.BukuHub

func NewService(repo Repository) BukuHub {
	return &bukuHubService{repo}
}

func (b *bukuHubService) CreateBukuHub(bukuHubRequest CreateRequest) {
	bukuhub := &model.BukuHub{
		Barcode:       bukuHubRequest.Barcode,
		BukuId:        bukuHubRequest.ListKondisiId,
		ListKondisiId: bukuHubRequest.ListKondisiId,
		AnggotaId:     bukuHubRequest.AnggotaId,
		RakId:         bukuHubRequest.RakId,
	}
	err := b.repo.Create(bukuhub)
	if err != nil {
		log.Println(err)
	}

}
func (b *bukuHubService) Get() []responseBukuHub {
	var bukuBukuHub []responseBukuHub
	bukuBukuHubRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return bukuBukuHub
	}
	// log.Println(bukuBukuHubRepo)
	for bukuBukuHubRepo.Next() {

		var bukuhub responseBukuHub
		if err := bukuBukuHubRepo.Scan(&bukuhub.Id, &bukuhub.Barcode, &bukuhub.BukuId, &bukuhub.ListKondisiId, &bukuhub.AnggotaId, &bukuhub.RakId); err != nil {
			return bukuBukuHub
		}
		bukuBukuHub = append(bukuBukuHub, bukuhub)
	}
	return bukuBukuHub
}

func (b *bukuHubService) Update(bukuHubRequest *CreateRequest, id int) {
	bukuhub := &model.BukuHub{
		ID:            bukuHubRequest.BukuId,
		Barcode:       bukuHubRequest.Barcode,
		BukuId:        bukuHubRequest.ListKondisiId,
		ListKondisiId: bukuHubRequest.ListKondisiId,
		AnggotaId:     bukuHubRequest.AnggotaId,
		RakId:         bukuHubRequest.RakId,
	}
	err := b.repo.Update(bukuhub)
	if err != nil {
		log.Println(err)
	}
}

func (b *bukuHubService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b *bukuHubService) GetById(id int) *responseBukuHub {
	barisBukuHub := b.repo.GetById(id)

	var bukuhub responseBukuHub
	if err := barisBukuHub.Scan(&bukuhub.Id, &bukuhub.Barcode, &bukuhub.BukuId, &bukuhub.ListKondisiId, &bukuhub.AnggotaId, &bukuhub.RakId); err != nil {
		log.Println(err)
	}
	return &bukuhub
}
