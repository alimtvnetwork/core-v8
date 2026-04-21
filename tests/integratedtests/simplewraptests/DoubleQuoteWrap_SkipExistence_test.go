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

package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/simplewrap"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

func Test_DoubleQuoteWrapElements_SkipOnExistence(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElements(true, "hello", "\"already\"")

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_DoubleQuoteWrapElements_NoSkip(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElements(false, "hello")

	// Act
	actual := args.Map{"result": len(result) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_DoubleQuoteWrapElements_Nil(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElements(false)

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_DoubleQuoteWrapElementsWithIndexes(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes("a", "b")

	// Act
	actual := args.Map{"result": len(result) != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_DoubleQuoteWrapElementsWithIndexes_Nil_FromDoubleQuoteWrapSkipE(t *testing.T) {
	// Arrange
	result := simplewrap.DoubleQuoteWrapElementsWithIndexes()

	// Act
	actual := args.Map{"result": len(result) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}
