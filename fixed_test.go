// Copyright (c) 2020-2022 The distuvx developers. All rights reserved.
// Project site: https://github.com/goinvest/distuvx
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package distuvx

import (
	"fmt"
	"testing"
)

func TestNewFixed(t *testing.T) {
	testCases := []struct {
		num           float64
		numIterations int
	}{
		{1.0, 5},
		{10.0, 5},
		{100.0, 5},
	}
	for i, tc := range testCases {
		name := fmt.Sprintf("new_fixed_%d", i)
		t.Run(name, func(t *testing.T) {
			fixed := NewFixed(tc.num)
			for i := 0; i < tc.numIterations; i++ {
				got := fixed.Rand()
				assertFloat64(t, name, got, tc.num, 0.0001)
			}
		})
	}
}
