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

package converterstests

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/alimtvnetwork/core-v8/converters"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/errcore"
)

// =============================================================================
// Tests: IntegersWithDefaults
// =============================================================================

func Test_StringsTo_IntegersWithDefaults_FromStringsTo(t *testing.T) {
	for caseIndex, testCase := range integersWithDefaultsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputRaw, rawFound := input.Get("input")
		if !rawFound {
			errcore.HandleErrMessage("input is required")
		}
		inputSlice := inputRaw.([]string)
		defaultInt, defFound := input.GetAsInt("defaultInt")
		if !defFound {
			errcore.HandleErrMessage("defaultInt is required")
		}

		// Act
		result := converters.StringsTo.IntegersWithDefaults(defaultInt, inputSlice...)
		actual := args.Map{
			"count":          result.Length(),
			"hadDefaultUsed": result.HasError(),
		}

		for i, v := range result.Values {
			actual[fmt.Sprintf("val%d", i)] = v
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
		)
	}
}

// =============================================================================
// Tests: BytesWithDefaults
// =============================================================================

func Test_StringsTo_BytesWithDefaults_FromStringsTo(t *testing.T) {
	for caseIndex, testCase := range bytesWithDefaultsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputRaw, rawFound := input.Get("input")
		if !rawFound {
			errcore.HandleErrMessage("input is required")
		}
		inputSlice := inputRaw.([]string)
		defaultByteRaw, defFound := input.Get("defaultByte")
		if !defFound {
			errcore.HandleErrMessage("defaultByte is required")
		}
		defaultByte := defaultByteRaw.(byte)

		// Act
		result := converters.StringsTo.BytesWithDefaults(defaultByte, inputSlice...)
		actual := args.Map{
			"count":          result.Length(),
			"hadDefaultUsed": result.HasError(),
		}

		for i, v := range result.Values {
			actual[fmt.Sprintf("val%d", i)] = int(v)
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
		)
	}
}

// =============================================================================
// Tests: CloneIf
// =============================================================================

func Test_StringsTo_CloneIf_FromStringsTo(t *testing.T) {
	for caseIndex, testCase := range cloneIfTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputRaw, rawFound := input.Get("input")
		if !rawFound {
			errcore.HandleErrMessage("input is required")
		}
		inputSlice := inputRaw.([]string)
		isCloneRaw, cloneFound := input.Get("isClone")
		if !cloneFound {
			errcore.HandleErrMessage("isClone is required")
		}
		isClone := isCloneRaw.(bool)

		// Act
		result := converters.StringsTo.CloneIf(isClone, inputSlice...)
		actual := args.Map{
			"count": len(result),
		}

		for i, v := range result {
			actual[fmt.Sprintf("item%d", i)] = v
		}

		// Check if it's a different slice (clone independence)
		isSamePointer := false
		if len(inputSlice) > 0 && len(result) > 0 {
			isSamePointer = unsafe.Pointer(&inputSlice[0]) == unsafe.Pointer(&result[0])
		}

		actual["isIndependent"] = isClone && !isSamePointer && len(inputSlice) > 0

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
		)
	}
}

// =============================================================================
// Tests: PtrOfPtrToPtrStrings
// =============================================================================

func Test_StringsTo_PtrOfPtrToPtrStrings_FromStringsTo(t *testing.T) {
	for caseIndex, testCase := range ptrOfPtrToPtrStringsTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)

		isNilRaw, _ := input.Get("isNil")
		isNil, _ := isNilRaw.(bool)
		isNilInnerRaw, _ := input.Get("isNilInner")
		isNilInner, _ := isNilInnerRaw.(bool)

		var result *[]string

		if isNil {
			// Act — nil outer pointer
			result = converters.StringsTo.PtrOfPtrToPtrStrings(nil)
		} else if isNilInner {
			// Act — nil inner pointer
			var nilInner []*string
			result = converters.StringsTo.PtrOfPtrToPtrStrings(&nilInner)
		} else {
			inputRaw, _ := input.Get("input")
			inputSlice := inputRaw.([]string)
			hasNilRaw, _ := input.Get("hasNil")
			hasNil, _ := hasNilRaw.(bool)

			// Build []*string
			ptrSlice := make([]*string, 0, len(inputSlice)+1)
			for i := range inputSlice {
				ptrSlice = append(ptrSlice, &inputSlice[i])
			}

			if hasNil {
				ptrSlice = append(ptrSlice, nil)
			}

			// Act
			result = converters.StringsTo.PtrOfPtrToPtrStrings(&ptrSlice)
		}

		// For cases with args.Map expected, use ShouldBeEqualMap
		if _, isMap := testCase.ExpectedInput.(args.Map); isMap {
			actual := args.Map{
				"count": len(*result),
			}

			for i, v := range *result {
				actual[fmt.Sprintf("item%d", i)] = v
			}

			testCase.ShouldBeEqualMap(
				t,
				caseIndex,
				actual,
			)
		} else {
			// Plain string expected
			actLines := []string{fmt.Sprintf("%d", len(*result))}

			testCase.ShouldBeEqual(t, caseIndex, actLines...)
		}
	}
}

// =============================================================================
// Tests: PtrOfPtrToMapStringBool
// =============================================================================

func Test_StringsTo_PtrOfPtrToMapStringBool_FromStringsTo(t *testing.T) {
	for caseIndex, testCase := range ptrOfPtrToMapStringBoolTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilRaw, _ := input.Get("isNil")
		isNil, _ := isNilRaw.(bool)

		var result map[string]bool

		if isNil {
			// Act — nil input
			result = converters.StringsTo.PtrOfPtrToMapStringBool(nil)
		} else {
			inputRaw, _ := input.Get("input")
			inputSlice := inputRaw.([]string)
			hasNilRaw, _ := input.Get("hasNil")
			hasNil, _ := hasNilRaw.(bool)

			ptrSlice := make([]*string, 0, len(inputSlice)+1)
			for i := range inputSlice {
				ptrSlice = append(ptrSlice, &inputSlice[i])
			}

			if hasNil {
				ptrSlice = append(ptrSlice, nil)
			}

			// Act
			result = converters.StringsTo.PtrOfPtrToMapStringBool(&ptrSlice)
		}

		// For cases with args.Map expected, use ShouldBeEqualMap
		if _, isMap := testCase.ExpectedInput.(args.Map); isMap {
			actual := args.Map{
				"count": len(result),
			}

			for k, v := range result {
				actual["has"+k] = v
			}

			testCase.ShouldBeEqualMap(
				t,
				caseIndex,
				actual,
			)
		} else {
			// Plain string expected
			actLines := []string{fmt.Sprintf("%d", len(result))}

			testCase.ShouldBeEqual(t, caseIndex, actLines...)
		}
	}
}
