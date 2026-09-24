package pinjaman

import (
	"fmt"
	"library/app/model"
	"library/app/service/anggota"
	"log"
	"time"
)

type Pinjaman interface {
	// Get() []ResponsePinjaman
	CreatePinjaman(pinjamanRequest CreateRequest) error
	// Update(pinjamanRequest *CreateRequest, id int)
	// Delete(id int)
	// GetById(id int) *ResponsePinjaman
	GetPinjamanByNoAnggota(noAnggota string) *ResponsePinjamanAnggota
	getPinjaman(anggotaId int, bukuId int) error
}

type pinjamanService struct {
	repo        Repository
	repoAnggota anggota.Repository
}

var data []model.Pinjaman

func NewService(repo Repository, repoAnggota anggota.Repository) Pinjaman {
	return &pinjamanService{repo, repoAnggota}
}
func (b *pinjamanService) GetPinjamanByNoAnggota(noAnggota string) *ResponsePinjamanAnggota {
	anggotaData := b.repoAnggota.GetAnggotaKlasifikasi(noAnggota)

	var dataAnggota TampungPinjamanAnggota
	if err := anggotaData.Scan(&dataAnggota.Id, &dataAnggota.NoAnggota, &dataAnggota.Nama, &dataAnggota.MaksBuku, &dataAnggota.MaksHari); err != nil {
		return &ResponsePinjamanAnggota{}
	}
	buku := b.repo.GetBukuPinjamanByAnggotaId(dataAnggota.Id)
	var dataBuku []ResponseBuku
	for buku.Next() {
		var bukuData ResponseBuku
		if err := buku.Scan(&bukuData.Id, &bukuData.Judul, &bukuData.TglPinjam); err != nil {
			fmt.Println(err)
			return &ResponsePinjamanAnggota{}
		}
		dataBuku = append(dataBuku, bukuData)
	}
	bolehPinjam := dataAnggota.MaksBuku >= len(dataBuku)
	return &ResponsePinjamanAnggota{
		Id:          dataAnggota.Id,
		NoAnggota:   dataAnggota.NoAnggota,
		Nama:        dataAnggota.Nama,
		BolehPinjam: &bolehPinjam,
		Buku:        dataBuku,
	}
}

func (b *pinjamanService) getPinjaman(anggotaId int, bukuId int) error {
	var tampungBuku []responseGetPinjamanByAnggotaId
	dataPinjaman := b.repo.GetPinjamanByAnggotaId(anggotaId)
	for dataPinjaman.Next() {
		var p responseGetPinjamanByAnggotaId
		if err := dataPinjaman.Scan(&p.AnggotaId, &p.BukuId, &p.MaksBuku); err != nil {
			return err
		}
		tampungBuku = append(tampungBuku, p)
	}
	if len(tampungBuku) == 0 {
		return nil
	} else if len(tampungBuku) >= tampungBuku[0].MaksBuku {
		return fmt.Errorf("melebihi batas pinjam buku!")
	}
	for _, buku := range tampungBuku {
		if buku.BukuId == bukuId {
			return fmt.Errorf("tidak boleh meminjam buku yang sama")
		}
	}
	return nil
}

func (b *pinjamanService) CreatePinjaman(pinjamanRequest CreateRequest) error {
	var err error
	if err = b.getPinjaman(pinjamanRequest.AnggotaId, pinjamanRequest.BukuId); err == nil {
		pinjaman := &model.Pinjaman{
			AnggotaId:       pinjamanRequest.AnggotaId,
			BukuId:          pinjamanRequest.BukuId,
			TglPinjam:       time.Now(),
			TglBalik:        nil,
			PetugasPinjamId: pinjamanRequest.PetugasPinjamId,
			PetugasBalikId:  pinjamanRequest.PetugasBalikId,
			KondisiAwalId:   pinjamanRequest.KondisiAwalId,
			KondisiAkhirId:  pinjamanRequest.KondisiAkhirId,
			Status:          pinjamanRequest.Status,
		}
		err := b.repo.Create(pinjaman)
		if err != nil {
			log.Println(err)
			return err
		}
		return err
	}
	return err
}

// func (b *pinjamanService) Get() []ResponsePinjaman {
// 	var pinjamanpinjaman []ResponsePinjaman
// 	pinjamanpinjamanRepo, err := b.repo.GetAll()

// 	if err != nil {
// 		log.Println(err)
// 		return pinjamanpinjaman
// 	}
// 	// log.Println(pinjamanpinjamanRepo)
// 	for pinjamanpinjamanRepo.Next() {

// 		var pinjaman ResponsePinjaman
// 		if err := pinjamanpinjamanRepo.Scan(&pinjaman.Id, &pinjaman.AnggotaId, &pinjaman.BukuId, &pinjaman.TglPinjam, &pinjaman.TglBalik,&pinjaman.PetugasPinjamId, &pinjaman.PetugasBalikId, &pinjaman.KondisiAwalId, &pinjaman.KondisiAkhirId, &pinjaman.Status); err != nil {
// 			return pinjamanpinjaman
// 		}
// 		pinjamanpinjaman = append(pinjamanpinjaman, pinjaman)
// 	}
// 	return pinjamanpinjaman
// }

// func (b *pinjamanService) Update(pinjamanRequest *CreateRequest, id int) {
// 	pinjaman := &model.Pinjaman{
// 		Id:              id,
// 		AnggotaId:       pinjamanRequest.AnggotaId,
// 		BukuId:          pinjamanRequest.BukuId,
// 		TglPinjam:       pinjamanRequest.TglPinjam,
// 		TglBalik:        pinjamanRequest.TglBalik,
// 		PetugasPinjamId: pinjamanRequest.PetugasPinjamId,
// 		PetugasBalikId: pinjamanRequest.PetugasBalikId,
// 		KondisiAwalId:   pinjamanRequest.KondisiAwalId,
// 		KondisiAkhirId:  pinjamanRequest.KondisiAkhirId,
// 		Status:            pinjamanRequest.Status,
// 	}
// 	err := b.repo.Update(pinjaman)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func (b *pinjamanService) Delete(id int) {
// 	if err := b.repo.Delete(id); err != nil {
// 		log.Println(err)
// 	}
// }

// func (b *pinjamanService) GetById(id int) *ResponsePinjaman {
// 	barispinjaman := b.repo.GetById(id)

//		var pinjaman ResponsePinjaman
//		if err := barispinjaman.Scan(&pinjaman.Id, &pinjaman.AnggotaId, &pinjaman.BukuId, &pinjaman.TglPinjam, &pinjaman.TglBalik, &pinjaman.PetugasPinjamId, &pinjaman.PetugasBalikId, &pinjaman.KondisiAwalId, &pinjaman.KondisiAkhirId, &pinjaman.Status); err != nil {
//			log.Println(err)
//		}
//		return &pinjaman
//	}
