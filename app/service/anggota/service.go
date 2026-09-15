package anggota

import (
	"fmt"
	"library/app/model"
	klasifikasianggotahub "library/app/service/klasifikasi_anggota_hub"
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
	repo                      Repository
	repoKlasifikasiAnggotaHub klasifikasianggotahub.Repository
}

func NewService(repo Repository, repoKah klasifikasianggotahub.Repository) Anggota {
	return &anggotaService{repo, repoKah}
}

func (b *anggotaService) CreateAnggota(anggotaRequest CreateRequest) (int, string) {
	getMaxNoAnggota, err := b.GetMaxNoAnggota()
	fmt.Println(anggotaRequest)
	if err != nil {
		fmt.Println("--- 1 ---")
		fmt.Println(err)
		return 405, "not allowed 1"

	}

	anggota := model.Anggota{
		NoAnggota: getMaxNoAnggota,
		Nama:      anggotaRequest.Nama,
		Alumni:    anggotaRequest.Alumni,
	}
	data, err := b.repo.Create(&anggota)
	lastId, err := data.LastInsertId()
	fmt.Println(data.LastInsertId())
	if err != nil {
		fmt.Println("--- 2 ---")
		return 405, "not allowed 2"
	}
	err = b.repoKlasifikasiAnggotaHub.Create(&model.KlasifikasiAnggotaHub{AnggotaId: int(lastId), KlasifikasiAnggotaId: anggotaRequest.KlasifikasiAnggotaId})
	if err != nil {
		fmt.Println("--- 3 ---")
		return 405, "not allowed 3"
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
	for anggotaanggotaRepo.Next() {

		var anggota ResponseAnggota
		if err := anggotaanggotaRepo.Scan(&anggota.Id, &anggota.NoAnggota, &anggota.Nama, &anggota.Alumni, &anggota.KlasifikasiAnggotaId); err != nil {
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
	if err := b.repoKlasifikasiAnggotaHub.Update(anggotaRequest.KlasifikasiAnggotaId, anggota.ID); err == nil {
		err = b.repo.Update(anggota)
		if err != nil {
			log.Println(err)
		}
	}
}

func (b *anggotaService) Delete(id int) {
	var err error
	if err = b.repoKlasifikasiAnggotaHub.Delete(id); err != nil {
		log.Println(err)
	}
	if err == nil {
		if err = b.repo.Delete(id); err != nil {
			log.Println(err)
		}
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
	var maxNoAnggota *string

	if err := noAnggota.Scan(&maxNoAnggota); err != nil {

		return "", err
	}
	fmt.Println(*maxNoAnggota)

	if maxNoAnggota == nil {
		return "A001", nil
	} else if maxNoAnggota != nil {
		noAnggotaTerakhirTampung := *maxNoAnggota
		noAnggotaTerakhir := noAnggotaTerakhirTampung[len(noAnggotaTerakhirTampung)-3:]
		noAnggotaTerakhirInt, err := strconv.Atoi(noAnggotaTerakhir)
		if err != nil {
			log.Println(err)
		}
		return fmt.Sprintf("A%03d", noAnggotaTerakhirInt+1), nil
	}
	return "", fmt.Errorf("error")
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
