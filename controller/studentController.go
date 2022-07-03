package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"student-manage/common"
	"student-manage/dto"
	"student-manage/model"
	"student-manage/response"
)

// GetStudents 获取所有学生信息 (Get)
func GetStudents(ctx *gin.Context) {
	db := common.GetDB()
	var students []model.Student
	db.Find(&students)
	response.Success(ctx, gin.H{"students": dto.ToStudentDtos(students...)}, "查询成功")
}

// GetStudentById 通过学号获取学生信息 (Get)
func GetStudentById(ctx *gin.Context) {
	db := common.GetDB()

	// 获取动态参数sid
	sid := ctx.Param("sid")

	var student model.Student
	db.Where("sid = ?", sid).First(&student)
	response.Success(ctx, gin.H{"students": dto.ToStudentDtos(student)}, "查询成功")
}

// GetStudentsByClass 通过班级获取学生信息 (Get)
func GetStudentsByClass(ctx *gin.Context) {
	db := common.GetDB()

	// 获取动态参数class
	class := ctx.Param("class")

	var students []model.Student
	db.Where("class = ?", class).Find(&students)
	response.Success(ctx, gin.H{"students": dto.ToStudentDtos(students...)}, "查询成功")
}

// CreateStudent 添加学生 (Post)
func CreateStudent(ctx *gin.Context) {
	db := common.GetDB()

	// 获取参数
	json := make(map[string]interface{})
	ctx.BindJSON(&json)
	sid := json["sid"].(string)
	name := json["name"].(string)
	sex := json["sex"].(string)
	major := json["major"].(string)
	class := json["class"].(string)
	age, _ := strconv.Atoi(json["age"].(string))

	for _, val := range json {
		s := val.(string)
		if len(s) == 0 {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "关键信息不能为空")
			return
		}
	}

	// 如果学生已存在，添加失败
	if isStudentExit(db, sid) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "学生已存在")
		return
	}

	student := &model.Student{
		Sid:   sid,
		Name:  name,
		Sex:   sex,
		Age:   age,
		Major: major,
		Class: class,
	}
	db.Create(student)

	response.Success(ctx, nil, "添加成功")
}

// UpdateStudent 通过学号修改学生信息 (Put)
func UpdateStudent(ctx *gin.Context) {
	db := common.GetDB()

	// 获取参数 (form)
	json := make(map[string]interface{})
	ctx.BindJSON(&json)
	sid := json["sid"].(string)
	name := json["name"].(string)
	sex := json["sex"].(string)
	major := json["major"].(string)
	class := json["class"].(string)
	age, _ := strconv.Atoi(json["age"].(string))

	// 如果学生不存在，修改失败
	if !isStudentExit(db, sid) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "学生不存在")
		return
	}

	db = db.Model(model.Student{}).Where("sid = ?", sid).Updates(model.Student{
		Name:  name,
		Sex:   sex,
		Age:   age,
		Major: major,
		Class: class,
	})

	response.Success(ctx, nil, "修改成功")
}

// DeleteStudent 通过学号删除学生 (Delete)
func DeleteStudent(ctx *gin.Context) {
	db := common.GetDB()
	sid := ctx.Param("sid")

	// 如果学生不存在，删除失败
	if !isStudentExit(db, sid) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "学生不存在")
		return
	}

	db = db.Unscoped().Where("sid = ?", sid).Delete(&model.Student{})
	fmt.Println("row affected ", db.RowsAffected)
	response.Success(ctx, nil, "删除成功")
}

func isStudentExit(db *gorm.DB, sid string) bool {
	var student *model.Student
	db.Where("sid = ?", sid).First(&student)
	if student.ID != 0 {
		return true
	}
	return false
}
