package buku

import (
	"database/sql"
	"fmt"
	"library/app/service/anggota"
	bukuhub "library/app/service/buku_hub"
	listkategori "library/app/service/list_kategori"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup, db *sql.DB) {
	repo := NewRepository(db)
	repoListKategori := listkategori.NewRepository(db)
	repoAnggota := anggota.NewRepository(db)
	repoBukuHub := bukuhub.NewRepository(db)

	service := NewService(repo, repoListKategori, repoAnggota, repoBukuHub)

	buku := g.Group("/buku")
	buku.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, service.Get())
	})

	buku.GET("/:id", func(ctx *gin.Context) {

		id, _ := strconv.Atoi(ctx.Param("id"))

		ctx.JSON(200, service.GetById(id))

	})
	buku.DELETE("/:id", func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		service.Delete(id)
		ctx.JSON(200, gin.H{"message": "success"})
	})

	buku.POST("", func(ctx *gin.Context) {
		var request *CreateRequest
		err := ctx.BindJSON(&request)
		if err != nil {
			ctx.JSON(400, gin.H{"message": "error bree"})
		}

		service.CreateBuku(*request)
		ctx.JSON(200, gin.H{
			"message": "created",
		})

	})

	buku.PUT("/:id", func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))

		var request *CreateRequest
		err := ctx.BindJSON(&request)
		if err != nil {
			ctx.JSON(400, gin.H{"message": "error bree"})
		}

		service.Update(request, id)
	})
	buku.GET("/search", func(ctx *gin.Context) {
		barcode := ctx.Query("q")
		buku, err := service.GetBukuByBarcode(barcode)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(400, gin.H{"message": "bad request"})
			return
		}
		ctx.JSON(200, buku)
	})
	buku.PUT("/:id/register", func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		var request *CreateRequestBarcode
		err := ctx.BindJSON(&request)
		if err != nil {
			ctx.JSON(400, gin.H{"message": "error update buku & create bukuHub"})
			fmt.Println("errrrrrrroooorororo", err)
			return
		}
		status, err := service.CreateBarcode(*request, id)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(status, err)
			return
		}
		ctx.JSON(200, gin.H{"message": "created"})
	})

}
