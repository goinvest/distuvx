// Copyright (c) 2020 The distuvx developers. All rights reserved.
// Project site: https://github.com/goinvest/distuvx
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package distuvx

import (
	"math"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

// Beta4 implements the four-parameter Beta distribution, which is like the
// two-parameter Beta distribution with ranges between the min and max instead
// of between 0 and 1.
type Beta4 struct {
	min float64
	max float64
	bd  distuv.Beta
}

// NewBeta4 constructs a new four-parameter Beta distribution.
func NewBeta4(alpha, beta, min, max float64, src rand.Source) Beta4 {
	checkBeta4Parameters(min, max)
	return Beta4{
		min: min,
		max: max,
		bd: distuv.Beta{
			Alpha: alpha,
			Beta:  beta,
			Src:   src,
		},
	}
}

func checkBeta4Parameters(min, max float64) {
	if min >= max {
		panic("pert: constraint of min < max violated")
	}
}

// CDF computes the value of the cumulative distribution function at y.
func (b Beta4) CDF(y float64) float64 {
	x := (y - b.min) / (b.max - b.min)
	return b.bd.CDF(x) / (b.max - b.min)
}

// Mean returns the mean of the Beta4 probability distribution.
func (b Beta4) Mean() float64 {
	return b.bd.Mean()*(b.max-b.min) + b.min
}

// Mode returns the mode of the Beta4 distribution.
//
// Mode returns NaN if either alpha or beta parameters are less than or equal
// to 1 as a special case.
func (b Beta4) Mode() float64 {
	if b.bd.Alpha <= 1 || b.bd.Beta <= 1 {
		return math.NaN()
	}
	return b.bd.Mode()*(b.max-b.min) + b.min
}

// NumParameters returns the number of parameters in the Beta4 distribution.
func (b Beta4) NumParameters() int {
	return 4
}

// Prob computes the value of the probability density function at y.
func (b Beta4) Prob(y float64) float64 {
	x := (y - b.min) / (b.max - b.min)
	return b.bd.Prob(x) / (b.max - b.min)
}

// Rand implements the Rander interface for the PERT distribution.
func (b Beta4) Rand() float64 {
	return b.bd.Rand()*(b.max-b.min) + b.min
}

// StdDev returns the standard deviation of the Beta4 probability distribution.
func (b Beta4) StdDev() float64 {
	return b.bd.StdDev() * (b.max - b.min)
}

// Variance returns the variance of the Beta4 probability distribution.
func (b Beta4) Variance() float64 {
	return b.bd.Variance() * (b.max - b.min) * (b.max - b.min)
}
