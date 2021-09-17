package structs

import "github.com/jinzhu/gorm"

type Items struct {
	gorm.Model
	ItemID      string `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
}

type Orders struct {
	gorm.Model
	OrderId      string  `json:"order_id"`
	CustomerName string  `json:"customer_name"`
	OrderedAt    string  `json:"ordered_at"`
	Items        []Items `json:"items"`
}
