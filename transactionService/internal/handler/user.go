package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"transactionService/pkg/proto/user_v1"
)

func (h *Handler) createUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.UsersService.Create(newUserToDto(user))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	user.Id = id
	c.JSON(http.StatusOK, user)
}

func (h *Handler) getUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := h.services.UsersService.GetById(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newUserFromDto(user))
}

type User struct {
	Id      string  `json:"id,omitempty"`
	Balance float64 `json:"balance,omitempty"`
}

func newUserToDto(user User) user_v1.User {
	return user_v1.User{
		Id:      user.Id,
		Balance: user.Balance,
	}
}

func newUserFromDto(user user_v1.User) User {
	return User{
		Id:      user.GetId(),
		Balance: user.GetBalance(),
	}
}
