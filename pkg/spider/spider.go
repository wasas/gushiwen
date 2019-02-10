package spider

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	"github.com/trytwice/gushiwen/pkg/db"
)

type Poetry struct {
	gorm.Model
	Title   string `gorm:"type:varchar(100)"`
	Content string `gorm:"type:varchar(3000)"`
	Author  string `gorm:"type:varchar(50)"`
	Dynasty string `gorm:"type:varchar(50)"`
	PoetURL string `gorm:"type:varchar(100)"`
	Liked   int    `gorm:"type:int;default(0)"`
}

type Poet struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:varchar(3000)"`
	ImageURL    string `gorm:"type:varchar(100)"`
	TotalPoetry int    `gorm:"type:int;default(0)"`
	Liked       int    `gorm:"type:int;default(0)"`
}

func GetPoetry(url string) error {
	resp, err := getHttpResponse(url, false)
	if err != nil {
		return err
	}
	db, err := db.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(resp)))
	if err != nil {
		return nil
	}
	doc.Find("div.left div.sons").Each(func(i int, s *goquery.Selection) {
		title := s.Find("p a b").Text()
		content := strings.TrimSpace(s.Find(".contson").Text())
		author := s.Find(".source a").Eq(1).Text()
		dynasty := s.Find(".source a").Eq(0).Text()
		poetryURL, _ := s.Find("p a").Eq(0).Attr("href")
		liked, _ := strconv.Atoi(strings.TrimSpace(s.Find(".good a span").Text()))

		poetry := Poetry{
			Title:   title,
			Content: content,
			Author:  author,
			Dynasty: dynasty,
			PoetURL: poetryURL,
			Liked:   liked,
		}
		db.Create(&poetry)
	})
	return nil
}

func GetPoet(url string) error {
	resp, err := getHttpResponse(url, false)
	if err != nil {
		return err
	}
	db, err := db.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(resp)))
	doc.Find(".left .sonspic").Each(func(i int, s *goquery.Selection) {
		name := s.Find("p a b").Text()
		description := strings.TrimSpace(s.Find("p").Eq(1).Text())
		imageURL, _ := s.Find(".divimg a img").Attr("src")
		liked, _ := strconv.Atoi(strings.TrimSpace(s.Find(".good a span").Text()))
		totalPoetry, _ := strconv.Atoi(strings.TrimPrefix(strings.TrimSuffix(s.Find("p a").Eq(2).Text(), "篇诗文"), "► "))

		poet := Poet{
			Name:        name,
			Description: description,
			ImageURL:    imageURL,
			Liked:       liked,
			TotalPoetry: totalPoetry,
		}
		db.Create(&poet)
	})
	return nil
}
