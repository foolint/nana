package admin

import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.String(200, "pong login")
}
