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

func Test_StringOnce_Core(t *testing.T) {
	for caseIndex, tc := range stringOnceCoreTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewStringOncePtr(func() string { return initVal })

		// Act
		actual := args.Map{
			"value":               once.Value(),
			"string":              once.String(),
			"isEmpty":             once.IsEmpty(),
			"isEmptyOrWhitespace": once.IsEmptyOrWhitespace(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringOnce_Caching(t *testing.T) {
	for caseIndex, tc := range stringOnceCachingTestCases {
		// Arrange
		callCount := 0
		initVal := tc.InitValue
		once := coreonce.NewStringOncePtr(func() string {
			callCount++

			return initVal
		})

		// Act
		r1 := once.Value()
		r2 := once.Value()
		r3 := once.Value()

		// Assert
		actual := args.Map{
			"r1":        r1,
			"r2":        r2,
			"r3":        r3,
			"callCount": callCount,
		}
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringOnce_Match(t *testing.T) {
	for caseIndex, tc := range stringOnceMatchTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewStringOncePtr(func() string { return initVal })

		// Act
		var actual args.Map

		hasPrefix := tc.MatchArg == "prefix"
		hasSuffix := tc.MatchArg == "suffix"

		if hasPrefix {
			actual = args.Map{
				"matchResult":   once.HasPrefix(tc.MatchArg),
				"noMatchResult": once.HasPrefix("data"),
			}
		} else if hasSuffix {
			actual = args.Map{
				"matchResult":   once.HasSuffix(tc.MatchArg),
				"noMatchResult": once.HasSuffix("data"),
			}
		} else if tc.MatchArg == tc.InitValue {
			actual = args.Map{
				"matchResult":   once.IsEqual(tc.MatchArg),
				"noMatchResult": once.IsEqual("xyz"),
			}
		} else {
			actual = args.Map{
				"matchResult":   once.IsContains(tc.MatchArg),
				"noMatchResult": once.IsContains("xyz"),
			}
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringOnce_Split(t *testing.T) {
	for caseIndex, tc := range stringOnceSplitTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewStringOncePtr(func() string { return initVal })

		// Act
		var actual args.Map

		switch tc.Method {
		case "splitBy":
			parts := once.SplitBy(tc.Splitter)
			actual = args.Map{
				"partsLength": len(parts),
				"firstPart":   parts[0],
				"lastPart":    parts[len(parts)-1],
			}
		case "splitLeftRightTrim":
			left, right := once.SplitLeftRightTrim(tc.Splitter)
			actual = args.Map{
				"left":  left,
				"right": right,
			}
		default: // splitLeftRight
			left, right := once.SplitLeftRight(tc.Splitter)
			actual = args.Map{
				"left":  left,
				"right": right,
			}
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_StringOnce_Json(t *testing.T) {
	for caseIndex, tc := range stringOnceJsonTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewStringOncePtr(func() string { return initVal })

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

// Ensure fmt is used
var _ = fmt.Sprintf
