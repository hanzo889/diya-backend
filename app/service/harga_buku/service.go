package hargabuku

import (
	"library/app/model"
	"log"
)

type HargaBuku interface {
	Get() []model.HargaBuku
	CreateHargaBuku(bukuID int, HargaB int)
	Update(hargaB model.HargaBuku)
	Delete(id int)
	GetById(id int) *model.HargaBuku
}

type HargaBukuService struct {
	repo Repository
}

var data []model.HargaBuku

func NewService(repo Repository) HargaBuku {
	return &HargaBukuService{repo}
}

func (b *HargaBukuService) CreateHargaBuku(bukuID int, hargaB int) {
	hargaBuku := &model.HargaBuku{
		BukuId: bukuID,
		Harga:  hargaB,
	}
	err := b.repo.CreateHargaBuku(hargaBuku)
	if err != nil {
		log.Println(err)
	}

}
func (b *HargaBukuService) Get() []model.HargaBuku {
	var hargaharga []model.HargaBuku
	hargahargaRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return hargaharga
	}
	// log.Println(hargahargaRepo)
	for hargahargaRepo.Next() {

		var hargaBuku model.HargaBuku
		if err := hargahargaRepo.Scan(&hargaBuku.Id, &hargaBuku.BukuId, &hargaBuku.Harga); err != nil {
			return hargaharga
		}
		hargaharga = append(hargaharga, hargaBuku)
	}
	return hargaharga
}

func (b *HargaBukuService) Update(hargaB model.HargaBuku) {
	err := b.repo.Update(&hargaB)
	if err != nil {
		log.Println(err)
	}
}

func (b *HargaBukuService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b HargaBukuService) GetById(id int) *model.HargaBuku {
	bariskondisi := b.repo.GetById(id)

	var hargaBuku model.HargaBuku
	if err := bariskondisi.Scan(&hargaBuku.Id, &hargaBuku.BukuId, &hargaBuku.Harga); err != nil {
		log.Println(err)
	}
	return &hargaBuku
}
