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

package coredynamic

import (
	"errors"
	"reflect"

	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/errcore"
	"github.com/alimtvnetwork/core-v8/internal/strutilinternal"
)

type SimpleResult struct {
	Dynamic
	Result  any
	Message string
	err     error
}

func InvalidSimpleResultNoMessage() *SimpleResult {
	return &SimpleResult{
		Result:  nil,
		Dynamic: NewDynamic(nil, false),
		Message: constants.EmptyString,
	}
}

func InvalidSimpleResult(
	invalidMessage string,
) *SimpleResult {
	return &SimpleResult{
		Result:  nil,
		Dynamic: NewDynamic(nil, false),
		Message: invalidMessage,
	}
}

func NewSimpleResultValid(
	result any,
) *SimpleResult {
	return &SimpleResult{
		Result:  result,
		Dynamic: NewDynamic(result, true),
		Message: constants.EmptyString,
	}
}

func NewSimpleResult(
	result any,
	isValid bool,
	invalidMessage string,
) *SimpleResult {
	return &SimpleResult{
		Result:  result,
		Dynamic: NewDynamic(result, isValid),
		Message: invalidMessage,
	}
}

func (it *SimpleResult) GetErrorOnTypeMismatch(
	typeMatch reflect.Type,
	isIncludeInvalidMessage bool,
) error {
	if it == nil {
		return nil
	}

	if it.IsReflectTypeOf(typeMatch) {
		return nil
	}

	typeMismatchMessage := errcore.CombineWithMsgTypeNoStack(
		errcore.TypeMismatchType,
		"Current type - ["+it.ReflectTypeName()+"], expected type",
		typeMatch,
	) + constants.NewLineUnix

	isExcludeMessage := !isIncludeInvalidMessage

	if isExcludeMessage {
		return errors.New(typeMismatchMessage)
	}

	return errors.New(typeMismatchMessage + it.Message)
}

func (it *SimpleResult) InvalidError() error {
	if it == nil {
		return nil
	}

	if it.err != nil {
		return it.err
	}

	if strutilinternal.IsEmptyOrWhitespace(it.Message) {
		return nil
	}

	if it.err == nil {
		it.err = errors.New(it.Message)
	}

	return it.err
}

func (it *SimpleResult) Clone() SimpleResult {
	if it == nil {
		return SimpleResult{}
	}

	return SimpleResult{
		Dynamic: it.Dynamic.Clone(),
		Result:  it.Result,
		Message: it.Message,
	}
}

func (it *SimpleResult) ClonePtr() *SimpleResult {
	if it == nil {
		return nil
	}

	return &SimpleResult{
		Dynamic: it.Dynamic.Clone(),
		Result:  it.Result,
		Message: it.Message,
	}
}
