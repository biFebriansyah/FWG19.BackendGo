package routers

import (
	"biFebriansyah/back/internal/handlers"
	"biFebriansyah/back/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func auth(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/auth")

	repo := repository.NewUser(d)
	handler := handlers.NewAuth(repo)

	route.POST("/", handler.Login)
}
