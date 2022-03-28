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

func TestNewBeta4One(t *testing.T) {
	c := struct {
		tolerance float64
	}{0.0001}
	testCases := []struct {
		alpha         float64
		beta          float64
		min           float64
		max           float64
		seed          uint64
		numIterations int
	}{
		{
			alpha:         2.0,
			beta:          2.0,
			min:           5.0,
			max:           10.0,
			seed:          12,
			numIterations: 5,
		},
	}
	for i, tc := range testCases {
		name := fmt.Sprintf("new_beta4_one_%d", i)
		t.Run(name, func(t *testing.T) {

			src := rand.New(rand.NewSource(tc.seed))
			beta4One := NewBeta4One(tc.alpha, tc.beta, tc.min, tc.max, src)
			num := beta4One.Rand()
			for i := 0; i < tc.numIterations; i++ {
				got := beta4One.Rand()
				assertFloat64(t, name, got, num, c.tolerance)
			}
		})
	}
}

func TestNewBeta4(t *testing.T) {
	c := struct {
		tolerance float64
	}{0.0001}
	testCases := []struct {
		alpha         float64
		beta          float64
		min           float64
		max           float64
		seed          uint64
		wantRand      float64
		wantStdDev    float64
		wantMode      float64
		wantMean      float64
		wantProb      float64
		yProb         float64
		wantCDF       float64
		xCDF          float64
		wantVariance  float64
		wantNumParams int
	}{
		{
			alpha:         2.0,
			beta:          2.0,
			min:           5.0,
			max:           10.0,
			seed:          12,
			wantRand:      8.298276,
			wantStdDev:    1.118034,
			wantMode:      7.5,
			wantMean:      7.5,
			wantProb:      0.3,
			yProb:         7.5,
			wantCDF:       0.1,
			xCDF:          7.5,
			wantVariance:  1.25,
			wantNumParams: 4,
		},
		{
			alpha:         1.1,
			beta:          1.3,
			min:           5.0,
			max:           10.0,
			seed:          1234,
			wantRand:      6.113136,
			wantStdDev:    1.351099,
			wantMode:      6.25,
			wantMean:      7.291667,
			wantProb:      0.231675,
			yProb:         6.0,
			wantCDF:       0.112213,
			xCDF:          7.5,
			wantVariance:  1.82547,
			wantNumParams: 4,
		},
		{
			alpha:         0.5,
			beta:          0.3,
			min:           5.0,
			max:           10.0,
			seed:          1234,
			wantRand:      5.016593,
			wantStdDev:    1.804220,
			wantMode:      math.NaN(),
			wantMean:      8.125,
			wantProb:      0.107666,
			yProb:         8.0,
			wantCDF:       0.072142,
			xCDF:          7.5,
			wantVariance:  3.255208,
			wantNumParams: 4,
		},
		{
			alpha:         0.5,
			beta:          2.0,
			min:           0.0,
			max:           1.0,
			seed:          1234,
			wantRand:      0.000321,
			wantStdDev:    0.213809,
			wantMode:      0.0,
			wantMean:      0.2,
			wantProb:      0.0,
			yProb:         8.0,
			wantCDF:       1.0,
			xCDF:          7.5,
			wantVariance:  0.045714,
			wantNumParams: 4,
		},
		{
			alpha:         2.5,
			beta:          0.5,
			min:           0.0,
			max:           1.0,
			seed:          1234,
			wantRand:      0.998835,
			wantStdDev:    0.186339,
			wantMode:      1.0,
			wantMean:      0.833333,
			wantProb:      0.0,
			yProb:         8.0,
			wantCDF:       1.0,
			xCDF:          7.5,
			wantVariance:  0.034722,
			wantNumParams: 4,
		},
		{
			alpha:         1.5,
			beta:          3.3,
			min:           5.0,
			max:           10.0,
			seed:          1234,
			wantRand:      5.772624,
			wantStdDev:    0.962315,
			wantMode:      5.892857,
			wantMean:      6.5625,
			wantProb:      0.141237,
			yProb:         8.0,
			wantCDF:       0.163862,
			xCDF:          7.5,
			wantVariance:  0.926051,
			wantNumParams: 4,
		},
		{
			alpha:         1.5,
			beta:          3.3,
			min:           0.0,
			max:           1.0,
			seed:          1234,
			wantRand:      0.154525,
			wantStdDev:    0.192463,
			wantMode:      0.178571,
			wantMean:      0.3125,
			wantProb:      1.808808,
			yProb:         0.3,
			wantCDF:       1.0,
			xCDF:          1.0,
			wantVariance:  0.037042,
			wantNumParams: 4,
		},
		{
			alpha:         2.0,
			beta:          2.0,
			min:           0.0,
			max:           1.0,
			seed:          1234,
			wantRand:      0.336436,
			wantStdDev:    0.223607,
			wantMode:      0.5,
			wantMean:      0.5,
			wantProb:      1.5,
			yProb:         0.5,
			wantCDF:       0.5,
			xCDF:          0.5,
			wantVariance:  0.05,
			wantNumParams: 4,
		},
		{
			alpha:         2.0,
			beta:          2.0,
			min:           9.0,
			max:           10.0,
			seed:          1234,
			wantRand:      9.336436,
			wantStdDev:    0.223607,
			wantMode:      9.5,
			wantMean:      9.5,
			wantProb:      1.5,
			yProb:         9.5,
			wantCDF:       0.5,
			wantVariance:  0.05,
			xCDF:          9.5,
			wantNumParams: 4,
		},
	}
	for i, tc := range testCases {
		name := fmt.Sprintf("new_beta4_%d", i)
		t.Run(name, func(t *testing.T) {
			src := rand.New(rand.NewSource(tc.seed))
			beta4 := NewBeta4(tc.alpha, tc.beta, tc.min, tc.max, src)
			gotAlpha := beta4.BetaDist.Alpha
			label := fmt.Sprintf("%s_alpha", name)
			assertFloat64(t, label, gotAlpha, tc.alpha, c.tolerance)
			gotBeta := beta4.BetaDist.Beta
			label = fmt.Sprintf("%s_beta", name)
			assertFloat64(t, label, gotBeta, tc.beta, c.tolerance)
			gotRand := beta4.Rand()
			label = fmt.Sprintf("%s_rand", name)
			assertFloat64(t, label, gotRand, tc.wantRand, c.tolerance)
			gotStdDev := beta4.StdDev()
			label = fmt.Sprintf("%s_std_dev", name)
			assertFloat64(t, label, gotStdDev, tc.wantStdDev, c.tolerance)
			gotMode := beta4.Mode()
			label = fmt.Sprintf("%s_mode", name)
			if math.IsNaN(gotMode) && !math.IsNaN(tc.wantMode) {
				t.Errorf("Got NaN but wanted %f", tc.wantMode)
			}
			assertFloat64(t, label, gotMode, tc.wantMode, c.tolerance)
			gotMean := beta4.Mean()
			label = fmt.Sprintf("%s_mean", name)
			assertFloat64(t, label, gotMean, tc.wantMean, c.tolerance)
			gotNumParams := beta4.NumParameters()
			label = fmt.Sprintf("%s_num_params", name)
			assertInt(t, label, gotNumParams, tc.wantNumParams)
			gotProb := beta4.Prob(tc.yProb)
			label = fmt.Sprintf("%s_prob", name)
			assertFloat64(t, label, gotProb, tc.wantProb, c.tolerance)
			gotCDF := beta4.CDF(tc.xCDF)
			label = fmt.Sprintf("%s_cdf", name)
			assertFloat64(t, label, gotCDF, tc.wantCDF, c.tolerance)
			gotVariance := beta4.Variance()
			label = fmt.Sprintf("%s_variance", name)
			assertFloat64(t, label, gotVariance, tc.wantVariance, c.tolerance)
		})
	}
}
