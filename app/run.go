package app

import (
	"library/app/service/buku"
	"library/database"

	"github.com/gin-gonic/gin"
)

func Run() {
	db := database.MysqlConnection()
	defer db.Close()
	router := gin.Default()

	api := router.Group("/api")
	
	buku.Router(api, db)

	router.Run()
}
