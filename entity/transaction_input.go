package entity

type (
	CartInput struct {
		UserID    int
		ProductID string `json:"product_id"`
		Qty       int    `json:"qty"`
	}
)
