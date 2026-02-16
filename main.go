package main

import (
	"fmt"
	"planets_observer/astro"
	cmd "planets_observer/cli"
	"time"
)

func main() {
	// ParsedDate, err := time.Parse(time.RFC3339, "2025-12-15T19:06:05Z")
	// if err != nil {
	// 	panic(err)
	// }
	// JulianDate := astro.ConvertToJulian(ParsedDate)
	// fmt.Println(JulianDate)
	// timePassedSinceReference := astro.CenturiesPassedSinceReference(JulianDate)
	// fmt.Println(timePassedSinceReference)
	// grid := astro.CreateGrid(40)
	// astro.Plot(grid, 0, 0, 'O')
	// Coordinates := astro.CalculateCoordinatesFromEpoch(timePassedSinceReference)
	// ScaledY, ScaledZ := astro.ScaleCoordinates(Coordinates.X, Coordinates.Y, 10)
	// astro.Plot(grid, ScaledZ, ScaledY, 'e')
	// astro.PrintGrid(grid)
	cmd.Execute()
}

func PrintLocationInfo(T time.Time, locationInfo astro.Location) {
	fmt.Printf("For date %s: \n - Mean longitude: %.6f,\n - Mean anomaly: %.6f,\n - Eccentricity: %.6f,\n - Equation of center: %.6f,\n - True anomaly: %.6f,\n - True longitude: %.6f,\n - Distance: %.6f",
		T, locationInfo.MeanLongitude, locationInfo.MeanAnomaly, locationInfo.Eccentricity, locationInfo.EquationOfCenter, locationInfo.TrueAnomaly, locationInfo.TrueLongitude, locationInfo.Distance)
}
