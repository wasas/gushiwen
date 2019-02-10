package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/trytwice/gushiwen/router"
)

func main() {
	routers := router.Router{}
	g := gin.New()
	routers.InitRouter(g)
	err := g.Run(":18080")
	if err != nil {
		log.Fatal(err)
	}
}
