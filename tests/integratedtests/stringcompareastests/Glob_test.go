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
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/enums/stringcompareas"
)

func Test_Glob_Match_Verification(t *testing.T) {
	for caseIndex, testCase := range globMatchTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		content, _ := input.GetAsString("content")
		isIgnoreCaseVal, _ := input.Get("isIgnoreCase")
		isIgnoreCase, _ := isIgnoreCaseVal.(bool)

		// Act
		isGlobMatch := stringcompareas.Glob.IsCompareSuccess(isIgnoreCase, content, pattern)
		isNonGlobMatch := stringcompareas.NonGlob.IsCompareSuccess(isIgnoreCase, content, pattern)

		actual := args.Map{
			"isMatch":   fmt.Sprintf("%v", isGlobMatch),
			"isInverse": fmt.Sprintf("%v", isNonGlobMatch),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Glob_IsGlob_ReturnsTrue(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.Glob.IsGlob()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Glob.IsGlob() should return true", actual)
}

func Test_Glob_IsNonGlob_ReturnsTrue(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.NonGlob.IsNonGlob()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NonGlob.IsNonGlob() should return true", actual)
}

func Test_NonGlob_IsNegativeCondition(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.NonGlob.IsNegativeCondition()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "NonGlob should be a negative condition", actual)
}

func Test_Glob_IsNotNegativeCondition(t *testing.T) {
	// Act
	actual := args.Map{"result": stringcompareas.Glob.IsNegativeCondition()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Glob should not be a negative condition", actual)
}

func Test_Glob_Name(t *testing.T) {
	// Arrange
	name := stringcompareas.Glob.Name()

	// Act
	actual := args.Map{"result": name != "Glob"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'Glob', got ''", actual)
}

func Test_NonGlob_Name(t *testing.T) {
	// Arrange
	name := stringcompareas.NonGlob.Name()

	// Act
	actual := args.Map{"result": name != "NonGlob"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'NonGlob', got ''", actual)
}
