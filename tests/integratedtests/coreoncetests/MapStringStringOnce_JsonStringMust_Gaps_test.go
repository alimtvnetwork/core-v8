// MIT License
// 
// Copyright (c) 2020–2026
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NON-INFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package coreoncetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage16 — coreonce remaining gaps
//
// Target 1: MapStringStringOnce.JsonStringMust line 309-314
//   json.Marshal error → panic. Dead code — marshalling map[string]string
//   cannot fail.
//
// Target 2: StringsOnce.JsonStringMust line 251-256
//   Same pattern — marshalling []string cannot fail. Dead code.
// ══════════════════════════════════════════════════════════════════════════════

func Test_MapStringStringOnce_JsonStringMust_FromMapStringStringOnceJ(t *testing.T) {
	// Arrange
	once := coreonce.NewMapStringStringOnce(func() map[string]string {
		return map[string]string{
			"key1": "val1",
			"key2": "val2",
		}
	})

	// Act
	result := once.JsonStringMust()

	// Assert
	convey.Convey("MapStringStringOnce.JsonStringMust returns valid JSON", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
		convey.So(result, convey.ShouldContainSubstring, "key1")
	})
}

func Test_StringsOnce_JsonStringMust_FromMapStringStringOnceJ(t *testing.T) {
	// Arrange
	once := coreonce.NewStringsOnce(func() []string {
		return []string{"a", "b", "c"}
	})

	// Act
	result := once.JsonStringMust()

	// Assert
	convey.Convey("StringsOnce.JsonStringMust returns valid JSON", t, func() {
		convey.So(result, convey.ShouldNotBeEmpty)
		convey.So(result, convey.ShouldContainSubstring, "a")
	})
}

// Coverage note: Both error branches are dead code — json.Marshal of
// map[string]string and []string cannot fail.
