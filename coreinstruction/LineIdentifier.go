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

package coreinstruction

import (
	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/reqtype"
)

type LineIdentifier struct {
	LineNumber   int             `json:"LineNumber,omitempty"`
	LineModifyAs reqtype.Request `json:"LineModifyAs,omitempty"`
}

func (it *LineIdentifier) IsInvalidLineNumber() bool {
	return it == nil || it.LineNumber < 0
}

func (it *LineIdentifier) IsInvalidLineNumberUsingLastLineNumber(lastLineNumber int) bool {
	return it == nil || it.LineNumber < 0 || it.LineNumber > lastLineNumber
}

func (it *LineIdentifier) HasLineNumber() bool {
	return it != nil && it.LineNumber > constants.InvalidValue
}

func (it *LineIdentifier) IsNewLineRequest() bool {
	if it == nil {
		return false
	}

	return it.LineModifyAs.IsCreate()
}

func (it *LineIdentifier) IsDeleteLineRequest() bool {
	if it == nil {
		return false
	}

	return it.HasLineNumber() &&
		(it.LineModifyAs.IsDelete() ||
			it.LineModifyAs.IsDrop())
}

func (it *LineIdentifier) IsModifyLineRequest() bool {
	if it == nil {
		return false
	}

	return it.HasLineNumber() &&
		it.LineModifyAs.IsUpdate()
}

func (it *LineIdentifier) IsAddNewOrModifyLineRequest() bool {
	if it == nil {
		return false
	}

	return it.IsNewLineRequest() || it.IsModifyLineRequest()
}

func (it *LineIdentifier) ToBaseLineIdentifier() *BaseLineIdentifier {
	if it == nil {
		return nil
	}

	return NewBaseLineIdentifier(it.LineNumber, it.LineModifyAs)
}

func (it *LineIdentifier) Clone() *LineIdentifier {
	if it == nil {
		return nil
	}

	return &LineIdentifier{
		LineNumber:   it.LineNumber,
		LineModifyAs: it.LineModifyAs,
	}
}
