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

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/errcore"
)

type LineValidator struct {
	LineNumber
	TextValidator
}

// IsMatch
//
// lineNumber == -1 mean no checking in line number,
//
// having LineValidator.LineNumber = -1 is also means the same
func (it *LineValidator) IsMatch(
	lineNumber int,
	content string,
	isCaseSensitive bool,
) bool {
	if !it.LineNumber.IsMatch(lineNumber) {
		return false
	}

	return it.TextValidator.IsMatch(
		content,
		isCaseSensitive)
}

func (it *LineValidator) IsMatchMany(
	isSkipOnContentsEmpty,
	isCaseSensitive bool,
	contentsWithLine ...corestr.TextWithLineNumber,
) bool {
	if it == nil {
		return true
	}

	if len(contentsWithLine) == 0 && isSkipOnContentsEmpty {
		return true
	}

	for _, textWithLine := range contentsWithLine {
		if !it.IsMatch(
			textWithLine.LineNumber,
			textWithLine.Text,
			isCaseSensitive) {
			return false
		}
	}

	return true
}

// VerifyError
//
// lineNumber == -1 mean no checking in line number,
//
// having LineValidator.LineNumber = -1 is also means the same
func (it *LineValidator) VerifyError(
	params *Parameter,
	processingLineNumber int,
	content string,
) error {
	if !it.LineNumber.IsMatch(processingLineNumber) {
		msg := errcore.GetSearchLineNumberExpectationMessage(
			params.CaseIndex,
			it.LineNumber.LineNumber,
			processingLineNumber,
			content,
			it.Search,
			*it)

		return errors.New(msg)
	}

	return it.TextValidator.verifyDetailErrorUsingLineProcessing(
		processingLineNumber,
		params,
		content)
}

func (it *LineValidator) VerifyMany(
	isContinueOnError bool,
	params *Parameter,
	contentsWithLine ...corestr.TextWithLineNumber,
) error {
	if isContinueOnError {
		return it.AllVerifyError(
			params,
			contentsWithLine...)
	}

	return it.VerifyFirstError(
		params,
		contentsWithLine...)
}

func (it *LineValidator) VerifyFirstError(
	params *Parameter,
	contentsWithLine ...corestr.TextWithLineNumber,
) error {
	if it == nil {
		return nil
	}

	length := len(contentsWithLine)
	if length == 0 && params.IsSkipCompareOnActualEmpty {
		return nil
	}

	for _, textWithLine := range contentsWithLine {
		err := it.VerifyError(
			params,
			textWithLine.LineNumber,
			textWithLine.Text,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *LineValidator) AllVerifyError(
	params *Parameter,
	contentsWithLine ...corestr.TextWithLineNumber,
) error {
	if it == nil {
		return nil
	}

	length := len(contentsWithLine)
	if length == 0 && params.IsSkipCompareOnActualEmpty {
		return nil
	}

	var sliceErr []string

	for _, textWithLine := range contentsWithLine {
		err := it.VerifyError(
			params,
			textWithLine.LineNumber,
			textWithLine.Text,
		)

		if err != nil {
			sliceErr = append(sliceErr, err.Error())
		}
	}

	return errcore.SliceToError(
		sliceErr)
}
