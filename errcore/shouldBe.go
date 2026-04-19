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
	"encoding/json"
	"errors"
	"fmt"
)

type shouldBe struct{}

func (it shouldBe) StrEqMsg(actual, expecting string) string {
	return fmt.Sprintf(
		ShouldBeMessageFormat,
		actual,
		expecting)
}

func (it shouldBe) StrEqErr(actual, expecting string) error {
	msg := it.StrEqMsg(expecting, actual)

	return errors.New(msg)
}

func (it shouldBe) AnyEqMsg(actual, expecting any) string {
	return fmt.Sprintf(
		ShouldBeMessageFormat,
		actual,
		expecting)
}

func (it shouldBe) AnyEqErr(actual, expecting any) error {
	msg := it.AnyEqMsg(expecting, actual)

	return errors.New(msg)
}

func (it shouldBe) JsonEqMsg(actual, expecting any) string {
	actualJson, err := json.Marshal(actual)
	MustBeEmpty(err)

	expectingJson, expectingErr := json.Marshal(expecting)
	MustBeEmpty(expectingErr)

	return fmt.Sprintf(
		ShouldBeMessageFormat,
		actualJson,
		expectingJson)
}

func (it shouldBe) JsonEqErr(actual, expecting any) error {
	msg := it.JsonEqMsg(expecting, actual)

	return errors.New(msg)
}
