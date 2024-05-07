package model

type ProcessRules struct {
	Id      int    `gorm:"type:int;primary_key;autoIncrement"`
	Path    string `gorm:"type:varchar(255)"`
	Process string `gorm:"type:varchar(255)"`
}
