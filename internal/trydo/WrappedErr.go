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

package trydo

import (
	"fmt"
	"reflect"

	"github.com/alimtvnetwork/core-v8/constants"
)

type WrappedErr struct {
	Error     error
	Exception Exception
	HasThrown bool
	HasError  bool
}

func (it *WrappedErr) ExceptionType() reflect.Type {
	if it.IsInvalidException() {
		return nil
	}

	return reflect.TypeOf(it.Exception)
}

func (it *WrappedErr) IsDefined() bool {
	return it != nil
}

func (it *WrappedErr) IsInvalid() bool {
	return it == nil
}

func (it *WrappedErr) IsInvalidException() bool {
	if it.IsInvalid() {
		return true
	}

	if it.HasThrown {
		return it.Exception == nil
	}

	return true
}

func (it *WrappedErr) HasErrorOrException() bool {
	return it != nil &&
		(it.HasError ||
			it.HasThrown)
}

func (it *WrappedErr) IsBothPresent() bool {
	return it != nil &&
		it.HasError &&
		it.HasThrown &&
		it.Exception != nil
}

func (it *WrappedErr) HasException() bool {
	return it != nil &&
		it.Exception != nil
}

func (it *WrappedErr) ErrorString() string {
	if it.IsInvalid() {
		return ""
	}

	if it.Error != nil {
		return it.Error.Error()
	}

	return ""
}

func (it *WrappedErr) ExceptionString() string {
	if it.IsInvalidException() {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintFullPropertyNameValueFormat,
		it.Exception)
}

func (it *WrappedErr) String() string {
	if it.IsInvalid() {
		return ""
	}

	if it.IsBothPresent() {
		return it.ErrorString() +
			constants.CommaSpace +
			it.ExceptionString()
	}

	if it.HasError {
		return it.ErrorString()
	}

	if it.HasException() {
		return it.ExceptionString()
	}

	return ""
}
