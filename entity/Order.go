package entity

import "time"

type Order struct {
	Id           string
	CustomerId   string
	DeliveryDate time.Time
	Items        []MenuItem
	State        string
	PaymentId    string
	CouponCode   string
	Date         time.Time
}