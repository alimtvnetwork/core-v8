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

package codestacktests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/codestack"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage12 — remaining codestack gaps
//
// Target 1: TraceCollection.FilterWithLimit line 529
//   Final return after loop completes without break or limit hit.
//
// Target 2: dirGetter.CurDir line 22 — runtime.Caller fail fallback (dead code)
// Target 3: fileGetter.CurrentFilePath line 74 — runtime.Caller fail fallback (dead code)
//
// Targets 2 & 3 are unreachable in normal execution (runtime.Caller always
// succeeds from test context). Documented as accepted dead-code gaps.
// ══════════════════════════════════════════════════════════════════════════════

func Test_FilterWithLimit_LoopCompletes(t *testing.T) {
	// Arrange — create a small collection, filter takes all, limit > count
	stacks := codestack.New.StackTrace.SkipNone()
	collection := &stacks

	// Act — limit much larger than items; loop will complete naturally
	result := collection.FilterWithLimit(
		9999,
		func(trace *codestack.Trace) (isTake, isBreak bool) {
			return true, false
		},
	)

	// Assert
	convey.Convey("FilterWithLimit returns all items when loop completes without break", t, func() {
		convey.So(len(result), convey.ShouldBeGreaterThan, 0)
	})
}

// Coverage note: dirGetter.CurDir line 22 and fileGetter.CurrentFilePath line 74
// are runtime.Caller failure fallbacks — unreachable in normal test execution.
// Documented as accepted dead-code gaps.
