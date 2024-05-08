package handlers

import (
	"biFebriansyah/back/config"
	"biFebriansyah/back/internal/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repoUserMock = repository.RepoMock{}
var reqBody = `{
	"user_id": "123",
	"username": "testing",
	"password": "abcd1234",
	"role": "user"
}`

func TestPostData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	w := httptest.NewRecorder()

	handler := NewUser(&repoUserMock)
	exptedResult := &config.Result{Message: "1 data user created"}
	repoUserMock.On("CreateUser", mock.Anything).Return(exptedResult, nil)

	r.POST("/create", handler.PostData)
	req := httptest.NewRequest("POST", "/create", strings.NewReader(reqBody))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"description": "1 data user created", "status": "OK"}`, w.Body.String())
}
