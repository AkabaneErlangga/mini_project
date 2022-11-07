package model

type OrderItems struct {
	ID 			int 	`gorm:"primary_key"`
	OrderId		int	 	`json:"orderId" form:"order_id"`
	ItemId		int 	`json:"itemId" form:"item_id"`
	Quantity	int 	`json:"quantity" form:"quantity"`
	SubPrice	int		`json:"subPrice" form:"sub_price"`
}