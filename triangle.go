// Copyright (c) 2020-2022 The distuvx developers. All rights reserved.
// Project site: https://github.com/goinvest/distuvx
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package distuvx

import (
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

// NewTriangleOne create a new triangle distribution one time. Thereafter, the
// same fixed value is always returned.
func NewTriangleOne(a, b, c float64, src rand.Source) Fixed {
	triangle := distuv.NewTriangle(a, b, c, src)
	val := triangle.Rand()
	return NewFixed(val)
}
