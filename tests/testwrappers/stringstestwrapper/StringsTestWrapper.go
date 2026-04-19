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

package stringstestwrapper

import (
	"github.com/alimtvnetwork/core/coretests"
)

// StringsTestWrapper wraps BaseTestCase with typed []string accessors
// for ArrangeInput and ExpectedInput. Use a type alias in package-local
// testWrapper.go files to keep test case definitions unchanged:
//
//	type testWrapper = stringstestwrapper.StringsTestWrapper
type StringsTestWrapper struct {
	coretests.BaseTestCase
}

// Arrange returns ArrangeInput cast to []string.
func (it StringsTestWrapper) Arrange() []string {
	return it.ArrangeInput.([]string)
}

// Expected returns ExpectedInput cast to []string.
func (it StringsTestWrapper) Expected() []string {
	return it.ExpectedInput.([]string)
}
