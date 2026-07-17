package klasifikasianggota

import (
	"library/app/model"
	// "library/service/klasifikasi_anggota/repository.go"
	"log"
)

type KlasifikasiAnggota interface {
	Get() []model.KlasifikasiAnggota
	CreateBuku(kl string, maksB int, maksH int)
	Update(KlA model.KlasifikasiAnggota)
	Delete(id int)
	GetById(id int) *model.KlasifikasiAnggota
}

type klasifikasianggotaService struct {
	repo Repository
}

var data []model.KlasifikasiAnggota

func NewService(repo Repository) KlasifikasiAnggota {
	return &klasifikasianggotaService{repo}
}

func (b *klasifikasianggotaService) CreateBuku(kl string, maksB int, maksH int) {
	Kla := &model.KlasifikasiAnggota{
		Klasifikasi: kl,
		MaksBuku:    maksB,
		MaksHari:    maksH,
	}
	err := b.repo.Create(Kla)
	if err != nil {
		log.Println(err)
	}

}
func (b *klasifikasianggotaService) Get() []model.KlasifikasiAnggota {
	var klA2 []model.KlasifikasiAnggota
	klA2Repo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return klA2
	}
	// log.Println(bukuBukuRepo)
	for klA2Repo.Next() {

		var kla model.KlasifikasiAnggota
		if err := klA2Repo.Scan(&kla.Id, &kla.Klasifikasi, &kla.MaksBuku, &kla.MaksHari); err != nil {
			return klA2
		}
		klA2 = append(klA2, kla)
	}
	return klA2
}

func (b *klasifikasianggotaService) Update(KlA model.KlasifikasiAnggota) {
	err := b.repo.Update(&KlA)
	if err != nil {
		log.Println(err)
	}
}

func (b *klasifikasianggotaService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b *klasifikasianggotaService) GetById(id int) *model.KlasifikasiAnggota {
	barisklA := b.repo.GetById(id)

	var klA model.KlasifikasiAnggota
	if err := barisklA.Scan(&klA.Id, &klA.Klasifikasi, &klA.MaksBuku, &klA.MaksHari); err != nil {
		log.Println(err)
	}
	return &klA
}
