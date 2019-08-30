package invasion

import (
	"AlienInvasion/aliens"
	"AlienInvasion/world"
	"fmt"
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
		fmt.Println("Day = ", i)
		// Sync day, move all aliens one city and then do fighting.

		// Move all aliens, one city step.
		for a := 0; a < anInv.aliensInvading.NumberOfAliens(); a++ {
			fmt.Println( "Alien = ",a )
			if anInv.aliensInvading.IsDead(a) {
				fmt.Println( "Alien is Dead = ",a )
				continue
			}
			aLoc := anInv.aliensInvading.Location(a)
			newCity := anInv.worldAttacked.RandomNeighboringCity(aLoc)
			anInv.aliensInvading.MoveAlienSync(a,newCity)
		}
		// Do sync fighting.
		//destroyedCities :=
		anInv.aliensInvading.FightingSync()
		// Iterate destroyed cities, erase cities from graph, and mark killed aliens as dead.

	}
}



