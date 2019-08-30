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
)

// Assumption: city names will never a single space " ".

// https://medium.com/@KeithAlpichi/gos-standard-library-by-example-encoding-csv-75f098169822

// https://golang.org/pkg/sync/
// Package sync provides basic synchronization primitives such as mutual exclusion locks. Other than the Once and WaitGroup types, most are intended for use by low-level library routines. Higher-level synchronization is better done via channels and communication.

type World struct {
	worldMap*   simple.UndirectedGraph
	//alienLocation map[int]int
	cityIds map[int]string
	invCityIds map[string]int
}

func New(worldMap simple.UndirectedGraph, cityIds map[int]string, invCityIds map[string]int) World {
	w := World{&worldMap, cityIds, invCityIds}
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

	world := World{g,nodeIds,invNodeIds}
	return world
}

func (w World) NumberOfCities() int {
	return len(w.cityIds)
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

