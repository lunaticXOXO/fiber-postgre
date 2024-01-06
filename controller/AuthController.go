package controller

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fiber-postgre/model"
	"github.com/gofiber/fiber/v2"
)


var tokenName = "auth_token"
var jwtKey = []byte("secret_key")

type Claims struct {
	Username string `json:"username"`
	User_type int `json:"user_type"`
	jwt.StandardClaims
}

func GenerateToken(c *fiber.Ctx, user model.Users) (string,error) { 
	tokenExpired := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Username: user.Username,
		User_type: user.User_type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpired.Unix(),
		},
	}

	token_new := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	token_string,err := token_new.SignedString(jwtKey)

	if err != nil{
		return token_string,err
	}

	c.Cookie(&fiber.Cookie{
		Name : tokenName,
		Value : token_string,
		Expires: tokenExpired,
		Secure: false,
		HTTPOnly: true,
	})

	return token_string,nil
}

func ResetToken(c *fiber.Ctx){

	c.Cookie(&fiber.Cookie{
		Name : tokenName,
		Value : "",
		Expires : time.Now(),
		Secure: false,
		HTTPOnly: true,
	})
}


func ValidateCookies(c *fiber.Ctx) (bool,string,int) {
	cookie := c.Cookies(tokenName)
	if cookie != ""{
		accessToken := cookie
		accessClaims := &Claims{}
		parseToken,err := jwt.ParseWithClaims(accessToken,accessClaims,func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey,nil
		})

		if err == nil && parseToken.Valid{
			return true,accessClaims.Username,accessClaims.User_type
		}
	}

	return false,"",-1
}

func ValidateUser(c *fiber.Ctx,accessType int) bool{
	isAccess,username,usertype := ValidateCookies(c)
	fmt.Println("access :",isAccess,"username : ",username,"type : ",usertype)
	if isAccess{
		userValid := accessType == usertype
		if userValid{
			return true
		}
	}

	return false
}


func Authorize(next fiber.Handler,accessType int) fiber.Handler{
	return func(c * fiber.Ctx) error{

		isValidToken := ValidateUser(c,accessType)
		if isValidToken{
			return next(c)
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : "unauthorized access",
		})
	}

}
