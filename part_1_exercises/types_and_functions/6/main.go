package main

import "fmt"

type Planet struct {
	name            string
	distanceFromSun float64
	orbitalPeriod   float64
}

type SolarSystem struct {
	planets []Planet
}

func (ss *SolarSystem) addPlanet(p Planet) {
	ss.planets = append(ss.planets, p)
}

func (ss *SolarSystem) displayPlanets() {
	for _, planet := range ss.planets {
		fmt.Println("Name:", planet.name)
		fmt.Println("Distance from Sun:", planet.distanceFromSun, "million km")
		fmt.Println("Orbital Period:", planet.orbitalPeriod, "days")
		fmt.Println("-----")
	}
}

func (ss *SolarSystem) simulateOrbits() {
	for _, planet := range ss.planets {
		fmt.Printf("%s completes its orbit in %f days.\n", planet.name, planet.orbitalPeriod)
	}
}

func main() {
	earth := Planet{name: "Earth", distanceFromSun: 149.6, orbitalPeriod: 365.25}
	mars := Planet{name: "Mars", distanceFromSun: 227.9, orbitalPeriod: 687.0}

	system := &SolarSystem{}
	system.addPlanet(earth)
	system.addPlanet(mars)

	system.displayPlanets()
	system.simulateOrbits()
}
