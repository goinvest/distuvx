// Copyright (c) 2020-2022 The distuvx developers. All rights reserved.
// Project site: https://github.com/goinvest/distuvx
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package distuvx

import (
	"fmt"
	"testing"

	"golang.org/x/exp/rand"
)

func TestNewTriangle(t *testing.T) {
	testCases := []struct {
		lowerLimit    float64
		upperLimit    float64
		mode          float64
		numIterations int
	}{
		{1.0, 10.0, 5.0, 5},
	}
	for i, tc := range testCases {
		name := fmt.Sprintf("new_triangle_%d", i)
		t.Run(name, func(t *testing.T) {

			triangleSrc := rand.New(rand.NewSource(1234))
			triangleNum := NewTriangleOne(tc.lowerLimit, tc.upperLimit, tc.mode, triangleSrc)
			num := triangleNum.Rand()
			for i := 0; i < tc.numIterations; i++ {
				got := triangleNum.Rand()
				assertFloat64(t, name, got, num, 0.0001)
			}
		})
	}
}
