package author

type AuthorPost struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Biography string `json:"biography" binding:""`
}

type AuthorUpdate struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Biography string `json:"biography" binding:""`
}
