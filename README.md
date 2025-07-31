# Go Units

[![Go Reference](https://pkg.go.dev/badge/github.com/google/go-units.svg)](https://pkg.go.dev/github.com/google/go-units) [![Go Build and Test](https://github.com/google/go-units/actions/workflows/go.yml/badge.svg)](https://github.com/google/go-units/actions/workflows/go.yml)

This Go library represents physical units like length, area, speed, and
temperature as typed floating point numbers. Methods to convert between metric and [US customary units](https://en.wikipedia.org/wiki/United_States_customary_units)
of measurement are provided on each type. Additional units can be
created by multiplying by const values and variables can be converted by
casting. See [this example snippet](unit/example_test.go) for the exact syntax.

Unit types are represented as a `float64` in the base SI measure: meters, square
meters, meters per second, and Kelvin. Working with very large or very small
values relative to the base SI unit will lose some precision: the distance from
the Sun to Alpha Centauri in meters is near the precision limit of 64-bit
floating point numbers, and the distance to the edge of universe cannot be
represented as a `unit.Length`. Likewise, subatomic measurements will not be
precise.

This is not an officially supported Google product. This project is not eligible
for the
[Google Open Source Software Vulnerability Rewards Program](https://bughunters.google.com/open-source-security).
