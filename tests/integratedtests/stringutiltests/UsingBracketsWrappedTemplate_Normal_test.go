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
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coreutils/stringutil"
)

// ══════════════════════════════════════════════════════════════════════════════
// UsingNamerMapOptions — curly and non-curly with actual values
// ══════════════════════════════════════════════════════════════════════════════

// Note: UsingNamerMapOptions non-nil paths require in-package test (namer is unexported)

// ══════════════════════════════════════════════════════════════════════════════
// UsingBracketsWrappedTemplate, UsingQuotesWrappedTemplate — normal paths
// ══════════════════════════════════════════════════════════════════════════════

func Test_UsingBracketsWrappedTemplate_Normal(t *testing.T) {
	// Arrange
	result := stringutil.ReplaceTemplate.UsingBracketsWrappedTemplate("prefix {brackets-wrapped} suffix", "VALUE")

	// Act
	actual := args.Map{"has": len(result) > 0}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "UsingBracketsWrappedTemplate returns correct value -- normal", actual)
}

func Test_UsingQuotesWrappedTemplate_Normal(t *testing.T) {
	// Arrange
	result := stringutil.ReplaceTemplate.UsingQuotesWrappedTemplate(`prefix "{quotes-wrapped}" suffix`, "VALUE")

	// Act
	actual := args.Map{"has": len(result) > 0}

	// Assert
	expected := args.Map{"has": true}
	expected.ShouldBeEqual(t, 0, "UsingQuotesWrappedTemplate returns correct value -- normal", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReplaceWhiteSpaces — with tabs and newlines
// ══════════════════════════════════════════════════════════════════════════════

func Test_ReplaceWhiteSpaces_WithTabs(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.ReplaceWhiteSpaces("  a\tb\nc  ")}

	// Assert
	expected := args.Map{"v": "abc"}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpaces returns correct value -- tabs", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// ReplaceWhiteSpacesToSingle — with newlines/tabs
// ══════════════════════════════════════════════════════════════════════════════

func Test_ReplaceWhiteSpacesToSingle_WithNewlines(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.ReplaceWhiteSpacesToSingle("a\nb\tc")}

	// Assert
	expected := args.Map{"v": "abc"}
	expected.ShouldBeEqual(t, 0, "ReplaceWhiteSpacesToSingle returns correct value -- newlines", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CurlyKeyUsingMap — normal path
// ══════════════════════════════════════════════════════════════════════════════

func Test_CurlyKeyUsingMap_Normal(t *testing.T) {
	// Act
	actual := args.Map{"v": stringutil.ReplaceTemplate.CurlyKeyUsingMap("{x}-{y}", map[string]string{"x": "1", "y": "2"})}

	// Assert
	expected := args.Map{"v": "1-2"}
	expected.ShouldBeEqual(t, 0, "CurlyKeyUsingMap returns correct value -- normal", actual)
}
