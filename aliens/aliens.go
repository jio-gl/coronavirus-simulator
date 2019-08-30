package aliens

import (
	"math/rand"
)

type Aliens struct {
	population int
	locations []int
	// https://stackoverflow.com/questions/12677934/create-a-map-of-lists
	aliensPerLocation map[int][]int
}

// Init aliens locations with random cities.
// Assumption: the alien can't fight on their initial city, they starting fighting after the first move.
func New(population int, numberOfLocations int) Aliens {
	locations := make([]int, population)
	aliensPerLocation := make(map[int][]int)
	for i := 0;  i < population; i++ {
		randLoc := rand.Intn(numberOfLocations)
		locations[i] = randLoc
		if _, ok := aliensPerLocation[randLoc]; !ok {
			aliensPerLocation[randLoc] = make([]int, 0)
		}
		aliensPerLocation[randLoc] = append(aliensPerLocation[randLoc], i)
	}
	return Aliens{population, locations,aliensPerLocation}
}

