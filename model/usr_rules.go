package model

type UsrRules struct {
	Id   int    `gorm:"type:int;primary_key;autoIncrement"`
	Path string `gorm:"type:varchar(255)"`
	Usr  string `gorm:"type:varchar(255)"`
}
