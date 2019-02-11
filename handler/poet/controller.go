package poet

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/trytwice/gushiwen/model"
	"github.com/trytwice/gushiwen/pkg/db"
)

func GetAllPoet(c *gin.Context) {
	db, err := db.OpenDB()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("open database failed: %s", err),
		})
		return
	}
	defer db.Close()
	params := c.Query("page")
	page, _ := strconv.Atoi(params)
	totalPoets, _ := countPoet()
	totalPages := int(math.Ceil(totalPoets / 10.0))
	if page <= 0 || page > totalPages {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("page should in [1, %d], your type page = %d", totalPages, page),
		})
		return
	}
	startID := 10 * (page - 1)
	poets := model.Poets{}
	db.Where("id BETWEEN ? AND ?", startID, startID+10).Find(&poets)
	poetHeader := model.PoetHeader{}
	poetHeader.TotalPoets = strconv.Itoa(int(totalPoets))
	poetHeader.TotalPages = strconv.Itoa(totalPages)
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
	totalPoets, _ := countPoet()
	if id < 1 || id > int(totalPoets) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("id should in [1, %d], your type id = %d", int(totalPoets), id),
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
