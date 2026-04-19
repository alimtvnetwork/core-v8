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
	"reflect"
)

type expected struct{}

func (it expected) But(
	title any,
	wasExpecting,
	actualReceived any,
) error {
	return ExpectingErrorSimpleNoType(
		title,
		wasExpecting,
		actualReceived)
}

func (it expected) ButFoundAsMsg(
	title any,
	wasExpecting,
	actualReceived any,
) string {
	return ExpectingSimpleNoType(
		title,
		wasExpecting,
		actualReceived)
}

func (it expected) ButFoundWithTypeAsMsg(
	title any,
	wasExpecting,
	actualReceived any,
) string {
	return ExpectingSimple(
		title,
		wasExpecting,
		actualReceived)
}

func (it expected) ButUsingType(
	title any,
	wasExpecting,
	actualReceived any,
) error {
	return errors.New(ExpectingSimple(
		title,
		wasExpecting,
		actualReceived))
}

func (it expected) ReflectButFound(
	expected, found reflect.Kind,
) error {
	return fmt.Errorf(
		"expected [%v] but found [%v]",
		expected.String(), found.String())
}

func (it expected) PrimitiveButFound(
	found reflect.Kind,
) error {
	return fmt.Errorf(
		"expected [primitive] but found [%v]",
		found.String())
}

func (it expected) ValueHasNoElements(
	typ reflect.Kind,
) error {
	return fmt.Errorf(
		"generic value [%v] is nil or has no element",
		typ.String())
}
