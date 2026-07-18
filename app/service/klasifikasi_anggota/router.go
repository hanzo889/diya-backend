package klasifikasianggota

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup, db *sql.DB) {
	repo := NewRepository(db)

	service := NewService(repo)

	klasifikasiAnggota := g.Group("/klasifikasi-anggota")
	klasifikasiAnggota.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, service.Get())
	})

	klasifikasiAnggota.GET("/:id", func(ctx *gin.Context) {

		id, _ := strconv.Atoi(ctx.Param("id"))

		ctx.JSON(200, service.GetById(id))

	})
	klasifikasiAnggota.DELETE("/:id", func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		service.Delete(id)
		ctx.JSON(200, gin.H{"message": "success"})
	})

	klasifikasiAnggota.POST("", func(ctx *gin.Context) {
		var request *CreateRequest
		err := ctx.BindJSON(&request)
		if err != nil {
			ctx.JSON(400, gin.H{"message": "error bree"})
		}

		service.CreateKlasifikasiAnggota(*request)
		ctx.JSON(200, gin.H{
			"message": "created",
		})

	})

	klasifikasiAnggota.PUT("/:id", func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))

		var request *CreateRequest
		err := ctx.BindJSON(&request)
		if err != nil {
			ctx.JSON(400, gin.H{"message": "error bree"})
		}

		service.Update(request, id)
	})

}
