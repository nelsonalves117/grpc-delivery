package delivery

import "time"

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type OrderStatus byte

const (
	OrderStatus_PENDING OrderStatus = iota + 1
	OrderStatus_ACCEPTED
	OrderStatus_IN_TRANSIT
	OrderStatus_DELIVERED
	OrderStatus_CANCELED
)

type Delivery struct {
	OrderID          string      `json:"order_id"`
	CustomerID       string      `json:"customer_id"`
	RestaurantID     string      `json:"restaurant_id"`
	DriverID         string      `json:"driver_id"`
	DeliveryLocation Location    `json:"delivery_location"`
	ItemIDs          []string    `json:"item_ids"`
	OrderTime        time.Time   `json:"order_time"`
	Status           OrderStatus `json:"status"`
	DeliveryTime     time.Time   `json:"delivery_time"`
	TotalAmount      float32     `json:"total_amount"`
}
