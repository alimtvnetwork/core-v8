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
	"strconv"
	"sync"
)

type BoolOnce struct {
	innerData       bool
	initializerFunc func() bool
	once            sync.Once
}

func NewBoolOnce(initializerFunc func() bool) BoolOnce {
	return BoolOnce{
		initializerFunc: initializerFunc,
	}
}

func NewBoolOncePtr(initializerFunc func() bool) *BoolOnce {
	return &BoolOnce{
		initializerFunc: initializerFunc,
	}
}

func (it *BoolOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.Value())
}

func (it *BoolOnce) UnmarshalJSON(data []byte) error {
	it.once.Do(func() {})
	return json.Unmarshal(data, &it.innerData)
}

func (it *BoolOnce) Value() bool {
	it.once.Do(func() {
		it.innerData = it.initializerFunc()
	})

	return it.innerData
}

func (it *BoolOnce) Execute() bool {
	return it.Value()
}

func (it *BoolOnce) String() string {
	return strconv.FormatBool(it.Value())
}

func (it *BoolOnce) Serialize() ([]byte, error) {
	value := it.Value()

	return json.Marshal(value)
}
