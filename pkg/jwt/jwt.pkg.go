package jwt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/raminfp/backend_gin/entity"
	"github.com/raminfp/backend_gin/serilizers"
	"io"
	"os"
	"time"
)

type Jwt struct {
}

func (j Jwt) CreateToken(user entity.User) (serilizers.Token, error) {
	var err error

	claims := jwt.MapClaims{}
	claims["user_id"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt := serilizers.Token{}

	jwt.AccessToken, err = token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return jwt, err
	}

	return j.createRefreshToken(jwt)
}

func (Jwt) ValidateToken(accessToken string) (entity.User, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	user := entity.User{}
	if err != nil {
		return user, err
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		user.Email = payload["user_id"].(string)

		return user, nil
	}

	return user, errors.New("invalid token")
}

func (Jwt) ValidateRefreshToken(model serilizers.Token) (entity.User, error) {
	sha1 := sha1.New()
	io.WriteString(sha1, os.Getenv("SECRET_KEY"))

	user := entity.User{}
	salt := string(sha1.Sum(nil))[0:16]
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		return user, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return user, err
	}

	data, err := base64.URLEncoding.DecodeString(model.RefreshToken)
	if err != nil {
		return user, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plain, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return user, err
	}

	if string(plain) != model.AccessToken {
		return user, errors.New("invalid token")
	}

	claims := jwt.MapClaims{}
	parser := jwt.Parser{}
	token, _, err := parser.ParseUnverified(model.AccessToken, claims)

	if err != nil {
		return user, err
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return user, errors.New("invalid token")
	}

	user.Email = payload["user_id"].(string)

	return user, nil
}

func (Jwt) createRefreshToken(token serilizers.Token) (serilizers.Token, error) {
	sha1 := sha1.New()
	io.WriteString(sha1, os.Getenv("SECRET_KEY"))

	salt := string(sha1.Sum(nil))[0:16]
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		fmt.Println(err.Error())

		return token, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return token, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return token, err
	}

	token.RefreshToken = base64.URLEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(token.AccessToken), nil))

	return token, nil
}
