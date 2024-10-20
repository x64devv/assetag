package controllers

import (
	"errors"

	"github.com/x64devv/assetag/database"
	"github.com/x64devv/assetag/models"
)

func AssignAsset(employeeId, assetId string) error {
	assignment := models.Assignmet{EmployeeId: employeeId, AssetId: assetId}
	result := database.DbInstanse.Create(&assignment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UnAssignAsset(assignmentId string) error {
	result := database.DbInstanse.Where("assignment_code = ?").Update("assignment_status", "complete")
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindAssignment(assignmentCode string) (models.Assignmet, models.Employee, models.Asset, error) {
	var assignment models.Assignmet
	var employee models.Employee
	var asset models.Asset

	result := database.DbInstanse.Where("assignment_code = ?", assignmentCode).First(&assignment)

	if result.Error != nil {
		return assignment, employee, asset, result.Error
	}

	if assignment.ID == 0 {
		return assignment, employee, asset, errors.New("Assignment not found.")
	}

	result = database.DbInstanse.Where("id = ?", assignment.EmployeeId).First(&employee)

	if result.Error != nil {
		return assignment, employee, asset, result.Error
	}

	result = database.DbInstanse.Where("id = ?", assignment.AssetId).First(&asset)

	if result.Error != nil {
		return assignment, employee, asset, result.Error
	}
	return assignment, employee, asset, nil
}
