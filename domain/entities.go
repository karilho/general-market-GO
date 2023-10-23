package domain

import "time"

type UserType string

type User struct {
	ID    int    `ksql:"id"`
	Name  string `ksql:"name"`
	Email string `ksql:"email"`
}

type UserData struct {
	UserDataID       int       `json:"user_data_id"`
	CurrentType      UserType  `json:"current_type"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	PasswordHash     string    `json:"password_hash"`
	FullName         string    `json:"full_name"`
	PhoneNumber      string    `json:"phone_number"`
	RegistrationDate time.Time `json:"registration_date"`
	StreetAddress    string    `json:"street_address"`
	PlaceNumber      string    `json:"place_number"`
	City             string    `json:"city"`
	StateProvince    string    `json:"state_province"`
	PostalCode       string    `json:"postal_code"`
}

type Buyer struct {
	BuyerID      int  `json:"buyer_id"`
	UserDataID   int  `json:"user_data_id"`
	HasPurchased bool `json:"has_purchased"`
}

type Product struct {
	ProductID    int     `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductValue float64 `json:"product_value"`
	Description  string  `json:"description"`
}

type BuyOrder struct {
	BuyOrderID    string    `json:"buy_order_id"`
	BuyerID       int       `json:"buyer_id"`
	OrderDate     time.Time `json:"order_date"`
	TotalValue    float64   `json:"total_value"`
	PaymentMethod string    `json:"payment_method"`
}

type PurchasedProduct struct {
	PurchasedProductID int     `json:"purchased_product_id"`
	BuyOrderID         int     `json:"buy_order_id"`
	ProductID          int     `json:"product_id"`
	Quantity           int     `json:"quantity"`
	ValuePerUnit       float64 `json:"value_per_unit"`
}
