package handlers

import (
	"biFebriansyah/back/internal/models"
	"biFebriansyah/back/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerMovie struct {
	*repository.RepoMovie
}

func NewMovie(r *repository.RepoMovie) *HandlerMovie {
	return &HandlerMovie{r}
}

func (h *HandlerMovie) GetMovie(ctx *gin.Context) {
	ctx.String(200, "hello worlds")
}

func (h *HandlerMovie) PostMovie(ctx *gin.Context) {
	var movie models.Movie

	if err := ctx.ShouldBind(&movie); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := h.CreateMovie(&movie)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, result)
}
