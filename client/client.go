package client

import(
	"github.com/gin-gonic/gin"
	"gopkg.in/gorp.v1"
	"net/http"
	//"os"
	"slot/config"
	"strings"
)

var(
	postMap = map[string]interface{}{
		"userregister": userRegister,
	}
	getMap = map[string]interface{}{
		"userregister": userRegister,
	}
)

func Post(c *gin.Context, db *gorp.DbMap) {
	newMethod, err := pathAndMethod(c)
	if err != nil {
		config.ErrorHandler(c, http.StatusInternalServerError, err)
	}

	// exec request
	result, err := config.Call(postMap[newMethod], c, db)
	if err != nil {
		config.ErrorHandler(c, http.StatusInternalServerError, err)
	}

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"data":result, "status":1, "code":0})
	}
}

func Get(c *gin.Context, db *gorp.DbMap) {
	newMethod, err := pathAndMethod(c)
	if err != nil {
		config.ErrorHandler(c, http.StatusInternalServerError, err)
	}

	// exec request
	result, err := config.Call(getMap[newMethod], c, db)

	if err != nil {
		config.ErrorHandler(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{"data":result, "status":1, "code":0})
}

func pathAndMethod(c *gin.Context) (newMethod string, err error){
	// get path, method from the query parameters
	path := c.DefaultPostForm("path", "nil")
	method := c.DefaultPostForm("method", "nil") // shortcut for c.Request.URL.Query().Get("method")

	err = config.ValidRequest("client", path)
	if err != nil {
		return
	}

	newMethod = strings.ToLower(path) + strings.ToLower(method)
	return
}
