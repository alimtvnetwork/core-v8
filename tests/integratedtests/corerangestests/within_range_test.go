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

package corerangestests

import (
	"testing"

	"github.com/alimtvnetwork/core/converters"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
)

func Test_RangeInt_Valid_WithInRange_Verification(t *testing.T) {
	// Arrange
	validCases := []int{
		5, 13, 5, 10, 25,
	}
	toString := converters.AnyTo.ValueString(validCases)

	// Act, Assert
	title := toString + " -- all these are valid for (range) : " + someRange.String()
	convey.Convey(
		title, t, func() {
			for _, v := range validCases {
				for name, inWithFunc := range isWithInFuncsMap {
					validationErr := rangeValidationError(
						name,
						true,
						inWithFunc,
						v,
					)

					convey.So(
						validationErr,
						should.BeNil,
					)
				}
			}
		},
	)
}

func rangeValidationError(
	name string,
	isExpectValid bool,
	isWithInFunc isWithInDefinitionFunc,
	v int,
) error {
	isInRange := isWithInFunc(v)

	if !isInRange && isExpectValid {
		return errcore.WasExpectingErrorF(
			true,
			false,
			"%s - should be valid and within range : %d",
			name,
			v,
		)
	}

	if isInRange && !isExpectValid {
		return errcore.WasExpectingErrorF(
			true,
			false,
			"%s - should be invalid and within range : %d",
			name,
			v,
		)
	}

	return nil
}

func Test_RangeInt_Invalid_WithInRange_Verification(t *testing.T) {
	// Arrange
	invalidCases := []int{
		265, 311, 4, 26, 100,
	}
	toString := converters.AnyTo.ValueString(invalidCases)

	// Act, Assert
	title := toString + " -- all these are invalid for (range) : " + someRange.String()
	convey.Convey(
		title, t, func() {
			for _, v := range invalidCases {
				for name, inWithFunc := range isWithInFuncsMap {
					validationErr := rangeValidationError(
						name,
						false,
						inWithFunc,
						v,
					)

					convey.So(
						validationErr,
						should.BeNil,
					)
				}
			}
		},
	)
}
