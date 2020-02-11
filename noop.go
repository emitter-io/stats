// +build !js
// Copyright (c) 2009-2020 Misakai Ltd. and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package stats

import (
	"time"
)

// Noop represents a no-op monitor
type Noop struct{}

// NewNoop creates a new no-op monitor.
func NewNoop() *Noop {
	return new(Noop)
}

// Assert contract compliance
var _ Measurer = NewNoop()

// Measure records a value in the queue
func (m *Noop) Measure(name string, value int32) {}

// MeasureElapsed measures elapsed time since the start
func (m *Noop) MeasureElapsed(name string, start time.Time) {}

// MeasureRuntime measures the runtime information
func (m *Noop) MeasureRuntime() {}

// Tag updates a tag.
func (m *Noop) Tag(name, tag string) {}

// Snapshot creates a snapshot
func (m *Noop) Snapshot() []byte { return []byte{} }
