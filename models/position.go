package models

import "gorm.io/gorm"

type Position struct {
	gorm.Model
	Name         string     `json:"name"`
	Code         string     `json:"code"`
	DepartmentID uint       `json:"department_id"`
	Department   Department `json:"department"`
	Employees    []Employee `json:"employee"`
}

type PositionResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	DepartmentID uint   `json:"department_id"`
}

type GetPositionResponse struct {
	ID           uint               `json:"id"`
	Name         string             `json:"name"`
	Code         string             `json:"code"`
	DepartmentID uint               `json:"department_id"`
	Department   DepartmentResponse `json:"department"`
}
