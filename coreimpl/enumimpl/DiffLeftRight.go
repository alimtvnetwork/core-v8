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
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/alimtvnetwork/core/constants"
)

type DiffLeftRight struct {
	Left, Right any
}

func (it *DiffLeftRight) Types() (l, r reflect.Type) {
	l = reflect.TypeOf(it.Left)
	r = reflect.TypeOf(it.Right)

	return l, r
}

func (it *DiffLeftRight) IsSameTypeSame() bool {
	l, r := it.Types()

	return l == r
}

func (it *DiffLeftRight) IsSame() bool {
	return reflect.DeepEqual(it.Left, it.Right)
}

func (it *DiffLeftRight) IsSameRegardlessOfType() bool {
	leftString := fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		it.Left)
	rightString := fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		it.Right)

	return leftString == rightString
}

func (it *DiffLeftRight) IsEqual(isRegardless bool) bool {
	if isRegardless {
		return it.IsSameRegardlessOfType()
	}

	return it.IsSame()
}

func (it *DiffLeftRight) HasMismatch(isRegardless bool) bool {
	if isRegardless {
		return it.HasMismatchRegardlessOfType()
	}

	return it.IsNotEqual()
}

func (it *DiffLeftRight) IsNotEqual() bool {
	return !it.IsSame()
}

func (it *DiffLeftRight) HasMismatchRegardlessOfType() bool {
	return !it.IsSameRegardlessOfType()
}

func (it *DiffLeftRight) String() string {
	return it.JsonString()
}

func (it *DiffLeftRight) JsonString() string {
	if it == nil {
		return ""
	}

	b, e := json.Marshal(*it)

	if e != nil {
		return "error : " + e.Error()
	}

	return string(b)
}

func (it *DiffLeftRight) SpecificFullString() (l, r string) {
	l = fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		it.Left)
	r = fmt.Sprintf(
		constants.SprintPropertyNameValueFormat,
		it.Right)

	return l, r
}

func (it *DiffLeftRight) DiffString() string {
	if it.IsSameRegardlessOfType() {
		return ""
	}

	return it.String()
}
