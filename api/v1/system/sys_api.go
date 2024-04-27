package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SystemApi struct{}

func (s *SystemApi) CreateApi(c *gin.Context) {
	c.String(http.StatusOK, "CreateApi")
}

func (s *SystemApi) Panic(c *gin.Context) {
	panic("panic")
}
