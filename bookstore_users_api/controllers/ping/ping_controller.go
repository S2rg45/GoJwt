package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	description := "pong"
	c.String(http.StatusOK, description)
}
