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
	cartesianCoordinates := CalculateCartesianCoordinates(location.trueLongitude, location.distance)
	fmt.Println(cartesianCoordinates)
}

func PrintLocationInfo(T time.Time, locationInfo Location) {
	fmt.Printf("For date %s: \n - Mean longitude: %.6f,\n - Mean anomaly: %.6f,\n - Eccentricity: %.6f,\n - Equation of center: %.6f,\n - True anomaly: %.6f,\n - True longitude: %.6f,\n - Distance: %.6f",
		T, locationInfo.meanLongitude, locationInfo.meanAnomaly, locationInfo.eccentricity, locationInfo.equationOfCenter, locationInfo.trueAnomaly, locationInfo.trueLongitude, locationInfo.distance)
}
