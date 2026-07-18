package petugas

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.RouterGroup, db *sql.DB) {
	repo := NewRepository(db)

	service := NewService(repo)

	petugas := g.Group("/petugas")
	petugas.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, service.Get())
	})

	petugas.GET("/:id", func(ctx *gin.Context) {

		id, _ := strconv.Atoi(ctx.Param("id"))

		ctx.JSON(200, service.GetById(id))

	})
	petugas.DELETE("/:id", func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))
		service.Delete(id)
		ctx.JSON(200, gin.H{"message": "success"})
	})

	petugas.POST("", func(ctx *gin.Context) {
		var request *CreateRequest
		err := ctx.BindJSON(&request)
		if err != nil {
			ctx.JSON(400, gin.H{"message": "error bree"})
		}

		service.CreatePetugas(*request)
		ctx.JSON(200, gin.H{
			"message": "created",
		})

	})

	petugas.PUT("/:id", func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))

		var request *CreateRequest
		err := ctx.BindJSON(&request)
		if err != nil {
			ctx.JSON(400, gin.H{"message": "error bree"})
		}
		
		service.Update(request,id)
	})

}
