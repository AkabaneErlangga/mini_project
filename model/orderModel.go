package model

import "time"

type Order struct {
	ID 				int 		`gorm:"primary_key"`
	TransactionId	string	 	`json:"transactionId" form:"transaction_id"`
	UserId			int	 		`json:"userId" form:"user_id"`
	Date			time.Time 	`json:"date" form:"date"`
	TotalPrice		int 		`json:"totalPrice" form:"total_price"`
	OrderItems		[]*OrderItems	`json:"orderItems" form:"order_items"`
}