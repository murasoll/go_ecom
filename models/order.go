package models

import "time"

type Order struct {
	ID              uint            `json:"id" gorm:"primaryKey"`
	UserID          uint            `json:"user_id" gorm:"not null"`
	User            User            `json:"user" gorm:"foreignKey:UserID"`
	Items           []OrderItem     `json:"items" gorm:"foreignKey:OrderID"`
	Status          string          `json:"status" gorm:"not null"`
	Total           float64         `json:"total" gorm:"not null"`
	ShippingAddress ShippingAddress `json:"shipping_address" gorm:"embedded"`
	ShippingCost    float64         `json:"shipping_cost" gorm:"not null"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type OrderItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	OrderID   uint    `json:"order_id" gorm:"not null"`
	ProductID uint    `json:"product_id" gorm:"not null"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity" gorm:"not null"`
	Price     float64 `json:"price" gorm:"not null"`
}
