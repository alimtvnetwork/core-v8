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

package enumimpl

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

type KeyAnyVal struct {
	Key      string
	AnyValue any
}

func (it KeyAnyVal) KeyString() string {
	return it.Key
}

func (it KeyAnyVal) AnyVal() any {
	return it.AnyValue
}

func (it KeyAnyVal) AnyValString() string {
	return convAnyValToString(it.AnyValue)
}

func (it KeyAnyVal) WrapKey() string {
	return toJsonName(it.Key)
}

func (it KeyAnyVal) WrapValue() string {
	return toJsonName(it.AnyValue)
}

func (it KeyAnyVal) IsString() bool {
	_, isString := it.AnyValue.(string)

	if isString {
		return true
	}

	castInt, isInt := it.AnyValue.(int)

	if isInt {
		return castInt == constants.MinInt
	}

	return false
}

func (it KeyAnyVal) ValInt() int {
	return ConvEnumAnyValToInteger(it.AnyValue)
}

func (it KeyAnyVal) KeyValInteger() KeyValInteger {
	return KeyValInteger{
		Key:          it.Key,
		ValueInteger: it.ValInt(),
	}
}

func (it KeyAnyVal) String() string {
	if it.IsString() {
		// stringer
		return fmt.Sprintf(
			constants.StringEnumNameValueFormat,
			it.Key,
		)
	}

	return fmt.Sprintf(
		constants.EnumNameValueFormat,
		it.Key,
		it.AnyValue)
}
