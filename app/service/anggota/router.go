package anggota

import (
	"database/sql"
	klasifikasianggotahub "library/app/service/klasifikasi_anggota_hub"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup, db *sql.DB) {
	repo := NewRepository(db)
	repoKlasifikasiAnggotaHub := klasifikasianggotahub.NewRepository(db)
	service := NewService(repo, repoKlasifikasiAnggotaHub)

	anggota := g.Group("/anggota")
	anggota.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, service.Get())
	})

	anggota.GET("/:id", func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		ctx.JSON(200, service.GetById(id))
	})
	anggota.DELETE("/:id", func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		service.Delete(id)
		ctx.JSON(200, gin.H{"message": "success"})
	})

	anggota.POST("", func(ctx *gin.Context) {
		var request *CreateRequest
		err := ctx.BindJSON(&request)
		if err != nil {
			ctx.JSON(400, gin.H{"message": "error bree"})
		}
		status, massage := service.CreateAnggota(*request)
		ctx.JSON(status, gin.H{
			"message": massage,
		})
	})
	anggota.PUT("/:id", func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))

		var request *CreateRequest
		err := ctx.BindJSON(&request)
		if err != nil {
			ctx.JSON(400, gin.H{"message": "error bree"})
		}

		service.Update(request, id)
	})
	anggota.GET("/search", func(ctx *gin.Context) {
		noAnggota := ctx.Query("q")
		anggota, err := service.GetAnggotaByQuery(noAnggota)
		if err != nil {
			ctx.JSON(400, gin.H{"message": "bad request"})
			return
		}
		ctx.JSON(200, anggota)
	})

}
