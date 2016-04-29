package config

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler(c *gin.Context, code int, err error) {
	if mode := gin.Mode(); "release" != mode {
		c.JSON(http.StatusOK, gin.H{"data": err.Error(), "status":0, "code":code})
	} else {
		c.Status(http.StatusNotFound)
	}
}
