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

package stringcompareastests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/enums/stringcompareas"
)

// ── VerifyMessage ignore-case positive mismatch ──

func Test_VerifyMessage_IgnoreCase_Positive(t *testing.T) {
	// Arrange
	msg := stringcompareas.StartsWith.VerifyMessage(true, "Hello", "world")

	// Act
	actual := args.Map{"nonEmpty": msg != ""}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "VerifyMessage returns correct value -- ignore case positive mismatch", actual)
}

// ── VerifyMessage ignore-case negative mismatch ──

func Test_VerifyMessage_IgnoreCase_Negative(t *testing.T) {
	// Arrange
	msg := stringcompareas.NotStartsWith.VerifyMessage(true, "Hello World", "hello")

	// Act
	actual := args.Map{
		"nonEmpty": msg != "",
		"isNeg": strings.Contains(msg, "negative"),
	}

	// Assert
	expected := args.Map{
		"nonEmpty": true,
		"isNeg": true,
	}
	expected.ShouldBeEqual(t, 0, "VerifyMessage returns correct value -- ignore case negative mismatch", actual)
}

// ── VerifyError case sensitive mismatch ──

func Test_VerifyErrorCaseSensitive_Mismatch(t *testing.T) {
	// Arrange
	err := stringcompareas.Equal.VerifyErrorCaseSensitive("hello", "world")

	// Act
	actual := args.Map{"hasError": err != nil}

	// Assert
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "VerifyErrorCaseSensitive returns error -- mismatch", actual)
}

// ── VerifyMessageCaseSensitive mismatch ──

func Test_VerifyMessageCaseSensitive_Mismatch(t *testing.T) {
	// Arrange
	msg := stringcompareas.Equal.VerifyMessageCaseSensitive("hello", "world")

	// Act
	actual := args.Map{"nonEmpty": msg != ""}

	// Assert
	expected := args.Map{"nonEmpty": true}
	expected.ShouldBeEqual(t, 0, "VerifyMessageCaseSensitive returns correct value -- mismatch", actual)
}

// ── NotAnyChars IsNegativeCondition ──

func Test_NotAnyChars_IsNegativeCondition(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.NotAnyChars.IsNegativeCondition()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotAnyChars returns correct value -- is negative", actual)
}

// ── NotEndsWith IsNegativeCondition ──

func Test_NotEndsWith_IsNegativeCondition(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.NotEndsWith.IsNegativeCondition()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotEndsWith returns non-empty -- is negative", actual)
}

// ── NotContains IsNegativeCondition ──

func Test_NotContains_IsNegativeCondition(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.NotContains.IsNegativeCondition()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotContains returns correct value -- is negative", actual)
}

// ── NotMatchRegex IsNegativeCondition ──

func Test_NotMatchRegex_IsNegativeCondition(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.NotMatchRegex.IsNegativeCondition()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NotMatchRegex returns correct value -- is negative", actual)
}

// ── IsAnyEnumsEqual match path ──

func Test_IsAnyEnumsEqual_Match(t *testing.T) {
	// Arrange
	a := stringcompareas.Equal
	b := stringcompareas.Equal

	// Act
	actual := args.Map{"result": a.IsAnyEnumsEqual(&b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "IsAnyEnumsEqual returns correct value -- match", actual)
}

// ── VerifyError ignore case match ──

func Test_VerifyError_IgnoreCase_Match(t *testing.T) {
	// Arrange
	err := stringcompareas.Equal.VerifyError(true, "Hello", "hello")

	// Act
	actual := args.Map{"isNil": err == nil}

	// Assert
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "VerifyError returns error -- ignore case match", actual)
}
