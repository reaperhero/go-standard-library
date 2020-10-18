package third

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

// 对称加密：加密和解密都在服务端
func Test_ExampleNewWithClaims_standardClaims(t *testing.T) {
	// 加密
	type User struct {
		Username string `json:"username"`
		jwt.StandardClaims
	}
	password := []byte("123456")
	user := User{
		Username:       "chenqiangjun",
		StandardClaims: jwt.StandardClaims{},
	}

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	token, err := tokenObj.SignedString(password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token) // eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImNoZW5xaWFuZ2p1biJ9.V_i3Yu11kutodD2KlJgxwxGqGrSePAs8_DpiSiHTKAI

	// 解密
	getUser := User{}
	getToken, _ := jwt.ParseWithClaims(token, &getUser, func(token *jwt.Token) (interface{}, error) {
		return password, nil
	})
	if getToken.Valid {
		fmt.Println(getToken.Claims.(*User).Username) // chenqiangjun
	}
}
