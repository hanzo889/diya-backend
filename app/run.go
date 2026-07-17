package app

import (
	"library/app/service/buku"
	bukuhub "library/app/service/buku_hub"
	"library/app/service/denda"
	hargabuku "library/app/service/harga_buku"
	listkondisi "library/app/service/list_kondisi"
	"library/app/service/rak"
	"library/database"

	"github.com/gin-gonic/gin"
)

func Run() {
	db := database.MysqlConnection()
	defer db.Close()
	router := gin.Default()

	api := router.Group("/api")

	buku.Router(api, db)
	bukuhub.Router(api,db)
	denda.Router(api,db)
	hargabuku.Router(api,db)
	rak.Router(api,db)
	listkondisi.Router(api,db)
	router.Run()
}
