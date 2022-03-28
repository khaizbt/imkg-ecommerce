package model

type (
	User struct {
		ID       int    `json:"id" bson:"id"`
		Name     string `json:"name" bson:"name"`
		Email    string `json:"email" bson:"email"`
		Username string `json:"username" bson:"username"`
		//Phone    string `json:"phone" bson:"phone"`
		//UserType int    `json:"user_type" bson:"user_type"`
		Password string `json:"password" bson:"password"`
	}
)
