package model

type Item struct {

	ID 				int 	`gorm:"primary_key"`
	Brand     		string 	`json:"brand" form:"brand"`
	Type     		string 	`json:"type" form:"type"`
	LaunchYear		string 	`json:"launch_year" form:"launch_year"`
	CategoryId 		int 	`json:"category_id" form:"categoryId"`
	Description    	string 	`json:"description" form:"description"`
	Stock			int 	`json:"stock" form:"stock"`
	Price 			int 	`json:"price" form:"price"`

}