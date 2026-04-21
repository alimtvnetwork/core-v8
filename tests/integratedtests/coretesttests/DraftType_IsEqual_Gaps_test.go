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

package coretesttests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage3 — coretests remaining gaps
//
// Target 1: BaseTestCaseAssertions.go:88-92 — isFailed log branch (test failure only)
// Target 2: DraftType.go:148 — isIncludingInnerFields && f1String differs (unexported)
// Target 3: DraftType.go:174,184 — json.Marshal panic (dead code)
// Target 4: SkipOnUnix.go:12-14 — platform-dependent (Unix only)
// ══════════════════════════════════════════════════════════════════════════════

func Test_DraftType_IsEqual_WithInnerFields_Same(t *testing.T) {
	// Arrange — exported fields match, inner fields are zero-valued and equal
	left := &coretests.DraftType{
		SampleString1: "sample1",
		SampleString2: "sample2",
		SampleInteger: 10,
	}
	right := &coretests.DraftType{
		SampleString1: "sample1",
		SampleString2: "sample2",
		SampleInteger: 10,
	}

	// Act
	result := left.IsEqual(true, right)

	// Assert
	convey.Convey("DraftType.IsEqual returns true when all fields match including inner", t, func() {
		convey.So(result, convey.ShouldBeTrue)
	})
}

// Coverage note:
// - BaseTestCaseAssertions.go:88-92 — only reachable on assertion failure inside convey.
//   Not safely testable. Documented as defensive logging gap.
// - DraftType.go:148,152 — requires different unexported fields (f1String, f2Integer).
//   Cannot be set from external tests. Needs internal test or is accepted gap.
// - DraftType.go:174,184 — json.Marshal panic. Dead code.
// - SkipOnUnix.go:12-14 — platform-dependent. Only runs on Unix.
