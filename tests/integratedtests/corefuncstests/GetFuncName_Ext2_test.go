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
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/corefuncs"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ============================================================================
// GetFuncName
// ============================================================================

func Test_GetFuncName_Ext2_Verification(t *testing.T) {
	for caseIndex, tc := range ext2GetFuncNameTestCases {
		// Act
		name := corefuncs.GetFuncName(sampleFunc)

		// Assert
		actual := args.Map{
			"notEmpty": name != "",
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// GetFuncFullName
// ============================================================================

func Test_GetFuncFullName_Ext2_Verification(t *testing.T) {
	for caseIndex, tc := range ext2GetFuncFullNameTestCases {
		// Act
		name := corefuncs.GetFuncFullName(sampleFunc)

		// Assert
		actual := args.Map{
			"notEmpty": name != "",
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// NewInOutErrWrapper
// ============================================================================

func Test_NewInOutErrWrapper_Success_Ext2_Verification(t *testing.T) {
	for caseIndex, tc := range ext2NewInOutErrWrapperTestCases {
		// Arrange
		wrapper := corefuncs.NewInOutErrWrapper[string, int](
			"parse",
			func(s string) (int, error) { return len(s), nil },
		)

		// Act
		result, err := wrapper.Exec("hello")

		// Assert
		actual := args.Map{
			"result": result,
			"isNil":  err == nil,
			"name":   wrapper.Name,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_NewInOutErrWrapper_Fail_Ext2_Verification(t *testing.T) {
	for caseIndex, tc := range ext2NewInOutErrWrapperFailTestCases {
		// Arrange
		wrapper := corefuncs.NewInOutErrWrapper[string, int](
			"parse",
			func(s string) (int, error) {
				if s == "" {
					return 0, errors.New("empty input")
				}
				return len(s), nil
			},
		)

		// Act
		_, err := wrapper.Exec("")

		// Assert
		actual := args.Map{
			"hasError": err != nil,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// InOutErrFuncWrapperOf.ToLegacy
// ============================================================================

func Test_InOutErrWrapper_ToLegacy_Ext2_Verification(t *testing.T) {
	for caseIndex, tc := range ext2InOutErrToLegacyTestCases {
		// Arrange
		wrapper := corefuncs.NewInOutErrWrapper[string, int](
			"legacy-parse",
			func(s string) (int, error) { return len(s), nil },
		)

		// Act
		legacy := wrapper.ToLegacy()
		result, err := legacy.Exec("hello")

		// Assert
		actual := args.Map{
			"name":   legacy.Name,
			"result": fmt.Sprintf("%v", result),
			"isNil":  err == nil,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// ============================================================================
// SerializeOutputFuncWrapperOf.AsActionReturnsErrorFunc
// ============================================================================

func Test_SerializeWrapper_AsErrFunc_Ext2_Verification(t *testing.T) {
	for caseIndex, tc := range ext2SerializeAsErrFuncTestCases {
		// Arrange
		wrapper := corefuncs.NewSerializeWrapper[string](
			"serialize",
			func(s string) ([]byte, error) { return []byte(s), nil },
		)

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc("test")
		err := errFunc()

		// Assert
		actual := args.Map{
			"isNil": err == nil,
		}
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
