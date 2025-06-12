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

/*
Package unit implements types and functions for working with units of
measurement.

Usage:

Use constants to define literal quantities:

	distanceToSun := 149.6e6 * unit.Kilometer

or cast to create typed quantities from variables:

	var altitudeInFeet = 29031
	altitude := Length(altitudeInFeet) * unit.Foot

Use the conversion functions to convert back to units, as needed:

	altitudeInMeters := altitude.Meters()

Types in this package implement the fmt.GoStringer interface, returning a scaled
value times an SI unit.  For example,

	fmt.Sprintf("%#v", unit.Meter / 4) == "25 * Centimeter"
*/
package unit
