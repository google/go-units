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
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

const areaEpsilon = 1e-15

func TestEmptyArea(t *testing.T) {
	var a Area
	if m := a.SquareMeters(); m != 0 {
		t.Errorf("Empty value of Area was %v, want 0", m)
	}
}

func TestAreaConversion(t *testing.T) {
	units := []struct {
		singular, plural string
		unit             Area
		convert          func(Area) float64
	}{
		{"SquareKilometer", "SquareKilometers", SquareKilometer, Area.SquareKilometers},
		{"Hectare", "Hectares", Hectare, Area.Hectares},
		{"SquareMeter", "SquareMeters", SquareMeter, Area.SquareMeters},
		{"SquareCentimeter", "SquareCentimeters", SquareCentimeter, Area.SquareCentimeters},
		{"SquareMillimeter", "SquareMillimeters", SquareMillimeter, Area.SquareMillimeters},
		{"SquareMile", "SquareMiles", SquareMile, Area.SquareMiles},
		{"Acre", "Acres", Acre, Area.Acres},
		{"SquareFoot", "SquareFeet", SquareFoot, Area.SquareFeet},
		{"SquareInch", "SquareInches", SquareInch, Area.SquareInches},
	}

	const mmPerIn = 645.16
	const cmPerIn = 6.4516
	const inPerFt = 144.0
	const ftPerAcre = 43560.0
	const ftPerMile = 27878400.0
	const ftPerM = 1 / 0.09290304
	const acrePerMile = 640.0
	// Areas of various orders of magnitude.  Slices are in the order of units
	// above, with Square km first and square inches last.
	areas := [][]float64{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		// One square millimeter
		{1e-12, 1e-10, 1e-6, 1e-2, 1, 1 / (mmPerIn * inPerFt * ftPerMile), 1 / (mmPerIn * inPerFt * ftPerAcre), 1 / (mmPerIn * inPerFt), 1 / mmPerIn},
		// 2.4 square centimeters, about the size of a postage stamp
		{2.4e-10, 2.4e-8, 0.00024, 2.4, 240, 2.4 / (cmPerIn * inPerFt * ftPerMile), 2.4 / (cmPerIn * inPerFt * ftPerAcre), 2.4 / (cmPerIn * inPerFt), 2.4 / cmPerIn},
		// 10 feet by 12 feet, typical American bedroom size
		{120 / ftPerM / 1e6, 120 / ftPerM / 1e4, 120 / ftPerM, 120 * inPerFt * cmPerIn, 120 * inPerFt * mmPerIn, 120 / ftPerMile, 120 / ftPerAcre, 120, 120 * inPerFt},
		// Luzon, 15th-largest island on Earth
		{109_965, 10_996_500, 109_965_000_000, 109_965 * 1e10, 109_965 * 1e12, 42_457.72386412005, 42_457.72386412005 * acrePerMile, 42_457.72386412005 * ftPerMile, 42_457.72386412005 * ftPerMile * inPerFt},
		// Approximate surface area of Earth
		{509_600_000, 50_960_000_000, 509_600_000 * 1e6, 509_600_000 * 1e10, 509_600_000 * 1e12, 196_757_659.9932304, 196_757_659.9932304 * acrePerMile, 196_757_659.9932304 * ftPerMile, 196_757_659.9932304 * ftPerMile * inPerFt},
	}

	for _, test := range areas {
		for i, from := range units {
			for j, to := range units {
				if got, want := to.convert(Area(test[i])*from.unit), test[j]; !cmp.Equal(got, want, cmpopts.EquateApprox(areaEpsilon, areaEpsilon)) {
					t.Errorf("(Area(%v) * %v).%v = %v, want %v", test[i], from.singular, to.plural, got, want)
				}
			}
		}
	}
}

func TestAreaString(t *testing.T) {
	tests := []struct {
		area Area
		s    string
	}{
		{0 * SquareMeter, "0m^2"},
		{123456.789 * SquareMeter, "123456.789m^2"},
		{1234.56 * SquareCentimeter, "1234.56cm^2"},
		{SquareMillimeter, "1mm^2"},
		{98.7654321 * SquareMillimeter, "98.7654321mm^2"},
		{SquareMile, "2.5899881103360003km^2"},
		// area of Moloka'i
		{673.4 * SquareKilometer, "673.4km^2"},
		// area of Malawi
		{118484 * SquareKilometer, "118484km^2"},
		// area of Earth
		{6371.01 * 6371.01 * math.Pi * 4 * SquareKilometer, "5.1006607311798856e+14m^2"},
	}
	for _, test := range tests {
		if got, want := test.area.String(), test.s; got != want {
			t.Errorf("(%#v).String() = %#v, want %#v", test.area, got, want)
		}
	}
}

func TestAreaGoString(t *testing.T) {
	tests := []struct {
		area Area
		s    string
	}{
		{0 * SquareMeter, "0 * SquareMeter"},
		{123456.789 * SquareMeter, "123456.789 * SquareMeter"},
		{1234.56 * SquareCentimeter, "1234.56 * SquareCentimeter"},
		{SquareMillimeter, "1 * SquareMillimeter"},
		{98.7654321 * SquareMillimeter, "98.7654321 * SquareMillimeter"},
		{SquareMile, "2.5899881103360003 * SquareKilometer"},
		// area of Moloka'i
		{673.4 * SquareKilometer, "673.4 * SquareKilometer"},
		// area of Malawi
		{118484 * SquareKilometer, "118484 * SquareKilometer"},
		// area of Earth
		{6371.01 * 6371.01 * math.Pi * 4 * SquareKilometer, "5.1006607311798856e+14 * SquareMeter"},
	}
	for _, test := range tests {
		if got, want := fmt.Sprintf("%#v", test.area), test.s; got != want {
			t.Errorf("(%#v).GoString() = %#v, want %#v", test.area, got, want)
		}
	}
}
