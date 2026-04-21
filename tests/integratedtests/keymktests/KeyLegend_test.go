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

package keymktests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/keymk"
)

func Test_KeyLegend_GroupIntRange_Verification(t *testing.T) {
	for caseIndex, testCase := range keyLegendGroupIntRangeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		root, _ := input.GetAsString("root")
		pkg, _ := input.GetAsString("package")
		group, _ := input.GetAsString("group")
		state, _ := input.GetAsString("state")
		startId := input.GetAsIntDefault("startId", 0)
		endId := input.GetAsIntDefault("endId", 0)

		// Act
		k := keymk.NewKeyWithLegend.All(
			keymk.JoinerOption,
			keymk.ShortLegends,
			false,
			root, pkg, group, state,
		)
		result := k.GroupIntRange(startId, endId)

		actual := args.Map{
			"count":    fmt.Sprintf("%d", len(result)),
			"firstKey": result[0],
			"lastKey":  result[len(result)-1],
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_KeyLegend_UserStringWithoutState_Verification(t *testing.T) {
	for caseIndex, testCase := range keyLegendUserStringWithoutStateTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		root, _ := input.GetAsString("root")
		pkg, _ := input.GetAsString("package")
		group, _ := input.GetAsString("group")
		state, _ := input.GetAsString("state")
		user, _ := input.GetAsString("user")

		// Act
		k := keymk.NewKeyWithLegend.All(
			keymk.JoinerOption,
			keymk.ShortLegends,
			false,
			root, pkg, group, state,
		)
		result := k.UserStringWithoutState(user)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_KeyLegend_UpToState_Verification(t *testing.T) {
	for caseIndex, testCase := range keyLegendUpToStateTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		root, _ := input.GetAsString("root")
		pkg, _ := input.GetAsString("package")
		group, _ := input.GetAsString("group")
		state, _ := input.GetAsString("state")
		user, _ := input.GetAsString("user")

		// Act
		k := keymk.NewKeyWithLegend.All(
			keymk.JoinerOption,
			keymk.ShortLegends,
			false,
			root, pkg, group, state,
		)
		result := k.UpToState(user)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}
