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
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/corecsv"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/isany"
)

// Test_Extended_Defined_TypedNil verifies isany.Defined with typed-nil error and *int.
func Test_Extended_Defined_TypedNil(t *testing.T) {
	for caseIndex, testCase := range extendedDefinedTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputs := input["inputs"].([]any)

		// Act
		actual := args.Map{}
		for i, v := range inputs {
			actual[fmt.Sprintf("result%d", i)] = fmt.Sprintf("%t", isany.Defined(v))
			actual[fmt.Sprintf("type%d", i)] = fmt.Sprintf("%T", v)
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// Test_Extended_Null_TypedNil verifies isany.Null with typed-nil error and *int.
func Test_Extended_Null_TypedNil(t *testing.T) {
	for caseIndex, testCase := range extendedNullTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputs := input["inputs"].([]any)

		// Act
		actual := args.Map{}
		for i, v := range inputs {
			actual[fmt.Sprintf("result%d", i)] = fmt.Sprintf("%t", isany.Null(v))
			actual[fmt.Sprintf("type%d", i)] = fmt.Sprintf("%T", v)
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// Test_Extended_DefinedBoth_TypedNil verifies isany.DefinedBoth with error and *int typed nils.
func Test_Extended_DefinedBoth_TypedNil(t *testing.T) {
	for caseIndex, testCase := range extendedDefinedBothTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pairs := input["pairs"].([]args.TwoAny)

		// Act
		actual := args.Map{}
		for i, pair := range pairs {
			f := pair.First
			s := pair.Second
			actual[fmt.Sprintf("result%d", i)] = fmt.Sprintf("%t", isany.DefinedBoth(f, s))
			actual[fmt.Sprintf("types%d", i)] = corecsv.AnyToTypesCsvDefault(f, s)
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// Test_Extended_NullBoth_TypedNil verifies isany.NullBoth with error and *int typed nils.
func Test_Extended_NullBoth_TypedNil(t *testing.T) {
	for caseIndex, testCase := range extendedNullBothTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pairs := input["pairs"].([]args.TwoAny)

		// Act
		actual := args.Map{}
		for i, pair := range pairs {
			f := pair.First
			s := pair.Second
			actual[fmt.Sprintf("result%d", i)] = fmt.Sprintf("%t", isany.NullBoth(f, s))
			actual[fmt.Sprintf("types%d", i)] = corecsv.AnyToTypesCsvDefault(f, s)
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
