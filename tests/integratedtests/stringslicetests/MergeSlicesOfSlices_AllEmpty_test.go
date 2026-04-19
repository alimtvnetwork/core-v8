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
	"github.com/alimtvnetwork/core/coretests/args"
)

// Test_MergeSlicesOfSlices_AllEmpty_FromMergeSlicesOfSlicesA tests MergeSlicesOfSlices with all-empty slices.
func Test_MergeSlicesOfSlices_AllEmpty_FromMergeSlicesOfSlicesA(t *testing.T) {
	// Arrange / Act
	result := stringslice.MergeSlicesOfSlices([]string{}, []string{})

	// Assert
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MergeSlicesOfSlices with all empty: got len, want 0", actual)
}

// Test_RegexTrimmedSplitNonEmptyAll_EmptyResult tests empty content with regex split.
func Test_RegexTrimmedSplitNonEmptyAll_EmptyResult(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`\s+`)

	// Act
	result := stringslice.RegexTrimmedSplitNonEmptyAll(re, "   ")

	// Assert
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "RegexTrimmedSplitNonEmptyAll with whitespace-only: got len, want 0", actual)
}

// Test_SplitTrimmedNonEmpty_ZeroCount tests SplitTrimmedNonEmpty with n=0.
func Test_SplitTrimmedNonEmpty_ZeroCount(t *testing.T) {
	// Arrange / Act
	result := stringslice.SplitTrimmedNonEmpty("a,b,c", ",", 0)

	// Assert
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmpty with n=0: got len, want 0", actual)
}

// Test_SplitTrimmedNonEmptyAll_EmptyContent tests with empty splitter result.
func Test_SplitTrimmedNonEmptyAll_EmptyContent(t *testing.T) {
	// Arrange / Act
	result := stringslice.SplitTrimmedNonEmptyAll("", "")

	// Assert
	actual := args.Map{"result": len(result) != 0}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "SplitTrimmedNonEmptyAll empty/empty: got len, want 0", actual)
}
