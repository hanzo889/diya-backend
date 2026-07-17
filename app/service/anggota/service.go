package anggota

import (
	"library/app/model"
	"log"
)

type Anggota interface {
	Get() []ResponseAnggota
	CreateAnggota(anggotaRequest CreateRequest)
	Update(anggota *CreateRequest, id int)
	Delete(id int)
	GetById(id int) *ResponseAnggota
}

type anggotaService struct {
	repo Repository
}

func NewService(repo Repository) Anggota {
	return &anggotaService{repo}
}

func (b *anggotaService) CreateAnggota(anggotaRequest CreateRequest) {
	anggota := model.Anggota{
		NoAnggota:            anggotaRequest.NoAnggota,
		Nama:                 anggotaRequest.Nama,
		KlasifikasiAnggotaId: anggotaRequest.KlasifikasiAnggotaId,
		Alumni:               anggotaRequest.Alumni,
	}
	err := b.repo.Create(&anggota)
	if err != nil {
		log.Println(err)
	}

}
func (b *anggotaService) Get() []ResponseAnggota {
	var anggotaAnggota []ResponseAnggota
	anggotaanggotaRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return anggotaAnggota
	}
	// log.Println(anggotaanggotaRepo)
	for anggotaanggotaRepo.Next() {

		var anggota ResponseAnggota
		if err := anggotaanggotaRepo.Scan(&anggota.Id, &anggota.NoAnggota, &anggota.Nama, &anggota.KlasifikasiAnggotaId, &anggota.Alumni); err != nil {
			return anggotaAnggota
		}
		anggotaAnggota = append(anggotaAnggota, anggota)
	}
	return anggotaAnggota
}

func (b *anggotaService) Update(anggotaRequest *CreateRequest, id int) {
	anggota := &model.Anggota{
		ID:                   id,
		NoAnggota:            anggotaRequest.NoAnggota,
		Nama:                 anggotaRequest.Nama,
		KlasifikasiAnggotaId: anggotaRequest.KlasifikasiAnggotaId,
		Alumni:               anggotaRequest.Alumni,
	}
	err := b.repo.Update(anggota)
	if err != nil {
		log.Println(err)
	}
}

func (b *anggotaService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b *anggotaService) GetById(id int) *ResponseAnggota {
	barisanggota := b.repo.GetById(id)

	var anggota ResponseAnggota
	if err := barisanggota.Scan(&anggota.Id, &anggota.NoAnggota, &anggota.Nama, &anggota.KlasifikasiAnggotaId, &anggota.Alumni); err != nil {
		log.Println(err)
	}
	return &anggota
}
