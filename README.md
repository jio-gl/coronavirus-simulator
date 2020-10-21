# AlienInvasion: 

Simulating the invasion by an alien coronavirus on a social network described by persons and connections. Synchronic and Asynchronic versions.

Features:

* Synchronic version, all fights happen at once.
* Asynchronic version, using concurrency.
* Command-line arguments.
* GoNum library for handling graph or cities and routes.
* Random initial locations.
* Custom data format for describing worlds with cities and routes.
* Testing.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine 
for development and testing purposes. See deployment for notes on how to deploy the project 
on a live system.

```
$ git clone https://github.com/manfr3d/AlienInvasion
```


### Prerequisites

What things you need to install the software and how to install them

```
$ cd AlienInvasion
$ go get ./...
```

### Running

A step by step series of examples that tell you how to get a development env running

Say what the step will be

```
$ go run main.go
Usage: executablename inputfile.map alienPopulation numberOfSteps [-s]   (-s for synchronic invasion, not asynchronic)
```

You can change the map in "maps/big.map" if you want to change the planet description.

## Running the tests

This is how to run the automated tests:

```
$ get test ./...
```

## Authors

* **Jose I Orlicki** 

## License

This project is licensed under the Apache License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* Golang Guru.
* Dinning Philosophers.
* Family.
* etc
