package poetry

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/trytwice/gushiwen/model"
	"github.com/trytwice/gushiwen/pkg/db"
)

func GetAllPoetry(c *gin.Context) {
	params := c.Query("page")
	page, _ := strconv.Atoi(params)
	if page <= 0 || page >= 996 {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("page should in [1, 995], your type page = %d", page),
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
	poetrys := model.Poetrys{}
	db.Where("id BETWEEN ? AND ?", startID, startID+10).Find(&poetrys)
	poetryHeader := model.PoetryHeader{}
	poetryHeader.TotalPoetrys = "9947"
	poetryHeader.TotalPages = "995"
	poetryHeader.CurrentPage = params
	poetryHeader.PageSize = "10"
	poetryHeader.Poetrys = poetrys
	c.JSON(http.StatusOK, gin.H{
		"poets_info": poetryHeader,
	})
}

func GetPoetryByAuthor(c *gin.Context) {
	author := c.Query("author")
	db, err := db.OpenDB()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("open database failed: %s", err),
		})
		return
	}
	defer db.Close()
	poetrys := model.Poetrys{}
	db.Where("author = ?", author).Find(&poetrys)
	if len(poetrys) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("can not find any poetry writed by %s", author),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"poetry": poetrys,
	})
}

func GetSamplePoetry(c *gin.Context) {
	id := rand.Intn(9947)
	poetry := model.Poetry{}
	db, err := db.OpenDB()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  fmt.Sprintf("open database failed: %s", err),
		})
		return
	}
	defer db.Close()
	db.First(&poetry, id)
	c.JSON(http.StatusOK, gin.H{
		"poetry": poetry,
	})
}
