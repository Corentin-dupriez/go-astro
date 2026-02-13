package astro

import (
	"math"
	"testing"
)

func TestCalculateMeanLongitude(T *testing.T) {
	tests := []struct {
		name       string
		julianDate float64
		want       float64
	}{
		{"Test J2000", 0.0, 280.46646},
		{"Test 2026-02-07", 0.2610347951048301, 317.920057},
	}

	for _, tt := range tests {
		T.Run(tt.name, func(t *testing.T) {
			got := CalculateMeanLongitude(tt.julianDate)
			if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("got %.6f, want %.6f", got, tt.want)
			}
		})
	}
}

func TestCalculateMeanAnomaly(t *testing.T) {
	tests := []struct {
		name       string
		julianDate float64
		want       float64
	}{
		{"Test J2000", 0.0, 357.52911},
		{"Test 2026-02-07", 0.2610347951048301, 34.533816},
	}
	for _, tt := range tests {
		got := CalculateMeanAnomaly(tt.julianDate)
		if math.Abs(got-tt.want) > 1e-6 {
			t.Errorf("got %.6f, want %.6f", got, tt.want)
		}
	}
}

func TestCalculateOrbitEccentricity(t *testing.T) {
	tests := []struct {
		name       string
		julianDate float64
		want       float64
	}{
		{"Test J2000", 0.0, 0.016708634},
		{"Test Test 2026-02-07", 0.2610347951048301, 0.016698},
	}

	for _, tt := range tests {
		got := CalculateOrbitEccentricity(tt.julianDate)
		if math.Abs(got-tt.want) > 1e-6 {
			t.Errorf("got %.6f, want %.6f", got, tt.want)
		}
	}
}

func TestCalculateEquationOfCenter(t *testing.T) {
	tests := []struct {
		name        string
		julianDate  float64
		meanAnomaly float64
		want        float64
	}{
		{"Test J2000", 0.0, 357.52911, -0.084301},
		{"Test 2026-02-07", 0.2610347951048301, 34.533816, 1.103590},
	}

	for _, tt := range tests {
		got := CalculateEquationOfCenter(tt.julianDate, tt.meanAnomaly)
		if math.Abs(got-tt.want) > 1e-6 {
			t.Errorf("got %.6f, want %.6f", got, tt.want)
		}
	}
}

func TestCalculateTrueLongitude(t *testing.T) {
	tests := []struct {
		name             string
		meanLongitude    float64
		equationOfCenter float64
		want             float64
	}{
		{"Test true longitude", 1.0, 1.0, 2.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateTrueLongitude(tt.meanLongitude, tt.equationOfCenter)
			if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("got %.6f, want %.6f", got, tt.want)
			}
		})
	}
}

func TestCalculateTrueAnomaly(t *testing.T) {
	tests := []struct {
		name             string
		meanAnomaly      float64
		equationOfCenter float64
		want             float64
	}{
		{"Test true anomaly", 1.0, 1.0, 2.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateTrueAnomaly(tt.meanAnomaly, tt.equationOfCenter)
			if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("got %.6f, want %.6f", got, tt.want)
			}
		})
	}
}

func TestCalculateSunDistance(t *testing.T) {
	tests := []struct {
		name         string
		eccentricity float64
		trueAnomaly  float64
		want         float64
	}{
		{"Test J2000", 0.016708634, 357.52911, 0.986238},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateSunDistance(tt.eccentricity, tt.trueAnomaly)
			if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("got %.6f, want %.6f", got, tt.want)
			}
		})
	}
}
