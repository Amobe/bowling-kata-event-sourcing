package main

import (
	"log"

	"github.com/amobe/bowling-kata-event-sourcing/src/handler"
	"github.com/amobe/bowling-kata-event-sourcing/src/service"
	"github.com/amobe/bowling-kata-event-sourcing/src/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	r := storage.NewInmemEventStorage()
	bs, err := service.NewBowlingService(r)
	if err != nil {
		log.Fatal(err)
	}
	bh, err := handler.NewBowlingHandler(bs)
	if err != nil {
		log.Fatal(err)
	}

	g := gin.Default()
	g.POST("/roll", handler.GetGinHandler(bh))
	log.Fatal(g.Run(":80"))
}
