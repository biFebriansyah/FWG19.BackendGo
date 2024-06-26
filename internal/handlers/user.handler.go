package handlers

import (
	"biFebriansyah/back/config"
	"biFebriansyah/back/internal/models"
	"biFebriansyah/back/internal/repository"
	"biFebriansyah/back/pkg"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	repository.RepoUserIF
}

func NewUser(r repository.RepoUserIF) *HandlerUser {
	return &HandlerUser{r}
}

func (h *HandlerUser) PostData(ctx *gin.Context) {
	var ers error
	data := models.User{
		Role: "user",
	}

	if ers = ctx.ShouldBind(&data); ers != nil {
		ctx.AbortWithError(http.StatusBadRequest, ers)
		return
	}

	_, ers = govalidator.ValidateStruct(&data)
	if ers != nil {
		pkg.NewRes(401, &config.Result{
			Data: ers.Error(),
		}).Send(ctx)
		return
	}

	data.Password, ers = pkg.HashPassword(data.Password)
	if ers != nil {
		pkg.NewRes(401, &config.Result{
			Data: ers.Error(),
		}).Send(ctx)
		return
	}

	result, ers := h.CreateUser(&data)
	if ers != nil {
		pkg.NewRes(401, &config.Result{
			Data: ers.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, result).Send(ctx)

}

func (h *HandlerUser) FetchAll(ctx *gin.Context) {
	data, err := h.GetAllUser()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pkg.NewRes(200, data).Send(ctx)
}
