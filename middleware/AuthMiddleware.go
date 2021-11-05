package middleware

import (
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取authorization header
		tokenString := c.GetHeader("Authorization")

		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			//c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			response.Response(c, http.StatusUnauthorized,401, nil, "权限不足")
			c.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Response(c, http.StatusUnauthorized,401, nil, "权限不足")
			c.Abort()
			return
		}

		// 验证通过后获取claim中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 用户
		if user.ID == 0 {
			response.Response(c, http.StatusUnauthorized,401, nil, "权限不足")
			c.Abort()
			return
		}

		//用户存在 讲user的信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}