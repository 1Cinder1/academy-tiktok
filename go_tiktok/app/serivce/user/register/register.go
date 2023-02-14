package register

import (
	"github.com/gin-gonic/gin"
	"go_tiktok/app/internal/model"
	"go_tiktok/app/serivce/user/register/internal"
	"net/http"
)

func Register(c *gin.Context) {
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
	if len(username) > 32 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  "username too long",
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
	if len(password) > 32 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  "password too long",
			"user_id":     -1,
			"token":       nil,
		})
		return
	}
	err := internal.CheckUserIsExist(c, username)

	if err != nil {
		if err.Error() == "internal err" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": http.StatusBadRequest,
				"status_msg":  err.Error(),
				"user_id":     -1,
				"token":       nil,
			})
		} else if err.Error() == "username already exist" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status_code": http.StatusBadRequest,
				"status_msg":  err.Error(),
				"user_id":     -1,
				"token":       nil,
			})
		}
		return
	}

	userSubject := &model.UserSubject{}

	encryptedPassword := internal.EncryptPassword(password)

	userSubject.Username = username
	userSubject.Password = encryptedPassword
	internal.CreateUser(c, userSubject)
	id := internal.GeeUserId(c, username)
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"status_msg":  "register successfully",
		"user_id":     id,
		"token":       nil,
	})
}
