package world

import (
	"bufio"
	"fmt"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
)

// Assumption: city names will never a single space " ".

// https://medium.com/@KeithAlpichi/gos-standard-library-by-example-encoding-csv-75f098169822

// https://golang.org/pkg/sync/
// Package sync provides basic synchronization primitives such as mutual exclusion locks. Other than the Once and WaitGroup types, most are intended for use by low-level library routines. Higher-level synchronization is better done via channels and communication.

// Foo north=Bar west=Baz south=Qu-ux
// Bar south=Foo west=Bee

type CityLock struct{ sync.Mutex }

type World struct {
	worldMap*   simple.UndirectedGraph
	//alienLocation map[int]int
	cityIds map[int]string
	invCityIds map[string]int
	cityLocks []*CityLock
}

func New(worldMap simple.UndirectedGraph, cityIds map[int]string, invCityIds map[string]int) World {
	// Create city locks
	cityLocks := make([]*CityLock, worldMap.Nodes().Len())
	for i := 0; i < worldMap.Nodes().Len(); i++ {
		cityLocks[i] = new(CityLock)
	}
	w := World{&worldMap, cityIds, invCityIds, cityLocks}
	return w
}

func getUndirected() graph.Graph {
	g := simple.NewUndirectedGraph()
	g.AddNode(simple.Node(-1))
	for _, e := range []simple.Edge{
		{F: simple.Node(0), T: simple.Node(1)},
		{F: simple.Node(0), T: simple.Node(3)},
		{F: simple.Node(1), T: simple.Node(2)},
	} {
		g.SetEdge(e)
	}
	return g
}

func loadGraph() graph.Graph {

	g := simple.NewUndirectedGraph()
	g.AddNode(simple.Node(-1))
	for _, e := range []simple.Edge{
		{F: simple.Node(0), T: simple.Node(1)},
		{F: simple.Node(0), T: simple.Node(3)},
		{F: simple.Node(1), T: simple.Node(2)},
	} {
		g.SetEdge(e)
	}
	return g
}

// Assumption: I dont assume is a planar graph, nor 4-regular.
// Assumption: I dont assume that if we go south from CityA and we reach CityB then from CityB going north we will reach City,
// maybe or maybe not, routes are just bidirectional in the direction might reach other cities.
func LoadWorld(mapFilename string) World {

	// Assuming undirected graph the Aliens can move in both directions.
	g := simple.NewUndirectedGraph()
	nodeId := 0
	var nodeIds map[int]string
	nodeIds = make(map[int]string)

	var invNodeIds map[string]int
	invNodeIds = make(map[string]int)

	file, err := os.Open(mapFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line," ")
		// Assumming we can have cities with no bridges.
		// Assuming no repetitions of cities, one line per city.
		if len(lineSplit) < 1 {
			continue
		}
		fmt.Println( lineSplit )
		g.AddNode(simple.Node( nodeId ))
		nodeIds[nodeId] = lineSplit[0]
		invNodeIds[lineSplit[0]] = nodeId
		nodeId++

		for _, road := range lineSplit[1:]{
			// Assume = describing road
			if strings.Index(road, "=") == -1 {
				panic(err)
			}
			roadSplit := strings.Split(road,"=")
			// Foo north=Bar west=Baz south=Qu-ux
			fmt.Println( roadSplit )
			if _, ok := invNodeIds[roadSplit[1]]; !ok {
				// New city, Add new city.
				nodeIds[nodeId] = roadSplit[1]
				invNodeIds[roadSplit[1]] = nodeId
				nodeId++
			}
			// Add Edge
			idSrc, idDst := invNodeIds[lineSplit[0]], invNodeIds[roadSplit[1]]
			g.SetEdge( simple.Edge{F: simple.Node(idSrc), T: simple.Node(idDst)} )
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Create city locks
	cityLocks := make([]*CityLock, g.Nodes().Len())
	for i := 0; i < g.Nodes().Len(); i++ {
		cityLocks[i] = new(CityLock)
	}

	world := World{g,nodeIds,invNodeIds, cityLocks}
	return world
}

func (w *World) LockCity(city int) {
	w.cityLocks[city].Lock()
}

func (w *World) UnlockCity(city int) {
	w.cityLocks[city].Unlock()
}

func (w World) NumberOfCities() int {
	return w.worldMap.Nodes().Len() //len(w.cityIds)
}

func (w World) NumberOfRoutes() int {
	return w.worldMap.Edges().Len() //len(w.cityIds)
}

func (w World) NumberOfRoutesOut(cityId int) int {
	neighbors := w.worldMap.From(int64(cityId))
	return neighbors.Len()
}

func (w World) RandomNeighboringCity(cityId int) int {
	neighbors := w.worldMap.From(int64(cityId))
	if neighbors.Len() == 0 {
		return cityId // Same city if no neighboring cities, trapped.
	}
	randNeighborIndex := rand.Intn(neighbors.Len())
	i :=0
	for neighbors.Next() {
		if i == randNeighborIndex {
			break
		}
		i++
	}
	return int(neighbors.Node().ID())
}

func (w *World) DestroyCity(city int) {
	//fmt.Println( "Before destruction: ",w.worldMap )
	w.worldMap.RemoveNode(int64(city))
	//fmt.Println( "Before destruction: ",w.worldMap )
	//fmt.Println( "   " )
}

func (w World) CityName(city int) string {
	return w.cityIds[city]
}

/*
func (w World) CityDestroyed(city int) bool {
	for w.worldMap.Nodes().Next() {
		if
	}
	return true
}
*/
