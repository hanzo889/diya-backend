package pinjaman

import (
	"fmt"
	"library/app/model"
	"library/app/service/anggota"
	"library/app/service/buku"
	"log"
	"time"
)

type Pinjaman interface {
	// Get() []ResponsePinjaman
	CreatePinjaman(pinjamanRequest CreateRequest) error
	UpdateTglBalik(id int) error
	// Delete(id int)
	// GetById(id int) *ResponsePinjaman
	GetPinjamanByNoAnggota(noAnggota string) *ResponsePinjamanAnggota
	getPinjaman(anggotaId int, bukuId int) error
}

type pinjamanService struct {
	repo        Repository
	repoAnggota anggota.Repository
	repoBuku    buku.Repository
}

var data []model.Pinjaman

func NewService(repo Repository, repoAnggota anggota.Repository, repoBuku buku.Repository) Pinjaman {
	return &pinjamanService{repo, repoAnggota, repoBuku}
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
	var tampungStock []responseStock

	dataPinjaman := b.repo.GetPinjamanByAnggotaId(anggotaId)
	dataBuku, err := b.repoBuku.GetAll()

	for dataBuku.Next() {
		var c responseStock

		if err := dataBuku.Scan(&c.Stock); err != nil {
			return err
		}
		tampungStock = append(tampungStock, c)
	}
	if len(tampungStock) == 0 {
		return err
	}
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

func (b *pinjamanService) UpdateTglBalik(id int) error {
	getBuku := b.repo.GetById(id)

	var c model.Pinjaman
	if err := getBuku.Scan(&c.Id, &c.AnggotaId, &c.BukuId, &c.TglPinjam, &c.TglBalik, &c.PetugasPinjamId, &c.PetugasBalikId, &c.KondisiAwalId, &c.KondisiAkhirId, &c.Status); err != nil {
		fmt.Println("####")
		fmt.Println(err)
		return err
	}
	if c.Id != 0 && c.TglBalik == nil {
		now := time.Now()
		err := b.repo.UpdateTglBalik(now, id)
		return err
	}

	return fmt.Errorf("Data Tidak Di Temukan")
}
