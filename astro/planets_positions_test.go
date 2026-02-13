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
	}
	for _, tt := range tests {
		got := CalculateMeanAnomaly(tt.julianDate)
		if math.Abs(got-tt.want) > 1e-6 {
			t.Errorf("got %.6f, want %.6f", got, tt.want)
		}
	}
}
