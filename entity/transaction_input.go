package entity

type (
	CartInput struct {
		UserID    int
		ProductID string `json:"product_id"`
		Qty       int    `json:"qty"`
		CartId    string `json:"cart_id"`
	}

	Checkout struct {
		CourierID      int    `json:"courier_id"`
		UserID         int    `json:"user_id"`
		PaymentType    string `json:"payment_type"`
		CheckoutDetail `json:"checkout_detail"`
	}

	CheckoutDetail struct {
		ProductID string `json:"product_id"`
		Qty       int    `json:"qty"`
	}
)
