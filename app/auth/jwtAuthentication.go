package auth

import (
	"fmt"
	"github.com/Fadhli12/go-gin-gorm-playground/model"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const tokenExpired = 24

//jwt service
type JWTService interface {
	GenerateToken(user model.User, isUser bool) (string, string)
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
}

//auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(user model.User, isUser bool) (string, string) {
	claims := &authCustomClaims{
		user.Email,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * tokenExpired).Unix(),
			Issuer:    user.Name,
			IssuedAt:  time.Now().Unix(),
		},
	}
	refreshClaims := &authCustomClaims{
		user.Email,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2 * tokenExpired).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	rt, err := refreshToken.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t, rt
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}
