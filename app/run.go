package app

import (
	"library/app/service/anggota"
	"library/app/service/buku"
	bukuhub "library/app/service/buku_hub"
	"library/app/service/denda"
	hargabuku "library/app/service/harga_buku"
	klasifikasianggota "library/app/service/klasifikasi_anggota"
	listkategori "library/app/service/list_kategori"
	listkondisi "library/app/service/list_kondisi"
	"library/app/service/petugas"
	"library/app/service/pinjaman"
	"library/app/service/rak"
	"library/app/service/reservasi"
	"library/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	db := database.MysqlConnection()
	defer db.Close()
	router := gin.Default()
	router.SetTrustedProxies(nil)
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = false
	router.Use(cors.New(config))
	api := router.Group("/api")

	buku.Router(api, db)
	bukuhub.Router(api, db)
	denda.Router(api, db)
	hargabuku.Router(api, db)
	rak.Router(api, db)
	listkondisi.Router(api, db)
	listkategori.Router(api, db)
	anggota.Router(api, db)
	klasifikasianggota.Router(api, db)
	reservasi.Router(api, db)
	pinjaman.Router(api, db)
	petugas.Router(api, db)
	router.Run()
}
