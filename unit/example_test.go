package unit_test

import (
	"fmt"

	"github.com/google/go-units/unit"
)

func ExampleLength() {
	// Define a length using a constant.
	distance := 500 * unit.Mile

	// Define a length from a variable.
	var altitudeInFeet = 29031
	altitude := unit.Length(altitudeInFeet) * unit.Foot

	// Convert to different units.
	fmt.Printf("I would walk %.0f miles, but also %.1f kilometers.\n", distance.Miles(), distance.Kilometers())
	fmt.Printf("Altitude: %.2f m\n", altitude.Meters())

	// Perform calculations.
	totalDistance := 2 * distance
	fmt.Printf("And after I walk %.0f miles more, that's a total of %.1f nautical miles.\n", distance.Miles(), totalDistance.NauticalMiles())

	// Output:
	// I would walk 500 miles, but also 804.7 kilometers.
	// Altitude: 8848.65 m
	// And after I walk 500 miles more, that's a total of 869.0 nautical miles.
}
