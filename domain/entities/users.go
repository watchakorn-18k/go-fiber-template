package entities


type UserFormat struct {
	UserID   string `json:"user_id" bson:"user_id,omitempty"`
	Username string `json:"username" bson:"username,omitempty"`
	Email    string `json:"email" bson:"email,omitempty"`
}
