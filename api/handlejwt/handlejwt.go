package handlejwt

import (
	"auth-server/api/db"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

type user struct {
	id       int
	name     string
	mail     string
	password string
}

func (u user) createToken() string {
	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = u.id
	claims["name"] = u.name
	claims["mail"] = u.mail
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// 電子署名
	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
	return tokenString
}

// GetTokenHandler get token
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	d, err := db.Connect()
	if err != nil {
		fmt.Println(err)
	}
	defer d.Close()

	mail, pass, _ := r.BasicAuth()
	fmt.Println("basic " + mail + " : " + pass)

	var u user
	err = d.QueryRow(`
			SELECT *
			FROM users
			WHERE mail = ? AND password = ?
		`, mail, pass).Scan(&(u.id), &(u.name), &(u.mail), &(u.password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized\n"))
		return
	}
	// headerのセット
	tokenString := u.createToken()

	// JWTを返却
	w.Write([]byte(tokenString))
})

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
