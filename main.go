package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/x64devv/assetag/helpers"
	"github.com/x64devv/assetag/models"
)

func init() {

	if err := helpers.LoadEnv(); err != nil {
		panic(err)
	}
	//
	// if err := database.InitConnetion(); err != nil {
	// 	panic(err)
	// }
	//
	// if err := database.RunMigrations(); err != nil {
	// 	panic(err)
	// }

}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()
	options := []models.Screen{models.EmployeesSc, models.AssetsSc, models.AsssignmentsSc}
	p := tea.NewProgram(models.TeaModel{
		CurrentScreen: 0,
		Cursor:        0,
		MenuScreen: models.MenuScreen{
			SelectedOption: 0,
			Cursor:         0,
			Options:        options,
		},
		EmpolyeesScreen: models.EmpolyeesScreen{
			Employees:   []models.Employee{},
			NewEmployee: models.Employee{},
		},
		AssetScreen: models.AssetScreen{
			SelectedAsset: 0,
			Cursor:        0,
			Assets:        []models.Asset{},
			NewAsset:      models.Asset{},
		},
		AssignmentsScreen: models.AssignmentsScreen{
			SelectedAsset: 0,
			Cursor:        0,
			Assignments:   []models.Assignmet{},
		},
	})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
