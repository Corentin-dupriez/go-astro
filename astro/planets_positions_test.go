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
		{"Test 2026-2-07", 0.2610347951048301, 317.920057},
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
		{"Test Test_2026-02-07", 0.2610347951048301, 0.016698},
	}

	for _, tt := range tests {
		got := CalculateOrbitEccentricity(tt.julianDate)
		if math.Abs(got-tt.want) > 1e-6 {
			t.Errorf("got %.6f, want %.6f", got, tt.want)
		}
	}
}
