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

package errcore

import (
	"errors"
	"fmt"
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

type stackTraceEnhance struct{}

func (it stackTraceEnhance) Error(err error) error {
	return it.ErrorSkip(1, err)
}

func (it stackTraceEnhance) ErrorSkip(skip int, err error) error {
	if err == nil {
		return nil
	}

	msg := err.Error()

	return errors.New(it.MsgSkip(1+skip, msg))
}

func (it stackTraceEnhance) MsgToErrSkip(skip int, msg string) error {
	if len(msg) == 0 {
		return nil
	}

	return errors.New(it.MsgSkip(1+skip, msg))
}

func (it stackTraceEnhance) FmtSkip(skip int, format string, v ...any) error {
	if len(format) == 0 {
		return nil
	}

	msg := fmt.Sprintf(format, v...)

	return errors.New(it.MsgSkip(1+skip, msg))
}

func (it stackTraceEnhance) Msg(msg string) string {
	if len(msg) == 0 {
		return ""
	}

	return it.MsgSkip(1, msg)
}

func (it stackTraceEnhance) MsgSkip(skip int, msg string) string {
	if len(msg) == 0 {
		return ""
	}

	if strings.Contains(msg, constants.StackTrace) {
		return msg
	}

	toTrace := it.trace(skip + 1)

	if len(toTrace) == 0 {
		return fmt.Sprintf(
			combineMsgWithErrorFormat,
			it.methodName(skip+1),
			msg,
		)
	}

	fullMessage := fmt.Sprintf(
		stackEnhanceFormat,
		it.methodName(skip+1),
		msg,
		constants.StackTrace,
		toTrace,
	)

	return fullMessage
}

func (it stackTraceEnhance) methodName(skip int) string {
	return reflectinternal.CodeStack.MethodName(1 + skip)
}

func (it stackTraceEnhance) trace(skip int) string {
	lines := reflectinternal.CodeStack.StacksStringsFiltered(2+skip, 4)

	if len(lines) == 0 {
		return ""
	}

	return strings.Join(lines, "\n  - ")
}

func (it stackTraceEnhance) MsgErrorSkip(skip int, msg string, err error) string {
	if err == nil {
		return ""
	}

	compiledMsg := fmt.Sprintf(
		"%s %s",
		msg,
		err,
	)

	if strings.Contains(msg, constants.StackTrace) {
		return compiledMsg
	}

	toTrace := it.trace(skip + 1)

	if len(toTrace) == 0 {
		return fmt.Sprintf(
			combineMsgWithErrorFormat,
			it.methodName(skip+1),
			msg,
		)
	}

	fullMessage := fmt.Sprintf(
		stackEnhanceFormat,
		it.methodName(skip+1),
		msg,
		constants.StackTrace,
		toTrace,
	)

	return fullMessage
}

func (it stackTraceEnhance) MsgErrorToErrSkip(skip int, msg string, err error) error {
	if err == nil {
		return nil
	}

	return errors.New(it.MsgErrorSkip(1+skip, msg, err))
}
