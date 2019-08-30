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

// Each alien can move only to one neighboring city per day.
func (anInv Invasion) RunInvasionSync(days int) {

	for i := 0; i < days; i++ {
		// Sync day, move all aliens one city and then do fighting.

		// Move all aliens, one city step.
		for a := 0; a < anInv.aliensInvading.NumberOfAliens(); a++ {
			if anInv.aliensInvading.IsDead(a) {
				continue
			}
			aLoc := anInv.aliensInvading.Location(a)
			newCity := anInv.worldAttacked.RandomNeighboringCity(aLoc)
			anInv.aliensInvading.MoveAlienSync(a,newCity)
		}
		// Do sync fighting.
		destroyedCities := anInv.aliensInvading.FightingSync()
		// Iterate destroyed cities, erase cities from graph, and mark killed aliens as dead.

	}
}



