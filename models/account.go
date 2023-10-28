package models

type Account struct {
	Id       string  `json:"id" bson:"_id,omitempty"`
	Email    string  `json:"email" bson:"email"`
	Password string  `json:"password" bson:"password"`
	GroupIds []int64 `json:"group_ids" bson:"group_ids"`
	Role     string  `json:"role" bson:"role"`
}
