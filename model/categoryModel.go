package model

type Category struct {
	ID 		int 	`gorm:"primary_key"`
	Name	string 	`json:"name" form:"name"`
	Items	[]*Item	`gorm:"ForeignKey:CategoryId"`
}