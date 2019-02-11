package poetry

import "github.com/gin-gonic/gin"

func Register(r *gin.RouterGroup) {
	r.GET("/poetries/all", GetAllPoetry)
	r.GET("/poetries/author", GetPoetryByAuthor)
	r.GET("/poetries/sample", GetSamplePoetry)
}
