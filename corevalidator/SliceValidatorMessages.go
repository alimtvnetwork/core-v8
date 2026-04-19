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
	"errors"
	"fmt"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/errcore"
)

func (it *SliceValidator) ActualInputWithExpectingMessage(
	caseIndex int,
	header string,
) string {
	actualInputMessage := it.ActualInputMessage(
		caseIndex,
		header,
	)
	userExpectingMessage := it.UserExpectingMessage(
		caseIndex,
		header,
	)
	finalMessage := actualInputMessage +
		constants.NewLineUnix +
		constants.NewLineUnix +
		userExpectingMessage

	return finalMessage
}

func (it *SliceValidator) ActualInputMessage(
	caseIndex int,
	header string,
) string {
	finalHeader := fmt.Sprintf(
		actualUserInputsV2MessageFormat,
		caseIndex,
		header,
	)

	return errcore.MsgHeaderPlusEnding(
		finalHeader,
		it.ActualLinesString(),
	)
}

func (it *SliceValidator) UserExpectingMessage(
	caseIndex int,
	header string,
) string {
	finalHeader := fmt.Sprintf(
		expectingLinesV2MessageFormat,
		caseIndex,
		header,
	)

	return errcore.MsgHeaderPlusEnding(
		finalHeader,
		it.ExpectingLinesString(),
	)
}

// UserInputsMergeWithError
//
//   - Returns a combine error of actual and expecting inputs.
//   - If all validation successful then no error.
func (it *SliceValidator) UserInputsMergeWithError(
	parameter *Parameter,
	err error,
) error {
	if !parameter.IsAttachUserInputs {
		return err
	}

	toStr := it.ActualInputWithExpectingMessage(
		parameter.CaseIndex,
		parameter.Header,
	)

	if err == nil && len(toStr) == 0 {
		return nil
	}

	if err == nil && len(toStr) >= 0 {
		return errors.New(toStr)
	}

	msg := err.Error() + toStr

	return errors.New(msg)
}
