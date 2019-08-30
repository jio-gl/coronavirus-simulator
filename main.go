package main

import (
	"AlienInvasion/invasion"
	"fmt"
	"os"
	"strconv"
)


func main() {

	if len(os.Args[1:]) < 3 {
		fmt.Println("Usage: %s inputfile.map alienPopulation numberOfSteps [-s]   (-s for synchronic invasion, not asynchornic)", os.Args[0])
		os.Exit(0)
	}

	// alienPopulation
	i2, _ := strconv.Atoi(os.Args[2])
	mainInvasion := invasion.New(os.Args[1], i2)

	// numberOfSteps
	i3, _ := strconv.Atoi(os.Args[3])

	// Use sync version or not
	syncVersion := len(os.Args[1:]) > 3 && os.Args[len(os.Args)-1]  == "-s"
	if syncVersion {
		mainInvasion.RunInvasionSync( i3 )
	} else {
		mainInvasion.RunInvasionAsync(i3)
	}

}