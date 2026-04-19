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
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
)

// caseV1Compat is a lightweight assertion helper used across many coverage files.
// Moved here from Iteration37_test.go so split-recovery subfolders can see it.
type caseV1Compat struct {
	Name     string
	Expected any
	Actual   any
	Args     args.Map
}

func (it caseV1Compat) ShouldBeEqual(t *testing.T) {
	expected := args.Map{"value": it.Expected}
	actual := args.Map{"value": it.Actual}
	expected.ShouldBeEqual(t, 0, it.Name, actual)
}

// testErr is a simple error type used in LinkedList/SimpleSlice tests.
// Moved here from Seg5_Hashset_test.go.
type testErr struct{}

func (e *testErr) Error() string { return "test" }
