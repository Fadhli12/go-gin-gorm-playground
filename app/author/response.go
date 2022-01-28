package author

type AuthorResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Biography string `json:"biography"`
}
