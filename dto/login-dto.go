package dto

//LoginDTO is used when client post from /login url
type LoginDTO struct {
	Email 		string `json:"email" form:"email" binding:"required" validate:"email"`
	Password 	string `json:"password,omitempty" form:"password,omitempty" validate:"min:6"`
}
