package spider

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var (
	baseURLPoet   = "https://so.gushiwen.org/authors/Default.aspx?p=%d"
	baseURLPoetry = "https://www.gushiwen.org/shiwen/default.aspx?page=%d&type=0&id=0"
)

func TestGetPoetry(t *testing.T) {
	for i := 1; i < 50; i++ {
		url := fmt.Sprintf(baseURLPoetry, i)
		err := GetPoetry(url)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestGetPoet(t *testing.T) {
	for i := 1; i < 50; i++ {
		url := fmt.Sprintf(baseURLPoet, i)
		err := GetPoet(url)
		if err != nil {
			t.Error(err)
		}
	}
}
