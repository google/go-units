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
)

// Temperature represents a thermodynamic temperature measurement in Kelvin as
// a float64.
type Temperature float64

// Common temperatures.
const (
	Kelvin  Temperature = 1
	Rankine             = 5.0 / 9.0 * Kelvin
)

// Degrees celsius and fahrenheit 0-points.
const (
	celsius0    Temperature = 273.15 * Kelvin
	fahrenheit0             = 459.67 * Rankine
)

// temperatureFromKelvin returns a Temperature from a measurement in Kelvin.
// Unexported and used only in tests; clients should use the n * Kelvin literal
// syntax.
func temperatureFromKelvin(k float64) Temperature {
	return Temperature(k) * Kelvin
}

// TemperatureFromDegreesCelsius returns a Temperature from a measurement in
// degrees Celsius.
func TemperatureFromDegreesCelsius(c float64) Temperature {
	return Temperature(c)*Kelvin + celsius0
}

// TemperatureFromDegreesFahrenheit returns a Temperature from a measurement in
// degrees Fahrenheit.
func TemperatureFromDegreesFahrenheit(f float64) Temperature {
	return Temperature(f)*Rankine + fahrenheit0
}

// temperatureFromDegreesRankine returns a Temperature from a measurement in
// degrees Rankine.
// Unexported and used only in tests; clients should use the n * Rankine literal
// syntax.
func temperatureFromDegreesRankine(r float64) Temperature {
	return Temperature(r) * Rankine
}

// Kelvin returns the temperature in Kelvin.
func (t Temperature) Kelvin() float64 {
	return float64(t / Kelvin)
}

// DegreesCelsius returns the temperature in degrees Celsius.
func (t Temperature) DegreesCelsius() float64 {
	return float64((t - celsius0) / Kelvin)
}

// DegreesFahrenheit returns the temperature in degrees Fahrenheit.
func (t Temperature) DegreesFahrenheit() float64 {
	return float64((t - fahrenheit0) / Rankine)
}

// DegreesRankine returns the temperature in degrees Rankine.
func (t Temperature) DegreesRankine() float64 {
	return float64(t / Rankine)
}

// String returns a string representation of the temperature in Kelvin using
// compact number syntax (e.g. "294.15 K", "5778 K", "1.57e+07 K").
func (t Temperature) String() string {
	return fmt.Sprintf("%g K", t.Kelvin())
}

// GoString returns a Go syntax expression of the temperature (e.g.
// "294.15 * Kelvin", "5778 * Kelvin", "1.5e+07 * Kelvin").
func (t Temperature) GoString() string {
	return fmt.Sprintf("%v * Kelvin", t.Kelvin())
}
