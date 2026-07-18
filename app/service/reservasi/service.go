package reservasi

import (
	"library/app/model"
	"log"
)

type Reservasi interface {
	Get() []responReservasi
	CreateReservasi(rsv CreateRequest)
	Update(rsv *CreateRequest,id int)
	Delete(id int)
	GetById(id int) *responReservasi
}

type reservasiService struct {
	repo Repository
}

var data []model.Reservasi

func NewService(repo Repository) Reservasi {
	return &reservasiService{repo}
}

func (b *reservasiService) CreateReservasi(rsv CreateRequest) {
	reservasi := &model.Reservasi{
		AnggotaId:        rsv.AnggotaId,
		BukuId:           rsv.BukuId,
		BatasPengambilan: rsv.BatasPengambilan,
	}
	err := b.repo.Create(reservasi)
	if err != nil {
		log.Println(err)
	}

}
func (b *reservasiService) Get() []responReservasi {
	var reservasiii []responReservasi
	reservasiiiRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return reservasiii
	}
	// log.Println(reservasiiiRepo)
	for reservasiiiRepo.Next() {

		var reservasi responReservasi
		if err := reservasiiiRepo.Scan(&reservasi.Id, &reservasi.AnggotaId, &reservasi.BukuId, &reservasi.BatasPengambilan); err != nil {
			return reservasiii
		}
		reservasiii = append(reservasiii, reservasi)
	}
	return reservasiii
}

func (b *reservasiService) Update(rsv *CreateRequest,id int) {
	reservasi := &model.Reservasi{
		AnggotaId:        rsv.AnggotaId,
		BukuId:           rsv.BukuId,
		BatasPengambilan: rsv.BatasPengambilan,
	}
	err := b.repo.Update(reservasi)
	if err != nil {
		log.Println(err)
	}
}

func (b *reservasiService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b *reservasiService) GetById(id int) *responReservasi {
	barisreservasi := b.repo.GetById(id)

	var reservasi responReservasi
	if err := barisreservasi.Scan(&reservasi.Id, &reservasi.AnggotaId, &reservasi.BukuId, &reservasi.BatasPengambilan); err != nil {
		log.Println(err)
	}
	return &reservasi
}
