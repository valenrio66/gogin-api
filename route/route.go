package route

import (
	"github.com/gin-gonic/gin"
	"github.com/valenrio66/gogin-api/controller"
	"github.com/valenrio66/gogin-api/repository" // Import the repository package
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Create an instance of the repository
	albumRepo := repository.NewAlbumDb(db)

	// Create an instance of the Service struct
	albumService := controller.NewService(albumRepo)

	// Routes for handling albums
	albumRoutes := r.Group("/albums")
	{
		albumRoutes.GET("/", albumService.GetAllAlbum)
		albumRoutes.GET("/:id_album", albumService.GetAlbum)
		albumRoutes.POST("/", albumService.CreateAlbum)
		albumRoutes.PUT("/:id_album", albumService.UpdateAlbum)
		albumRoutes.DELETE("/:id_album", albumService.DeleteAlbum)
	}

	return r
}
