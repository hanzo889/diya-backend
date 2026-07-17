package pinjaman

import (
	"library/app/model"
	"log"
	"time"
)

type Pinjaman interface {
	Get() []model.Pinjaman
	CreatePinjaman(anggota_id int, buku_id int, tgl_pinjam time.Time, tgl_balik time.Time, petugas_pinjam_id int, list_kondisi_id int, petugas_balik_id int)
	Update(pinjaman model.Pinjaman)
	Delete(id int)
	GetById(id int) *model.Pinjaman
}

type pinjamanService struct {
	repo Repository
}

var data []model.Pinjaman

func NewService(repo Repository) Pinjaman {
	return &pinjamanService{repo}
}

func (b *pinjamanService) CreatePinjaman(anggota_id int, buku_id int, tgl_pinjam time.Time, tgl_balik time.Time, petugas_pinjam_id int, list_kondisi_id int, petugas_balik_id int) {
	pinjaman := &model.Pinjaman{
		AnggotaId:       anggota_id,
		BukuId:          buku_id,
		TglPinjam:       tgl_pinjam,
		TglBalik:        tgl_balik,
		PetugasPinjamId: petugas_pinjam_id,
		ListKondisiId:   list_kondisi_id,
		PetuagasBalikId: petugas_balik_id,
	}
	err := b.repo.Create(pinjaman)
	if err != nil {
		log.Println(err)
	}

}
func (b *pinjamanService) Get() []model.Pinjaman {
	var pinjamanpinjaman []model.Pinjaman
	pinjamanpinjamanRepo, err := b.repo.GetAll()

	if err != nil {
		log.Println(err)
		return pinjamanpinjaman
	}
	// log.Println(pinjamanpinjamanRepo)
	for pinjamanpinjamanRepo.Next() {

		var pinjaman model.Pinjaman
		if err := pinjamanpinjamanRepo.Scan(&pinjaman.Id, &pinjaman.AnggotaId, &pinjaman.BukuId, &pinjaman.TglPinjam, &pinjaman.TglBalik, &pinjaman.ListKondisiId, &pinjaman.ListKondisiId, &pinjaman.PetuagasBalikId); err != nil {
			return pinjamanpinjaman
		}
		pinjamanpinjaman = append(pinjamanpinjaman, pinjaman)
	}
	return pinjamanpinjaman
}

func (b *pinjamanService) Update(pinjaman model.Pinjaman) {
	err := b.repo.Update(&pinjaman)
	if err != nil {
		log.Println(err)
	}
}

func (b *pinjamanService) Delete(id int) {
	if err := b.repo.Delete(id); err != nil {
		log.Println(err)
	}
}

func (b *pinjamanService) GetById(id int) *model.Pinjaman {
	barispinjaman := b.repo.GetById(id)

	var pinjaman model.Pinjaman
	if err := barispinjaman.Scan(&pinjaman.Id, &pinjaman.AnggotaId, &pinjaman.BukuId, &pinjaman.TglPinjam, &pinjaman.TglBalik, &pinjaman.ListKondisiId, &pinjaman.ListKondisiId, &pinjaman.PetuagasBalikId); err != nil {
		log.Println(err)
	}
	return &pinjaman
}
