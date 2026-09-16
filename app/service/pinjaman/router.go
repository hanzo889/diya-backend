package pinjaman

import (
	"database/sql"
	"library/app/service/anggota"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup, db *sql.DB) {
	repo := NewRepository(db)
	repoAnggota := anggota.NewRepository(db)
	service := NewService(repo, repoAnggota)

	pinjaman := g.Group("/pinjaman")

	pinjaman.GET("/:no_anggota", func(ctx *gin.Context) {

		noAnggota := ctx.Param("no_anggota")

		ctx.JSON(200, service.GetPinjamanByNoAnggota(noAnggota))

	})

	// pinjaman.GET("", func(ctx *gin.Context) {
	// 	ctx.JSON(200, service.Get())
	// })

	// pinjaman.DELETE("/:id", func(ctx *gin.Context) {
	// 	id, _ := strconv.Atoi(ctx.Param("id"))
	// 	service.Delete(id)
	// 	ctx.JSON(200, gin.H{"message": "success"})
	// })

	// pinjaman.POST("", func(ctx *gin.Context) {
	// 	var request *CreateRequest
	// 	err := ctx.BindJSON(&request)
	// 	if err != nil {
	// 		ctx.JSON(400, gin.H{"message": "error bree"})
	// 	}

	// 	service.CreatePinjaman(*request)
	// 	ctx.JSON(200, gin.H{
	// 		"message": "created",
	// 	})

	// })

	// pinjaman.PUT("/:id", func(ctx *gin.Context) {
	// 	id, _ := strconv.Atoi(ctx.Param("id"))

	// 	var request *CreateRequest
	// 	err := ctx.BindJSON(&request)
	// 	if err != nil {
	// 		ctx.JSON(400, gin.H{"message": "error bree"})
	// 	}

	// 	service.Update(request, id)
	// })

}
