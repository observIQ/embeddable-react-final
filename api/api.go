package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func createTodo(ctx *gin.Context, s Store) {
	payload := &CreatePayload{}
	if err := ctx.BindJSON(payload); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, CreateResponse{
		Todo: s.Create(payload.Description),
	})
}

func deleteTodo(ctx *gin.Context, s Store) {
	id := ctx.Param("id")
	s.Delete(id)
}

func checkTodo(ctx *gin.Context, s Store) {
	id := ctx.Param("id")

	payload := &CheckPayload{}
	if err := ctx.BindJSON(payload); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	t, err := s.Check(id, payload.Completed)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, CheckResponse{
		Todo: t,
	})
}

func listTodos(ctx *gin.Context, s Store) {
	ctx.JSON(http.StatusOK, ListResponse{
		Todos: s.List(),
	})
}
