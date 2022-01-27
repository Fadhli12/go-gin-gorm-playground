package author

import "database/sql"

type AuthorResponse struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Email     sql.NullString `json:"email"`
	Biography sql.NullString `json:"biography"`
}
