// Copyright (c) 2020 The distuvx developers. All rights reserved.
// Project site: https://github.com/cumulusware/distuvx
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package distuvx

import (
	"golang.org/x/exp/rand"

	"gonum.org/v1/gonum/stat/distuv"
)

// PERT represents a PERT distribution, which is a four parameter Beta
// distribution described by the parameters min, max, and mode, as well as the
// requirement that the mean = (max + 4 * mod + min) / 6.
// (https://en.wikipedia.org/wiki/PERT_distribution)
type PERT struct {
	min  float64
	max  float64
	mode float64
	bd   distuv.Beta
}

// NewPERT constructs a new PERT distribution using the given min, max, and
// mode. Constraints are min < max and min ≤ mode ≤ max.
func NewPERT(min, max, mode float64, src rand.Source) PERT {
	checkPERTParameters(min, max, mode)
	alpha := 1 + 4*(mode-min)/(max-min)
	beta := 1 + 4*(max-mode)/(max-min)
	return PERT{
		min:  min,
		max:  max,
		mode: mode,
		bd: distuv.Beta{
			Alpha: alpha,
			Beta:  beta,
			Src:   src,
		},
	}
}

func checkPERTParameters(min, max, mode float64) {
	if min >= max {
		panic("pert: constraint of min < max violated")
	}
	if min > mode {
		panic("pert: constraint of min <= mode violated")
	}
	if mode > max {
		panic("pert: constraint of mode <= max violated")
	}
}

// CDF computes the value of the cumulative distribution function at y.
func (p PERT) CDF(y float64) float64 {
	x := (y - p.min) / (p.max - p.min)
	return p.bd.CDF(x) / (p.max - p.min)
}

// Mean returns the mean of the PERT probability distribution.
func (p PERT) Mean() float64 {
	return (p.min + 4*p.mode + p.max) / 6
}

// Mode returns the mode of the PERT probability distribution.
func (p PERT) Mode() float64 {
	return p.mode
}

// NumParameters returns the number of parameters in the PERT distribution.
func (p PERT) NumParameters() int {
	return 3
}

// Prob computes the value of the probability density function at y.
func (p PERT) Prob(y float64) float64 {
	x := (y - p.min) / (p.max - p.min)
	return p.bd.Prob(x) / (p.max - p.min)
}

// Rand implements the Rander interface for the PERT distribution.
func (p PERT) Rand() float64 {
	return p.bd.Rand()*(p.max-p.min) + p.min
}
