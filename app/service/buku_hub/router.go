package bukuhub

// import (
// 	"database/sql"
// 	"library/app/service/buku"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// func Router(g *gin.RouterGroup, db *sql.DB) {
// 	repo := NewRepository(db)
// 	interfaceBuku := buku.Buku.GetById()
// 	service := NewService(repo, interfaceBuku)

// 	hargaBuku := g.Group("/buku-hub")
// 	hargaBuku.GET("", func(ctx *gin.Context) {
// 		ctx.JSON(200, service.Get())
// 	})

// 	hargaBuku.GET("/:id", func(ctx *gin.Context) {

// 		id, _ := strconv.Atoi(ctx.Param("id"))

// 		ctx.JSON(200, service.GetById(id))

// 	})
// 	hargaBuku.DELETE("/:id", func(ctx *gin.Context) {
// 		id, _ := strconv.Atoi(ctx.Param("id"))
// 		service.Delete(id)
// 		ctx.JSON(200, gin.H{"message": "success"})
// 	})

// 	// hargaBuku.POST("", func(ctx *gin.Context) {
// 	// 	var id int
// 	// 	var request *CreateRequest
// 	// 	err := ctx.BindJSON(&request)
// 	// 	if err != nil {
// 	// 		ctx.JSON(400, gin.H{"message": "error bree"})
// 	// 	}

// 	// 	status, massage := service.CreateBukuHub(*request,id)
// 	// 	ctx.JSON(status, gin.H{
// 	// 		"message": massage,
// 	// 	})

// 	// })

// 	hargaBuku.PUT("/:id", func(ctx *gin.Context) {
// 		id, _ := strconv.Atoi(ctx.Param("id"))

// 		var request *CreateRequest
// 		err := ctx.BindJSON(&request)
// 		if err != nil {
// 			ctx.JSON(400, gin.H{"message": "error bree"})
// 		}

// 		service.Update(request, id)
// 	})

// }
