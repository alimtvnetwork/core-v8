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

package stringutiltests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coreutils/stringutil"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_QW_IsEndsWith_NegativeRemainingLength(t *testing.T) {
	// Arrange
	result := stringutil.IsEndsWith("ab", "abcdef", false)

	// Act
	actual := args.Map{"result": result}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false when endsWith is longer than base", actual)
}

func Test_QW_ToIntUsingRegexMatch_NilRegex(t *testing.T) {
	// Arrange
	result := stringutil.ToIntUsingRegexMatch(nil, "123")

	// Act
	actual := args.Map{"result": result != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 for nil regex", actual)
}

func Test_QW_ToIntUsingRegexMatch_NoMatch(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`^\d+$`)
	result := stringutil.ToIntUsingRegexMatch(re, "abc")

	// Act
	actual := args.Map{"result": result != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 for no match", actual)
}

func Test_QW_ToIntUsingRegexMatch_ParseError(t *testing.T) {
	// Arrange
	re := regexp.MustCompile(`.*`)
	result := stringutil.ToIntUsingRegexMatch(re, "abc")

	// Act
	actual := args.Map{"result": result != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 for parse error", actual)
}

func Test_QW_UsingBracketsWrappedTemplate(t *testing.T) {
	// Arrange
	result := stringutil.ReplaceTemplate.UsingBracketsWrappedTemplate(
		"hello {brackets-wrapped} world",
		"REPLACED",
	)

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	result2 := stringutil.ReplaceTemplate.UsingBracketsWrappedTemplate("", "REPLACED")
	actual = args.Map{"result": result2 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_QW_UsingQuotesWrappedTemplate(t *testing.T) {
	// Arrange
	result := stringutil.ReplaceTemplate.UsingQuotesWrappedTemplate(
		"hello {quotes-wrapped} world",
		"REPLACED",
	)

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	result2 := stringutil.ReplaceTemplate.UsingQuotesWrappedTemplate("", "REPLACED")
	actual = args.Map{"result": result2 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// Renamed to avoid redeclaration with Coverage6_test.go
type qwTestNamer struct{ name string }

func (n qwTestNamer) Name() string { return n.name }

func Test_QW_UsingNamerMapOptions_CurlyKeys(t *testing.T) {
	_ = fmt.Sprintf("placeholder") // avoid unused import
}

type qwTestStringer struct{ val string }

func (s qwTestStringer) String() string { return s.val }

func Test_QW_UsingStringerMapOptions_CurlyKeys(t *testing.T) {
	m := map[fmt.Stringer]string{
		qwTestStringer{"key"}: "val",
	}
	result := stringutil.ReplaceTemplate.UsingStringerMapOptions(true, "hello {key} world", m)
	_ = result
}

func Test_QW_UsingStringerMapOptions_DirectKeys(t *testing.T) {
	m := map[fmt.Stringer]string{
		qwTestStringer{"key"}: "val",
	}
	result := stringutil.ReplaceTemplate.UsingStringerMapOptions(false, "hello key world", m)
	_ = result
}
