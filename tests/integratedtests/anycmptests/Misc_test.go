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

package anycmptests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/anycmp"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// Coverage test for anycmp.Cmp to ensure all branches reach 100%.
// Existing CmpBranch tests cover the main cases; these cover edge paths
// in compound conditions.

var extCmpEdgeCaseTestCases = []coretestcases.CaseV1{
	{
		Title: "Cmp nil slice vs nil slice returns Equal (both reflection-null)",
		ArrangeInput: args.Map{
			"pair": args.TwoAny{First: ([]int)(nil), Second: ([]string)(nil)},
		},
		ExpectedInput: args.Map{"name": "Equal"},
	},
	{
		Title: "Cmp nil map vs nil map returns Equal (both reflection-null)",
		ArrangeInput: args.Map{
			"pair": args.TwoAny{First: (map[string]int)(nil), Second: (map[int]int)(nil)},
		},
		ExpectedInput: args.Map{"name": "Equal"},
	},
	{
		Title: "Cmp nil func vs nil func returns Equal (both reflection-null)",
		ArrangeInput: args.Map{
			"pair": args.TwoAny{First: (func())(nil), Second: (func(int))(nil)},
		},
		ExpectedInput: args.Map{"name": "Equal"},
	},
	{
		Title: "Cmp bool true vs true returns Equal (== match)",
		ArrangeInput: args.Map{
			"pair": args.TwoAny{First: true, Second: true},
		},
		ExpectedInput: args.Map{"name": "Equal"},
	},
	{
		Title: "Cmp bool true vs false returns Inconclusive",
		ArrangeInput: args.Map{
			"pair": args.TwoAny{First: true, Second: false},
		},
		ExpectedInput: args.Map{"name": "Inconclusive"},
	},
	{
		Title: "Cmp float64 same returns Equal",
		ArrangeInput: args.Map{
			"pair": args.TwoAny{First: 3.14, Second: 3.14},
		},
		ExpectedInput: args.Map{"name": "Equal"},
	},
}

func Test_Cmp_EdgeCases(t *testing.T) {
	for caseIndex, testCase := range extCmpEdgeCaseTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pairRaw, _ := input.Get("pair")
		pair := pairRaw.(args.TwoAny)

		// Act
		result := anycmp.Cmp(pair.First, pair.Second)

		actual := args.Map{
			"name": result.String(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
