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

package coretests

import (
	"reflect"

	"github.com/alimtvnetwork/core-v8/issetter"
)

// VerifyTypeOf
//
// # Verify type of
//
//   - BaseTestCase.ArrangeInput : What we set for initial input
//   - BaseTestCase.ActualInput : must set BaseTestCase.SetActual
//   - BaseTestCase.ExpectedInput : What we expect in the test case
type VerifyTypeOf struct {
	IsVerify      issetter.Value // Only false makes it verify, creating and attaching it makes it verify true by default.
	ArrangeInput  reflect.Type   // Verify type for the BaseTestCase.ArrangeInput
	ActualInput   reflect.Type   // Verify type for the BaseTestCase.ActualInput, must set BaseTestCase.SetActual
	ExpectedInput reflect.Type   // Verify type for the BaseTestCase.ExpectedInput
}

func NewVerifyTypeOf(arrange any) *VerifyTypeOf {
	return &VerifyTypeOf{
		IsVerify:      issetter.True,
		ArrangeInput:  reflect.TypeOf(arrange),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}
}

func (it *VerifyTypeOf) IsDefined() bool {
	return it != nil
}

func (it *VerifyTypeOf) IsInvalid() bool {
	return it != nil
}

func (it *VerifyTypeOf) IsInvalidOrSkipVerify() bool {
	return it == nil || it.IsVerify.IsFalse()
}
