package astro

import (
	"math"
	"time"
)

const (
	DaysPerYearApprox     = 365.25
	DaysPerMonthCoeff     = 30.6001
	JulianCenturyBaseline = 2451545.0
)

// ConvertToJulian converts a UTC time to Julian Day (JD).
//
// Julian Day is a continuous count of days since
// 4713 BC January 1 at 12:00 UTC.
//
// The returned value includes the fractional day
// and is suitable for astronomical calculations.
func ConvertToJulian(parsedDate time.Time) float64 {
	parsedDate = parsedDate.UTC()

	year := parsedDate.Year()
	month := int(parsedDate.Month())
	day := parsedDate.Day()
	hour := parsedDate.Hour()
	minutes := parsedDate.Minute()
	seconds := parsedDate.Second()

	if month <= 2 {
		year -= 1
		month += 12
	}

	A := year / 100
	B := 2 - A + (A / 4)

	dayFrac := (float64(hour) +
		float64(minutes)/60.0 +
		float64(seconds)/3600.0) / 24.0

	D := float64(day) + dayFrac

	JD := math.Floor(float64(DaysPerYearApprox)*float64(year+4716)) +
		math.Floor(DaysPerMonthCoeff*float64(month+1)) +
		float64(D) + float64(B) - 1524.5
	return JD
}

// CenturiesPassedSinceReference counts the number of Julian centuries since the astronomical reference date
//
// Currently, the astronomical reference date in 1st January 2000, 12PM.
//
// Both the JulianDays value and JulianCenturyBaseline are Julian dates.
//
// The returned value is suitable for astronomical calculations
func CenturiesPassedSinceReference(JulianDays float64) float64 {
	return (JulianDays - JulianCenturyBaseline) / 36525
}
