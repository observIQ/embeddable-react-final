package api

import "github.com/gin-gonic/gin"

func addRoutes(r *gin.Engine, store Store) {
	api := r.Group("/api")
	api.GET("/todos", func(ctx *gin.Context) { listTodos(ctx, store) })
	api.POST("/todos", func(ctx *gin.Context) { createTodo(ctx, store) })
	api.DELETE("/todos/:id", func(ctx *gin.Context) { deleteTodo(ctx, store) })
	api.PUT("/todos/:id", func(ctx *gin.Context) { checkTodo(ctx, store) })
}
