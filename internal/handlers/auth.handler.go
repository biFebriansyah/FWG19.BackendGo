package handlers

import (
	"biFebriansyah/back/config"
	"biFebriansyah/back/internal/repository"
	"biFebriansyah/back/pkg"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `db:"username" json:"username" form:"username"`
	Password string `db:"password" json:"password,omitempty"`
}

type HandlerAuth struct {
	*repository.RepoUser
}

func NewAuth(r *repository.RepoUser) *HandlerAuth {
	return &HandlerAuth{r}
}

func (h *HandlerAuth) Login(ctx *gin.Context) {
	var data User

	if err := ctx.ShouldBind(&data); err != nil {
		pkg.NewRes(500, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	users, err := h.GetAuthData(data.Username)
	if err != nil {
		pkg.NewRes(401, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	if err := pkg.VerifyPassword(users.Password, data.Password); err != nil {
		pkg.NewRes(401, &config.Result{
			Data: "Password salah",
		}).Send(ctx)
		return
	}

	jwtt := pkg.NewToken(users.User_id, users.Role)
	tokens, err := jwtt.Genrate()
	if err != nil {
		pkg.NewRes(500, &config.Result{
			Data: err.Error(),
		}).Send(ctx)
		return
	}

	pkg.NewRes(200, &config.Result{Data: tokens}).Send(ctx)
}
