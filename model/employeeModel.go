package model

import "time"

type Employee struct {
	ID 			int 		`gorm:"primary_key"`
	Name		string 		`json:"name" form:"name"`
	Position	string 		`json:"position" form:"position"`
	HireDate	time.Time 	`json:"hire_date" form:"hire_date"`
	Phone		string 		`json:"phone" form:"phone"`
}