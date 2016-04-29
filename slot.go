package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"slot/client"
	"slot/config"
)

func init() {
	//gin.SetMode(gin.ReleaseMode)
}

func main() {
	// Create a gin router with default middleware.
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	dbMap := config.InitDb()
	defer dbMap.Db.Close()

	// Delete any existings rows.
	//err := dbMap.TruncateTables()
	//if err != nil {
	//	panic(err) // TODO
	//}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusNotFound, "Oops...")
	})
	router.POST("/", func(c *gin.Context) {
		c.String(http.StatusNotFound, "Oops...")
	})

	router.GET("/client", func(c *gin.Context) {
		client.Get(c, dbMap)
	})
	router.POST("/client", func(c *gin.Context) {
		client.Post(c, dbMap)
	})
	router.PUT("/client", func(c *gin.Context) {})
	router.DELETE("/client", func(c *gin.Context) {})
	router.PATCH("/client", func(c *gin.Context) {})
	router.HEAD("/client", func(c *gin.Context) {})


	// By default it serves on :8080 unless a PORT environment variable was defined.
	router.Run(":80")
	// router.Run(":8000") for a hard coded port
}
