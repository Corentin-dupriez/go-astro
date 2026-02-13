package main

import (
	"fmt"
	"planets_observer/astro"
	"time"
)

func main() {
	ParsedDate, err := time.Parse(time.RFC3339, "2026-02-07T19:06:05Z")
	if err != nil {
		return
	}
	JulianDate := astro.ConvertToJulian(ParsedDate)
	fmt.Println(JulianDate)
	timePassedSinceReference := astro.CenturiesPassedSinceReference(JulianDate)
	fmt.Println(timePassedSinceReference)
	meanLongitude := astro.CalculateMeanLongitude(timePassedSinceReference)
	fmt.Println(meanLongitude)
	location := astro.ComputeLocation(timePassedSinceReference)
	PrintLocationInfo(ParsedDate, location)
	cartesianCoordinates := astro.CalculateCartesianCoordinates(location.TrueLongitude, location.Distance)
	fmt.Println(cartesianCoordinates)
}

func PrintLocationInfo(T time.Time, locationInfo astro.Location) {
	fmt.Printf("For date %s: \n - Mean longitude: %.6f,\n - Mean anomaly: %.6f,\n - Eccentricity: %.6f,\n - Equation of center: %.6f,\n - True anomaly: %.6f,\n - True longitude: %.6f,\n - Distance: %.6f",
		T, locationInfo.MeanLongitude, locationInfo.MeanAnomaly, locationInfo.Eccentricity, locationInfo.EquationOfCenter, locationInfo.TrueAnomaly, locationInfo.TrueLongitude, locationInfo.Distance)
}
