// Copyright (c) 2009-2020 Misakai Ltd. and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleZero(t *testing.T) {
	var s sample
	assert.Zero(t, s.Min())
	assert.Zero(t, s.Mean())
	assert.Zero(t, s.StdDev())
	assert.Zero(t, s.Variance())
	assert.Zero(t, s.Quantile(50)[0])
}

func TestQuantileZero(t *testing.T) {
	var s sample
	for i := int64(0); i < 10000; i++ {
		s = append(s, 0)
	}

	assert.Zero(t, s.Quantile(0.0001)[0])
	assert.Zero(t, s.Quantile(50)[0])
	assert.Zero(t, s.Quantile(5000)[0])
}

func TestQuantiles(t *testing.T) {
	var s sample
	for i := int32(0); i < 10000; i++ {
		s = append(s, i/100)
	}

	assert.Equal(t, float64(49.5), s.Quantile(50)[0])
}

func TestSample_Histogram(t *testing.T) {
	s := sample{1, 2, 5, 54, 34, 9, 3, 2, 1, 1}
	p := []int{0, 10, 100}

	assert.Len(t, s.Histogram(p), 2)
	assert.Equal(t, 8, s.Histogram(p)[0].Count)
	assert.Equal(t, 2, s.Histogram(p)[1].Count)
}

func TestSample_HistogramEmpty(t *testing.T) {
	s := sample{1, 2, 5, 54, 34, 9, 3, 2, 1, 1}
	assert.Equal(t, 10, s.Histogram([]int{})[0].Count)
}
