package reservasi

import (
	"library/app/model"
	"log"
	"time"
)

type Reservasi interface {
	Get() []model.Reservasi
	CreateReservasi(anggota_id int, buku_id int, batas_pengembaliaan time.Time)
	Update(reservasi model.Reservasi)
	Delete(id int)
	GetById(id int) *model.Reservasi
}

type reservasiService struct {
	repo Repository
}

var data []model.Reservasi

func NewService(repo Repository) Reservasi {
	return &reservasiService{repo}
}

func (b *reservasiService) CreateReservasi(anggota_id int, buku_id int, batas_pengembaliaan time.Time) {
	reservasi := &model.Reservasi{
		AnggotaId:        anggota_id,
		BukuId:           buku_id,
		BatasPengambilan: batas_pengembaliaan,
	}
	err := b.repo.Create(reservasi)
	if err != nil {
		log.Println(err)
	}

}
func (b *reservasiService) Get() []model.Reservasi {
	var reservasiii []model.Reservasi
	reservasiiiRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return reservasiii
	}
	// log.Println(reservasiiiRepo)
	for reservasiiiRepo.Next() {

		var reservasi model.Reservasi
		if err := reservasiiiRepo.Scan(&reservasi.Id, &reservasi.AnggotaId, &reservasi.BukuId, &reservasi.BatasPengambilan); err != nil {
			return reservasiii
		}
		reservasiii = append(reservasiii, reservasi)
	}
	return reservasiii
}

func (b *reservasiService) Update(reservasi model.Reservasi) {
	err := b.repo.Update(&reservasi)
	if err != nil {
		log.Println(err)
	}
}

func (b *reservasiService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b *reservasiService) GetById(id int) *model.Reservasi {
	barisreservasi := b.repo.GetById(id)

	var reservasi model.Reservasi
	if err := barisreservasi.Scan(&reservasi.Id, &reservasi.AnggotaId, &reservasi.BukuId, &reservasi.BatasPengambilan); err != nil {
		log.Println(err)
	}
	return &reservasi
}
