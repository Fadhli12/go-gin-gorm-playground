package auth

import (
	"github.com/Fadhli12/go-gin-gorm-playground/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type LoginService interface {
	LoginUser(email string, password string) bool
}

type auth struct {
	db *gorm.DB
}

func NewLoginService(db *gorm.DB) LoginService {
	return &auth{db}
}
func (auth *auth) LoginUser(email string, password string) bool {
	var user model.User
	err := auth.db.Where("email = ?", email).First(&user).Error
	pwd := []byte(password)
	return err == nil && comparePasswords(user.Password, pwd)
}

func HashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
