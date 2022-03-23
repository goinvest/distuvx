// Copyright (c) 2020-2022 The distuvx developers. All rights reserved.
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
	Min      float64
	Max      float64
	BetaDist distuv.Beta
}

// NewBeta4 constructs a new four-parameter Beta distribution.
func NewBeta4(alpha, beta, min, max float64, src rand.Source) Beta4 {
	checkBeta4Parameters(min, max)
	return Beta4{
		Min: min,
		Max: max,
		BetaDist: distuv.Beta{
			Alpha: alpha,
			Beta:  beta,
			Src:   src,
		},
	}
}

// NewBeta4One creates a new four-parameter Beta distribution one time.
// Thereafter, the same fixed value is always returned.
func NewBeta4One(alpha, beta, min, max float64, src rand.Source) Fixed {
	beta4 := NewBeta4(alpha, beta, min, max, src)
	val := beta4.Rand()
	return NewFixed(val)
}

func checkBeta4Parameters(min, max float64) {
	if min >= max {
		panic("pert: constraint of min < max violated")
	}
}

// CDF computes the value of the cumulative distribution function at y.
func (b Beta4) CDF(y float64) float64 {
	x := (y - b.Min) / (b.Max - b.Min)
	return b.BetaDist.CDF(x) / (b.Max - b.Min)
}

// Mean returns the mean of the Beta4 probability distribution.
func (b Beta4) Mean() float64 {
	return b.BetaDist.Mean()*(b.Max-b.Min) + b.Min
}

// Mode returns the mode of the Beta4 distribution.
//
// Mode returns NaN if both parameters are less than or equal to 1 as a special
// case, 0 if only Alpha <= 1 and 1 if only Beta <= 1.
func (b Beta4) Mode() float64 {
	if b.BetaDist.Alpha <= 1 && b.BetaDist.Beta <= 1 {
		return math.NaN()
	} else if b.BetaDist.Alpha <= 1 && b.BetaDist.Beta > 1 {
		return 0.0
	} else if b.BetaDist.Beta <= 1 && b.BetaDist.Alpha > 1 {
		return 1.0
	}
	return b.BetaDist.Mode()*(b.Max-b.Min) + b.Min
}

// NumParameters returns the number of parameters in the Beta4 distribution.
func (b Beta4) NumParameters() int {
	return 4
}

// Prob computes the value of the probability density function at y.
func (b Beta4) Prob(y float64) float64 {
	x := (y - b.Min) / (b.Max - b.Min)
	return b.BetaDist.Prob(x) / (b.Max - b.Min)
}

// Rand implements the Rander interface for the Beta4 distribution.
func (b Beta4) Rand() float64 {
	return b.BetaDist.Rand()*(b.Max-b.Min) + b.Min
}

// StdDev returns the standard deviation of the Beta4 probability distribution.
func (b Beta4) StdDev() float64 {
	return b.BetaDist.StdDev() * (b.Max - b.Min)
}

// Variance returns the variance of the Beta4 probability distribution.
func (b Beta4) Variance() float64 {
	return b.BetaDist.Variance() * (b.Max - b.Min) * (b.Max - b.Min)
}
