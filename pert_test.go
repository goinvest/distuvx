// Copyright (c) 2020-2022 The distuvx developers. All rights reserved.
// Project site: https://github.com/goinvest/distuvx
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package distuvx

import (
	"fmt"
	"math"
	"testing"

	"golang.org/x/exp/rand"
)

func TestNewPERTOne(t *testing.T) {
	c := struct {
		tolerance float64
	}{0.0001}
	testCases := []struct {
		min           float64
		max           float64
		mode          float64
		seed          uint64
		numIterations int
	}{
		{
			min:           5.0,
			max:           10.0,
			mode:          5.0,
			seed:          12,
			numIterations: 5,
		},
	}
	for i, tc := range testCases {
		name := fmt.Sprintf("new_pert_one_%d", i)
		t.Run(name, func(t *testing.T) {

			src := rand.New(rand.NewSource(tc.seed))
			pertOne := NewPERTOne(tc.min, tc.max, tc.mode, src)
			num := pertOne.Rand()
			for i := 0; i < tc.numIterations; i++ {
				got := pertOne.Rand()
				assertFloat64(t, name, got, num, c.tolerance)
			}
		})
	}
}
func TestNewPERT(t *testing.T) {
	c := struct {
		tolerance float64
	}{0.0001}
	testCases := []struct {
		min           float64
		max           float64
		mode          float64
		seed          uint64
		wantCDF       float64
		yCDF          float64
		wantMean      float64
		wantMode      float64
		wantNumParams int
		wantProb      float64
		yProb         float64
		wantRand      float64
		wantStdDev    float64
		wantVariance  float64
	}{
		{
			min:           0.0,
			max:           10.0,
			mode:          5.0,
			seed:          12,
			wantCDF:       0.089648,
			yCDF:          7.5,
			wantMean:      5.0,
			wantMode:      5.0,
			wantNumParams: 3,
			wantProb:      0.105469,
			yProb:         7.5,
			wantRand:      6.186076,
			wantStdDev:    1.889822,
			wantVariance:  3.571429,
		},
	}
	for i, tc := range testCases {
		name := fmt.Sprintf("new_pert_%d", i)
		t.Run(name, func(t *testing.T) {
			src := rand.New(rand.NewSource(tc.seed))
			pert := NewPERT(tc.min, tc.max, tc.mode, src)
			gotCDF := pert.CDF(tc.yCDF)
			label := fmt.Sprintf("%s_cdf", name)
			assertFloat64(t, label, gotCDF, tc.wantCDF, c.tolerance)
			gotMean := pert.Mean()
			label = fmt.Sprintf("%s_mean", name)
			assertFloat64(t, label, gotMean, tc.wantMean, c.tolerance)
			gotMode := pert.Mode()
			label = fmt.Sprintf("%s_mode", name)
			if math.IsNaN(gotMode) && !math.IsNaN(tc.wantMode) {
				t.Errorf("Got NaN but wanted %f", tc.wantMode)
			}
			assertFloat64(t, label, gotMode, tc.wantMode, c.tolerance)
			gotNumParams := pert.NumParameters()
			label = fmt.Sprintf("%s_num_params", name)
			assertInt(t, label, gotNumParams, tc.wantNumParams)
			gotProb := pert.Prob(tc.yProb)
			label = fmt.Sprintf("%s_prob", name)
			assertFloat64(t, label, gotProb, tc.wantProb, c.tolerance)
			gotRand := pert.Rand()
			label = fmt.Sprintf("%s_rand", name)
			assertFloat64(t, label, gotRand, tc.wantRand, c.tolerance)
			gotStdDev := pert.StdDev()
			label = fmt.Sprintf("%s_std_dev", name)
			assertFloat64(t, label, gotStdDev, tc.wantStdDev, c.tolerance)
			gotVariance := pert.Variance()
			label = fmt.Sprintf("%s_variance", name)
			assertFloat64(t, label, gotVariance, tc.wantVariance, c.tolerance)
		})
	}
}
