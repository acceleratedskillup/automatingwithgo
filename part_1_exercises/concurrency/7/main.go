package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DataPacket struct {
	Type  string
	Value float64
}

func sendData(c chan DataPacket) {
	dataTypes := []string{"Temperature", "Pressure", "Humidity"}

	for {
		dataType := dataTypes[rand.Intn(len(dataTypes))]
		value := rand.Float64() * 100
		c <- DataPacket{Type: dataType, Value: value}
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
}

func main() {
	c := make(chan DataPacket, 5)
	go sendData(c)

	for i := 0; i < 10; i++ {
		select {
		case data := <-c:
			fmt.Printf("Received %s data: %.2f\n", data.Type, data.Value)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout: No data received from spacecraft!")
		}
	}
}
