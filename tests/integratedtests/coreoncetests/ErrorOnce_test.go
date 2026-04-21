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
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/coreonce"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

func newErrorOnce(initError string) *coreonce.ErrorOnce {
	if initError == "" {
		return coreonce.NewErrorOncePtr(func() error { return nil })
	}

	if initError == "empty-marker" {
		return coreonce.NewErrorOncePtr(func() error { return errors.New("") })
	}

	return coreonce.NewErrorOncePtr(func() error { return errors.New(initError) })
}

func Test_ErrorOnce_Core(t *testing.T) {
	for caseIndex, tc := range errorOnceCoreTestCases {
		// Arrange
		once := newErrorOnce(tc.InitError)

		// Act
		actual := args.Map{
			"hasError":   once.HasError(),
			"isValid":    once.IsValid(),
			"isSuccess":  once.IsSuccess(),
			"isEmpty":    once.IsEmpty(),
			"isInvalid":  once.IsInvalid(),
			"isFailed":   once.IsFailed(),
			"hasAnyItem": once.HasAnyItem(),
			"isDefined":  once.IsDefined(),
			"message":    once.Message(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ErrorOnce_Caching(t *testing.T) {
	for caseIndex, tc := range errorOnceCachingTestCases {
		// Arrange
		callCount := 0
		initErr := tc.InitError
		once := coreonce.NewErrorOncePtr(func() error {
			callCount++

			return errors.New(initErr)
		})

		// Act
		r1 := once.Value()
		r2 := once.Value()
		r3 := once.Value()

		// Assert
		actual := args.Map{
			"r1":        r1.Error(),
			"r2":        r2.Error(),
			"r3":        r3.Error(),
			"callCount": callCount,
		}
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ErrorOnce_NullOrEmpty(t *testing.T) {
	for caseIndex, tc := range errorOnceNullOrEmptyTestCases {
		// Arrange
		once := newErrorOnce(tc.InitError)

		// Act
		actual := args.Map{
			"isNullOrEmpty": once.IsNullOrEmpty(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ErrorOnce_MessageEqual(t *testing.T) {
	for caseIndex, tc := range errorOnceMessageEqualTestCases {
		// Arrange
		once := newErrorOnce(tc.InitError)

		// Act
		actual := args.Map{
			"isMessageEqualMatch": once.IsMessageEqual(tc.MatchMsg),
			"isMessageEqualOther": once.IsMessageEqual("other"),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ErrorOnce_ConcatNew_FromErrorOnce(t *testing.T) {
	for caseIndex, tc := range errorOnceConcatTestCases {
		// Arrange
		once := newErrorOnce(tc.InitError)

		// Act
		result := once.ConcatNewString(tc.ExtraMsg)

		var actual args.Map

		isNilError := tc.InitError == ""
		if isNilError {
			actual = args.Map{
				"result": result,
			}
		} else {
			actual = args.Map{
				"containsBase":  strings.Contains(result, tc.InitError),
				"containsExtra": strings.Contains(result, tc.ExtraMsg),
			}
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ErrorOnce_Json(t *testing.T) {
	for caseIndex, tc := range errorOnceJsonTestCases {
		// Arrange
		once := newErrorOnce(tc.InitError)

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
