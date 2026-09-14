package anggota

import (
	"fmt"
	"library/app/model"
	"log"
	"strconv"
)

type Anggota interface {
	Get() []ResponseAnggota
	CreateAnggota(anggotaRequest CreateRequest) (int, string)
	Update(anggota *CreateRequest, id int)
	Delete(id int)
	GetById(id int) *ResponseAnggota
	GetMaxNoAnggota() (string, error)
	GetAnggotaByQuery(noAnggota string) (*ResponseAnggota, error)
}

type anggotaService struct {
	repo Repository
}

func NewService(repo Repository) Anggota {
	return &anggotaService{repo}
}

func (b *anggotaService) CreateAnggota(anggotaRequest CreateRequest) (int, string) {
	getMaxNoAnggota, err := b.GetMaxNoAnggota()
	if err != nil {
		return 405, "not allowed"
	}
	anggota := model.Anggota{
		NoAnggota: getMaxNoAnggota,
		Nama:      anggotaRequest.Nama,
		Alumni:    anggotaRequest.Alumni,
	}
	err = b.repo.Create(&anggota)
	if err != nil {
		return 405, "not allowed"
	}
	return 200, "created"

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
		if err := anggotaanggotaRepo.Scan(&anggota.Id, &anggota.NoAnggota, &anggota.Nama, &anggota.Alumni); err != nil {
			return anggotaAnggota
		}
		anggotaAnggota = append(anggotaAnggota, anggota)
	}
	return anggotaAnggota
}

func (b *anggotaService) Update(anggotaRequest *CreateRequest, id int) {
	anggota := &model.Anggota{
		ID:     id,
		Nama:   anggotaRequest.Nama,
		Alumni: anggotaRequest.Alumni,
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
	if err := barisanggota.Scan(&anggota.Id, &anggota.NoAnggota, &anggota.Nama, &anggota.Alumni); err != nil {
		log.Println(err)
	}
	return &anggota
}
func (b *anggotaService) GetMaxNoAnggota() (string, error) {
	noAnggota := b.repo.GetMaxNoAnggota()
	var maxNoAnggota string
	if err := noAnggota.Scan(&maxNoAnggota); err != nil {
		return "", err
	}
	if maxNoAnggota == "" {
		return "A001", nil
	}
	noAnggotaTerakhir := maxNoAnggota[len(maxNoAnggota)-3:]
	noAnggotaTerakhirInt, err := strconv.Atoi(noAnggotaTerakhir)
	if err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("A%03d", noAnggotaTerakhirInt+1), nil
}
func (b *anggotaService) GetAnggotaByQuery(noAnggota string) (*ResponseAnggota, error) {
	anggotaRepo := b.repo.GetAnggotaByNoAnggota(noAnggota)
	var anggota ResponseAnggota
	if err := anggotaRepo.Scan(&anggota.Id, &anggota.NoAnggota, &anggota.Nama, &anggota.Alumni); err != nil {
		log.Println(err)
		return nil, err
	}
	return &anggota, nil
}
