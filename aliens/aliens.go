package aliens

import (
	"fmt"
	"math/rand"
	"sync"
)

type Aliens struct {
	population int
	locations []int
	dead []bool
	// https://stackoverflow.com/questions/12677934/create-a-map-of-lists
	aliensPerLocation []map[int]bool
}

// Init aliens locations with random cities.
// Assumption: we can have more than one aline per city initially, but the start fighting after the first move.
func New(population int, numberOfLocations int) Aliens {

	locations := make([]int, population)
	dead := make([]bool, population)
	aliensPerLocation := make([]map[int]bool,numberOfLocations)
	for i := 0;  i < population; i++ {
		randLoc := rand.Intn(numberOfLocations)
		locations[i] = randLoc
		if len(aliensPerLocation[randLoc]) == 0 {
			aliensPerLocation[randLoc] = make(map[int]bool)
		}
		aliensPerLocation[randLoc][i] = true // append(aliensPerLocation[randLoc], i)
	}
	return Aliens{population, locations, dead,aliensPerLocation}
}

// Move Alien Sync
func (a *Aliens) MoveAlienSync (alien int, dst int) {
	src := a.locations[alien]
	a.locations[alien] = dst
	// https://stackoverflow.com/questions/34018908/golang-why-dont-we-have-a-set-datastructure
	delete(a.aliensPerLocation[src], alien) // remove alien from original loc
	if len(a.aliensPerLocation[dst]) == 0 {
		a	.aliensPerLocation[dst] = make(map[int]bool)
	}
	a.aliensPerLocation[dst][alien] = true
	a.aliensPerLocation[src][alien] = false // add alien to new destination
}

// Move Alien Sync
func (a *Aliens) MoveAlienAsync (alien int, dst int, movementLock sync.Mutex) {
	movementLock.Lock()
	src := a.locations[alien]
	a.locations[alien] = dst
	// https://stackoverflow.com/questions/34018908/golang-why-dont-we-have-a-set-datastructure
	delete(a.aliensPerLocation[src], alien) // remove alien from original loc
	if len(a.aliensPerLocation[dst]) == 0 {
		a	.aliensPerLocation[dst] = make(map[int]bool)
	}
	a.aliensPerLocation[dst][alien] = true
	a.aliensPerLocation[src][alien] = false // add alien to new destination
	movementLock.Unlock()
}


func (a Aliens) NumberOfAliens() int {
	return a.population
}

func (a Aliens) NumberOfAliensAlive() int {
	ret := 0
	for _, deadBool := range a.dead {
		if !deadBool {
			ret++
		}
	}
	return ret
}

func (a Aliens) IsDead(alien int) bool {
	return a.dead[alien]
}

func (a *Aliens) SetDead(alien int) {
	a.dead[alien] = true
}

func (a Aliens) Location(alien int) int {
	return a.locations[alien]
}

// Fighting of Aliens plus Killing and return list of destroyed Cities.
func (a *Aliens) FightingSync () map[int][]int {

	destroyedCities := make(map[int][]int)
	for location := 0; location < len(a.aliensPerLocation); location++ {
		alienLocCount := 0
		fightingAliens := make([]int, 0)
		for alien, insitu := range a.aliensPerLocation[location] {
			if !a.dead[alien] && insitu {
				alienLocCount++
				fightingAliens = append(fightingAliens, alien)
			}
		}
		if alienLocCount > 1 {
			// Return destroyed Cities
			destroyedCities[location] = fightingAliens
			// Mark fighters as dead.
			for fa := range fightingAliens {
				a.dead[fa] = true
			}
		}

	}
	return destroyedCities
}

// Fighting of Aliens plus Killing and return bool for city destroyed.
func (a *Aliens) SingleFight (city int, cityName string) bool {

	wasDestroyed := false
	alienLocCount := 0
	fightingAliens := make([]int, 0)
	for alien, insitu := range a.aliensPerLocation[city] {
		if !a.dead[alien] && insitu {
			alienLocCount++
			fightingAliens = append(fightingAliens, alien)
		}
	}
	if alienLocCount > 1 {
		wasDestroyed = true
		// Mark fighters as dead.
		for fa := range fightingAliens {
			a.dead[fa] = true
		}
		fmt.Printf("ALERT: %s has been destroyed by the following aliens: %s\n",cityName,fightingAliens)
	}

	return wasDestroyed
}

func (a Aliens) GetLocation(alien int) int {
	return a.locations[alien]
}
