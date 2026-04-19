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
)

type ExpectingRecord struct {
	ExpectingTitle string
	WasExpecting   any
}

// Message
// Expecting
//
// returns
//
//	"%s - expecting (type:[%T]) : [\"%v\"], but received or actual (type:[%T]) : [\"%v\"]"
func (it *ExpectingRecord) Message(actual any) string {
	if it == nil {
		return ""
	}

	return fmt.Sprintf(
		expectingMessageFormat,
		it.ExpectingTitle,
		it.WasExpecting, it.WasExpecting,
		actual, actual)
}

// MessageSimple
// Expecting
//
// returns
//
//	"%s - Expect (type:\"%T\")[\"%v\"] != [\"%v\"](type:\"%T\") Actual"
func (it *ExpectingRecord) MessageSimple(actual any) string {
	if it == nil {
		return ""
	}

	return ExpectingSimple(
		it.ExpectingTitle,
		it.WasExpecting,
		actual)
}

// MessageSimpleNoType
//
// returns
//
//	"%s - Expect [\"%v\"] != [\"%v\"] Actual"
func (it *ExpectingRecord) MessageSimpleNoType(actual any) string {
	if it == nil {
		return ""
	}

	return ExpectingSimpleNoType(
		it.ExpectingTitle,
		it.WasExpecting,
		actual)
}

// Error
// Expecting
//
// returns
//
//	"%s - expecting (type:[%T]) : [\"%v\"], but received or actual (type:[%T]) : [\"%v\"]"
func (it *ExpectingRecord) Error(actual any) error {
	if it == nil {
		return nil
	}

	return errors.New(it.Message(actual))
}

// ErrorSimple
// Expecting
//
// returns
//
//	"%s - Expect (type:\"%T\")[\"%v\"] != [\"%v\"](type:\"%T\") Actual"
func (it *ExpectingRecord) ErrorSimple(actual any) error {
	if it == nil {
		return nil
	}

	return errors.New(it.MessageSimple(actual))
}

// ErrorSimpleNoType
// Expecting
//
// returns
//
//	"%s - Expect [\"%v\"] != [\"%v\"] Actual"
func (it *ExpectingRecord) ErrorSimpleNoType(actual any) error {
	if it == nil {
		return nil
	}

	return errors.New(it.MessageSimpleNoType(actual))
}
