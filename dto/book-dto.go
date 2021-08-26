package dto

//BookUpdateDTO is a model that client use when updating a book
type BookUpdateDTO struct {
	ID 			uint64 `json:"id" form:"id" binding:"required"`
	Title 		string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserId 		uint64 `json:"user_id,omitempty" form:"user_id,omitempty" binding:"required"`
}

//BookCreateDTO is a model that client use when create a new book
type BookCreateDTO struct {
	Title 		string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserId 		uint64 `json:"user_id,omitempty" form:"user_id,omitempty" binding:"required"`
}