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

package corevalidator

import (
	"fmt"

	"github.com/alimtvnetwork/core-v8/coredata/corestr"
)

type RangeSegmentsValidator struct {
	actual           *corestr.SimpleSlice
	Title            string
	VerifierSegments []RangesSegment
}

func (it *RangeSegmentsValidator) LengthOfVerifierSegments() int {
	return len(it.VerifierSegments)
}

func (it *RangeSegmentsValidator) SetActual(
	lines []string,
) *RangeSegmentsValidator {
	it.actual = corestr.New.SimpleSlice.Direct(
		false,
		lines,
	)

	return it
}

func (it *RangeSegmentsValidator) Validators() HeaderSliceValidators {
	validators := make([]HeaderSliceValidator, 0, it.LengthOfVerifierSegments())

	for _, segment := range it.VerifierSegments {
		expectedSegments := segment.ExpectedLines
		start := segment.RangeInt.Start
		end := segment.RangeInt.End
		actualSegments := it.actual.Strings()[start:end]
		totalItems := end - start + 1
		header := fmt.Sprintf(
			"%s - validate for range %d to %d (total: %d lines)",
			it.Title,
			start,
			end,
			totalItems,
		)
		validator := HeaderSliceValidator{
			Header: header,
			SliceValidator: SliceValidator{
				Condition:     segment.Condition,
				CompareAs:     segment.CompareAs,
				ActualLines:   actualSegments,
				ExpectedLines: expectedSegments,
			},
		}

		validators = append(validators, validator)
	}

	return validators
}

func (it *RangeSegmentsValidator) VerifyAll(
	header string,
	actual []string,
	params *Parameter,
	isPrintError bool,
) error {
	it.SetActual(actual)

	return it.Validators().VerifyAll(
		header,
		params,
		isPrintError,
	)
}

func (it *RangeSegmentsValidator) VerifySimple(
	actual []string,
	params *Parameter,
	isPrintError bool,
) error {
	return it.VerifyAll(
		it.Title,
		actual,
		params,
		isPrintError,
	)
}

func (it *RangeSegmentsValidator) VerifyFirst(
	header string,
	actual []string,
	params *Parameter,
	isPrintError bool,
) error {
	params.Header = header
	it.SetActual(actual)

	return it.Validators().VerifyFirst(
		params,
		isPrintError,
	)
}

func (it *RangeSegmentsValidator) VerifyUpto(
	header string,
	actual []string,
	params *Parameter,
	length int,
	isPrintError bool,
) error {
	params.Header = header
	it.SetActual(actual)

	return it.Validators().VerifyUpto(
		isPrintError,
		false,
		length,
		params,
	)
}

func (it *RangeSegmentsValidator) VerifyFirstDefault(
	actual []string,
	params *Parameter,
	isPrintError bool,
) error {
	return it.VerifyFirst(
		it.Title,
		actual,
		params,
		isPrintError,
	)
}

func (it *RangeSegmentsValidator) VerifyUptoDefault(
	actual []string,
	params *Parameter,
	length int,
	isPrintError bool,
) error {
	return it.VerifyUpto(
		it.Title,
		actual,
		params,
		length,
		isPrintError,
	)
}
