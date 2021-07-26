package progressupdate

import (
	"fmt"
	"time"
)

const secondsInMinute = 60

func PrintProgress(fileCount int, startTime time.Time) {
	endTime := time.Now()

	// Calculate duration
	duration := endTime.Sub(startTime)

	totalSeconds := int(duration.Seconds())

	// Check if totalSeconds is 59 or less
	if totalSeconds <= secondsInMinute-1 {
		// Display only seconds
		fmt.Printf("\rprocessing %d files, elapsed time: %d seconds", fileCount, totalSeconds)
	} else {
		// Calculate minutes and seconds
		minutes := totalSeconds / secondsInMinute
		seconds := totalSeconds % secondsInMinute

		// Display the duration in minutes and seconds
		fmt.Printf("\rprocessing %d files, elapsed time: %d minutes and %d seconds", fileCount, minutes, seconds)
	}

}
