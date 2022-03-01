package controller

import (
	"github.com/gin-gonic/gin"
	Stores "go-api-friends/Stores"
)

// PhoneServer contains function which will work on each routes
type PhoneServer struct {
	store          Stores.UserStore
	authMiddleWare *AuthService
}

// NewPhoneServer constructs PhoneServer
func NewPhoneServer() *PhoneServer {
	store := Stores.NewStore()
	authMW := NewJwt(store)
	return &PhoneServer{store, authMW}
}

func (ps *PhoneServer) LoginHandler(c *gin.Context) {
	ps.authMiddleWare.authMW.LoginHandler(c)
}

func (ps *PhoneServer) AuthMiddleWare() gin.HandlerFunc {
	return ps.authMiddleWare.authMW.MiddlewareFunc()
}

func (ps *PhoneServer) RefreshHandler(c *gin.Context) {
	ps.authMiddleWare.authMW.RefreshHandler(c)
}
