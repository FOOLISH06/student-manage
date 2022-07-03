package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"student-manage/common"
	"student-manage/dto"
	"student-manage/model"
	"student-manage/response"
)

// Login 管理员登录
func Login(ctx *gin.Context) {
	db := common.GetDB()

	// 获取参数
	//telephone := ctx.PostForm("telephone")
	//password := ctx.PostForm("password")
	json := make(map[string]interface{})
	ctx.BindJSON(&json)
	telephone := json["telephone"].(string)
	password := json["password"].(string)
	fmt.Println("telephone: ", telephone)
	fmt.Println("password: ", password)

	// 参数验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}

	// 判断用户（手机号）是否存在
	var manager model.Manager
	db.Where("telephone = ?", telephone).First(&manager)
	if manager.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	// 密码验证
	if err := bcrypt.CompareHashAndPassword([]byte(manager.Password), []byte(password)); err != nil {
		response.Fail(ctx, nil, "密码错误")
		return
	}

	// 发放token
	token, err := common.ReleaseToken(manager)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Printf("token generate error: %v\n", err)
		return
	}

	// 返回结果：token + manager(name+telephone)
	response.Success(ctx, gin.H{"token": token, "manager": dto.ToManagerDto(manager)}, "登录成功")
}

// Register 管理员注册（只有超级管理员才能注册其他管理员）
func Register(ctx *gin.Context) {
	db := common.GetDB()

	// 获取参数
	json := make(map[string]interface{})
	ctx.BindJSON(&json)
	telephone := json["telephone"].(string)
	password := json["password"].(string)
	name := json["name"].(string)
	var isSuper bool = json["isSuper"].(bool)
	//if _, status := ctx.GetPostForm("isSuper"); status {
	//	isSuper = true
	//}

	// 数据验证，失败则返回“请求格式错误”
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能小于6位")
		return
	}
	if len(name) == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户名不能为空")
		return
	}

	// 判断手机号是否存在
	if isTelephoneExit(db, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号已存在")
		return
	}

	// 创建用户（需要加密）
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}

	newManager := &model.Manager{
		Name:           name,
		Telephone:      telephone,
		Password:       string(hashedPassword),
		IsSuperManager: isSuper,
	}
	db.Create(&newManager)

	// 返回结果
	response.Success(ctx, nil, "注册成功")
}

func isTelephoneExit(db *gorm.DB, telephone string) bool {
	var manager *model.Manager
	db.Where("telephone = ?", telephone).First(&manager)
	if manager.ID != 0 {
		return true
	}
	return false
}
