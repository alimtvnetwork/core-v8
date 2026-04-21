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
	"github.com/alimtvnetwork/core-v8/issetter"
)

type SimpleRequest struct {
	Dynamic
	message string
	err     error
}

func InvalidSimpleRequestNoMessage() *SimpleRequest {
	return &SimpleRequest{
		Dynamic: NewDynamic(nil, false),
		message: constants.EmptyString,
	}
}

func InvalidSimpleRequest(
	message string,
) *SimpleRequest {
	return &SimpleRequest{
		Dynamic: NewDynamic(nil, false),
		message: message,
	}
}

func NewSimpleRequest(
	request any,
	isValid bool,
	message string,
) *SimpleRequest {
	return &SimpleRequest{
		Dynamic: NewDynamic(request, isValid),
		message: message,
	}
}

func NewSimpleRequestValid(
	request any,
) *SimpleRequest {
	return &SimpleRequest{
		Dynamic: NewDynamic(request, true),
		message: constants.EmptyString,
	}
}

func (receiver *SimpleRequest) Message() string {
	if receiver == nil {
		return constants.EmptyString
	}

	return receiver.message
}

func (receiver *SimpleRequest) Request() any {
	if receiver == nil {
		return nil
	}

	return receiver.Dynamic.Data()
}

func (receiver *SimpleRequest) Value() any {
	if receiver == nil {
		return nil
	}

	return receiver.Dynamic.Data()
}

func (receiver *SimpleRequest) GetErrorOnTypeMismatch(
	typeMatch reflect.Type,
	isIncludeInvalidMessage bool,
) error {
	if receiver == nil {
		return nil
	}

	if receiver.IsReflectTypeOf(typeMatch) {
		return nil
	}

	typeMismatchMessage := errcore.CombineWithMsgTypeNoStack(
		errcore.TypeMismatchType,
		"Current type - ["+receiver.ReflectTypeName()+"], expected type",
		typeMatch,
	) + constants.NewLineUnix

	isExcludeMessage := !isIncludeInvalidMessage

	if isExcludeMessage {
		return errors.New(typeMismatchMessage)
	}

	return errors.New(typeMismatchMessage + receiver.message)
}

func (receiver *SimpleRequest) IsReflectKind(checkingKind reflect.Kind) bool {
	if receiver == nil {
		return false
	}

	return receiver.ReflectKind() == checkingKind
}

func (receiver *SimpleRequest) IsPointer() bool {
	if receiver == nil {
		return false
	}

	if receiver.isPointer.IsUninitialized() {
		receiver.isPointer = issetter.GetBool(
			receiver.IsReflectKind(reflect.Ptr),
		)
	}

	return receiver.isPointer.IsTrue()
}

func (receiver *SimpleRequest) InvalidError() error {
	if receiver == nil {
		return nil
	}

	if receiver.err != nil {
		return receiver.err
	}

	if strutilinternal.IsEmptyOrWhitespace(receiver.message) {
		return nil
	}

	if receiver.err == nil {
		receiver.err = errors.New(receiver.message)
	}

	return receiver.err
}
