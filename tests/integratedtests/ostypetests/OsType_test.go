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

package ostypetests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/ostype"
)

func Test_GetVariant_Verification(t *testing.T) {
	for caseIndex, testCase := range getVariantTestCases {
		// Arrange
		input := testCase.ArrangeInput.(string)

		// Act
		variant := ostype.GetVariant(input)
		name := variant.Name()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, name)
	}
}

func Test_GetGroup_Verification(t *testing.T) {
	for caseIndex, testCase := range getGroupTestCases {
		// Arrange
		input := testCase.ArrangeInput.(string)

		// Act
		group := ostype.GetGroup(input)
		name := group.Name()

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, name)
	}
}

func Test_Variation_Group_Verification(t *testing.T) {
	for caseIndex, testCase := range variationGroupTestCases {
		// Arrange
		input := testCase.ArrangeInput.(ostype.Variation)

		// Act
		group := input.Group()
		actual := args.Map{
			"groupName": group.Name(),
			"isUnix":    fmt.Sprintf("%v", group.IsUnix()),
			"isWindows": fmt.Sprintf("%v", group.IsWindows()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Variation_Identity_Verification(t *testing.T) {
	for caseIndex, testCase := range variationIdentityTestCases {
		// Arrange
		input := testCase.ArrangeInput.(ostype.Variation)

		// Act
		actual := args.Map{
			"isWindows": fmt.Sprintf("%v", input.IsWindows()),
			"isLinux":   fmt.Sprintf("%v", input.IsLinux()),
			"isDarwin":  fmt.Sprintf("%v", input.IsDarwinOrMacOs()),
			"isValid":   fmt.Sprintf("%v", input.IsValid()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
