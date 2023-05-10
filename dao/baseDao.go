package dao

import "gorm.io/gorm"

type PageQueryInput struct {
	Offset int 
	Limit  int
	Keyword    string
}

type Dao struct {
	Db *gorm.DB
}
