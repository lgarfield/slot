package client

import(
	"github.com/gin-gonic/gin"
	"gopkg.in/gorp.v1"
	"net/http"
	//"os"
	"slot/config"
	"strings"
)

func Post(c *gin.Context, db *gorp.DbMap) {
	newMethod, err := pathAndMethod(c)
	if err != nil {
		config.ErrorHandler(c, http.StatusInternalServerError, err)
	}

	PostFuncs := map[string]interface{} {
		"userRegister": userRegister,
	}

	// exec request
	result, err := config.Call(PostFuncs, newMethod, c, db)

	if err != nil {
		config.ErrorHandler(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func Get(c *gin.Context, db *gorp.DbMap) {
	newMethod, err := pathAndMethod(c)
	if err != nil {
		config.ErrorHandler(c, http.StatusInternalServerError, err)
	}

	GetFuncs := map[string]interface{} {
		"userRegister": userRegister,
	}

	// exec request
	result, err := config.Call(GetFuncs, newMethod, c, db)

	if err != nil {
		config.ErrorHandler(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func pathAndMethod(c *gin.Context) (newMethod string, err error){
	// get path, method from the query parameters
	path := c.DefaultPostForm("path", "nil")
	method := c.DefaultPostForm("method", "nil") // shortcut for c.Request.URL.Query().Get("method")

	err = config.ValidRequest(path)

	newMethod = path + strings.Replace(method, "", strings.ToUpper(string(method[0])), 1)
	return
}
