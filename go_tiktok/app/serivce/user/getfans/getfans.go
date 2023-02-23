package getfans

import (
	"github.com/gin-gonic/gin"
	"go_tiktok/app/serivce/user/getfans/internal"
	"go_tiktok/utils/jwt"
	"net/http"
	"strconv"
)

func GetFans(c *gin.Context) {

	token := c.Query("token")
	userId := c.Query("user_id")

	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg":  "token or user_id cannot be null",
			"user":        "null",
		})
		return
	}
	if token != "" {
		id, err := jwt.GetUserId(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": -1,
				"status_msg":  err.Error(),
				"user":        "null",
			})
			return

		}
		userid, _ := strconv.ParseInt(userId, 10, 64)
		err, user := internal.GetFans(c, int(userid), id)
		if err != nil {

			if err.Error() == "internal err" {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status_code": -1,
					"status_msg":  err.Error(),
					"user":        "null",
				})
				return
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  http.StatusOK,
			"user":        user,
		})
		return
	} else {
		id := -1
		userid, _ := strconv.ParseInt(userId, 10, 64)
		err, user := internal.GetFans(c, int(userid), id)
		if err != nil {

			if err.Error() == "internal err" {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status_code": -1,
					"status_msg":  err.Error(),
					"user":        "null",
				})
				return
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  http.StatusOK,
			"user":        user,
		})
		return
	}

}
