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

import "fmt"

// Area represents a two-dimensional measurement in square meters as a float64.
type Area float64

// Common area units.
const (
	SquareMeter      Area = 1
	SquareKilometer       = 1e6 * SquareMeter
	Hectare               = 1e4 * SquareMeter
	SquareCentimeter      = 1e-4 * SquareMeter
	SquareMillimeter      = 1e-6 * SquareMeter
	SquareFoot            = 0.3048 * 0.3048 * SquareMeter
	SquareMile            = 5280 * 5280 * SquareFoot
	Acre                  = 66 * 660 * SquareFoot
	SquareInch            = SquareFoot / (12 * 12)
)

// Abs returns the area as an absolute value.
func (a Area) Abs() Area {
	if a < 0 {
		return -a
	}
	return a
}

// SquareMeters returns the area in square meters.
func (a Area) SquareMeters() float64 { return float64(a) }

// SquareKilometers returns the area in square kilometers.
func (a Area) SquareKilometers() float64 { return float64(a / SquareKilometer) }

// Hectares returns the area in hectares.
func (a Area) Hectares() float64 { return float64(a / Hectare) }

// SquareCentimeters returns the area in square centimeters.
func (a Area) SquareCentimeters() float64 { return float64(a / SquareCentimeter) }

// SquareMillimeters returns the area in square millimeters.
func (a Area) SquareMillimeters() float64 { return float64(a / SquareMillimeter) }

// SquareFeet returns the area in square feet.
func (a Area) SquareFeet() float64 { return float64(a / SquareFoot) }

// SquareMiles returns the area in square miles.
func (a Area) SquareMiles() float64 { return float64(a / SquareMile) }

// Acres returns the area in acres.
func (a Area) Acres() float64 { return float64(a / Acre) }

// SquareInches returns the area in square inches.
func (a Area) SquareInches() float64 { return float64(a / SquareInch) }

// String returns a string representation of the area in square meters.
//
// If possible, the area will be formatted with an appropriate SI prefix,
// e.g. 1.2km^2, 2.3m^2, 3.4cm^2, 4.5mm^2.  Otherwise the distance will be
// formatted as a scientific representation in square meters, e.g. 123.45e+09m^2.
// Except in the case where the length is 0, leading zeroes will always be
// omitted, e.g. 0.2km^2 will be returned as "200000m^2" and 0.9mm^2 will be
// returned as "9e-7m^2".
func (a Area) String() string {
	value, desc := a.format()
	return fmt.Sprintf("%v%v", value, desc.symbol)
}

// GoString returns a Go syntax expression of the area.  For example:
//
// "0 * SquareMeter"
// "1234.56 * SquareCentimeter"
// "118484 * SquareKilometer"
// "5.10066e+15 * SquareMeter"
func (a Area) GoString() string {
	value, desc := a.format()
	return fmt.Sprintf("%v * %v", value, desc.name)
}

func (a Area) format() (string, areaUnitDesc) {
	if a.Abs() > 1_000_000*squareKmDesc.area {
		return fmt.Sprintf("%g", a.SquareMeters()), squareMeterDesc
	}
	for _, u := range areaUnitThresholds {
		if a.Abs() >= u.area {
			return fmt.Sprintf("%v", float64(a/u.area)), u
		}
	}
	return fmt.Sprintf("%g", a.SquareMeters()), squareMeterDesc
}

type areaUnitDesc struct {
	area   Area
	name   string
	symbol string
}

var (
	squareKmDesc    = areaUnitDesc{SquareKilometer, "SquareKilometer", "km^2"}
	squareMeterDesc = areaUnitDesc{SquareMeter, "SquareMeter", "m^2"}

	// areaUnitThresholds contains the thresholds for SI prefixed area values to
	// be used when returning a string representation of an area.
	areaUnitThresholds = []areaUnitDesc{
		squareKmDesc,
		squareMeterDesc,
		{SquareCentimeter, "SquareCentimeter", "cm^2"},
		{SquareMillimeter, "SquareMillimeter", "mm^2"},
	}
)
