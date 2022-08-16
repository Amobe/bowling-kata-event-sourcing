package main

import (
	"log"

	handler2 "github.com/amobe/bowling-kata-event-sourcing/src/v0/handler"
	"github.com/amobe/bowling-kata-event-sourcing/src/v0/service"
	"github.com/amobe/bowling-kata-event-sourcing/src/v0/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	r := storage.NewInmemEventStorage()
	bs, err := service.NewBowlingService(r)
	if err != nil {
		log.Fatal(err)
	}
	bh, err := handler2.NewBowlingHandler(bs)
	if err != nil {
		log.Fatal(err)
	}

	g := gin.Default()
	g.POST("/roll", handler2.GetGinHandler(bh))
	log.Fatal(g.Run(":80"))
}
