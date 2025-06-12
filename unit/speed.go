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

// Speed represents the magnitude of velocity of an object in meters per second
// as a float64.
type Speed float64

// Common durations.
const (
	MeterPerSecond   Speed = 1
	KilometerPerHour       = Speed(Kilometer/Meter) * MeterPerSecond / hourInSeconds
	FootPerSecond          = Speed(Foot/Meter) * MeterPerSecond
	MilePerHour            = Speed(Mile/Foot) * FootPerSecond / hourInSeconds
	Knot                   = Speed(NauticalMile/Kilometer) * KilometerPerHour

	// Not really a speed, used for conversion above.
	hourInSeconds = Speed(time.Hour / time.Second)
)

// MetersPerSecond returns the speed in meters per second.
func (s Speed) MetersPerSecond() float64 {
	return float64(s / MeterPerSecond)
}

// KilometersPerHour returns the speed in kilometers per hour.
func (s Speed) KilometersPerHour() float64 {
	return float64(s / KilometerPerHour)
}

// MilesPerHour returns the speed in miles per hour.
func (s Speed) MilesPerHour() float64 {
	return float64(s / MilePerHour)
}

// FeetPerSecond returns the speed in feet per second.
func (s Speed) FeetPerSecond() float64 {
	return float64(s / FootPerSecond)
}

// Knots returns the speed in nautical miles per hour.
func (s Speed) Knots() float64 {
	return float64(s / Knot)
}

// String returns a string representation of the speed in meters per second
// using compact number syntax. For example:
//
//	"0 m/s"
//	"0.001 m/s"
//	"2.99792458e+08 m/s"
func (s Speed) String() string {
	return fmt.Sprintf("%g m/s", s.MetersPerSecond())
}

// GoString returns a Go syntax expression of the speed. For example:
//
//	"0 * MeterPerSecond"
//	"0.001 * MeterPerSecond"
//	"2.99792458e+08 * MeterPerSecond"
func (s Speed) GoString() string {
	return fmt.Sprintf("%g * MeterPerSecond", s.MetersPerSecond())
}
