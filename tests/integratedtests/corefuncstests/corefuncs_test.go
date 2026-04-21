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

package corefuncstests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core-v8/corefuncs"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// sampleFunc is a helper for GetFuncName tests.
func sampleFunc() {}

func Test_GetFuncName_Verification(t *testing.T) {
	for caseIndex, tc := range getFuncNameTestCases {
		// Act
		shortName := corefuncs.GetFuncName(sampleFunc)
		fullName := corefuncs.GetFuncFullName(sampleFunc)

		// Assert
		actual := args.Map{
			"hasShortName":        len(shortName) > 0,
			"fullLongerThanShort": len(fullName) > len(shortName),
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ActionReturnsErrorFuncWrapper_Success_Verification(t *testing.T) {
	for caseIndex, tc := range actionErrWrapperSuccessTestCases {
		// Arrange
		wrapper := corefuncs.New.ActionErr("cleanup", func() error {
			return nil
		})

		// Act
		err := wrapper.Exec()

		// Assert
		actual := args.Map{
			"isNil": err == nil,
			"name":  wrapper.Name,
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ActionReturnsErrorFuncWrapper_Failure_Verification(t *testing.T) {
	for caseIndex, tc := range actionErrWrapperFailureTestCases {
		// Arrange
		wrapper := corefuncs.New.ActionErr("cleanup", func() error {
			return errors.New("cleanup failed")
		})

		// Act
		err := wrapper.Exec()

		// Assert
		actual := args.Map{
			"isNil":    err == nil,
			"hasError": err != nil,
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsSuccessFuncWrapper_Verification(t *testing.T) {
	results := []bool{true, false}

	for caseIndex, tc := range isSuccessWrapperTestCases {
		// Arrange
		expectedResult := results[caseIndex]
		wrapper := corefuncs.New.IsSuccess("checker", func() bool {
			return expectedResult
		})

		// Act & Assert
		actual := args.Map{
			"result": wrapper.Exec(),
			"name":   wrapper.Name,
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_InOutErrFuncWrapperOf_Success_Verification(t *testing.T) {
	for caseIndex, tc := range inOutErrWrapperOfSuccessTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		wrapper := corefuncs.NewInOutErrWrapper[string, int](
			"strlen",
			func(s string) (int, error) {
				return len(s), nil
			},
		)

		// Act
		result, err := wrapper.Exec(inputStr)

		// Assert
		actual := args.Map{
			"result": result,
			"isNil":  err == nil,
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_InOutErrFuncWrapperOf_Failure_Verification(t *testing.T) {
	for caseIndex, tc := range inOutErrWrapperOfFailureTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		wrapper := corefuncs.NewInOutErrWrapper[string, int](
			"strlen",
			func(s string) (int, error) {
				if s == "" {
					return 0, errors.New("empty input")
				}
				return len(s), nil
			},
		)

		// Act
		result, err := wrapper.Exec(inputStr)

		// Assert
		actual := args.Map{
			"result": result,
			"isNil":  err == nil,
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NewCreator_ActionErr_Verification(t *testing.T) {
	for caseIndex, tc := range newCreatorActionErrTestCases {
		// Arrange
		wrapper := corefuncs.New.ActionErr("my-action", func() error {
			return nil
		})

		// Assert
		actual := args.Map{
			"name":      wrapper.Name,
			"hasAction": wrapper.Action != nil,
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NewCreator_IsSuccess_Verification(t *testing.T) {
	for caseIndex, tc := range newCreatorIsSuccessTestCases {
		// Arrange
		wrapper := corefuncs.New.IsSuccess("my-check", func() bool {
			return true
		})

		// Assert
		actual := args.Map{
			"name":      wrapper.Name,
			"hasAction": wrapper.Action != nil,
		}

		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
