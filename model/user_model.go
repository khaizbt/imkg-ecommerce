package model

import "time"

type (
	User struct {
		ID       int    `json:"id" bson:"id"`
		Name     string `json:"name" bson:"name"`
		Email    string `json:"email" bson:"email"`
		Username string `json:"username" bson:"username"`
		//Phone    string `json:"phone" bson:"phone"`
		IdUserType int    `json:"id_user_type" bson:"id_user_type"`
		Password   string `json:"password" bson:"password"`
	}

	UserType struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	Feature struct {
		ID        int    `json:"id"`
		Key       string `json:"key"`
		Value     string `json:"value"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	UserTypeFeature struct {
		ID         int       `json:"id"`
		IDUserType int       `json:"user_type_id"`
		IDFeature  int       `json:"feature_id"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		Feature    Feature   `gorm:"foreignKey:IDFeature"`
		UserType   UserType  `gorm:"foreignKey:IDUserType"`
	}
)
