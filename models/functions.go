package models

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

// import "gorm.io/gorm"

func (a *Assignmet) GetEmployee() Employee {
	var employee Employee
	return employee
}

func (a *Assignmet) GetAsset() Asset {
	var asset Asset
	return asset
}

func (e *Employee) GenerateCode() string {
	return fmt.Sprintf("%s%d", e.EmployementType, e.ID)
}

func (a *Asset) GenerateCode() string {
	return fmt.Sprintf("%s%d", a.AssetType, a.ID)
}

func (m TeaModel) Init() tea.Cmd {
	return nil
}

//

func (m TeaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		} else if msg.String() == "m" {
			m.CurrentScreen = 0
		} else {
			switch m.CurrentScreen {
			case MainMenuSc:
				switch msg.String() {
				case "j", "down":
					if m.MenuScreen.Cursor < len(m.MenuScreen.Options) {
						m.MenuScreen.Cursor++
						m.Cursor++
					}
				case "k", "up":
					if m.MenuScreen.Cursor > 0 {
						m.MenuScreen.Cursor--
						m.Cursor--
					}
				case "enter":
					m.CurrentScreen = m.MenuScreen.Options[m.Cursor]

				}
			case EmployeesSc:
				switch msg.String() {
				case "n":
					log.Panicln("wants to add new employee")
				}
			case AssetsSc:
				switch msg.String() {
				case "n":
					log.Panicln("wants to add new asset")
				case "a":
					log.Panicln("wants to assign asset")
				case "u":
					log.Panicln("wants to unassign asset")
				case "j", "down":
					if m.AssetScreen.Cursor < len(m.AssetScreen.Assets) {
						m.AssetScreen.Cursor++
					}
				case "k", "up":
					if m.AssetScreen.Cursor > 0 {
						m.AssetScreen.Cursor--
					}
				}
			case AsssignmentsSc:
				switch msg.String() {
				case "v":
					log.Println("want to view more about assignment")
				case "j", "down":
					if m.AssignmentsScreen.Cursor < len(m.AssignmentsScreen.Assignments) {
						m.AssignmentsScreen.Cursor++
					}
				case "k", "up":
					if m.AssignmentsScreen.Cursor > 0 {
						m.AssignmentsScreen.Cursor--
					}
				}
			}
		}
	}
	return m, nil
}

func (m TeaModel) View() string {
	s := "=========================================\n"
	s += "================ ASSETAG ================\n"
	s += "=========================================\n\n\n"

	switch m.CurrentScreen {
	case MainMenuSc:
		s += "\t > Main Menu\n\n"
		for i, option := range m.MenuScreen.Options {
			selected := " "
			if i == m.MenuScreen.Cursor {
				selected = "*"
			}

			s += fmt.Sprintf("\t[%s] %s\n", selected, option)
		}

		s += "\n\npress: "
	case EmployeesSc:
		s += "\t > Employes\n\n"
		s += fmt.Sprintf("# \t| Name\t\t\t| Email\t\t\t| Contract type\t\t\t| Employee Code\n")
		for i, emp := range m.EmpolyeesScreen.Employees {
			s += fmt.Sprintf("[%d] \t| %s %s\t\t\t| %s\t\t\t| %s\t\t\t| %s\n", i+1, emp.Firstname, emp.Lastname, emp.Email, emp.EmployementType, emp.EmployeeCode)
		}

		s += "\n\npress: n => new employee"
	case AssetsSc:
		s += "\t > Assets\n\n"
		s += fmt.Sprintf("[-] [#]\tName \t\t\t| SN\t\t\t|Type \t\t\t| Status\t\t\t| Asset Code\n")
		for i, asset := range m.AssetScreen.Assets {
			selected := " "
			if m.AssetScreen.Cursor == i {
				selected = "*"
			}
			s += fmt.Sprintf("[%s] [%d]\t%s \t\t\t|%s \t\t\t|%s \t\t\t|%s \t\t\t|%s\n", selected, i+1, asset.Name, asset.SerialNumber, asset.AssetType, asset.AssetStatus, asset.AssetCode)
		}

		s += "\n\npress n => add asset a => assign asset u => unassign asset"
	case AsssignmentsSc:
		s += "\t > Assignents:\n\n"
		s += fmt.Sprintf("[-] [#]\t Assignment Code\t\t\t| Start Date\t\t\t| End Date\t\t\t| Status\n")

		for i, assignment := range m.AssignmentsScreen.Assignments {
			selected := " "
			if m.AssignmentsScreen.Cursor == i {
				selected = "*"
			}

			s += fmt.Sprintf("[%s] [%d]\t %s\t\t\t| %s\t\t\t| %s\t\t\t| %s \n", selected, i+1, assignment.AssignmetCode, assignment.AssignmetStart, assignment.AssignmetEnd, assignment.AssignmetStatus)
		}

		s += "\n\npress v => view more info"

	}

	return s + " m => main menu q => quit\n"
}
