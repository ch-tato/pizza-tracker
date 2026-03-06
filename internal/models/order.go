package models

import (
	"time"

	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

var (
	OrderStatuses = []string{
		"Order Placed",
		"Preparing",
		"Baking",
		"Out for Delivery",
		"Delivered",
	}

	PizzaTypes = []string{
		"Margherita",
		"Pepperoni",
		"BBQ Chicken",
		"Meat Lovers",
		"Supreme",
		"Four Cheese",
		"Hawaiian",
		"Buffalo Chicken",
		"Truffle Mushroom",
	}

	PizzaSizes = []string{
		"Small",
		"Medium",
		"Large",
	}

	PizzaCrusts = []string{
		"Thin Crust",
		"Thick Crust",
		"Stuffed Crust",
		"Gluten-Free Crust",
	}

	PizzaAddOns = []string{
		"Extra Cheese",
		"Olives",
		"Mushrooms",
		"Peppers",
		"Onions",
	}
)

type OrderModel struct {
	DB *gorm.DB
}

type OrderItem struct {
	ID           string `gorm:"primaryKey;size:14" json:"id"`
	OrderID      string `gorm:"not null" json:"orderId"`
	PizzaType    string `gorm:"not null" json:"pizzaType"`
	PizzaSize    string `gorm:"not null" json:"pizzaSize"`
	PizzaCrust   string `gorm:"not null" json:"pizzaCrust"`
	AddOns       string `gorm:"type:text" json:"addOns"`
	Instructions string `gorm:"type:text" json:"instructions"`
}

type Order struct {
	ID           string      `gorm:"primaryKey;size:14" json:"id"`
	Status       string      `gorm:"not null" json:"status"`
	CustomerName string      `gorm:"not null" json:"customerName"`
	Phone        string      `gorm:"not null" json:"phone"`
	Address      string      `gorm:"not null" json:"address"`
	Items        []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt    time.Time   `gorm:"autoCreateTime" json:"createdAt"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	if o.ID == "" {
		o.ID = shortid.MustGenerate()
	}
	return nil
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) error {
	if oi.ID == "" {
		oi.ID = shortid.MustGenerate()
	}
	return nil
}

func (o *OrderModel) CreateOrder(order *Order) error {
	return o.DB.Create(order).Error
}

func (o *OrderModel) GetOrderByID(id string) (*Order, error) {
	var order Order
	err := o.DB.Preload("Items").First(&order, "id = ?", id).Error
	return &order, err
}
