package author

import (
	"database/sql"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name      string
	Email     sql.NullString
	Biography sql.NullString
}
