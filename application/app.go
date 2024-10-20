package application

import (
	"github.com/x64devv/assetag/models"
)

func initialModlel() models.TeaModel {
	return models.TeaModel{
		CurrentScreen:     0,
		EmpolyeesScreen:   models.EmpolyeesScreen{},
		AssetScreen:       models.AssetScreen{},
		AssignmentsScreen: models.AssignmentsScreen{},
	}
}

/*
main screen
  [*] employees
  [ ] assets
  [ ] assignments

employees screen
  [1] john doe | john@mail.com | FT | FT0001
  [2] john doe | john@mail.com | FT | FT0001
  [3] john doe | john@mail.com | FT | FT0001

  n => new employee  m=> main menu q => quit

asset screen
 [*] [1] Dell XPS | JHLJK97AKJ | LAPTOP | INUSE | LT0001
 [ ] [2] Dell XPS | JHLJK97AKJ | LAPTOP | INUSE | LT0001
 [ ] [3] Dell XPS | JHLJK97AKJ | LAPTOP | INUSE | LT0001

  n => add asset a => assign asset u => unassign asset m=> main menu q => quit


assignment screen
  [*] [1] FT0001-LT0001-aldko-aklj | 2002-02-23 | Current | active
  [ ] [2] FT0001-LT0001-aldko-aklj | 2002-02-23 | 2003-09-12 | completed
  [ ] [3] FT0001-LT0001-aldko-aklj | 2002-02-23 | 2003-09-12 | completed

  v => view more info m=> main menu q => quit

*/
