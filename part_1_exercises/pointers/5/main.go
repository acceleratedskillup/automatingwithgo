package main

import "fmt"

func updateAltitude(altitude *float64) {
	*altitude = 350.5
}

func main() {
	satelliteAltitude := 300.0
	updateAltitude(&satelliteAltitude)
	fmt.Println("Updated Satellite Altitude:", satelliteAltitude)
}
