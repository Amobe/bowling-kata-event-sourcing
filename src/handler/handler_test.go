package handler

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetGinHandler(t *testing.T) {
	handler := &fakeHandler{}
	handerFunc := GetGinHandler(handler)
	handerFunc(&gin.Context{})
	assert.True(t, handler.isHandled)
}

type fakeHandler struct {
	isHandled bool
}

func (h *fakeHandler) Handle(c Context) {
	h.isHandled = true
}
