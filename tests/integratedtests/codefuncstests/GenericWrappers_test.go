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

package codefuncstests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core-v8/corefuncs"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// =============================================================================
// InOutErrFuncWrapperOf — Exec
// =============================================================================

func Test_InOutErrFuncWrapperOf_Exec_Verification(t *testing.T) {
	for caseIndex, tc := range inOutErrOfExecTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsString("input")
		hasErr, _ := input.GetAsBool("hasActionErr")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.NewInOutErrWrapper[string, int](
			name,
			makeStrToIntErrFunc(hasErr),
		)

		// Act
		output, err := wrapper.Exec(inputVal)
		actual := args.Map{
			"output":   output,
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// InOutErrFuncWrapperOf — AsActionReturnsErrorFunc
// =============================================================================

func Test_InOutErrFuncWrapperOf_AsActionReturnsErrorFunc_Verification(t *testing.T) {
	for caseIndex, tc := range inOutErrOfAsActionReturnsErrorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsString("input")
		hasErr, _ := input.GetAsBool("hasActionErr")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.NewInOutErrWrapper[string, int](
			name,
			makeStrToIntErrFunc(hasErr),
		)

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc(inputVal)
		err := errFunc()
		actual := args.Map{
			"hasError": err != nil,
		}
		if err != nil {
			actual["containsName"] = strings.Contains(err.Error(), name)
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// InOutErrFuncWrapperOf — ToLegacy
// =============================================================================

func Test_InOutErrFuncWrapperOf_ToLegacy_Verification(t *testing.T) {
	for caseIndex, tc := range inOutErrOfToLegacyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsString("input")
		hasErr, _ := input.GetAsBool("hasActionErr")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.NewInOutErrWrapper[string, int](
			name,
			makeStrToIntErrFunc(hasErr),
		)
		legacy := wrapper.ToLegacy()

		// Act
		output, err := legacy.Exec(inputVal)
		actual := args.Map{
			"output":   output.(int),
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// InOutFuncWrapperOf — Exec
// =============================================================================

func Test_InOutFuncWrapperOf_Exec_Verification(t *testing.T) {
	for caseIndex, tc := range inOutFuncOfExecTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsString("input")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.NewInOutWrapper[string, string](
			name,
			makeStrToStrFunc(),
		)

		// Act
		output := wrapper.Exec(inputVal)
		actual := args.Map{
			"output": output,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// InOutFuncWrapperOf — AsActionReturnsErrorFunc
// =============================================================================

func Test_InOutFuncWrapperOf_AsActionReturnsErrorFunc_Verification(t *testing.T) {
	for caseIndex, tc := range inOutFuncOfAsActionReturnsErrorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsString("input")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.NewInOutWrapper[string, string](
			name,
			makeStrToStrFunc(),
		)

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc(inputVal)
		err := errFunc()
		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// InActionReturnsErrFuncWrapperOf — Exec
// =============================================================================

func Test_InActionReturnsErrFuncWrapperOf_Exec_Verification(t *testing.T) {
	for caseIndex, tc := range inActionReturnsErrOfExecTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsString("input")
		hasErr, _ := input.GetAsBool("hasActionErr")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.NewInActionErrWrapper[string](
			name,
			makeStrErrFunc(hasErr),
		)

		// Act
		err := wrapper.Exec(inputVal)
		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// InActionReturnsErrFuncWrapperOf — AsActionReturnsErrorFunc
// =============================================================================

func Test_InActionReturnsErrFuncWrapperOf_AsActionReturnsErrorFunc_Verification(t *testing.T) {
	for caseIndex, tc := range inActionReturnsErrOfAsActionReturnsErrorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsString("input")
		hasErr, _ := input.GetAsBool("hasActionErr")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.NewInActionErrWrapper[string](
			name,
			makeStrErrFunc(hasErr),
		)

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc(inputVal)
		err := errFunc()
		actual := args.Map{
			"hasError": err != nil,
		}
		if err != nil {
			actual["containsName"] = strings.Contains(err.Error(), name)
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// InActionReturnsErrFuncWrapperOf — ToLegacy
// =============================================================================

func Test_InActionReturnsErrFuncWrapperOf_ToLegacy_Verification(t *testing.T) {
	for caseIndex, tc := range inActionReturnsErrOfToLegacyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsString("input")
		hasErr, _ := input.GetAsBool("hasActionErr")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.NewInActionErrWrapper[string](
			name,
			makeStrErrFunc(hasErr),
		)
		legacy := wrapper.ToLegacy()

		// Act
		_, err := legacy.Exec(inputVal)
		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// ResultDelegatingFuncWrapperOf — Exec
// =============================================================================

func Test_ResultDelegatingFuncWrapperOf_Exec_Verification(t *testing.T) {
	for caseIndex, tc := range resultDelegatingOfExecTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		hasErr, _ := input.GetAsBool("hasActionErr")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.NewResultDelegatingWrapper[*fillTarget](
			name,
			makeResultDelegatingFunc(hasErr),
		)
		target := &fillTarget{}

		// Act
		err := wrapper.Exec(target)
		actual := args.Map{
			"hasError": err != nil,
			"filled":   target.Filled,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// ResultDelegatingFuncWrapperOf — AsActionReturnsErrorFunc
// =============================================================================

func Test_ResultDelegatingFuncWrapperOf_AsActionReturnsErrorFunc_Verification(t *testing.T) {
	for caseIndex, tc := range resultDelegatingOfAsActionReturnsErrorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		hasErr, _ := input.GetAsBool("hasActionErr")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.NewResultDelegatingWrapper[*fillTarget](
			name,
			makeResultDelegatingFunc(hasErr),
		)
		target := &fillTarget{}

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc(target)
		err := errFunc()
		actual := args.Map{
			"hasError": err != nil,
		}
		if err != nil {
			actual["containsName"] = strings.Contains(err.Error(), name)
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// ResultDelegatingFuncWrapperOf — ToLegacy
// =============================================================================

func Test_ResultDelegatingFuncWrapperOf_ToLegacy_Verification(t *testing.T) {
	for caseIndex, tc := range resultDelegatingOfToLegacyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		hasErr, _ := input.GetAsBool("hasActionErr")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.NewResultDelegatingWrapper[*fillTarget](
			name,
			makeResultDelegatingFunc(hasErr),
		)
		target := &fillTarget{}
		legacy := wrapper.ToLegacy()

		// Act
		err := legacy.Exec(target)
		actual := args.Map{
			"hasError": err != nil,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// SerializeOutputFuncWrapperOf — Exec
// =============================================================================

func Test_SerializeOutputFuncWrapperOf_Exec_Verification(t *testing.T) {
	for caseIndex, tc := range serializeOutputOfExecTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsString("input")
		hasErr, _ := input.GetAsBool("hasActionErr")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.NewSerializeWrapper[string](
			name,
			makeSerializeFunc(hasErr),
		)

		// Act
		output, err := wrapper.Exec(inputVal)
		actual := args.Map{
			"hasError":  err != nil,
			"hasOutput": len(output) > 0,
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

// =============================================================================
// SerializeOutputFuncWrapperOf — AsActionReturnsErrorFunc
// =============================================================================

func Test_SerializeOutputFuncWrapperOf_AsActionReturnsErrorFunc_Verification(t *testing.T) {
	for caseIndex, tc := range serializeOutputOfAsActionReturnsErrorTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsString("input")
		hasErr, _ := input.GetAsBool("hasActionErr")
		name, _ := input.GetAsString("name")
		wrapper := corefuncs.NewSerializeWrapper[string](
			name,
			makeSerializeFunc(hasErr),
		)

		// Act
		errFunc := wrapper.AsActionReturnsErrorFunc(inputVal)
		err := errFunc()
		actual := args.Map{
			"hasError": err != nil,
		}
		if err != nil {
			actual["containsName"] = strings.Contains(err.Error(), name)
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
