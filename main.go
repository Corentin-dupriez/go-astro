package main

import (
	"fmt"
	"time"
)

func main() {
	parsedDate, err := time.Parse(time.RFC3339, "2026-02-07T19:06:05Z")
	if err != nil {
		return
	}
	julianDate := ConvertToJulian(parsedDate)
	fmt.Println(julianDate)
	timePassedSinceReference := CenturiesPassedSinceReference(julianDate)
	fmt.Println(timePassedSinceReference)
	meanLongitude := CalculateMeanLongitude(timePassedSinceReference)
	fmt.Println(meanLongitude)
}
