// Copyright (c) 2009-2020 Misakai Ltd. and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetricSnaphsot(t *testing.T) {
	s := NewMetric("x")
	s.UpdateTag("test")
	for i := int32(0); i < 100; i++ {
		s.Update(i)
	}

	// Create a snapshot
	h := s.Snapshot()
	assert.Equal(t, 100, h.Count())
	assert.Equal(t, 99, h.Max())
	assert.Equal(t, 0, h.Min())
	assert.True(t, h.Mean() > 49)
	assert.True(t, h.StdDev() > 28)
	assert.Equal(t, "x", h.Name())
	assert.Equal(t, float64(49.5), h.Quantile(50)[0])
	assert.Equal(t, 833.25, h.Variance())
	assert.Equal(t, "test", h.Tag())
	assert.Equal(t, 4950, h.Sum())

	h.T0 = 0
	h.T1 = 10
	assert.Equal(t, float64(10), h.Rate())

	t0, t1 := h.Window()
	assert.NotEqual(t, 0, t0)
	assert.NotEqual(t, 0, t1)
}

func TestSnapshots(t *testing.T) {
	snapshots := Snapshots{
		{Metric: "a"},
		{Metric: "b"},
	}

	m := snapshots.ToMap()
	assert.Equal(t, "a", m["a"].Metric)
	assert.Equal(t, "b", m["b"].Metric)
}

func TestMergeSnapshots(t *testing.T) {
	s1 := Snapshots{
		{Metric: "a"},
		{Metric: "b", T0: 10, T1: 20},
	}

	s2 := Snapshots{
		{Metric: "c"},
		{Metric: "b", T0: 5, T1: 40},
	}

	s1.Merge(s2)

	m := s1.ToMap()
	assert.Equal(t, "a", m["a"].Metric)
	assert.Equal(t, "b", m["b"].Metric)
	assert.Equal(t, "c", m["c"].Metric)
}
