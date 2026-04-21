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

	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/errcore"
)

type LineNumber struct {
	LineNumber int `json:"LineNumber,omitempty"` // -1 means no checking in line
}

func (it *LineNumber) HasLineNumber() bool {
	return it.LineNumber > constants.InvalidValue
}

func (it *LineNumber) IsMatch(lineNumber int) bool {
	if lineNumber == constants.InvalidValue ||
		it.LineNumber == constants.InvalidValue {
		return true
	}

	return it.LineNumber == lineNumber
}

func (it *LineNumber) VerifyError(
	lineNumber int,
) error {
	if it.IsMatch(lineNumber) {
		return nil
	}

	msg := errcore.Expecting(
		"Line Number didn't match",
		it.LineNumber,
		lineNumber)

	return errors.New(msg)
}
