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
	"log/slog"
	"testing"

	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/coredata/corestr"
	"github.com/alimtvnetwork/core-v8/errcore"
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
)

type HeaderSliceValidators []HeaderSliceValidator

func (it HeaderSliceValidators) Length() int {
	if it == nil {
		return 0
	}

	return len(it)
}

func (it HeaderSliceValidators) IsEmpty() bool {
	return it == nil || len(it) == 0
}

func (it HeaderSliceValidators) SetActualOnAll(actualLines ...string) {
	if it.IsEmpty() {
		return
	}

	for _, sliceValidator := range it {
		sliceValidator.SetActual(actualLines)
	}
}

func (it HeaderSliceValidators) IsValid(
	isCaseSensitive bool,
) bool {
	return it.IsMatch(isCaseSensitive)
}

func (it HeaderSliceValidators) IsMatch(
	isCaseSensitive bool,
) bool {
	if it.IsEmpty() {
		return true
	}

	for _, sliceValidator := range it {
		if !sliceValidator.IsValid(isCaseSensitive) {
			return false
		}
	}

	return true
}

func (it HeaderSliceValidators) VerifyAll(
	header string,
	params *Parameter,
	isPrintError bool,
) error {
	if it.IsEmpty() {
		return nil
	}

	errs := corestr.New.SimpleSlice.Cap(constants.Capacity2)

	for _, sliceValidator := range it {
		err := sliceValidator.AllVerifyError(params)

		if err != nil {
			diffMsg := errcore.LineDiffToString(
				params.CaseIndex,
				params.Header,
				sliceValidator.ActualLines,
				sliceValidator.ExpectedLines,
			)

			errs.AddError(err)

			if len(diffMsg) > 0 {
				errs.Add(diffMsg)
			}
		}
	}

	if errs.IsEmpty() {
		return nil
	}

	errs.InsertAt(0, header)
	err := errs.AsDefaultError()

	if isPrintError {
		slog.Error("verification failed", "error", err)
	}

	return err
}

func (it HeaderSliceValidators) AssertVerifyAll(
	t *testing.T,
	params *Parameter,
) {
	if it.IsEmpty() {
		return
	}

	finalError := it.VerifyAllError(params)

	convey.Convey(params.Header, t, func() {
		convey.So(
			finalError,
			should.BeNil)
	})
}

func (it HeaderSliceValidators) VerifyAllError(
	params *Parameter,
) error {
	if it.IsEmpty() {
		return nil
	}

	errs := corestr.New.SimpleSlice.Cap(constants.Capacity2)

	for _, sliceValidator := range it {
		err := sliceValidator.AllVerifyError(params)

		if err != nil {
			diffMsg := errcore.LineDiffToString(
				params.CaseIndex,
				params.Header,
				sliceValidator.ActualLines,
				sliceValidator.ExpectedLines,
			)

			errs.AddError(err)

			if len(diffMsg) > 0 {
				errs.Add(diffMsg)
			}
		}
	}

	header := params.Header

	errs.InsertAt(0, header)

	return errs.AsDefaultError()
}

func (it HeaderSliceValidators) AssertVerifyAllUsingActual(
	t *testing.T,
	params *Parameter,
	actualLines ...string,
) {
	if it.IsEmpty() {
		return
	}

	finalError := it.VerifyAllErrorUsingActual(
		params,
		actualLines...)

	convey.Convey(params.Header, t, func() {
		convey.So(
			finalError,
			should.BeNil)
	})
}

func (it HeaderSliceValidators) VerifyAllErrorUsingActual(
	params *Parameter,
	actualLines ...string,
) error {
	if it.IsEmpty() {
		return nil
	}

	errs := corestr.New.SimpleSlice.Cap(constants.Capacity2)

	for _, sliceValidator := range it {
		sliceValidator.SetActual(actualLines)
		err := sliceValidator.AllVerifyError(params)

		if err != nil {
			diffMsg := errcore.LineDiffToString(
				params.CaseIndex,
				params.Header,
				sliceValidator.ActualLines,
				sliceValidator.ExpectedLines,
			)

			errs.AddError(err)

			if len(diffMsg) > 0 {
				errs.Add(diffMsg)
			}
		}
	}

	if errs.IsEmpty() {
		return nil
	}

	header := params.Header

	errs.InsertAt(0, header)

	return errs.AsDefaultError()
}

// VerifyFirst
//
// Only collect using the SliceValidator.VerifyFirstError
func (it HeaderSliceValidators) VerifyFirst(
	params *Parameter,
	isPrintError bool,
) error {
	if it.IsEmpty() {
		return nil
	}

	errs := corestr.New.SimpleSlice.Cap(constants.Capacity2)

	for _, sliceValidator := range it {
		err := sliceValidator.VerifyFirstError(params)

		if err != nil {
			diffMsg := errcore.LineDiffToString(
				params.CaseIndex,
				params.Header,
				sliceValidator.ActualLines,
				sliceValidator.ExpectedLines,
			)

			errs.AddError(err)

			if len(diffMsg) > 0 {
				errs.Add(diffMsg)
			}
		}
	}

	if errs.IsEmpty() {
		return nil
	}

	header := params.Header

	errs.InsertAt(0, header)
	err := errs.AsDefaultError()

	if isPrintError {
		slog.Error("verification failed", "error", err)
	}

	return err
}

func (it HeaderSliceValidators) VerifyUpto(
	isPrintErr,
	isFirstOnly bool,
	length int,
	params *Parameter,
) error {
	if it.IsEmpty() {
		return nil
	}

	errs := corestr.New.SimpleSlice.Cap(constants.Capacity2)

	for _, sliceValidator := range it {
		err := sliceValidator.AllVerifyErrorUptoLength(
			isFirstOnly,
			params,
			length)

		if err != nil {
			diffMsg := errcore.LineDiffToString(
				params.CaseIndex,
				params.Header,
				sliceValidator.ActualLines,
				sliceValidator.ExpectedLines,
			)

			errs.AddError(err)

			if len(diffMsg) > 0 {
				errs.Add(diffMsg)
			}
		}
	}

	if errs.IsEmpty() {
		return nil
	}

	errs.InsertAt(0, params.Header)
	err := errs.AsDefaultError()

	if isPrintErr {
		slog.Error("verification failed", "error", err)
	}

	return err
}
