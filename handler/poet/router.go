package poet

import "github.com/gin-gonic/gin"

func Register(r *gin.RouterGroup) {
	r.GET("/poet/all", GetAllPoet)
	r.GET("/poet/name", GetPoetByName)
	r.GET("/poet/id", GetPoetByID)
}
