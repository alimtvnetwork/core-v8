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

	"github.com/alimtvnetwork/core-v8/coredata/corestr"
	"github.com/alimtvnetwork/core-v8/coretests/args")

// ══════════════════════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength
// ══════════════════════════════════════════════════════════════════════════════

func Test_TextWithLineNumber_HasLineNumber_FromSeg1(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_HasLineNumber_FromSeg1", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}

		// Act
		actual := args.Map{
			"has": tw.HasLineNumber(),
			"invalid": tw.IsInvalidLineNumber(),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"invalid": false,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber HasLineNumber true -- valid line", actual)
	})
}

func Test_TextWithLineNumber_Invalid_FromSeg1(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_Invalid_FromSeg1", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hello"}

		// Act
		actual := args.Map{
			"has": tw.HasLineNumber(),
			"invalid": tw.IsInvalidLineNumber(),
		}

		// Assert
		expected := args.Map{
			"has": false,
			"invalid": true,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber IsInvalidLineNumber -- invalid line", actual)
	})
}

func Test_TextWithLineNumber_Length_FromSeg1(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_Length_FromSeg1", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: "hello"}

		// Act
		actual := args.Map{"len": tw.Length()}

		// Assert
		expected := args.Map{"len": 5}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber Length -- 5 chars", actual)
	})
}

func Test_TextWithLineNumber_NilLength_FromSeg1(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_NilLength_FromSeg1", func() {
		// Arrange
		var tw *corestr.TextWithLineNumber

		// Act
		actual := args.Map{"len": tw.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber Length -- nil", actual)
	})
}

func Test_TextWithLineNumber_IsEmpty_FromSeg1(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_IsEmpty_FromSeg1", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: -1, Text: "hello"}

		// Act
		actual := args.Map{"empty": tw.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber IsEmpty -- invalid line number", actual)
	})
}

func Test_TextWithLineNumber_IsEmptyText_FromSeg1(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_IsEmptyText_FromSeg1", func() {
		// Arrange
		tw := &corestr.TextWithLineNumber{LineNumber: 1, Text: ""}

		// Act
		actual := args.Map{
			"emptyText": tw.IsEmptyText(),
			"emptyBoth": tw.IsEmptyTextLineBoth(),
		}

		// Assert
		expected := args.Map{
			"emptyText": true,
			"emptyBoth": true,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber IsEmptyText -- empty text", actual)
	})
}

func Test_TextWithLineNumber_NilEmpty_FromSeg1(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_NilEmpty_FromSeg1", func() {
		// Arrange
		var tw *corestr.TextWithLineNumber

		// Act
		actual := args.Map{
			"empty": tw.IsEmpty(),
			"emptyText": tw.IsEmptyText(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"emptyText": true,
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber IsEmpty -- nil receiver", actual)
	})
}

