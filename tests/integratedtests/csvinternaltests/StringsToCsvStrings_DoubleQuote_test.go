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

package csvinternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/csvinternal"
)

// ── StringsToCsvStrings double quote branch ──

func Test_StringsToCsvStrings_DoubleQuote(t *testing.T) {
	// Arrange
	result := csvinternal.StringsToCsvStrings(true, false, "a")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Strings_DoubleQuote returns correct value -- with args", actual)
}

// ── AnyItemsToCsvStrings no quote ──

func Test_AnyItemsToCsvStrings_NoQuote(t *testing.T) {
	// Arrange
	result := csvinternal.AnyItemsToCsvStrings(false, false, "a")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "Any_NoQuote returns correct value -- with args", actual)
}

// ── AnyItemsToCsvString all branches ──

func Test_AnyItemsToCsvString_DoubleQuote(t *testing.T) {
	// Arrange
	result := csvinternal.AnyItemsToCsvString(", ", true, false, "a")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyItems_CsvString_DoubleQuote returns correct value -- with args", actual)
}

func Test_AnyItemsToCsvString_NoQuote(t *testing.T) {
	// Arrange
	result := csvinternal.AnyItemsToCsvString(", ", false, false, "a")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyItems_CsvString_NoQuote returns correct value -- with args", actual)
}

func Test_AnyItemsToCsvString_Empty(t *testing.T) {
	// Arrange
	result := csvinternal.AnyItemsToCsvString(", ", false, false)

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "AnyItems_CsvString_Empty returns empty -- with args", actual)
}

// ── StringsToCsvString all branches ──

func Test_StringsToCsvString_DoubleQuote(t *testing.T) {
	// Arrange
	result := csvinternal.StringsToCsvString(", ", true, false, "a")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Strings_CsvString_DoubleQuote returns correct value -- with args", actual)
}

func Test_StringsToCsvString_NoQuote(t *testing.T) {
	// Arrange
	result := csvinternal.StringsToCsvString(", ", false, false, "a")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Strings_CsvString_NoQuote returns correct value -- with args", actual)
}

func Test_StringsToCsvString_Empty(t *testing.T) {
	// Arrange
	result := csvinternal.StringsToCsvString(", ", false, false)

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "Strings_CsvString_Empty returns empty -- with args", actual)
}

// ── StringersToString all branches ──

type cov3Stringer struct{ v string }

func (s cov3Stringer) String() string { return s.v }

func Test_StringersToString_SingleQuote(t *testing.T) {
	// Arrange
	s := cov3Stringer{v: "x"}
	result := csvinternal.StringersToString(", ", true, true, s)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringersToString_SingleQuote returns correct value -- with args", actual)
}

func Test_StringersToString_DoubleQuote(t *testing.T) {
	// Arrange
	s := cov3Stringer{v: "x"}
	result := csvinternal.StringersToString(", ", true, false, s)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringersToString_DoubleQuote returns correct value -- with args", actual)
}

func Test_StringersToString_NoQuote(t *testing.T) {
	// Arrange
	s := cov3Stringer{v: "x"}
	result := csvinternal.StringersToString(", ", false, false, s)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringersToString_NoQuote returns correct value -- with args", actual)
}

func Test_StringersToString_Empty(t *testing.T) {
	// Arrange
	result := csvinternal.StringersToString(", ", false, false)

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "StringersToString_Empty returns empty -- with args", actual)
}

// ── CompileStringersToString all branches ──

func Test_CompileStringersToString_SingleQuote(t *testing.T) {
	// Arrange
	f := func() string { return "x" }
	result := csvinternal.CompileStringersToString(", ", true, true, f)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CompileStringersToString_SingleQuote returns correct value -- with args", actual)
}

func Test_CompileStringersToString_DoubleQuote(t *testing.T) {
	// Arrange
	f := func() string { return "x" }
	result := csvinternal.CompileStringersToString(", ", true, false, f)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CompileStringersToString_DoubleQuote returns correct value -- with args", actual)
}

func Test_CompileStringersToString_NoQuote(t *testing.T) {
	// Arrange
	f := func() string { return "x" }
	result := csvinternal.CompileStringersToString(", ", false, false, f)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "CompileStringersToString_NoQuote returns correct value -- with args", actual)
}

func Test_CompileStringersToString_Empty(t *testing.T) {
	// Arrange
	result := csvinternal.CompileStringersToString(", ", false, false)

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "CompileStringersToString_Empty returns empty -- with args", actual)
}

// ── RangeNamesWithValuesIndexes ──

func Test_RangeNamesWithValuesIndexes_FromStringsToCsvStringsD(t *testing.T) {
	// Arrange
	result := csvinternal.RangeNamesWithValuesIndexes("A", "B")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "RangeNamesWithValuesIndexes returns non-empty -- with args", actual)
}

func Test_RangeNamesWithValuesIndexes_Empty(t *testing.T) {
	// Arrange
	result := csvinternal.RangeNamesWithValuesIndexes()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "RangeNamesWithValuesIndexes returns empty -- empty", actual)
}

func Test_RangeNamesWithValuesIndexesCsvString_Empty(t *testing.T) {
	// Arrange
	result := csvinternal.RangeNamesWithValuesIndexesCsvString()

	// Act
	actual := args.Map{"empty": result == ""}

	// Assert
	expected := args.Map{"empty": true}
	expected.ShouldBeEqual(t, 0, "RangeNamesCsvString returns empty -- empty", actual)
}
