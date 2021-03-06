// Copyright (c) 2019-2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package geomys

// Point -- represents a pair of geographic coordinates
// (latitude,longitude).
type Point struct {
	lat, lon float64
}

// Geo -- returns a point with the geographic latitude `lat`
// and the geographic longitude `lon`.
// This function causes a runtime panic when lat∉[-90,90]
// or lon∉[-180,180].
func Geo(lat, lon float64) Point {
	if !(-90 <= lat && lat <= 90) {
		panic("geomys.Geo: lat not in [-90,90]")
	}
	if !(-180 <= lon && lon <= 180) {
		panic("geomys.Geo: lon not in [-180,180]")
	}
	return Point{lat, lon}
}

// Geo -- returns the geographic coordinates of `p`.
func (p Point) Geo() (lat, lon float64) {
	lat, lon = p.lat, p.lon
	return
}
