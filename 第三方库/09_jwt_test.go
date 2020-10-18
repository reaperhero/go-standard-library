package third

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

// 对称加密：加密和解密都在服务端

// 1、不过期
func Test_ExampleNewWithClaims_standardClaims_01(t *testing.T) {
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

// 2、会过期
func Test_ExampleNewWithClaims_standardClaims_02(t *testing.T) {
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
	user.ExpiresAt = time.Now().Add(time.Hour * 1).Unix() // 设置一小时后过期

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	token, err := tokenObj.SignedString(password)
	if err != nil {
		fmt.Println(err)
	}
	// 由于设置了时间，这里每次生成的token都不一致
	fmt.Println(token)

	// 解密
	getUser := User{}
	getToken, _ := jwt.ParseWithClaims(token, &getUser, func(token *jwt.Token) (interface{}, error) {
		return password, nil
	})
	if getToken.Valid { // 1小时后为false
		fmt.Println(getToken.Claims.(*User).Username)
	}
}

func Test_ExampleNewWithClaims_standardClaims_03(t *testing.T) {
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
	user.ExpiresAt = time.Now().Add(time.Second * 3).Unix() // 设置3秒后过期

	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	token, err := tokenObj.SignedString(password)
	if err != nil {
		fmt.Println(err)
	}
	// 由于设置了时间，这里每次生成的token都不一致
	fmt.Println(token)

	// 解密
	getUser := User{}
	time.Sleep(time.Second * 6)
	getToken, err := jwt.ParseWithClaims(token, &getUser, func(token *jwt.Token) (interface{}, error) {
		return password, nil
	})
	if getToken.Valid {
		fmt.Println("通过了")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
}
