package handlers

import (
	"biFebriansyah/back/internal/repository"
)

type HandlerAuth struct {
	*repository.RepoUser
}

func NewAuth(r *repository.RepoUser) *HandlerAuth {
	return &HandlerAuth{r}
}
