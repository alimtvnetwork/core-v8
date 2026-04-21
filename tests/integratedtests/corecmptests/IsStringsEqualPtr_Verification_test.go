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

package corecmptests

import (
	"testing"
	"time"

	"github.com/alimtvnetwork/core-v8/corecmp"
	"github.com/alimtvnetwork/core-v8/corecomparator"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

func Test_IsStringsEqualPtr_Verification(t *testing.T) {
	for caseIndex, testCase := range isStringsEqualPtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftNil := input.GetAsBoolDefault("leftNil", false)
		rightNil := input.GetAsBoolDefault("rightNil", false)

		var left, right []string
		if !leftNil {
			rawLeft, has := input.Get("left")
			if has {
				if sl, ok := rawLeft.([]string); ok {
					left = sl
				} else {
					left = []string{}
				}
			} else {
				left = []string{}
			}
		}
		if !rightNil {
			rawRight, has := input.Get("right")
			if has {
				if sl, ok := rawRight.([]string); ok {
					right = sl
				} else {
					right = []string{}
				}
			} else {
				right = []string{}
			}
		}

		// Act
		result := corecmp.IsStringsEqualPtr(left, right)

		actual := args.Map{"result": result}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_TimePtr_Verification(t *testing.T) {
	for caseIndex, testCase := range timePtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftNil := input.GetAsBoolDefault("leftNil", false)
		rightNil := input.GetAsBoolDefault("rightNil", false)

		var left, right *time.Time
		now := time.Now()

		if !leftNil {
			left = &now
		}
		if !rightNil {
			sameTime := input.GetAsBoolDefault("sameTime", false)
			if sameTime {
				right = &now
			} else {
				later := now.Add(time.Hour)
				right = &later
			}
		}

		// Act
		result := corecmp.TimePtr(left, right)

		actual := args.Map{
			"isEqual": result == corecomparator.Equal,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
