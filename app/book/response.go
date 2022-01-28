package book

type BookResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
	Discount    int    `json:"discount"`
	AuthorID    uint   `json:"author_id"`
	Author      AuthorResponse
}

type AuthorResponse struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Biography string `json:"biography"`
}
