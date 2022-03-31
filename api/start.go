package api

import "github.com/gin-gonic/gin"

func Start() {
	store := newStore()
	router := gin.Default()
	addRoutes(router, store)

	router.Run(":4000")
}
