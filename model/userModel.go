package model

type User struct {
	ID			int 	`gorm:"primary_key"`
	Username	string 	`json:"username" form:"username"`
	Password	string 	`json:"password" form:"password"`
	EmployeeId	int 	`gorm:"ForeignKey:id" json:"employeeId" form:"employee_id"`
	Employee Employee
}