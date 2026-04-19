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

package coreonce

import (
	"encoding/json"
	"errors"
	"strings"
	"sync"

	"github.com/alimtvnetwork/core/constants"
)

type StringOnce struct {
	innerData       string
	initializerFunc func() string
	once            sync.Once
}

func NewStringOnce(initializerFunc func() string) StringOnce {
	return StringOnce{
		initializerFunc: initializerFunc,
	}
}

func NewStringOncePtr(initializerFunc func() string) *StringOnce {
	return &StringOnce{
		initializerFunc: initializerFunc,
	}
}

func (it *StringOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.Value())
}

func (it *StringOnce) UnmarshalJSON(data []byte) error {
	it.once.Do(func() {})
	return json.Unmarshal(data, &it.innerData)
}

func (it *StringOnce) ValuePtr() *string {
	val := it.Value()
	return &val
}

func (it *StringOnce) Value() string {
	it.once.Do(func() {
		it.innerData = it.initializerFunc()
	})

	return it.innerData
}

func (it *StringOnce) Execute() string {
	return it.Value()
}

func (it *StringOnce) IsEqual(equalString string) bool {
	return it.Value() == equalString
}

func (it *StringOnce) HasPrefix(prefix string) bool {
	return strings.HasPrefix(
		it.Value(), prefix)
}

func (it *StringOnce) IsStartsWith(startsWith string) bool {
	return strings.HasPrefix(
		it.Value(), startsWith)
}

func (it *StringOnce) HasSuffix(suffix string) bool {
	return strings.HasSuffix(
		it.Value(), suffix)
}

func (it *StringOnce) IsEndsWith(
	endsWith string,
) bool {
	return strings.HasSuffix(
		it.Value(), endsWith)
}

func (it *StringOnce) SplitBy(
	splitter string,
) []string {
	return strings.Split(it.Value(), splitter)
}

func (it *StringOnce) SplitLeftRightTrim(
	splitter string,
) (left, right string) {
	left, right = it.SplitLeftRight(splitter)

	return strings.TrimSpace(left), strings.TrimSpace(right)
}

func (it *StringOnce) SplitLeftRight(
	splitter string,
) (left, right string) {
	items := strings.SplitN(
		it.Value(),
		splitter,
		constants.Two)

	if len(items) == 2 {
		return items[0], items[1]
	}

	// len <= 1
	return items[0], ""
}

func (it *StringOnce) IsContains(equalString string) bool {
	return strings.Contains(it.Value(), equalString)
}

func (it *StringOnce) IsEmpty() bool {
	return it.Value() == ""
}

func (it *StringOnce) IsEmptyOrWhitespace() bool {
	return strings.TrimSpace(it.Value()) == ""
}

func (it *StringOnce) Bytes() []byte {
	return []byte(it.Value())
}

func (it *StringOnce) Error() error {
	return errors.New(it.Value())
}

func (it *StringOnce) String() string {
	return it.Value()
}

func (it *StringOnce) Serialize() ([]byte, error) {
	value := it.Value()

	return json.Marshal(value)
}
