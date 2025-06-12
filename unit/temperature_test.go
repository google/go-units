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
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// 1°C to Kelvin conversion is two orders of magnitude, so use a larger epsilon
// value than the length-derived units.
const tempEpsilon = 1e-13

func TestTemperatures(t *testing.T) {
	units := []struct {
		name    string
		unit    func(float64) Temperature
		convert func(Temperature) float64
	}{
		{"Kelvin", temperatureFromKelvin, Temperature.Kelvin},
		{"DegreesCelsius", TemperatureFromDegreesCelsius, Temperature.DegreesCelsius},
		{"DegreesFahrenheit", TemperatureFromDegreesFahrenheit, Temperature.DegreesFahrenheit},
		{"DegreesRankine", temperatureFromDegreesRankine, Temperature.DegreesRankine},
	}

	// Example temperatures modified from
	// http://en.wikipedia.org/wiki/Temperature_conversion_formulas
	//
	// Keep values in the same order as the units above; they are the
	// temperatures in the corresponding units. First value is in Kelvin, last
	// one in degrees Rankine.
	temperatures := [][]float64{
		// Absolute zero
		{0, -273.15, -459.67, 0},
		// Lowest recorded earth surface temperature on Earth
		{183.95, -89.2, -128.56, 331.11},
		// Fahrenheit's ice/salt mixture
		{255.37222222222223, -17.777777777777743, 0, 459.67},
		// Ice melts (at standard pressure)
		{273.15, 0, 32, 491.67},
		// Triple point of water
		{273.16, 0.01, 32.018, 491.688},
		// Average surface temperature on Earth
		{288.15, 15, 59, 518.67},
		// Average human body temperature
		{309.95, 36.8, 98.24, 557.91},
		// Highest recorded surface temperature on Earth
		{331.15, 58, 136.4, 596.07},
		// Water boils (at standard pressure)
		{373.1339, 99.9839, 211.97102, 671.64102},
		// Titanium melts
		{1941, 1667.85, 3034.13, 3493.80},
		// The surface of the Sun
		{5778, 5504.85, 9940.73, 10400.4},
	}

	for _, test := range temperatures {
		for i, from := range units {
			for j, to := range units {
				t.Run(fmt.Sprintf("%0.f K from %v to %v", test[0], from.name, to.name), func(t *testing.T) {
					if got, want := to.convert(from.unit(test[i])), test[j]; !cmp.Equal(got, want, cmpopts.EquateApprox(tempEpsilon, tempEpsilon)) {
						t.Errorf(
							"TemperatureFrom%v(%#v).%v() = %#v, want %#v",
							from.name, test[i], to.name, got, want)
					}
				})
			}
		}
	}
}

var temperatureStringTests = []struct {
	in Temperature
	s  string // String()
	gs string // GoString()
}{
	{0 * Kelvin, "0 K", "0 * Kelvin"},
	{294.15 * Kelvin, "294.15 K", "294.15 * Kelvin"},
	{459.67 * Rankine, "255.37222222222223 K", "255.37222222222223 * Kelvin"},
	{5778 * Kelvin, "5778 K", "5778 * Kelvin"},
	{15.7e6 * Kelvin, "1.57e+07 K", "1.57e+07 * Kelvin"},
}

func TestTemperatureString(t *testing.T) {
	for _, test := range stringTests {
		t.Run(test.s, func(t *testing.T) {
			if got, want := test.in.String(), test.s; got != want {
				t.Errorf("(%#v).String() = %#v, want %#v", test.in, got, want)
			}
		})
	}
}

func TestTemperatureGoString(t *testing.T) {
	for _, test := range stringTests {
		t.Run(test.gs, func(t *testing.T) {
			if got, want := test.in.GoString(), test.gs; got != want {
				t.Errorf("(%#v).GoString() = %#v, want %#v", test.in, got, want)
			}
		})
	}
}
