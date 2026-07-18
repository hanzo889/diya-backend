package klasifikasianggota

import (
	"library/app/model"
	// "library/service/klasifikasi_anggota/repository.go"
	"log"
)

type KlasifikasiAnggota interface {
	Get() []responseKlasifikasiAnggota
	CreateKlasifikasiAnggota(klasifikasiAnggotaRequest CreateRequest)
	Update(klasifikasiAnggotaRequest *CreateRequest, id int)
	Delete(id int)
	GetById(id int) *responseKlasifikasiAnggota
}

type klasifikasianggotaService struct {
	repo Repository
}

var data []model.KlasifikasiAnggota

func NewService(repo Repository) KlasifikasiAnggota {
	return &klasifikasianggotaService{repo}
}

func (b *klasifikasianggotaService) CreateKlasifikasiAnggota(klasifikasiAnggotaRequest CreateRequest) {
	klasifikasianggota := &model.KlasifikasiAnggota{
		Klasifikasi: klasifikasiAnggotaRequest.Klasifikasi,
		MaksBuku:    klasifikasiAnggotaRequest.MaksBuku,
		MaksHari:    klasifikasiAnggotaRequest.MaksHari,
	}
	err := b.repo.Create(klasifikasianggota)
	if err != nil {
		log.Println(err)
	}

}
func (b *klasifikasianggotaService) Get() []responseKlasifikasiAnggota {
	var klasifikasiAnggotaAnggota []responseKlasifikasiAnggota
	klA2Repo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return klasifikasiAnggotaAnggota
	}
	// log.Println(bukuBukuRepo)
	for klA2Repo.Next() {

		var klasifikasiAnggota responseKlasifikasiAnggota
		if err := klA2Repo.Scan(&klasifikasiAnggota.Id, &klasifikasiAnggota.Klasifikasi, &klasifikasiAnggota.MaksBuku, &klasifikasiAnggota.MaksHari); err != nil {
			return klasifikasiAnggotaAnggota
		}
		klasifikasiAnggotaAnggota = append(klasifikasiAnggotaAnggota, klasifikasiAnggota)
	}
	return klasifikasiAnggotaAnggota
}

func (b *klasifikasianggotaService) Update(klasifikasiAnggotaRequest *CreateRequest, id int) {
	klasifikasiAnggota := &model.KlasifikasiAnggota{
		Id:          id,
		Klasifikasi: klasifikasiAnggotaRequest.Klasifikasi,
		MaksBuku:    klasifikasiAnggotaRequest.MaksBuku,
		MaksHari:    klasifikasiAnggotaRequest.MaksHari,
	}
	err := b.repo.Update(klasifikasiAnggota)
	if err != nil {
		log.Println(err)
	}
}

func (b *klasifikasianggotaService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b *klasifikasianggotaService) GetById(id int) *responseKlasifikasiAnggota {
	barisKlasifikasiAnggota := b.repo.GetById(id)

	var klasifikasiAnggota responseKlasifikasiAnggota
	if err := barisKlasifikasiAnggota.Scan(&klasifikasiAnggota.Id, &klasifikasiAnggota.Klasifikasi, &klasifikasiAnggota.MaksBuku, &klasifikasiAnggota.MaksHari); err != nil {
		log.Println(err)
	}
	return &klasifikasiAnggota
}
