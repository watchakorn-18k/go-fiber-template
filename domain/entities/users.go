package entities

type NewUserBody struct {
	UserID   string `json:"user_id" bson:"user_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
}

type UserDataFormat struct {
	UserID   string `json:"user_id" bson:"user_id,omitempty"`
	Username string `json:"username" bson:"username,omitempty"`
	Email    string `json:"email" bson:"email,omitempty"`
}
