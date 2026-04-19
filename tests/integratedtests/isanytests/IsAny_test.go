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

package isanytests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

func Test_IsAny_Defined_Null_Verification(t *testing.T) {
	for caseIndex, testCase := range isAnyDefinedNullTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		var item any

		useError, _ := input.Get("useError")
		if useError == true {
			item = errors.New("test error")
		} else {
			item = input.GetDirectLower("input")
		}

		// Act
		actual := args.Map{
			"isDefined": fmt.Sprintf("%v", isany.Defined(item)),
			"isNull":    fmt.Sprintf("%v", isany.Null(item)),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsAny_Both_Verification(t *testing.T) {
	for caseIndex, testCase := range isAnyBothTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		first := input.GetDirectLower("first")
		second := input.GetDirectLower("second")

		// Act
		actual := args.Map{
			"definedBoth": fmt.Sprintf("%v", isany.DefinedBoth(first, second)),
			"nullBoth":    fmt.Sprintf("%v", isany.NullBoth(first, second)),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
