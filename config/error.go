package config

import(
	"github.com/gin-gonic/gin"
	"os"
)

func ErrorHandler(c *gin.Context, code int, err error) {
	if env := os.Getenv("ENV"); "development" == env {
		c.JSON(code, gin.H{"err": err.Error()})
	} else {
		c.Status(code)
	}
}
