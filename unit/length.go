// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package unit

import (
	"fmt"
	"time"
)

// Length represents a linear dimension measurement in meters as a float64.
type Length float64

// Common length units.
const (
	Meter        Length = 1
	Kilometer           = 1e3 * Meter
	Centimeter          = 1e-2 * Meter
	Millimeter          = 1e-3 * Meter
	Micrometer          = 1e-6 * Meter
	Foot                = 0.3048 * Meter
	Mile                = 5280 * Foot
	Inch                = Foot / 12
	NauticalMile        = 1852 * Meter
)

// PerTime returns a speed from a length travelled in a given amount of time.
func (l Length) PerTime(d time.Duration) Speed {
	return Speed(l.Meters()/d.Seconds()) * MeterPerSecond
}

// Abs returns the length as an absolute value.
func (l Length) Abs() Length {
	if l < 0 {
		return -l
	}
	return l
}

// Meters returns the length in meters.
func (l Length) Meters() float64 { return float64(l) }

// Kilometers returns the length in kilometers.
func (l Length) Kilometers() float64 { return float64(l / Kilometer) }

// Centimeters returns the length in centimeters.
func (l Length) Centimeters() float64 { return float64(l / Centimeter) }

// Millimeters returns the length in millimeters.
func (l Length) Millimeters() float64 { return float64(l / Millimeter) }

// Micrometers returns the length in micrometers.
func (l Length) Micrometers() float64 { return float64(l / Micrometer) }

// Feet returns the length in feet.
func (l Length) Feet() float64 { return float64(l / Foot) }

// Miles returns the length in miles.
func (l Length) Miles() float64 { return float64(l / Mile) }

// Inches returns the length in inches.
func (l Length) Inches() float64 { return float64(l / Inch) }

// NauticalMiles returns the length in nautical miles.
func (l Length) NauticalMiles() float64 { return float64(l / NauticalMile) }

// String returns a string representation of the length in meters.
//
// If possible, the length will be returned with an appropriate SI prefix
// (e.g. 1.2km, 2.3m, 3.4cm, 4.5mm, 5.6µm), otherwise the distance will be
// returned as a scientific representation in meters (e.g. 149.6e+09m). Except
// in the case where the length is 0, leading zeros will always be omitted (e.g
// 0.2km will be returned as "200m", 0.9e9m will be returned as "9e+08m").
func (l Length) String() string {
	value, desc := l.format()
	return fmt.Sprintf("%v%v", value, desc.symbol)
}

func (l Length) GoString() string {
	value, desc := l.format()
	return fmt.Sprintf("%v * %v", value, desc.name)
}

type unitDesc struct {
	length Length
	name   string
	symbol string
}

var (
	kmDesc    = unitDesc{Kilometer, "Kilometer", "km"}
	meterDesc = unitDesc{Meter, "Meter", "m"}

	// unitThresholds contains the thresholds for SI prefixed meter values to be
	// used when returning a string representation of the length.
	unitThresholds = []unitDesc{
		// NOTE: keep in descending order so that format() works correctly.
		kmDesc,
		meterDesc,
		{Centimeter, "Centimeter", "cm"},
		{Millimeter, "Millimeter", "mm"},
		{Micrometer, "Micrometer", "µm"},
	}
)

func (l Length) format() (string, unitDesc) {
	if l.Abs() >= 1000*kmDesc.length {
		// %g instead of %e for variable precision
		return fmt.Sprintf("%g", l/meterDesc.length), meterDesc
	}
	for _, unitThreshold := range unitThresholds {
		if l.Abs() >= unitThreshold.length {
			return fmt.Sprintf("%v", float64(l/unitThreshold.length)), unitThreshold
		}
	}
	// %g instead of %e for variable precision
	return fmt.Sprintf("%g", l/meterDesc.length), meterDesc
}
