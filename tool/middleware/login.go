package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/houyanzu/cache"
	"net/http"
	"strings"
)

func Login() gin.HandlerFunc {
	return loginHandler
}

func loginHandler(c *gin.Context) {
	token := c.GetHeader("token")
	account := c.GetHeader("wallet")
	account = strings.ToLower(account)
	lang := c.GetHeader("Language")
	if token == "" {
		if lang == "zh" {
			c.JSON(http.StatusOK, gin.H{
				"code": 3,
				"msg":  "亲，登陆过期了，需要重新登录哟",
				"data": gin.H{},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 3,
				"msg":  "Session expired, Please Login",
				"data": gin.H{},
			})
		}
		c.Abort()
		return
	}

	userId := cache.GetInt(token)
	if userId <= 0 {
		if lang == "zh" {
			c.JSON(http.StatusOK, gin.H{
				"code": 3,
				"msg":  "亲，登陆过期了，需要重新登录哟",
				"data": gin.H{},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 3,
				"msg":  "Session expired, Please Login",
				"data": gin.H{},
			})
		}
		c.Abort()
		return
	}

	tokenAccount := cache.GetString(token + "_address")
	if account != tokenAccount {
		if lang == "zh" {
			c.JSON(http.StatusOK, gin.H{
				"code": 3,
				"msg":  "亲，登陆过期了，需要重新登录哟",
				"data": gin.H{},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 3,
				"msg":  "Session expired, Please Login",
				"data": gin.H{},
			})
		}
		c.Abort()
		return
	}
	c.Set("userId", userId)
	c.Next()
	return
}
