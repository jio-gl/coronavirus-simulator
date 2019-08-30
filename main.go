package main

import (
	"AlienInvasion/world"
	"AlienInvasion/aliens"
	"math/rand"
	)


func main() {

	rand.Seed(42)

	var mapFilename = "x.map"
	someworld := world.LoadWorld(mapFilename)

}