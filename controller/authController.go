package controller

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go-api-friends/Stores"
	"go-api-friends/model"
	"log"
	"os"
	"time"
)

type AuthService struct {
	authMW *jwt.GinJWTMiddleware
}

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func NewJwt(us Stores.UserStore) *AuthService {
	authMW, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(os.Getenv("secret_key")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					"id": v.UserID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			id := int(claims["id"].(float64))
			user, _ := us.GetUser(id)
			return user
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginValues login
			if err := c.ShouldBind(&loginValues); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginValues.Username
			password := loginValues.Password

			if user := us.ValidateUser(userID, password); user != nil {
				return user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*model.User); ok {
				return true
			}

			return false
		},

		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	return &AuthService{authMW}
}
