package middlerware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"student-manage/common"
	"student-manage/model"
)

// AuthMiddleware 中间件：验证token
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 Authorization Header
		tokenString := ctx.GetHeader("Authorization")

		// invalid token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			log.Println("invalid token format, token:  ", tokenString)
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			log.Printf("authorize token failed, err: %v\n", err.Error())
			ctx.Abort()
			return
		}

		// 验证通过，获取 claims 中的userId，在数据库中查找用户
		userId := claims.UserId
		db := common.GetDB()
		var manager model.Manager
		db.First(&manager, userId)

		// 如果用户不存在
		if manager.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			log.Println("user doesn't exist, userId: ", userId)
			ctx.Abort()
			return
		}

		// 用户存在，将 user 的信息写入上下文
		ctx.Set("user", manager)

		ctx.Next()
	}
}
