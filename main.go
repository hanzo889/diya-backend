package main

import (
	"library/app"
)

// import (
// 	"github.com/gin-gonic/gin"
// )

// type Siswa struct{
// 	Nisn int `json:"nisn"`
//  	Nama string `json:"nama"`
// }

// "fmt"

// klasifikasianggota "library/service/klasifikasi_anggota"

// "library/app/model"

// "library/app/model"
// listkondisi "library/service/list_kondisi"

// "library/app/model"
// "library/service/buku"github.com/gin-gonic/gin

// "library/app/model"

// "library/app/model"
// "library/service/buku"

// "log"

// func Stop(){
// 	time.Sleep(2*time.Second)
// 	println("stop")
// }

func main() {

	app.Run()
	// defer Stop()
	// fmt.Println("halo")

	// siswa:=Siswa{
	// 	Nisn: 134,
	// 	Nama: "royyan",
	// }

	// r := gin.Default()
	// r.GET("",func (c *gin.Context)  {
	// 	c.JSON(200,gin.H{
	// 		"message":"hello bree",
	// 		"naji":"ganteng",
	// 		"royyan":"ganteng sekedik",
	// 		"hanzo":"ganteng banget",
	// 		"ardista":"ndak paham2",
	// 		"zardista":"ndak paham2",
	// 		"bang zardista":"ndak paham2",
	// 	})

	// })

	// r.GET("/siswa",func (c *gin.Context)  {
	// 	c.JSON(200,siswa)

	// })

	// r.Run()

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

	//---------------/kondisi//---------------------

	// db := database.MysqlConnection()
	// defer db.Close()
	// log.Println("database terkoneksi ")
	// kondisiRepo := listkondisi.NewRepository(db)
	// kondisiService := listkondisi.NewService(kondisiRepo)
	// kondisiService.CreateListKondisi("burukkk")

	// kondisikondisi := kondisiService.Get()
	// fmt.Println(kondisikondisi)

	// kondisiService.Update(model.ListKondisi{
	// 	Id:      1,
	// 	Kondisi: "jelek",
	// })
	// kondisikondisi = kondisiService.Get()
	// fmt.Println(kondisikondisi)
	// kondisiService.Delete(1)
	// kondisikondisi = kondisiService.Get()
	// fmt.Println(kondisikondisi)
	// kondisi := kondisiService.GetById(2)
	// fmt.Println(kondisi)

	//---------------------buku/-------------/

	// db := database.MysqlConnection()
	// defer db.Close()
	// log.Println("database terkoneksi ")
	// bukuRepo := buku.NewRepository(db)
	// bukuService := buku.NewService(bukuRepo)
	// // bukuService.CreateBuku("BUMI MANUSIA", 2, 5)

	// bukuBuku := bukuService.Get()
	// fmt.Println(bukuBuku)

	// bukuService.Update(model.Buku{
	// 	Id: 5,
	// 	Judul: "atomic habit",
	// 	ListKategoriId: 2,
	// 	Stock: 10,
	// })
	// bukuBuku = bukuService.Get()
	// fmt.Println(bukuBuku)
	// bukuService.Delete(2)
	// bukuBuku = bukuService.Get()
	// fmt.Println(bukuBuku)
	// buku := bukuService.GetById(6)
	// fmt.Println("idddd", buku)

	//##############klasifikasi_anggota####################//

	// db := database.MysqlConnection()
	// defer db.Close()
	// log.Println("database terkoneksi ")
	// klARepo := klasifikasianggota.NewRepository(db)
	// klAService := klasifikasianggota.NewService(klARepo)
	// klAService.CreateBuku("guruuuu",100,7)

	// klA2 := klAService.Get()
	// fmt.Println("get:",klA2)

	// klAService.Update(model.KlasifikasiAnggota{
	// 	Id: 2,
	// 	Klasifikasi: "muriddddd",
	// 	MaksBuku: 4,
	// 	MaksHari: 6,
	// })
	// klA2 = klAService.Get()
	// fmt.Println(klA2)
	// klAService.Delete(1)
	// klA2 = klAService.Get()
	// fmt.Println(klA2)
	// kla := klAService.GetById(2)
	// fmt.Println(kla)

	//############         #################////
}
