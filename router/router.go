package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trytwice/gushiwen/handler/poet"
	"github.com/trytwice/gushiwen/handler/poetry"
)

type Router struct {
}

func (r *Router) InitRouter(g *gin.Engine, handler ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery(), gin.Logger())
	rg := g.Group("/api")
	{
		poet.Register(rg)
		poetry.Register(rg)
	}
	g.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, fetchPath(g))
	})
	g.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return g
}

func fetchPath(g *gin.Engine) []string {
	routers := g.Routes()
	path := []string{}
	for _, router := range routers {
		path = append(path, router.Path)
	}
	return path
}
