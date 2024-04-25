package routers

import (
	"biFebriansyah/back/internal/handlers"
	"biFebriansyah/back/internal/middleware"
	"biFebriansyah/back/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func movie(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/movie")

	repo := repository.NewMovie(d)
	handler := handlers.NewMovie(repo)

	route.POST("/", middleware.AuthJwt("admin"), middleware.UploadFile, handler.PostData)
	route.PATCH("/", middleware.UploadFile, handler.PatchData)
	route.DELETE("/:id", handler.RemoveData)
	route.GET("/", middleware.AuthJwt("admin", "user"), handler.FetchData)
	route.GET("/all", handler.FetchAllData)
}
