package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/trytwice/gushiwen/model"
)

type dbInfo struct {
	User   string `json:"user"`
	Passwd string `json:"passwd"`
	Table  string `json:"table"`
}

var (
	user   string
	passwd string
	table  string
)

func CreatDB() error {
	db, err := OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()
	db.AutoMigrate(&model.Poet{})
	db.AutoMigrate(&model.Poetry{})
	return nil
}

func getDB() (string, string, string, error) {
	gopath := os.Getenv("GOPATH")
	file, err := ioutil.ReadFile(gopath + "/src/github.com/trytwice/gushiwen/config/conf.json")
	if err != nil {
		return "", "", "", err
	}
	info := dbInfo{}
	err = json.Unmarshal(file, &info)
	if err != nil {
		return "", "", "", err
	}
	return info.User, info.Passwd, info.Table, nil
}

func OpenDB() (*gorm.DB, error) {
	user, passwd, table, err := getDB()
	if err != nil {
		panic(err)
	}
	return gorm.Open("mysql", fmt.Sprintf(`%s:%s@/%s?parseTime=True`, user, passwd, table))
}
