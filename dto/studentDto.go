package dto

import (
	"strconv"
	"student-manage/model"
)

type StudentDto struct {
	Sid   string `json:"sid"`
	Name  string `json:"name"`
	Sex   string `json:"sex"`
	Age   string `json:"age"`
	Major string `json:"major"`
	Class string `json:"class"`
}

func ToStudentDtos(students ...model.Student) []StudentDto {
	var studentDtos []StudentDto
	for _, student := range students {
		studentDto := StudentDto{
			Sid:   student.Sid,
			Name:  student.Name,
			Sex:   student.Sex,
			Age:   strconv.Itoa(student.Age),
			Major: student.Major,
			Class: student.Class,
		}
		studentDtos = append(studentDtos, studentDto)
	}
	return studentDtos
}
