// +build !js
// Copyright (c) 2009-2020 Misakai Ltd. and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package stats

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func BenchmarkMetricSnapshot(b *testing.B) {
	h := NewMetric("x")
	for i := int32(0); i < 50000; i++ {
		h.Update(i)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Snapshot()
	}
}

func BenchmarkMetricUpdate(b *testing.B) {
	m := NewMetric("x")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Update(15423)
	}
}

func TestMetricConcurrency(t *testing.T) {
	h := NewMetric("x")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := int32(0); i < 50000; i++ {
			h.Update(i)
		}
		wg.Done()
	}()

	go func() {
		for i := int32(0); i < 50000; i++ {
			h.Reset()
		}
		wg.Done()
	}()

	wg.Wait()
}

func TestMetric(t *testing.T) {
	h := NewMetric("x")
	for i := int32(0); i < 100; i++ {
		h.Update(i)
	}

	h.UpdateTag("test")
	assert.Equal(t, "test", h.Tag())

	// Create a snapshot
	assert.Equal(t, 100, h.Count())
	assert.Equal(t, 99, h.Max())
	assert.Equal(t, 0, h.Min())
	assert.True(t, h.Mean() > 49)
	assert.True(t, h.StdDev() > 28)
	assert.Equal(t, "x", h.Name())
	assert.Equal(t, float64(49.5), h.Quantile(50)[0])
	assert.Equal(t, 833.25, h.Variance())
	assert.NotZero(t, h.Rate())

	t0, t1 := h.Window()
	assert.NotEqual(t, time.Unix(0, 0), t0)
	assert.NotEqual(t, time.Unix(0, 0), t1)

	assert.Len(t, h.Histogram(0, 50, 100), 2)

	h.Reset()
	assert.Equal(t, 0, h.Count())
	assert.Equal(t, 0, h.Max())

}

func TestSampleClamp(t *testing.T) {
	h := NewMetric("x")
	for i := int32(0); i < 2000; i++ {
		h.Update(i)
	}

	sample := h.sample()
	assert.Equal(t, reservoirSize, len(sample))
}
