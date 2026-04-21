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
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/coreonce"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

func Test_IntegerOnce_Core(t *testing.T) {
	for caseIndex, tc := range integerOnceCoreTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewIntegerOncePtr(func() int { return initVal })

		// Act
		actual := args.Map{
			"value":          once.Value(),
			"string":         once.String(),
			"isZero":         once.IsZero(),
			"isEmpty":        once.IsEmpty(),
			"isAboveZero":    once.IsAboveZero(),
			"isPositive":     once.IsPositive(),
			"isLessThanZero": once.IsLessThanZero(),
			"isNegative":     once.IsNegative(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerOnce_Caching(t *testing.T) {
	for caseIndex, tc := range integerOnceCachingTestCases {
		// Arrange
		callCount := 0
		initVal := tc.InitValue
		once := coreonce.NewIntegerOncePtr(func() int {
			callCount++

			return initVal
		})

		// Act
		r1 := once.Value()
		r2 := once.Value()

		// Assert
		actual := args.Map{
			"r1":        r1,
			"r2":        r2,
			"callCount": callCount,
		}
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerOnce_Compare(t *testing.T) {
	for caseIndex, tc := range integerOnceCompareTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewIntegerOncePtr(func() int { return initVal })
		cmpVal := tc.CompareValue

		// Act
		var actual args.Map

		isLessThanCase := tc.InitValue < tc.CompareValue
		if isLessThanCase {
			actual = args.Map{
				"isLessThanCompare":   once.IsLessThan(cmpVal),
				"isLessThanSelf":      once.IsLessThan(initVal),
				"isLessThanEqualSelf": once.IsLessThanEqual(initVal),
			}
		} else {
			actual = args.Map{
				"isAboveCompare":   once.IsAbove(cmpVal),
				"isAboveSelf":      once.IsAbove(initVal),
				"isAboveEqualSelf": once.IsAboveEqual(initVal),
			}
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegerOnce_Json(t *testing.T) {
	for caseIndex, tc := range integerOnceJsonTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewIntegerOncePtr(func() int { return initVal })

		// Act
		data, err := once.MarshalJSON()

		// Assert
		actual := args.Map{
			"noError":        err == nil,
			"marshaledValue": string(data),
		}
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// Ensure fmt is used (for potential future use)
var _ = fmt.Sprintf
