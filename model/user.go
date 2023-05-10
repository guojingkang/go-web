package model

import "time"

type User struct {
	Id           int        `json:"id,omitempty"`
	Account      string     `json:"account,omitempty"`
	Name         string     `json:"name"`
	Password     string     `json:"password,omitempty"`
	Phone        string     `json:"phone"`
	Detail       string     `json:"detail,omitempty"`
	CreateTime   time.Time  `json:"createTime,omitempty" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdateTime   time.Time  `json:"updateTime,omitempty" gorm:"default:CURRENT_TIMESTAMP()"`
	CreateId     int        `json:"creatorId,omitempty"`
	UpdateId     int        `json:"updateId,omitempty"`
	DetailStruct UserDetail `json:"detailStruct" gorm:"-"`
}

type UserDetail struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
