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

package coredata

import (
	"errors"
	"fmt"

	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/coreindexes"
	"github.com/alimtvnetwork/core-v8/internal/csvinternal"
)

type BytesError struct {
	toString *string
	Bytes    []byte
	Error    error
}

func (it *BytesError) String() string {
	return *it.StringPtr()
}

func (it *BytesError) StringPtr() *string {
	if it == nil {
		strEmpty := ""

		return &strEmpty
	}

	if it.toString != nil {
		return it.toString
	}

	if it.HasBytes() {
		jsonString := string(it.Bytes)
		it.toString = &jsonString
	} else {
		emptyStr := ""
		it.toString = &emptyStr
	}

	return it.toString
}

func (it *BytesError) CombineErrorWithRef(references ...string) string {
	if it.IsEmptyError() {
		return ""
	}

	csv := csvinternal.StringsToStringDefault(references...)

	return fmt.Sprintf(
		constants.MessageReferenceWrapFormat,
		it.Error.Error(),
		csv)
}

func (it *BytesError) CombineErrorWithRefError(references ...string) error {
	if it.IsEmptyError() {
		return nil
	}

	errorString := it.CombineErrorWithRef(
		references...)

	return errors.New(errorString)
}

func (it *BytesError) HasError() bool {
	return it != nil && it.Error != nil
}

func (it *BytesError) IsEmptyError() bool {
	return it == nil || it.Error == nil
}

func (it *BytesError) HandleError() {
	if it.IsEmptyError() {
		return
	}

	panic(it.Error)
}

func (it *BytesError) HandleErrorWithMsg(msg string) {
	if it.IsEmptyError() {
		return
	}

	if msg != "" {
		panic(msg + it.Error.Error())
	}

	panic(it.Error)
}

func (it *BytesError) HasBytes() bool {
	return !it.IsEmptyOrErrorBytes()
}

// IsEmptyOrErrorBytes len == 0, nil, {} returns as empty true
func (it *BytesError) IsEmptyOrErrorBytes() bool {
	isEmptyFirst := it.HasError() ||
		it.Bytes == nil

	if isEmptyFirst {
		return isEmptyFirst
	}

	length := len(it.Bytes)

	if length == 0 {
		return true
	}

	if length == 2 {
		// empty json
		return it.Bytes[coreindexes.First] == 123 &&
			it.Bytes[coreindexes.Second] == 125
	}

	return false
}

func (it *BytesError) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Bytes)
}

func (it *BytesError) IsEmpty() bool {
	return it == nil || len(it.Bytes) == 0
}
