package main

import (
	"fmt"
	"library/database"
	// "library/model"
	"library/service/buku"
	"log"
)

func main() {
	// anggotaService := anggota.New() //anggota.AnggotaService{}
	// anggotaService.CreateAnggota(model.Anggota{ID: 1, NoAnggota: "1", Nama: "kkj", KlasifikasiAnggotaId: 1, Alumni: true})
	// anggotaService.CreateAnggota(model.Anggota{ID: 2, NoAnggota: "2", Nama: "jeki", KlasifikasiAnggotaId: 2, Alumni: false})
	// anggotaService.CreateAnggota(model.Anggota{ID: 3, NoAnggota: "3", Nama: "nowa", KlasifikasiAnggotaId: 1, Alumni: false})
	// fmt.Println(anggotaService.Get())
	// bukuService := buku.New()
	// bukuService.CreateBuku(model.Buku{Id: 1, Judul: "1001 malam", ListKategoriId: 1, Stock: 2})
	// bukuService.CreateBuku(model.Buku{Id: 2, Judul: "atomic habits", ListKategoriId: 2, Stock: 1})
	// bukuService.CreateBuku(model.Buku{Id: 3, Judul: "seporsi mie ayam sebelum mati", ListKategoriId: 4, Stock: 4})
	// fmt.Println(bukuService.Get())
	// klasifikasiService := klasifikasianggota.New()
	// klasifikasiService.CreateKlasifikasiAnggota(model.KlasifikasiAnggota{Id: 1, Klasifikasi: "siswa", MaksBuku: 3, MaksHari: 7})
	// klasifikasiService.CreateKlasifikasiAnggota(model.KlasifikasiAnggota{Id: 2, Klasifikasi: "guru", MaksBuku: 5, MaksHari: 14})
	// klasifikasiService.CreateKlasifikasiAnggota(model.KlasifikasiAnggota{Id: 3, Klasifikasi: "staf TU", MaksBuku: 2, MaksHari: 7})
	// fmt.Println(klasifikasiService.Get())

	db := database.MysqlConnection()
	defer db.Close()
	log.Println("database terkoneksi ")
	bukuRepo := buku.NewRepository(db)
	bukuService := buku.NewService(bukuRepo)
	// bukuService.CreateBuku("BUMI MANUSIA",2,5)

	bukuBuku := bukuService.Get()
	fmt.Println(bukuBuku)

	// bukuService.Update(model.Buku{
	// 	Id:             1,
	// 	Judul:          "bumiii",
	// 	ListKategoriId: 3,
	// 	Stock:          1,
	// })
	bukuBuku=bukuService.Get()
	fmt.Println(bukuBuku)
	// bukuService.Delete(1)
	// bukuBuku=bukuService.Get()
	// fmt.Println(bukuBuku)
	buku:=bukuService.GetById(2)
	fmt.Println(buku)

}
