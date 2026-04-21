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

package corestr

import "github.com/alimtvnetwork/core-v8/constants"

type TextWithLineNumber struct {
	LineNumber int
	Text       string
}

func (it *TextWithLineNumber) HasLineNumber() bool {
	return it != nil && it.LineNumber > constants.InvalidValue
}

func (it *TextWithLineNumber) IsInvalidLineNumber() bool {
	return it == nil || it.LineNumber == constants.InvalidValue
}

func (it *TextWithLineNumber) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Text)
}

func (it *TextWithLineNumber) IsEmpty() bool {
	if it == nil {
		return true
	}

	return it.IsEmptyText() || it.IsInvalidLineNumber()
}

func (it *TextWithLineNumber) IsEmptyText() bool {
	if it == nil {
		return true
	}

	return len(it.Text) == 0
}

func (it *TextWithLineNumber) IsEmptyTextLineBoth() bool {
	return it.IsEmpty()
}
