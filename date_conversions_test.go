package main

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestConvertToJulian(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want float64
	}{
		{"convert 2026-01-01", time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC), 2461042.000000},
		{"convert 2000-01-01", time.Date(2000, time.January, 1, 12, 0, 0, 0, time.UTC), 2451545.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertToJulian(tt.date)
			if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("got %.6f, want %.6f", got, tt.want)
			}
		})
	}
}

func TestTimeElapsedSinceJ2000(t *testing.T) {
	tests := []struct {
		name       string
		julianDate float64
		want       float64
	}{
		{"J2000 Epoch", 2451545.000000, 0},
		{"2461042.000000", 2461042.000000, 0.26001369},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CenturiesPassedSinceReference(tt.julianDate)
			if math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("got %.6f, want %.6f", got, tt.want)
			}
		})
	}
}

func ExampleConvertToJulian() {
	t := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println(ConvertToJulian(t))
	// Output: 2.451545e+06
}
