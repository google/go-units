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
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

const lengthEpsilon = 1e-15

func TestEmptyLength(t *testing.T) {
	var l Length
	if m := l.Meters(); m != 0 {
		t.Errorf("Empty value of Length was %v, want 0", m)
	}
}

func TestPerTime(t *testing.T) {
	tests := []struct {
		l    Length
		d    time.Duration
		want Speed
	}{
		// Speed of a common snail (1 millimeter per second).
		{1 * Millimeter, 1 * time.Second, 0.001 * MeterPerSecond},
		// Taipei 101 observatory elevator
		{1010 * Meter, 1 * time.Minute, (16.5 + 1.0/3) * MeterPerSecond},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v per %v", test.l, test.d), func(t *testing.T) {
			if got := test.l.PerTime(test.d); !cmp.Equal(got, test.want, cmpopts.EquateApprox(0, lengthEpsilon)) {
				t.Errorf("(%#v).PerTime(%#v) = %#v, want %#v", test.l, test.d, got, test.want)
			}
		})
	}
}

func TestLengthConversion(t *testing.T) {
	units := []struct {
		singular, plural string
		unit             Length
		convert          func(Length) float64
	}{
		{"Kilometer", "Kilometers", Kilometer, Length.Kilometers},
		{"Meter", "Meters", Meter, Length.Meters},
		{"Centimeter", "Centimeters", Centimeter, Length.Centimeters},
		{"Millimeter", "Millimeters", Millimeter, Length.Millimeters},
		{"Micrometer", "Micrometers", Micrometer, Length.Micrometers},
		{"Mile", "Miles", Mile, Length.Miles},
		{"Foot", "Feet", Foot, Length.Feet},
		{"Inch", "Inches", Inch, Length.Inches},
		{"NauticalMile", "NauticalMiles", NauticalMile, Length.NauticalMiles},
	}

	// Keep values in the same order as the units above; they are the lengths in
	// the corresponding units. First value is in kilometers, last one in inches.
	lengths := [][]float64{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1e-3, 1, 1e2, 1e3, 1e6, 6.2137119223733391e-4, 3.280839895013123, 39.37007874015748, 0.0005399568034557236},
		{-3.048e-4, -0.3048, -30.48, -304.8, -3.048e5, -1.893939393939394e-4, -1, -12, -0.00016457883369330455},
	}

	for _, test := range lengths {
		for i, from := range units {
			for j, to := range units {
				t.Run(fmt.Sprintf("%0.f from %v to %v", test[i], from.singular, to.plural), func(t *testing.T) {
					if got, want := to.convert(Length(test[i])*from.unit), test[j]; !cmp.Equal(got, want, cmpopts.EquateApprox(lengthEpsilon, lengthEpsilon)) {
						t.Errorf(
							"(Length(%v) * %v).%v = %v, want %v",
							test[i], from.singular, to.plural, got, want)
					}
				})
			}
		}
	}
}

var stringTests = []struct {
	in Length
	s  string // String()
	gs string // GoString()
}{
	{1e6 * Meter, "1e+06m", "1e+06 * Meter"},
	{9.99e5 * Meter, "999km", "999 * Kilometer"},
	{1e3 * Meter, "1km", "1 * Kilometer"},
	{999.99 * Meter, "999.99m", "999.99 * Meter"},
	{1 * Meter, "1m", "1 * Meter"},
	{0.998 * Meter, "99.8cm", "99.8 * Centimeter"},
	{0.01 * Meter, "1cm", "1 * Centimeter"},
	{0.0099 * Meter, "9.9mm", "9.9 * Millimeter"},
	{0.001 * Meter, "1mm", "1 * Millimeter"},
	{9.999e-4 * Meter, "999.9µm", "999.9 * Micrometer"},
	{1e-6 * Meter, "1µm", "1 * Micrometer"},
	{9.99e-7 * Meter, "9.99e-07m", "9.99e-07 * Meter"},
}

func TestLengthString(t *testing.T) {
	for _, test := range stringTests {
		t.Run(fmt.Sprintf("%v", test.s), func(t *testing.T) {
			if got, want := test.in.String(), test.s; got != want {
				t.Errorf("Length(%v).String() = %q, want %q", test.in, got, want)
			}
		})
	}
	for _, test := range stringTests {
		t.Run(fmt.Sprintf("-%v", test.s), func(t *testing.T) {
			if got, want := (-test.in).String(), "-"+test.s; got != want {
				t.Errorf("Length(%v).String() = %q, want %q", test.in, got, want)
			}
		})
	}
}

func TestLengthGoString(t *testing.T) {
	for _, test := range stringTests {
		t.Run(fmt.Sprintf("%v", test.gs), func(t *testing.T) {
			if got, want := test.in.GoString(), test.gs; got != want {
				t.Errorf("Length(%v).GoString() = %q, want %q", test.in, got, want)
			}
		})
	}
	for _, test := range stringTests {
		t.Run(fmt.Sprintf("-%v", test.gs), func(t *testing.T) {
			if got, want := (-test.in).GoString(), "-"+test.gs; got != want {
				t.Errorf("Length(%v).GoString() = %q, want %q", test.in, got, want)
			}
		})
	}
}
