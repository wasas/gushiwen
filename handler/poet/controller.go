package poet

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/trytwice/gushiwen/model"
	"github.com/trytwice/gushiwen/pkg/db"
)

func GetAllPoet(c *gin.Context) {
	params := c.Query("page")
	page, _ := strconv.Atoi(params)
	if page <= 0 || page >= 1000 {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("page should in [1, 999], your type page = %d", page),
		})
		return
	}
	db, err := db.OpenDB()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("open database failed: %s", err),
		})
		return
	}
	defer db.Close()
	startID := 10 * (page - 1)
	poets := model.Poets{}
	db.Where("id BETWEEN ? AND ?", startID, startID+10).Find(&poets)
	poetHeader := model.PoetHeader{}
	poetHeader.TotalPoets = "9989"
	poetHeader.TotalPages = "999"
	poetHeader.CurrentPage = params
	poetHeader.PageSize = "10"
	poetHeader.Poets = poets
	c.JSON(http.StatusOK, gin.H{
		"poets_info": poetHeader,
	})
}

func GetPoetByName(c *gin.Context) {
	name := c.Query("name")
	db, err := db.OpenDB()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("open database failed: %s", err),
		})
		return
	}
	defer db.Close()
	poet := model.Poet{}
	db.Where("name = ?", name).First(&poet)
	if poet == (model.Poet{}) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("poet %s not found", name),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"poet": poet,
	})
}

func GetPoetByID(c *gin.Context) {
	params := c.Query("id")
	id, _ := strconv.Atoi(params)
	if id < 1 || id > 9989 {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("id should in [1, 9989], your type id = %d", id),
		})
		return
	}
	db, err := db.OpenDB()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("open database failed: %s", err),
		})
		return
	}
	poet := model.Poet{}
	db.First(&poet, id)
	c.JSON(http.StatusOK, gin.H{
		"poet": poet,
	})
}
