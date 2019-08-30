package invasion

import (
	"AlienInvasion/aliens"
	"AlienInvasion/world"
)

type Invasion struct {
	aliensInvading *aliens.Aliens
	worldAttacked *world.World
}

func New(worldFilename string, alienPopulation int) Invasion {
	worldAttacked := world.LoadWorld(worldFilename)
	aliensInvading := aliens.New(alienPopulation, worldAttacked.NumberOfCities())
	return Invasion{&aliensInvading, &worldAttacked}
}

func (i Invasion) NumberOfCities() int {
	return len(w.cityIds)
}



