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

package coretests

import (
	"testing"
)

// These tests use unexported symbols and must remain in source package.

func Test_MessagePrinter_FailedExpected(t *testing.T) {
	p := printMessage{}
	// Arrange + Act: not failed
	p.FailedExpected(false, "when", "actual", "expected", 0)
	// Arrange + Act: failed
	p.FailedExpected(true, "when", "actual", "expected", 1)
}

func Test_MessagePrinter_NameValueAndValue(t *testing.T) {
	p := printMessage{}
	// Act
	p.NameValue("header", map[string]int{"a": 1})
	p.Value("header2", []int{1, 2, 3})
}

func Test_SkipOnUnix(t *testing.T) {
	t.Run("sub", func(st *testing.T) {
		SkipOnUnix(st)
	})
}
