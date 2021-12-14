package handler

import (
	"io"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -destination mocks/mock_context.go -package mocks -source context.go
type Context interface {
	Writer() io.Writer
	Query(key string) string
}

var _ Context = &context{}

type context struct {
	ctx *gin.Context
}

func NewHandlerContext(ctx *gin.Context) Context {
	return &context{
		ctx: ctx,
	}
}

func (c *context) Writer() io.Writer {
	return c.ctx.Writer
}

func (c *context) Query(key string) string {
	return c.ctx.Query(key)
}
