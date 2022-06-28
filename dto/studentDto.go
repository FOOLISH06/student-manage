package dto

import "student-manage/model"

type StudentDto struct {
	Sid   string `json:"学号"`
	Name  string `json:"姓名"`
	Sex   string `json:"性别"`
	Age   int    `json:"年龄"`
	Major string `json:"专业"`
	Class string `json:"班级"`
}

func ToStudentDtos(students ...model.Student) []StudentDto {
	var studentDtos []StudentDto
	for _, student := range students {
		studentDto := StudentDto{
			Sid:   student.Sid,
			Name:  student.Name,
			Sex:   student.Sex,
			Age:   student.Age,
			Major: student.Major,
			Class: student.Class,
		}
		studentDtos = append(studentDtos, studentDto)
	}
	return studentDtos
}
