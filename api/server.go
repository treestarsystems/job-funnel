package api

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// StartServer creates a new server instance
func StartServer() *gin.Engine {

	// This should not be hardcoded. It should be set in the environment but for some reason it is not working
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	// Pass routes to the router
	// RoutesRecipe(router)
	fmt.Printf("Starting server on port :%s\n", os.Getenv("PORT"))
	router.Run()
	return router
}
