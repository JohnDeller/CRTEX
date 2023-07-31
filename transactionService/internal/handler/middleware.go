package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const userCtx = "userId"

func getUserId(c *gin.Context) (string, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return "", errors.New("user id not found")
	}

	idString, ok := id.(string)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id nis of invalid type")
		return "", errors.New("user id not found")
	}

	return idString, nil
}
