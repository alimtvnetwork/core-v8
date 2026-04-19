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

package stringslicetests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/stringslice"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage13 — stringslice remaining gaps
//
// Target 1: MergeSlicesOfSlices.go:13-15 — duplicate len==0 check (dead code,
//   line 7 already catches len==0)
//
// Target 2: RegexTrimmedSplitNonEmptyAll.go:17-19 — split returns empty
//   regexp.Split always returns at least one element, so len==0 is unreachable.
//   Dead code.
// ══════════════════════════════════════════════════════════════════════════════

func Test_MergeSlicesOfSlices_AllEmpty(t *testing.T) {
	// Arrange — pass non-nil but empty slices
	slice1 := []string{}
	slice2 := []string{}

	// Act
	result := stringslice.MergeSlicesOfSlices(slice1, slice2)

	// Assert
	convey.Convey("MergeSlicesOfSlices returns empty when all input slices empty", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

func Test_RegexTrimmedSplitNonEmptyAll_NoMatch(t *testing.T) {
	// Arrange — regex that doesn't match
	re := regexp.MustCompile(`\d+`)

	// Act
	result := stringslice.RegexTrimmedSplitNonEmptyAll(re, "hello world")

	// Assert
	convey.Convey("RegexTrimmedSplitNonEmptyAll returns trimmed result when no match", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

// Coverage note: Both uncovered branches are dead code:
// MergeSlicesOfSlices line 13: duplicate of line 7
// RegexTrimmedSplitNonEmptyAll line 17: regexp.Split always returns ≥1 element
