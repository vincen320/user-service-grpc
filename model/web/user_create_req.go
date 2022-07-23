package web

type CreateUserRequest struct {
	Name     string `validate:"required,min=6,max=20" json:"name,omitempty"`
	Username string `validate:"required,min=6,max=20" json:"username,omitempty"`
	Password string `validate:"required,min=6,max=20" json:"password,omitempty"`
	Email    string `validate:"required,email" json:"email,omitempty"`
}
