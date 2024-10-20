package models

import "gorm.io/gorm"

type Screen int

const (
	MainMenuSc Screen = iota
	EmployeesSc
	AssetsSc
	AsssignmentsSc
)

func (s Screen) String() string {
	return [...]string{"Employees", "Assets", "Assignments"}[s-1]
}

type Employee struct {
	gorm.Model
	Firstname       string
	Lastname        string
	Email           string `gorm:"unique"`
	EmployeeCode    string `gorm:"unique"`
	NationalId      string `gorm:"unique"`
	EmployementType string
}

type Asset struct {
	gorm.Model
	Name         string
	SerialNumber string `gorm:"unique"`
	AssetType    string
	AssetCode    string `gorm:"unique"`
	AssetStatus  string
}

type Assignmet struct {
	gorm.Model
	EmployeeId      string
	AssetId         string
	AssignmetCode   string `gorm:"unique"`
	AssignmetStatus string
	AssignmetStart  string
	AssignmetEnd    string
}

type TeaModel struct {
	CurrentScreen     Screen
	Cursor            int
	MenuScreen        MenuScreen
	EmpolyeesScreen   EmpolyeesScreen
	AssetScreen       AssetScreen
	AssignmentsScreen AssignmentsScreen
}
type MenuScreen struct {
	SelectedOption int
	Cursor         int
	Options        []Screen
}

type EmpolyeesScreen struct {
	Employees   []Employee
	NewEmployee Employee
}

type AssetScreen struct {
	SelectedAsset int
	Cursor        int
	Assets        []Asset
	NewAsset      Asset
}

type AssignmentsScreen struct {
	SelectedAsset int
	Cursor        int
	Assignments   []Assignmet
}
