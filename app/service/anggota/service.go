package anggota

import (
	"library/app/model"
	"log"
)

type Anggota interface {
	Get() []model.Anggota
	CreateBuku(no_anggota int, nama string, klAID int, alumni bool)
	Update(anggota model.Anggota)
	Delete(id int)
	GetById(id int) *model.Anggota
}

type anggotaService struct {
	repo Repository
}

var data []model.Anggota

func NewService(repo Repository) Anggota {
	return &anggotaService{repo}
}

func (b *anggotaService) CreateBuku(no_anggota int, nama string, klAID int, alumni bool) {
	anggota := model.Anggota{}
	err := b.repo.Create(&anggota)
	if err != nil {
		log.Println(err)
	}

}
func (b *anggotaService) Get() []model.Anggota {
	var anggotaanggota []model.Anggota
	anggotaanggotaRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return anggotaanggota
	}
	// log.Println(anggotaanggotaRepo)
	for anggotaanggotaRepo.Next() {

		var anggota model.Anggota
		if err := anggotaanggotaRepo.Scan(&anggota.ID, &anggota.NoAnggota, &anggota.Nama, &anggota.KlasifikasiAnggotaId, &anggota.Alumni); err != nil {
			return anggotaanggota
		}
		anggotaanggota = append(anggotaanggota, anggota)
	}
	return anggotaanggota
}

func (b *anggotaService) Update(anggota model.Anggota) {
	err := b.repo.Update(&anggota)
	if err != nil {
		log.Println(err)
	}
}

func (b *anggotaService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b *anggotaService) GetById(id int) *model.Anggota {
	barisanggota := b.repo.GetById(id)

	var anggota model.Anggota
	if err := barisanggota.Scan(&anggota.ID, &anggota.NoAnggota, &anggota.Nama, &anggota.KlasifikasiAnggotaId, &anggota.Alumni); err != nil {
		log.Println(err)
	}
	return &anggota
}
