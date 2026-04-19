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

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/msgcreator"
)

// Cover Assert.SortedArray with isPrint=true
func Test_Assert_SortedArray_Print(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.SortedArray(true, true, "c b a")

	// Act
	actual := args.Map{
		"first": result[0],
		"len": len(result),
	}

	// Assert
	expected := args.Map{
		"first": "a",
		"len": 3,
	}
	expected.ShouldBeEqual(t, 0, "SortedArray returns non-empty -- isPrint=true", actual)
}

func Test_Assert_SortedArray_NoSort(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.SortedArray(false, false, "c b a")

	// Act
	actual := args.Map{
		"first": result[0],
		"len": len(result),
	}

	// Assert
	expected := args.Map{
		"first": "c",
		"len": 3,
	}
	expected.ShouldBeEqual(t, 0, "SortedArray returns non-empty -- isSort=false", actual)
}

func Test_Assert_SortedMessage_Print(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.SortedMessage(true, "c b a", ",")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "SortedMessage returns non-empty -- isPrint=true", actual)
}

// Cover ToStrings with various types
func Test_Assert_ToStrings_Slice(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.ToStrings([]string{"a", "b"})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 2}
	expected.ShouldBeEqual(t, 0, "ToStrings returns correct value -- slice", actual)
}

func Test_Assert_ToStrings_Int(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.ToStrings(42)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ToStrings returns correct value -- int", actual)
}

func Test_Assert_ToStrings_Bool(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.ToStrings(true)

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 1}
	expected.ShouldBeEqual(t, 0, "ToStrings returns correct value -- bool", actual)
}

func Test_Assert_ToStrings_MapStringAny(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.ToStrings(map[string]any{"k": "v"})

	// Act
	actual := args.Map{"hasItems": len(result) > 0}

	// Assert
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "ToStrings returns correct value -- map[string]any", actual)
}

func Test_Assert_ToStringsWithSpace_Empty(t *testing.T) {
	// Arrange
	result := msgcreator.Assert.ToStringsWithSpace(4, []string{})

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "ToStringsWithSpace returns empty -- empty", actual)
}
