// Copyright (c) 2020-2022 The distuvx developers. All rights reserved.
// Project site: https://github.com/goinvest/distuvx
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package distuvx

import (
	"fmt"
	"math"
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

func assertInt(t *testing.T, label string, got, want int) {
	if got != want {
		t.Errorf("\t got = %d %s\n\t\twant = %d", got, label, want)
	}
}

func assertFloat64(t *testing.T, label string, got, want, tolerance float64) {
	if diff := math.Abs(want - got); diff >= tolerance {
		t.Errorf("\t got = %f %s\n\t\t\twant = %f", got, label, want)
	}
}

func assertBool(t *testing.T, label string, got, want bool) {
	if got != want {
		t.Errorf("\t got = %t %s\n\t\t\twant = %t", got, label, want)
	}
}

func assertString(t *testing.T, label string, got, want string) {
	if got != want {
		t.Errorf("\t got = %s %s\n\t\t\twant = %s", got, label, want)
	}
}
