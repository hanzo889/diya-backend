package hargabuku

import (
	"library/app/model"
	"log"
)

type HargaBuku interface {
	Get() []responseHargaBuku
	CreateHargaBuku(hbr CreateRequest)
	Update(hbr *CreateRequest, id int)
	Delete(id int)
	GetById(id int) *responseHargaBuku
}

type HargaBukuService struct {
	repo Repository
}

var data []model.HargaBuku

func NewService(repo Repository) HargaBuku {
	return &HargaBukuService{repo}
}

func (b *HargaBukuService) CreateHargaBuku(hbr CreateRequest) {
	hargaBuku := &model.HargaBuku{
		BukuId: hbr.BukuId,
		Harga:  hbr.Harga,
	}
	err := b.repo.CreateHargaBuku(hargaBuku)
	if err != nil {
		log.Println(err)
	}

}
func (b *HargaBukuService) Get() []responseHargaBuku {
	var hargaharga []responseHargaBuku
	hargahargaRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return hargaharga
	}
	// log.Println(hargahargaRepo)
	for hargahargaRepo.Next() {

		var hargaBuku responseHargaBuku
		if err := hargahargaRepo.Scan(&hargaBuku.Id, &hargaBuku.BukuId, &hargaBuku.Harga); err != nil {
			return hargaharga
		}
		hargaharga = append(hargaharga, hargaBuku)
	}
	return hargaharga
}

func (b *HargaBukuService) Update(hbr *CreateRequest, id int) {
	hargaB := &model.HargaBuku{
		Id: id,
		BukuId: hbr.BukuId,
		Harga:  hbr.Harga,
	}

	err := b.repo.Update(hargaB)
	if err != nil {
		log.Println(err)
	}
}

func (b *HargaBukuService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b HargaBukuService) GetById(id int) *responseHargaBuku {
	bariskondisi := b.repo.GetById(id)

	var hargaBuku responseHargaBuku
	if err := bariskondisi.Scan(&hargaBuku.Id, &hargaBuku.BukuId, &hargaBuku.Harga); err != nil {
		log.Println(err)
	}
	return &hargaBuku
}
