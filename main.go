package main

import (
	"AlienInvasion/invasion"
)


func main() {


	mainInvasion := invasion.New("x.map", 3)
	mainInvasion.RunInvasionSync( 16 )
}