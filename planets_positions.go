package main

import "math"

func CalculateMeanLongitude(T float64) float64 {
	mean := 280.46646 +
		36000.76983*T +
		0.0003032*(T*T)
	if mean > 360 {
		mean = math.Mod(mean, 360)
	}
	return mean
}

func CalculateMeanAnomaly(T float64) float64 {
	return 357.52911 + 35999.05029*T - 0.0001537*(T*T)
}

func CalculateOrbitEccentricity(T float64) float64 {
	return 0.016708634 - 0.000042037*T - 0.0000001267*(T*T)
}

func CalculateEquationOfCenter(T float64, M float64) float64 {
	C := (1.914602-0.004817*T-0.000014*(T*T))*math.Sin(M) +
		(0.019993-0.000101*T)*math.Sin(2*M) +
		0.000289*math.Sin(3*M)
	return C
}

func CalculateTrueLongitude(meanLongitude float64, equationOfCenter float64) float64 {
	return meanLongitude + equationOfCenter
}

func CalculateTrueAnomaly(meanAnomaly float64, equationOfCenter float64) float64 {
	return meanAnomaly + equationOfCenter
}

func CalculateSunDistance(eccentricity float64, trueAnomaly float64) float64 {
	return (1.000001018 * (1 - (eccentricity * eccentricity))) / (1 + eccentricity*math.Cos(trueAnomaly))
}

type Location struct {
	meanLongitude    float64
	meanAnomaly      float64
	eccentricity     float64
	equationOfCenter float64
	trueLongitude    float64
	trueAnomaly      float64
	distance         float64
}

func ComputeLocation(T float64) Location {
	meanLongitude := CalculateMeanLongitude(T)
	meanAnomaly := CalculateMeanAnomaly(T)
	eccentricity := CalculateOrbitEccentricity(T)
	equationOfCenter := CalculateEquationOfCenter(T, meanAnomaly)
	trueLongitude := CalculateTrueLongitude(meanLongitude, equationOfCenter)
	trueAnomaly := CalculateTrueAnomaly(meanAnomaly, equationOfCenter)
	distance := CalculateSunDistance(eccentricity, trueAnomaly)

	return Location{
		meanLongitude:    meanLongitude,
		meanAnomaly:      meanAnomaly,
		eccentricity:     eccentricity,
		equationOfCenter: equationOfCenter,
		trueLongitude:    trueLongitude,
		trueAnomaly:      trueAnomaly,
		distance:         distance,
	}
}
