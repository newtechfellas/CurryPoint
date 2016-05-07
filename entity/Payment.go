package entity

var Payment struct {
	Id          string
	FirstName   string
	LastName    string
	Number      string
	ExpiryMonth int
	ExpiryYear  int
	Code        int
	Zip         string
	OrderId     string
}