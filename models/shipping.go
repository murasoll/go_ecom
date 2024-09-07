package models

type ShippingAddress struct {
	Street string `json:"street" gorm:"not null"`
	City   string `json:"city" gorm:"not null"`
	CityID uint   `json:"city_id" gorm:"not null"`
}

type City struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	Name         string  `json:"name" gorm:"not null;unique"`
	ShippingCost float64 `json:"shipping_cost" gorm:"not null"`
}
