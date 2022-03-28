package entity

type (
	User struct {
		ID    int    `json:"id" bson:"id"`
		Name  string `json:"name" bson:"name"`
		Email string `json:"email" bson:"email"`
		Phone string `json:"phone" bson:"phone"`
	}
)
