package entity

type (
	LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)