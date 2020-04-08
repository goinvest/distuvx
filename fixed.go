// Copyright (c) 2020 The distuvx developers. All rights reserved.
// Project site: https://github.com/goinvest/distuvx
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package distuvx

// Fixed is a fixed value regardless of the period.
type Fixed struct {
	num float64
}

// Rand implements the Rander interface for Fixed.
func (f Fixed) Rand() float64 {
	return f.num
}

// NewFixed creates a new fixed number.
func NewFixed(num float64) Fixed {
	return Fixed{num}
}
