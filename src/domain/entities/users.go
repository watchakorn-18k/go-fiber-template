package entities

type NewUserBody struct {
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
}

type UserDataFormat struct {
	UserID   string `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
}
