package main

import (
	"AlienInvasion/invasion"
)


func main() {


	mainInvasion := invasion.New("maps/big.map", 16)
	//mainInvasion.RunInvasionSync( 16 )
	mainInvasion.RunInvasionAsync(20)
	mainInvasion.GetAliens().NumberOfAliensAlive()
}