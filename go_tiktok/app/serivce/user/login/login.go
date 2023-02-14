package login

import (
	"github.com/gin-gonic/gin"
	"go_tiktok/app/internal/model"
	"go_tiktok/app/serivce/user/login/internal"
	"net/http"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  "username cannot be null",
			"user_id":     -1,
			"token":       nil,
		})
		return
	}

	if password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  "password cannot be null",
			"user_id":     -1,
			"token":       nil,
		})
		return
	}

	userSubject := &model.UserSubject{
		Username: username,
		Password: internal.EncryptPassword(password),
	}

	err := internal.CheckPassword(c, userSubject)
	if err != nil {
		switch err.Error() {
		case "internal err":
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": http.StatusBadRequest,
				"status_msg":  err.Error(),
				"user_id":     -1,
				"token":       nil,
			})
		case "invalid username or password":
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": http.StatusBadRequest,
				"status_msg":  err.Error(),
				"user_id":     -1,
				"token":       nil,
			})
		}

		return
	}

	tokenString, err := internal.GenerateToken(c, userSubject)
	if err != nil {
		switch err.Error() {
		case "internal err":
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": http.StatusBadRequest,
				"status_msg":  err.Error(),
				"user_id":     -1,
				"token":       nil,
			})
		}

	}

	id := internal.GeeUserId(c, username)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"status_msg":  "login successfully",
		"user_id":     id,
		"token":       tokenString,
	})
}
