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

	"github.com/alimtvnetwork/core-v8/coredata/corerange"
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
)

func Test_Int_Ranges_ValidCases(t *testing.T) {
	for _, testCase := range validIntRangeTestCases {
		// Arrange
		arrangeInputs := testCase.Arrange()
		first := arrangeInputs[0]
		rest := arrangeInputs[1:]

		// Act
		actualRanges := first.CreateRanges(rest...)

		// Assert
		convey.Convey(
			testCase.Title, t, func() {
				convey.So(
					actualRanges,
					should.Equal,
					testCase.ExpectedInput,
				)
			},
		)

		convey.Convey(
			testCase.Title+" - type verify", t, func() {
				convey.So(
					testCase.TypeValidationError(),
					should.BeNil,
				)
			},
		)
	}
}

func Test_Int_ExceptRanges_Verify(t *testing.T) {
	// Arrange
	arrangeInput := corerange.MinMaxInt{
		Min: 1,
		Max: 15,
	}

	// Act
	actualRanges := arrangeInput.RangesExcept(
		3, 4, 5,
	)

	// Assert
	convey.Convey(
		"Ranges 1-15, RangesExcept(3, 4, 5), should not contain 3,4,5", t, func() {
			convey.So(
				actualRanges, should.Equal, []int{
					1, 2, 6,
					7, 8, 9,
					10, 11,
					12, 13,
					14, 15,
				},
			)
		},
	)
}

func Test_Int8_Ranges_ValidCases(t *testing.T) {
	for _, testCase := range validInt8RangeTestCases {
		// Arrange
		arrangeInputs := testCase.ArrangeInput.([]corerange.MinMaxInt8)
		first := arrangeInputs[0]
		rest := arrangeInputs[1:]

		// Act
		actualRanges := first.CreateRanges(rest...)

		// Assert
		convey.Convey(
			testCase.Title, t, func() {
				convey.So(
					actualRanges,
					should.Equal,
					testCase.ExpectedInput,
				)
			},
		)

		convey.Convey(
			testCase.Title+" - type verify", t, func() {
				convey.So(
					testCase.TypeValidationError(),
					should.BeNil,
				)
			},
		)
	}
}
