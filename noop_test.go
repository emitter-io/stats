// +build !js
// Copyright (c) 2009-2020 Misakai Ltd. and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package stats

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNoop(t *testing.T) {
	m := NewNoop()
	m.Measure("a", 1)
	m.MeasureElapsed("b", time.Now())
	m.MeasureRuntime()
	m.Tag("a", "b")
	assert.NotNil(t, m)

	b := m.Snapshot()
	assert.Empty(t, b)
}
