package follow

import (
	"github.com/gin-gonic/gin"
	"go_tiktok/app/serivce/user/follow/internal"
	"go_tiktok/utils/jwt"
	"net/http"
	"strconv"
)

func Follow(c *gin.Context) {

	token := c.Query("token")
	userId := c.Query("to_user_id")
	actionType := c.Query("actio_type")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg":  "token or user_id cannot be null",
		})
		return
	}
	if actionType == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg":  "actionType cannot be null",
		})
		return
	}
	id, err := jwt.GetUserId(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": -1,
			"status_msg":  err.Error(),
		})
		return

	}
	userid, _ := strconv.ParseInt(userId, 10, 64)

	if actionType == "1" {
		flag := internal.CheckIsFollower(c, id, int(userid))
		if !flag {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": -1,
				"status_msg":  "has add follow",
			})
			return
		}

		err := internal.AddFollower(c, id, int(userid))

		if err != nil {

			if err.Error() == "internal err" {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status_code": -1,
					"status_msg":  err.Error(),
				})
				return
			}
			return
		}
	}

	if actionType == "2" {
		flag := internal.CheckIsFollower(c, id, int(userid))
		if flag {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": -1,
				"status_msg":  "didn't add follow",
			})
			return
		}

		err := internal.DeleteFollower(c, id, int(userid))

		if err != nil {

			if err.Error() == "internal err" {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status_code": -1,
					"status_msg":  err.Error(),
				})
				return
			}
			return
		}
	}

}
