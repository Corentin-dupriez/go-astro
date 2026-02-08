package main

import "math"

func CalculateMeanLongitude(JulianCentury float64) float64 {
	mean := 280.46646 +
		36000.76983*JulianCentury +
		0.0003032*(JulianCentury*JulianCentury)
	if mean > 360 {
		mean = math.Mod(mean, 360)
	}
	return mean
}
