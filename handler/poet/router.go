package poet

import "github.com/gin-gonic/gin"

func Register(r *gin.RouterGroup) {
	r.GET("/poets/all", GetAllPoet)
	r.GET("/poets/name", GetPoetByName)
	r.GET("/poets/id", GetPoetByID)
}
