package domain

import "time"

type UserType string

type User struct {
	ID    int    `ksql:"id"`
	Name  string `ksql:"name"`
	Email string `ksql:"email"`
}

type UserData struct {
	UserDataID       int       `ksql:"user_data_id"`
	CurrentType      string    `ksql:"current_type"`
	Username         string    `ksql:"username"`
	Email            string    `ksql:"email"`
	PasswordHash     string    `ksql:"password_hash"`
	FullName         string    `ksql:"full_name"`
	PhoneNumber      string    `ksql:"phone_number"`
	RegistrationDate time.Time `ksql:"registration_date,timeNowUTC/skipUpdates"`
	StreetAddress    string    `ksql:"street_address"`
	PlaceNumber      string    `ksql:"place_number"`
	City             string    `ksql:"city"`
	StateProvince    string    `ksql:"state_province"`
	PostalCode       string    `ksql:"postal_code"`
}

type Buyers struct {
	BuyerID      int  `ksql:"buyer_id"`
	UserDataID   int  `ksql:"user_data_id"`
	HasPurchased bool `ksql:"has_purchased"`
}

type Product struct {
	ProductID    int     `ksql:"product_id"`
	ProductName  string  `ksql:"product_name"`
	ProductValue float64 `ksql:"product_value"`
	Description  string  `ksql:"description"`
}

type BuyOrder struct {
	BuyOrderID    string    `ksql:"buy_order_id"`
	BuyerID       int       `ksql:"buyer_id"`
	OrderDate     time.Time `ksql:"order_date"`
	TotalValue    float64   `ksql:"total_value"`
	PaymentMethod string    `ksql:"payment_method"`
}

type PurchasedProduct struct {
	PurchasedProductID int     `ksql:"purchased_product_id"`
	BuyOrderID         int     `ksql:"buy_order_id"`
	ProductID          int     `ksql:"product_id"`
	Quantity           int     `ksql:"quantity"`
	ValuePerUnit       float64 `ksql:"value_per_unit"`
}
