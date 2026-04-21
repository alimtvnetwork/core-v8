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

package corestrtests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/internal/trydo"
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
)

// safeTest wraps a test body with trydo.Block so that any panic is
// caught and reported as a GoConvey assertion failure instead of
// crashing the entire test binary (which would wipe the coverage profile).
//
// Usage:
//
//	func Test_Something(t *testing.T) {
//	    safeTest(t, "Something", func() {
//	        // ... original test body ...
//	    })
//	}
func safeTest(t *testing.T, name string, fn func()) {
	t.Helper()

	trydo.Block{
		Try: fn,
		Catch: func(e trydo.Exception) {
			msg := fmt.Sprintf("[PANIC RECOVERED] %s: %v", name, e)

			convey.Convey(
				name+" (panic recovered)", t, func() {
					convey.So(
						msg,
						should.BeEmpty,
					)
				},
			)
		},
	}.Do()
}
