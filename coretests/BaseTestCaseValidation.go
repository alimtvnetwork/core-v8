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

package coretests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

// =============================================================================
// Type Validation
// =============================================================================

func (it *BaseTestCase) TypesValidationMustPasses(t *testing.T) {
	err := it.TypeValidationError()

	if err != nil {
		t.Error(
			"any one of the type validation failed",
			err.Error(),
		)
	}
}

// TypeValidationError
//
// must use SetActual to set the actual,
// what received from the act method,
// set it using SetActual
func (it *BaseTestCase) TypeValidationError() error {
	if it.IsTypeInvalidOrSkipVerify() {
		return nil
	}

	var sliceErr []string
	arrangeInputActualType := reflect.TypeOf(it.ArrangeInput)
	actualInputActualType := reflect.TypeOf(it.ActualInput)
	expectedInputActualType := reflect.TypeOf(it.ExpectedInput)
	verifyOf := it.VerifyTypeOf

	if reflectinternal.Is.Defined(it.ArrangeInput) && arrangeInputActualType != verifyOf.ArrangeInput {
		sliceErr = append(
			sliceErr,
			errcore.ExpectingSimpleNoType(
				"Arrange Type Mismatch",
				verifyOf.ArrangeInput,
				arrangeInputActualType,
			),
		)
	}

	if reflectinternal.Is.Defined(it.ActualInput) && actualInputActualType != verifyOf.ActualInput {
		sliceErr = append(
			sliceErr,
			errcore.ExpectingSimpleNoType(
				"Actual Type Mismatch",
				verifyOf.ActualInput,
				actualInputActualType,
			),
		)
	}

	if reflectinternal.Is.Defined(it.ExpectedInput) && expectedInputActualType != verifyOf.ExpectedInput {
		sliceErr = append(
			sliceErr,
			errcore.ExpectingSimpleNoType(
				"Expected Type Mismatch",
				verifyOf.ExpectedInput,
				expectedInputActualType,
			),
		)
	}

	if len(sliceErr) > 0 {
		var newSlice []string

		newSlice = append(
			newSlice,
			it.Title,
		)
		sliceErr = append(
			newSlice,
			sliceErr...,
		)
	}

	return errcore.SliceToError(sliceErr)
}
