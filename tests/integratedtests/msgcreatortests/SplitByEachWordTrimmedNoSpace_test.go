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

package msgcreatortests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/internal/msgcreator"
)

// ── SplitByEachWordTrimmedNoSpace ──

func Test_SplitByEachWordTrimmedNoSpace(t *testing.T) {
	// Arrange
	result := msgcreator.SplitByEachWordTrimmedNoSpace("hello world foo", true)
	emptyResult := msgcreator.SplitByEachWordTrimmedNoSpace("", false)
	unsorted := msgcreator.SplitByEachWordTrimmedNoSpace("b a", false)

	// Act
	actual := args.Map{
		"resultLen": len(result),
		"first":     result[0],
		"emptyLen":  len(emptyResult),
		"unsorted0": unsorted[0],
	}

	// Assert
	expected := args.Map{
		"resultLen": 3,
		"first":     "foo",
		"emptyLen":  0,
		"unsorted0": "b",
	}
	expected.ShouldBeEqual(t, 0, "SplitByEachWord returns correct value -- with args", actual)
}

// ── Assert.Quick ──

func Test_Assert_Quick(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.Quick("input", "actual", "expected", 0)

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Quick returns correct value -- with args", actual)
}

// ── Assert message methods ──

func Test_Assert_Messages(t *testing.T) {
	// Act
	actual := args.Map{
		"isEqual":    msgcreator.Assert.IsEqualMessage("w", "a", "e") != "",
		"isNotEqual": msgcreator.Assert.IsNotEqualMessage("w", "a", "e") != "",
		"isTrue":     msgcreator.Assert.IsTrueMessage("w", "a") != "",
		"isFalse":    msgcreator.Assert.IsFalseMessage("w", "a") != "",
		"isNil":      msgcreator.Assert.IsNilMessage("w", "a") != "",
		"isNotNil":   msgcreator.Assert.IsNotNilMessage("w", "a") != "",
		"shouldBe":   msgcreator.Assert.ShouldBeMessage("t", "a", "e") != "",
		"shouldNot":  msgcreator.Assert.ShouldNotBeMessage("t", "a", "e") != "",
	}

	// Assert
	expected := args.Map{
		"isEqual":    true,
		"isNotEqual": true,
		"isTrue":     true,
		"isFalse":    true,
		"isNil":      true,
		"isNotNil":   true,
		"shouldBe":   true,
		"shouldNot":  true,
	}
	expected.ShouldBeEqual(t, 0, "Messages returns correct value -- with args", actual)
}

// ── Assert.SortedMessage ──

func Test_Assert_SortedMessage(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.SortedMessage(false, "hello world", " ")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SortedMessage returns correct value -- with args", actual)
}

// ── Assert.SortedArrayNoPrint ──

func Test_Assert_SortedArrayNoPrint(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.SortedArrayNoPrint("c b a")

	// Act
	actual := args.Map{"first": result[0]}

	// Assert
	expected := args.Map{"first": "a"}
	expected.ShouldBeEqual(t, 0, "SortedArrayNoPrint returns correct value -- with args", actual)
}

// ── Assert.ToStrings ──

func Test_Assert_ToStrings(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.ToStrings("hello")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ToStrings returns correct value -- with args", actual)
}

// ── Assert.ToStringsWithSpace ──

func Test_Assert_ToStringsWithSpace(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.ToStringsWithSpace(2, "hello")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ToStringsWithSpace returns non-empty -- with args", actual)
}

func Test_Assert_ToStringsWithSpaceDefault(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.ToStringsWithSpaceDefault("hello")

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ToStringsWithSpaceDefault returns non-empty -- with args", actual)
}

func Test_Assert_ToStringWithSpace(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.ToStringWithSpace(2, "hello")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ToStringWithSpace returns non-empty -- with args", actual)
}

// ── Assert.StringsToWithSpaceLines ──

func Test_Assert_StringsToWithSpaceLines(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.StringsToWithSpaceLines(2, "a", "b")
	empty := msgcreator.Assert.StringsToWithSpaceLines(2)

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(empty),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "StringsToWithSpaceLines returns non-empty -- with args", actual)
}

// ── Assert.StringsToSpaceStringUsingFunc ──

func Test_Assert_StringsToSpaceStringUsingFunc(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.StringsToSpaceStringUsingFunc(2, func(i int, prefix, line string) string {
		return prefix + line
	}, "a", "b")
	empty := msgcreator.Assert.StringsToSpaceStringUsingFunc(2, func(i int, prefix, line string) string {
		return ""
	})

	// Act
	actual := args.Map{
		"len": len(result),
		"emptyLen": len(empty),
	}

	// Assert
	expected := args.Map{
		"len": 2,
		"emptyLen": 0,
	}
	expected.ShouldBeEqual(t, 0, "StringsToSpaceStringUsingFunc returns correct value -- with args", actual)
}
