package dto

import "student-manage/model"

type ManagerDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToManagerDto(manager model.Manager) ManagerDto {
	return ManagerDto{
		Name:      manager.Name,
		Telephone: manager.Telephone,
	}
}
