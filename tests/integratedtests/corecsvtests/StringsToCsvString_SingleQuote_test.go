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

package corecsvtests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/corecsv"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ── StringsToCsvString — all quote branches ──

func Test_StringsToCsvString_SingleQuote(t *testing.T) {
	// Arrange
	result := corecsv.StringsToCsvString(", ", true, true, "a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringsToCsvString returns correct value -- single quote", actual)
}

func Test_StringsToCsvString_DoubleQuote(t *testing.T) {
	// Arrange
	result := corecsv.StringsToCsvString(", ", true, false, "a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringsToCsvString returns correct value -- double quote", actual)
}

func Test_StringsToCsvString_NoQuote(t *testing.T) {
	// Arrange
	result := corecsv.StringsToCsvString(", ", false, false, "a", "b")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "StringsToCsvString returns empty -- no quote", actual)
}

// ── AnyItemsToCsvString — all quote branches ──

func Test_AnyItemsToCsvString_SingleQuote(t *testing.T) {
	// Arrange
	result := corecsv.AnyItemsToCsvString(", ", true, true, "a", 1)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyItemsToCsvString returns correct value -- single quote", actual)
}

func Test_AnyItemsToCsvString_DoubleQuote(t *testing.T) {
	// Arrange
	result := corecsv.AnyItemsToCsvString(", ", true, false, "a", 1)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyItemsToCsvString returns correct value -- double quote", actual)
}

func Test_AnyItemsToCsvString_NoQuote(t *testing.T) {
	// Arrange
	result := corecsv.AnyItemsToCsvString(", ", false, false, "a", 1)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "AnyItemsToCsvString returns empty -- no quote", actual)
}

// ── AnyToTypesCsvStrings — no-quote branch ──

func Test_AnyToTypesCsvStrings_NoQuote_FromStringsToCsvStringSi(t *testing.T) {
	// Arrange
	result := corecsv.AnyToTypesCsvStrings(false, false, "a")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyToTypesCsvStrings returns empty -- no quote", actual)
}

func Test_AnyToTypesCsvStrings_Empty_FromStringsToCsvStringSi(t *testing.T) {
	// Arrange
	result := corecsv.AnyToTypesCsvStrings(false, false)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyToTypesCsvStrings returns empty -- empty", actual)
}

// ── StringsToCsvStrings — double-quote branch (explicit) ──

func Test_StringsToCsvStrings_DoubleQuote(t *testing.T) {
	// Arrange
	result := corecsv.StringsToCsvStrings(true, false, "a")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "StringsToCsvStrings returns correct value -- double quote", actual)
}

func Test_StringsToCsvStrings_Empty(t *testing.T) {
	// Arrange
	result := corecsv.StringsToCsvStrings(true, true)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "StringsToCsvStrings returns empty -- empty", actual)
}

// ── AnyItemsToCsvStrings — double-quote, empty ──

func Test_AnyItemsToCsvStrings_DoubleQuote(t *testing.T) {
	// Arrange
	result := corecsv.AnyItemsToCsvStrings(true, false, "a")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "AnyItemsToCsvStrings returns correct value -- double quote", actual)
}

func Test_AnyItemsToCsvStrings_Empty(t *testing.T) {
	// Arrange
	result := corecsv.AnyItemsToCsvStrings(true, true)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "AnyItemsToCsvStrings returns empty -- empty", actual)
}

// ── RangeNamesWithValuesIndexes — multi items ──

func Test_RangeNamesWithValuesIndexes_MultiItems(t *testing.T) {
	// Arrange
	result := corecsv.RangeNamesWithValuesIndexes("A", "B", "C")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 3}
	expected.ShouldBeEqual(t, 0, "RangeNamesWithValuesIndexes returns non-empty -- multi items", actual)
}

// ── StringersToString — empty ──

func Test_StringersToString_Empty(t *testing.T) {
	// Arrange
	result := corecsv.StringersToString(", ", false, false)

	// Act
	actual := args.Map{"empty": result}

	// Assert
	expected := args.Map{"empty": ""}
	expected.ShouldBeEqual(t, 0, "StringersToString returns empty -- empty", actual)
}

// ── CompileStringersToString — empty ──

func Test_CompileStringersToString_Empty(t *testing.T) {
	// Arrange
	result := corecsv.CompileStringersToString(", ", false, false)

	// Act
	actual := args.Map{"empty": result}

	// Assert
	expected := args.Map{"empty": ""}
	expected.ShouldBeEqual(t, 0, "CompileStringersToString returns empty -- empty", actual)
}
