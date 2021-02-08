// Copyright (c) 2021 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

package geomys

import (
	"github.com/reconditematter/kgeo"
)

// Geodesic -- geodesic solver for a spheroidal model of the Earth.
//
// Reference: Karney, C.F.F. Algorithms for geodesics. J Geod 87, 43–55 (2013).
//
// DOI: https://doi.org/10.1007/s00190-012-0578-z
type Geodesic struct {
	sph Spheroid
	geo kgeo.Geodesic
}

// NewGeodesic -- returns a geodesic solver for the spheroid `sph`.
func NewGeodesic(sph Spheroid) Geodesic {
	return Geodesic{sph: sph, geo: kgeo.NewGeodesic(sph.A(), sph.F())}
}

// Spheroid -- returns the spheroid of `g`.
func (g Geodesic) Spheroid() Spheroid {
	return g.sph
}

// Direct -- solves the direct problem: given the source point `p1`, the azimuth `α1` (degrees),
// and the geodesic distance `s12` (meters), find the target point `p2`, also find the azimuth
// `α2` (degrees) at `p2`.
func (g Geodesic) Direct(p1 Point, α1 float64, s12 float64) (p2 Point, α2 float64) {
	lat1, lon1 := p1.Geo()
	sol := g.geo.Direct(lat1, lon1, α1, s12)
	p2 = Geo(sol.Lat2, sol.Lon2)
	α2 = sol.Azi2
	return
}
