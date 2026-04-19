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

package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/simplewrap"
	"github.com/alimtvnetwork/core/coretests/args"
)

// Test_Cov8_DoubleQuoteWrapElements_EmptyNonNilSlice tests the length==0 branch
// with a non-nil empty slice.
func Test_DoubleQuoteWrapElements_EmptyNonNilSlice(t *testing.T) {
	// Arrange
	input := []string{}

	// Act
	result := simplewrap.DoubleQuoteWrapElements(false, input...)

	// Assert
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElements with empty slice: got len, want 0", actual)
}

// Test_Cov8_DoubleQuoteWrapElementsWithIndexes_EmptyNonNilSlice tests length==0 branch.
func Test_DoubleQuoteWrapElementsWithIndexes_EmptyNonNilSlice(t *testing.T) {
	// Arrange
	input := []string{}

	// Act
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes(input...)

	// Assert
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "DoubleQuoteWrapElementsWithIndexes with empty slice: got len, want 0", actual)
}
