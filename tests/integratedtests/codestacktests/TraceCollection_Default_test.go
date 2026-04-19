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

	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── TraceCollection via StackTrace ──

func Test_TraceCollection_Default(t *testing.T) {
	// Arrange & Act — use exported StackTrace to get a TraceCollection
	tc := codestack.New.StackTrace.Default(0, 5)

	// Assert
	actual := args.Map{"hasItems": tc.Length() > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "StackTrace Default returns non-empty collection", actual)
}

func Test_TraceCollection_Empty(t *testing.T) {
	// Arrange & Act — use zero-value TraceCollection
	tc := codestack.TraceCollection{}

	// Assert
	actual := args.Map{"length": tc.Length()}
	expected := args.Map{"length": 0}
	expected.ShouldBeEqual(t, 0, "TraceCollection zero-value returns empty -- default", actual)
}

func Test_TraceCollection_Add_FromTraceCollectionDefau(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	trace := codestack.New.Default()

	// Act
	tc.Add(trace)

	// Assert
	actual := args.Map{"length": tc.Length()}
	expected := args.Map{"length": 1}
	expected.ShouldBeEqual(t, 0, "TraceCollection Add increases length -- single trace", actual)
}

func Test_GetSinglePageCollection_ZeroPagePanic(t *testing.T) {
	// Arrange
	tc := codestack.TraceCollection{}
	tc.Add(codestack.New.Default())
	tc.Add(codestack.New.Default())
	tc.Add(codestack.New.Default())

	// Act
	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		tc.GetSinglePageCollection(2, 0)
	}()

	// Assert
	actual := args.Map{"didPanic": didPanic}
	expected := args.Map{"didPanic": true}
	expected.ShouldBeEqual(t, 0, "GetSinglePageCollection panics -- zero pageIndex", actual)
}
