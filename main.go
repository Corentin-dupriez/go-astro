package main

import (
	"fmt"
	"time"
)

func main() {
	ParsedDate, err := time.Parse(time.RFC3339, "2026-02-07T19:06:05Z")
	if err != nil {
		return
	}
	JulianDate := ConvertToJulian(ParsedDate)
	fmt.Println(JulianDate)
	timePassedSinceReference := CenturiesPassedSinceReference(JulianDate)
	fmt.Println(timePassedSinceReference)
	meanLongitude := CalculateMeanLongitude(timePassedSinceReference)
	fmt.Println(meanLongitude)
	location := ComputeLocation(timePassedSinceReference)
	PrintLocationInfo(ParsedDate, location)
}

func PrintLocationInfo(T time.Time, locationInfo Location) {
	fmt.Printf(`For date %s: Mean longitude: %.6f, Mean anomaly: %.6f, Eccentricity: %.6f`, T, locationInfo.meanLongitude, locationInfo.meanAnomaly, locationInfo.eccentricity)
}
