package domain

type User struct {
	Id        string `json:"id,omitempty" bson:"_id,omitempty"`
	Username  string `json:"username,omitempty" bson:"username,omitempty"`
	Password  string `json:"password,omitempty" bson:"password,omitempty"`
	Email     string `json:"email,omitempty" bson:"email,omitempty"`
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
