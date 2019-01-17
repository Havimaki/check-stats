package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// regular web server
// func main() {
// 	http.HandleFunc("/", testHandler)
// 	http.ListenAndServe(":3000", nil)
// }

// func testHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello world!")
// }

// GIN
func main() {
	// set router as gin's default router
	router := gin.Default()

	// serve the frontend static files
	router.Use(static.Serve("/", static.LocalFile("../views", true)))

	// setup route group for apis
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	api.GET("/links", getLinksHandler)
	api.POST("add-link", addLinkHandler)

	router.Run(":3000")
}

func getLinksHandler(c *gin.Context) {
	fmt.Println(c)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "get link functionality not yet available",
	})
}

func addLinkHandler(c *gin.Context) {
	fmt.Println(c)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Add Link functionality not yet available",
	})
}
