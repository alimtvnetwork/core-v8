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

package coreoncetests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreonce"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_IntegersOnce_Core(t *testing.T) {
	for caseIndex, tc := range integersOnceCoreTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewIntegersOncePtr(func() []int { return initVal })

		// Act
		actual := args.Map{
			"length":  len(once.Value()),
			"isEmpty": once.IsEmpty(),
			"isZero":  once.IsZero(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegersOnce_Sorted_FromIntegersOnce(t *testing.T) {
	for caseIndex, tc := range integersOnceSortedTestCases {
		// Arrange
		initVal := make([]int, len(tc.InitValue))
		copy(initVal, tc.InitValue)
		once := coreonce.NewIntegersOncePtr(func() []int { return initVal })

		// Act
		sorted := once.Sorted()

		actual := args.Map{
			"first": sorted[0],
			"last":  sorted[len(sorted)-1],
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegersOnce_Ranges(t *testing.T) {
	for caseIndex, tc := range integersOnceRangesTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewIntegersOncePtr(func() []int { return initVal })

		// Act
		rangesMap := once.RangesMap()
		rangesBoolMap := once.RangesBoolMap()
		uniqueMap := once.UniqueMap()

		_, has10 := rangesMap[10]
		actual := args.Map{
			"rangesMapLength":    len(rangesMap),
			"rangesBoolMapLen":   len(rangesBoolMap),
			"uniqueMapLen":       len(uniqueMap),
			"rangesMapHas10":     has10,
			"rangesBoolMapHas20": rangesBoolMap[20],
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegersOnce_IsEqual_FromIntegersOnce(t *testing.T) {
	for caseIndex, tc := range integersOnceIsEqualTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewIntegersOncePtr(func() []int { return initVal })

		// Act
		actual := args.Map{
			"isEqualSame":    once.IsEqual(1, 2, 3),
			"isEqualDiff":    once.IsEqual(1, 2, 4),
			"isEqualDiffLen": once.IsEqual(1, 2),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegersOnce_Caching(t *testing.T) {
	for caseIndex, tc := range integersOnceCachingTestCases {
		// Arrange
		callCount := 0
		initVal := tc.InitValue
		once := coreonce.NewIntegersOncePtr(func() []int {
			callCount++

			return initVal
		})

		// Act
		_ = once.Value()
		_ = once.Value()

		actual := args.Map{
			"callCount": callCount,
			"length":    len(once.Value()),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegersOnce_Json(t *testing.T) {
	for caseIndex, tc := range integersOnceJsonTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewIntegersOncePtr(func() []int { return initVal })

		// Act
		data, err := once.MarshalJSON()

		actual := args.Map{
			"noError":        err == nil,
			"marshaledValue": string(data),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegersOnce_Constructor(t *testing.T) {
	for caseIndex, tc := range integersOnceConstructorTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewIntegersOnce(func() []int { return initVal })

		// Act
		actual := args.Map{
			"length": len(once.Value()),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
