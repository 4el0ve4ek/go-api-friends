package controller

import (
	"github.com/gin-gonic/gin"
	"go-api-friends/model"
	"net/http"
	"strconv"
)

// GetUsers returns to sender all saved users using json format.
func (ps *PhoneServer) GetUsers(c *gin.Context) {
	c.JSON(200, ps.store.GetAllUser())
}

// GetUserById returns to sender user with id from request.
func (ps *PhoneServer) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := ps.store.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

// AddUser adds new user with name and city from body.
func (ps *PhoneServer) AddUser(c *gin.Context) {
	type RequestUser struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	var user RequestUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ps.store.AddUser(user.Name, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(200, "OK")
}

// GetUserFromCity returns users from certain city.
func (ps *PhoneServer) GetUserFromCity(c *gin.Context) {
	users := ps.store.GetAllUser()
	requiredCity := c.Param("city")
	result := make([]*model.User, 0)
	for _, user := range users {
		if user.City == requiredCity {
			result = append(result, user)
		}
	}
	c.JSON(200, result)
}

// ChangeCity updates city of certain user.
func (ps *PhoneServer) ChangeCity(c *gin.Context) {
	type RequestedCity struct {
		City string `json:"city"`
	}
	var field RequestedCity
	if err := c.ShouldBindJSON(&field); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := ps.authMiddleWare.authMW.IdentityHandler(c).(*model.User)
	user.City = field.City
	ps.store.UpdateUser(user)

	c.JSON(200, user)
}

// ChangeStatus updates status of certain user.
func (ps *PhoneServer) ChangeStatus(c *gin.Context) {
	type RequestedStatus struct {
		Status string `json:"status"`
	}
	var field RequestedStatus
	if err := c.ShouldBindJSON(&field); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := ps.authMiddleWare.authMW.IdentityHandler(c).(*model.User)
	user.Status = field.Status
	ps.store.UpdateUser(user)

	c.JSON(200, user)
}

// AddRelations subscribe sender to user in json in body.
func (ps *PhoneServer) AddRelations(c *gin.Context) {
	type RequestedId struct {
		GoalId uint `json:"goal_id"`
	}
	var field RequestedId
	if err := c.ShouldBindJSON(&field); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := ps.authMiddleWare.authMW.IdentityHandler(c).(*model.User)
	err := ps.store.AddFollower(int(user.UserID), int(field.GoalId))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

// GetSubsNews return statuses of users, followed by sender.
func (ps *PhoneServer) GetSubsNews(c *gin.Context) {
	user := ps.authMiddleWare.authMW.IdentityHandler(c).(*model.User)
	subs := ps.store.GetSubs(int(user.UserID))
	headlines := make(map[string]string)
	for _, user := range subs {
		headlines[user.Name] = user.Status
	}
	c.JSON(200, headlines)
}
