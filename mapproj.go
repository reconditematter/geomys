// Copyright (c) 2019-2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package geomys

// MapProjection -- a method to transfrom a globe's surface into a plane.
type MapProjection interface {
	// Spheroid -- returns the spheroid of the map projection.
	Spheroid() Spheroid
	// Params -- returns the parameters of the map projection.
	Params() map[string]float64
	// Project -- transforms a geographic point from
	// the spheroid into a location on the plane.
	Project(Point) [2]float64
}
