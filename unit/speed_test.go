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
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

const speedEpsilon = 1e-15

func TestEmptySpeed(t *testing.T) {
	var s Speed
	if m := s.MetersPerSecond(); m != 0 {
		t.Errorf("Empty value of Speed was %v, want 0", m)
	}
}

func TestSpeed(t *testing.T) {
	units := []struct {
		singular_name, plural_name string
		unit                       Speed
		convert                    func(Speed) float64
	}{
		{"MeterPerSecond", "MetersPerSecond", MeterPerSecond, Speed.MetersPerSecond},
		{"FootPerSecond", "FeetPerSecond", FootPerSecond, Speed.FeetPerSecond},
		{"KilometerPerHour", "KilometersPerHour", KilometerPerHour, Speed.KilometersPerHour},
		{"MilePerHour", "MilesPerHour", MilePerHour, Speed.MilesPerHour},
		{"Knot", "Knots", Knot, Speed.Knots},
	}

	// Example speeds modified from
	// http://en.wikipedia.org/wiki/Speed
	//
	// Keep values in the same order as the units above; they are the speeds in
	// the corresponding units. First value is in meters per second, last one in
	// miles per hour.
	speeds := [][]float64{
		{0, 0, 0, 0, 0},
		// Speed of a common snail (1 millimeter per second)
		{0.001, 0.0032808398950131233, 0.0036, 0.002236936292054402, 0.0019438444924406047},
		// British National Speed Limit (single carriageway)
		{26.822400000000002, 88, 96.56064, 60, 52.13857451403888},
		// Speed of light in vacuum (exactly 299,792,458 m/s).
		{299792458, 983571056.4304461, 1079252848.8, 670616629.384395, 582749918.3585312},
	}

	for _, test := range speeds {
		for i, from := range units {
			for j, to := range units {
				if got, want := to.convert(Speed(test[i])*from.unit), test[j]; !cmp.Equal(got, want, cmpopts.EquateApprox(speedEpsilon, speedEpsilon)) {
					t.Errorf(
						"(Speed(%v) * %v).%v = %v, want %v",
						test[i], from.singular_name, to.plural_name, got, want)
				}
			}
		}
	}
}

var speedStringTests = []struct {
	in Speed
	s  string // String()
	gs string // GoString()
}{
	{0 * MeterPerSecond, "0 m/s", "0 * MeterPerSecond"},
	{0.001 * MeterPerSecond, "0.001 m/s", "0.001 * MeterPerSecond"},
	{299792458 * MeterPerSecond, "2.99792458e+08 m/s", "2.99792458e+08 * MeterPerSecond"},
}

func TestSpeedString(t *testing.T) {
	for _, test := range stringTests {
		if got, want := test.in.String(), test.s; got != want {
			t.Errorf("(%#v).String() = %#v, want %#v", test.in, got, want)
		}
	}
}

func TestSpeedGoString(t *testing.T) {
	for _, test := range stringTests {
		if got, want := test.in.GoString(), test.gs; got != want {
			t.Errorf("(%#v).GoString() = %#v, want %#v", test.in, got, want)
		}
	}
}
