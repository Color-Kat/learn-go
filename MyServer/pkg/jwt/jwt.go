package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTData struct {
	Email string
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{Secret: secret}
}

func (j *JWT) Create(data JWTData) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": data.Email,
		},
	)

	signedString, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func (j *JWT) Parse(token string) (*JWTData, bool) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, false
	}

	email := t.Claims.(jwt.MapClaims)["email"].(string)
	return &JWTData{
		Email: email,
	}, t.Valid
}
