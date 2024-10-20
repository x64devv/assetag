package controllers

import (
	"github.com/x64devv/assetag/database"
	"github.com/x64devv/assetag/models"
)

func AddEmployee(employee models.Employee) error {
	result := database.DbInstanse.Create(&employee)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func SearchEmployee(searchTerm string) ([]models.Employee, error) {
	var employees []models.Employee
	searchTerm = "%" + searchTerm + "%"
	result := database.DbInstanse.Where("name like ? or serial_number like ? or asset_code like ? ", searchTerm, searchTerm, searchTerm).Find(&employees)
	if result != nil {
		return employees, result.Error
	}
	return employees, nil
}
