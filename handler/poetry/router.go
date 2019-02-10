package poetry

import "github.com/gin-gonic/gin"

func Register(r *gin.RouterGroup) {
	r.GET("/poetry/all", GetAllPoetry)
	r.GET("/poetry/author", GetPoetryByAuthor)
	r.GET("/poetry/sample", GetSamplePoetry)
}
