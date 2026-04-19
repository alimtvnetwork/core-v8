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

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// TextWithLineNumber — line-number predicates + emptiness checks (from S01)
// ══════════════════════════════════════════════════════════════════════════════

func Test_TextWithLineNumber_HasLineNumber_TrueWhenPositive_FalseWhenInvalidOrNil(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_HasLineNumber_TrueWhenPositive_FalseWhenInvalidOrNil", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hi"}
		tw2 := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hi"}
		var tw3 *corestr.TextWithLineNumber

		// Act
		actual := args.Map{
			"valid": tw.HasLineNumber(),
			"invalid": tw2.HasLineNumber(),
			"nil": tw3.HasLineNumber(),
		}

		// Assert
		expected := args.Map{
			"valid": true,
			"invalid": false,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "HasLineNumber returns correct value -- valid, invalid, nil", actual)
	})
}


func Test_TextWithLineNumber_IsInvalidLineNumber_TrueWhenInvalidOrNil(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_IsInvalidLineNumber_TrueWhenInvalidOrNil", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hi"}
		tw2 := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hi"}
		var tw3 *corestr.TextWithLineNumber

		// Act
		actual := args.Map{
			"valid": tw.IsInvalidLineNumber(),
			"invalid": tw2.IsInvalidLineNumber(),
			"nil": tw3.IsInvalidLineNumber(),
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"invalid": true,
			"nil": true,
		}
		expected.ShouldBeEqual(t, 0, "IsInvalidLineNumber returns correct value -- valid, invalid, nil", actual)
	})
}


func Test_TextWithLineNumber_Length_ReturnsTextLengthOrZeroForNil(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_Length_ReturnsTextLengthOrZeroForNil", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}
		var tw2 *corestr.TextWithLineNumber

		// Act
		actual := args.Map{
			"len": tw.Length(),
			"nilLen": tw2.Length(),
		}

		// Assert
		expected := args.Map{
			"len": 5,
			"nilLen": 0,
		}
		expected.ShouldBeEqual(t, 0, "Length returns correct value -- normal and nil", actual)
	})
}


func Test_TextWithLineNumber_IsEmpty_TrueWhenAnyTextOrLineMissing(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_IsEmpty_TrueWhenAnyTextOrLineMissing", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hi"}
		tw2 := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hi"}
		tw3 := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
		var tw4 *corestr.TextWithLineNumber

		// Act
		actual := args.Map{
			"validText": tw.IsEmpty(),
			"invalidLn": tw2.IsEmpty(),
			"emptyText": tw3.IsEmpty(),
			"nil":       tw4.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"validText": false,
			"invalidLn": true,
			"emptyText": true,
			"nil":       true,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty returns correct value -- various cases", actual)
	})
}


func Test_TextWithLineNumber_IsEmptyText_TrueWhenTextEmptyOrNil(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_IsEmptyText_TrueWhenTextEmptyOrNil", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hi"}
		tw2 := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}
		var tw3 *corestr.TextWithLineNumber

		// Act
		actual := args.Map{
			"hasText": tw.IsEmptyText(),
			"empty": tw2.IsEmptyText(),
			"nil": tw3.IsEmptyText(),
		}

		// Assert
		expected := args.Map{
			"hasText": false,
			"empty": true,
			"nil": true,
		}
		expected.ShouldBeEqual(t, 0, "IsEmptyText returns correct value -- text, empty, nil", actual)
	})
}


func Test_TextWithLineNumber_IsEmptyTextLineBoth_DelegatesToIsEmpty(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_IsEmptyTextLineBoth_DelegatesToIsEmpty", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hi"}

		// Act
		actual := args.Map{"both": tw.IsEmptyTextLineBoth()}

		// Assert
		expected := args.Map{"both": false}
		expected.ShouldBeEqual(t, 0, "IsEmptyTextLineBoth returns correct value -- delegates to IsEmpty", actual)
	})
}
