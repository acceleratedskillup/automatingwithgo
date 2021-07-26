package main

import "fmt"

func main() {
	melodies := []string{"Dawn's Whisper", "Twilight's Echo", "Midday's Hum"}
	melodies = append(melodies, "Night's Lullaby")
	fmt.Println(melodies)
}
