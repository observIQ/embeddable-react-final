package api

import (
	"github.com/gin-gonic/gin"

	"github.com/observiq/embeddable-react/ui"
)

func Start() {
	store := newStore()
	router := gin.Default()

	addRoutes(router, store)
	ui.AddRoutes(router)

	router.Run(":4000")
}
