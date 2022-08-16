package handler

import "github.com/gin-gonic/gin"

type Handler interface {
	Handle(Context)
}

func GetGinHandler(h Handler) func(*gin.Context) {
	return func(ctx *gin.Context) {
		h.Handle(NewHandlerContext(ctx))
	}
}
