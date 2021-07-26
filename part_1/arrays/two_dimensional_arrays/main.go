package main

import "fmt"

func main() {

	//list of 2 sensor readings
	listOfReadings := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	firstSensorReadings := listOfReadings[0]
	secondSensorReadings := listOfReadings[1]

	fmt.Println("listOfReadings:", listOfReadings)
	fmt.Println("firstSensorReadings:", firstSensorReadings)
	fmt.Println("secondSensorReadings:", secondSensorReadings)
	fmt.Println("listOfReadings[0][0]:", listOfReadings[0][0])
	fmt.Println("firstSensorReadings[0]:", firstSensorReadings[0])
	firstSensorReadings[0] = 1000
	fmt.Println("listOfReadings[0][0]:", listOfReadings[0][0])
	fmt.Println("firstSensorReadings[0]:", firstSensorReadings[0])

	//to avoid a copy, use the 2D array directly
	listOfReadings[0][0] = 2000
	fmt.Println("listOfReadings[0][0]:", listOfReadings[0][0])
	/*
		//second example of initializing a 2d array

		firstSensorReadings := [3]int{1, 2, 3}
		secondSensorReadings := [3]int{4, 5, 6}
		var listOfReadings [2][3]int
		listOfReadings[0] = firstSensorReadings
		listOfReadings[1] = secondSensorReadings
		fmt.Println("listOfReadings:", listOfReadings)

		//show that the assignment above is a copy of the array
		listOfReadings[0][0] = 1000
		fmt.Println("listOfReadings:", listOfReadings)
		fmt.Println("firstSensorReadings:", firstSensorReadings)
	*/
	var sum int
	for _, singleSensorReadings := range listOfReadings {
		for _, singleReading := range singleSensorReadings {
			sum += singleReading
		}
	}
	fmt.Println("sum:", sum)
}
