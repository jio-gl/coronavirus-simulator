package invasion

import (
	"AlienInvasion/aliens"
	"AlienInvasion/world"
	"fmt"
	"sync"
	"time"
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

func (inv Invasion) GetWorld() *world.World {
	return inv.worldAttacked
}

func (inv Invasion) GetAliens() *aliens.Aliens {
	return inv.aliensInvading
}


// Assumption: Each alien can move only to one neighboring city per day.
func (anInv *Invasion) RunInvasionSync(days int) {

	fmt.Println( anInv.worldAttacked)
	fmt.Println( anInv.aliensInvading)

	fmt.Println("Initial Number of Cities = ",anInv.worldAttacked.NumberOfCities())


	for i := 0; i < days; i++ {
		fmt.Println("Day = ", i)

		// Sync day, move all aliens one city and then do fighting.

		// Move all aliens, one city step.
		anyMovement := false
		for a := 0; a < anInv.aliensInvading.NumberOfAliens(); a++ {
			if !anInv.aliensInvading.IsDead(a) {
				fmt.Println( "    Alien = ",a )
				aLoc := anInv.aliensInvading.Location(a)
				newCity := anInv.worldAttacked.RandomNeighboringCity(aLoc)
				fmt.Printf( "    Moving alien %d from location %d to location %d\n",a,aLoc,newCity)

				if aLoc != newCity {
					anyMovement = true
				}
				anInv.aliensInvading.MoveAlienSync(a,newCity)
			} else {
				fmt.Println( "    Alien is Dead = ",a )
			}
		}
		// Do sync fighting.
		destroyedCities := anInv.aliensInvading.FightingSync()
		// Iterate destroyed cities, erase cities from graph, and mark killed aliens as dead.
		for loc, aliensDead := range destroyedCities {
			fmt.Printf("location[%d] aliensDead[%d]\n", loc, aliensDead)
			fmt.Println("Destroying city =",anInv.worldAttacked.CityName(loc))
			anInv.worldAttacked.DestroyCity(loc)
			for a := range aliensDead {
				anInv.aliensInvading.SetDead(a)
			}
		}
		fmt.Println("Number of Cities = ",anInv.worldAttacked.NumberOfCities())
		if anInv.worldAttacked.NumberOfCities() == 0 {
			fmt.Println("WARNING: All cities were destroyed!!! Stopping simulation...")
			break
		}
		if anInv.worldAttacked.NumberOfCities() == 1 {
			fmt.Println("WARNING: Only one city remaining! Stopping simulation...")
			break
		}
		if anInv.worldAttacked.NumberOfRoutes() == 0 {
			fmt.Printf("WARNING: No more routes, aliens are trapped and cant moved! Remaining number of aliens = %d  Stopping simulation...\n", anInv.aliensInvading.NumberOfAliensAlive())
			break
		}
		if !anyMovement {
			fmt.Println("WARNING: No movements detected, aliens are trapped or all dead! Stopping simulation...")
			break
		}

	}
}


// Each alien can move only to one neighboring city per day.
func (anInv *Invasion) RunInvasionAsync(days int) {

	fmt.Println(anInv.worldAttacked)
	fmt.Println(anInv.aliensInvading)

	fmt.Println("Initial Number of Cities = ", anInv.worldAttacked.NumberOfCities())

	var endOfInvasion sync.WaitGroup
	var movementLock sync.Mutex
	endOfInvasion.Add( anInv.aliensInvading.NumberOfAliens() )
	for a := 0; a < anInv.aliensInvading.NumberOfAliens(); a++ {
		go anInv.startAlien(a,days, &endOfInvasion, movementLock)
	}

	endOfInvasion.Wait()
	fmt.Println("End of Async Invasion reached.")

}

func (anInv *Invasion) startAlien(alien int, days int, endOfInvasion *sync.WaitGroup, movementLock sync.Mutex ) {

	for i := 0; i < days; i++ {

		fmt.Printf("Day %d for alien %d, sleeping and then move ...\n",alien,i)
		time.Sleep(1 * time.Second)

		aLoc := anInv.aliensInvading.Location(alien)
		newCity := anInv.worldAttacked.RandomNeighboringCity(aLoc)

		// Lock destination city before moving.
		anInv.worldAttacked.LockCity(newCity)

		// Do the movement, fight and destruction
		fmt.Printf("    Moving alien %d from location %d %s to location %d %s\n ", alien, aLoc, anInv.worldAttacked.CityName(aLoc), newCity,anInv.worldAttacked.CityName(newCity))
		anInv.aliensInvading.MoveAlienAsync(alien, newCity, movementLock)

		wasDestroyed := anInv.aliensInvading.SingleFight(newCity, anInv.worldAttacked.CityName(newCity))
		if wasDestroyed {
			anInv.worldAttacked.DestroyCity(newCity)
		}
		// End of movement, fight and destruction

		fmt.Println("Number of Cities = ",anInv.worldAttacked.NumberOfCities())
		fmt.Println("Number of Aliens Alive = ",anInv.aliensInvading.NumberOfAliensAlive())
		if anInv.aliensInvading.IsDead(alien) {
			fmt.Printf("WARNING: Alien is dead!!! Stopping alien %d ...\n", alien)
			anInv.worldAttacked.UnlockCity(newCity)
			break
		}
		if anInv.worldAttacked.NumberOfRoutesOut(newCity) == 0 {
			fmt.Printf("WARNING: Alien is trapped in city %s!!! Stopping alien %d ...\n", anInv.worldAttacked.CityName(newCity),alien)
			anInv.worldAttacked.UnlockCity(newCity)
			break
		}
		if anInv.worldAttacked.NumberOfCities() == 0 {
			fmt.Printf("WARNING: All cities were destroyed!!! Stopping alien %d ...\n", alien)
			anInv.worldAttacked.UnlockCity(newCity)
			break
		}
		if anInv.worldAttacked.NumberOfCities() == 1 {
			fmt.Printf("WARNING: Only one city remaining! Stopping alien %d ...\n", alien)
			anInv.worldAttacked.UnlockCity(newCity)
			break
		}
		if anInv.worldAttacked.NumberOfRoutes() == 0 {
			fmt.Printf("WARNING: No more routes, aliens are trapped and cant moved! Remaining number of aliens = %d  Stopping simulation...\n", anInv.aliensInvading.NumberOfAliensAlive())
			anInv.worldAttacked.UnlockCity(newCity)
			break
		}

		// Unlock destination city after moving.
		anInv.worldAttacked.UnlockCity(newCity)

	}

	// Signal end of alien activity
	endOfInvasion.Done()

}