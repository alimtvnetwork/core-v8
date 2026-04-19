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
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

type SimpleSliceValidator struct {
	Expected *corestr.SimpleSlice
	actual   *corestr.SimpleSlice
	Condition
	CompareAs stringcompareas.Variant
}

func (it *SimpleSliceValidator) SetActual(lines []string) *SimpleSliceValidator {
	it.actual = corestr.New.SimpleSlice.Direct(
		false,
		lines,
	)

	return it
}

func (it *SimpleSliceValidator) SliceValidator() *SliceValidator {
	var actualLines []string
	if it.actual != nil {
		actualLines = it.actual.Strings()
	}

	sliceValidator := SliceValidator{
		CompareAs:     it.CompareAs,
		Condition:     it.Condition,
		ActualLines:   actualLines,
		ExpectedLines: it.Expected.Strings(),
	}

	return &sliceValidator
}

func (it *SimpleSliceValidator) VerifyAll(
	actual []string,
	params *Parameter,
) error {
	sliceValidator := it.SliceValidator()
	sliceValidator.ActualLines = actual

	return sliceValidator.AllVerifyError(params)
}

func (it *SimpleSliceValidator) VerifyFirst(
	actual []string,
	params *Parameter,
) error {
	sliceValidator := it.SliceValidator()
	sliceValidator.ActualLines = actual

	return sliceValidator.VerifyFirstError(params)
}

func (it *SimpleSliceValidator) VerifyUpto(
	actual []string,
	params *Parameter,
	length int,
) error {
	sliceValidator := it.SliceValidator()
	sliceValidator.ActualLines = actual

	return sliceValidator.AllVerifyErrorUptoLength(
		false,
		params,
		length,
	)
}
