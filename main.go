package main

import (
	"AlienInvasion/invasion"
	"math/rand"
)


func main() {

	rand.Seed(42)

	mainInvasion := invasion.New("x.map", 2)
	mainInvasion.RunInvasionSync( 100 )
}