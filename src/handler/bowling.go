package handler

import (
	"fmt"
	"strconv"

	"github.com/amobe/bowling-kata-event-sourcing/src/service"
	"github.com/gin-gonic/gin"
)

type BowlingHandler interface {
	Handle(*gin.Context)
}

var _ BowlingHandler = &bowlingHandler{}

type bowlingHandler struct {
	bs service.Bowling
}

func NewBowlingHandler(bs service.Bowling) (BowlingHandler, error) {
	if bs == nil {
		return nil, fmt.Errorf("bowling service is nil")
	}
	return &bowlingHandler{
		bs: bs,
	}, nil
}

func (h *bowlingHandler) Handle(ctx *gin.Context) {
	hitParam := ctx.Query("hit")
	if len(hitParam) < 1 {
		fmt.Fprint(ctx.Writer, "Url param 'hit' is missing")
	}
	hit, err := strconv.ParseUint(hitParam, 10, 32)
	if err != nil {
		fmt.Fprint(ctx.Writer, fmt.Errorf("roll action: invalid hit number: %w", err))
		return
	}
	if err := h.bs.Throw("0", uint32(hit)); err != nil {
		fmt.Fprint(ctx.Writer, fmt.Errorf("roll action: handler error: %w", err))
		return
	}
}
