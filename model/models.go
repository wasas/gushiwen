package model

import "github.com/jinzhu/gorm"

type Poetry struct {
	gorm.Model
	Title   string `gorm:"type:varchar(100)"`
	Content string `gorm:"type:varchar(3000)"`
	Author  string `gorm:"type:varchar(50)"`
	Dynasty string `gorm:"type:varchar(50)"`
	PoetURL string `gorm:"type:varchar(100)"`
	Liked   int    `gorm:"type:int;default(0)"`
}

type Poetrys []Poetry

type Poet struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100)" json:"name"`
	Description string `gorm:"type:varchar(3000)" json:"description"`
	ImageURL    string `gorm:"type:varchar(100)" json:"image_url"`
	TotalPoetry int    `gorm:"type:int;default(0)" json:"total_poetry"`
	Liked       int    `gorm:"type:int;default(0)" json:"liked"`
}

type Poets []Poet

type PoetHeader struct {
	TotalPoets  string `json:"total_poets"`
	TotalPages  string `json:"total_pages"`
	CurrentPage string `json:"current_page"`
	PageSize    string `json:"page_size"`
	Poets       `json:"poets"`
}

type PoetryHeader struct {
	TotalPoetrys string `json:"total_poetrys"`
	TotalPages   string `json:"total_pages"`
	CurrentPage  string `json:"current_page"`
	PageSize     string `json:"page_size"`
	Poetrys      `json:"poetrys"`
}
