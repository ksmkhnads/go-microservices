package models

type Customer struct {
	CustomerID string `gorm:"column:customer_id;primaryKey"`
	FirstName  string `gorm:"column:firstName"`
	LastName   string `gorm:"column:lastName"`
	Email      string `gorm:"column:email"`
	Phone      string `gorm:"column:phone"`
	Address    string `gorm:"column:address"`
}
