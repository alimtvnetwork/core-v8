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

type IntegerOnce struct {
	innerData       int
	initializerFunc func() int
	once            sync.Once
}

func NewIntegerOnce(initializerFunc func() int) IntegerOnce {
	return IntegerOnce{
		initializerFunc: initializerFunc,
	}
}

func NewIntegerOncePtr(initializerFunc func() int) *IntegerOnce {
	return &IntegerOnce{
		initializerFunc: initializerFunc,
	}
}

func (it *IntegerOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.Value())
}

func (it *IntegerOnce) UnmarshalJSON(data []byte) error {
	it.once.Do(func() {})
	return json.Unmarshal(data, &it.innerData)
}

func (it *IntegerOnce) Value() int {
	it.once.Do(func() {
		it.innerData = it.initializerFunc()
	})

	return it.innerData
}

func (it *IntegerOnce) Execute() int {
	return it.Value()
}

// IsEmpty returns true if zero
func (it *IntegerOnce) IsEmpty() bool {
	return it.Value() == 0
}

func (it *IntegerOnce) IsZero() bool {
	return it.Value() == 0
}

func (it *IntegerOnce) IsAboveZero() bool {
	return it.Value() > 0
}

func (it *IntegerOnce) IsAboveEqualZero() bool {
	return it.Value() >= 0
}

func (it *IntegerOnce) IsLessThanZero() bool {
	return it.Value() < 0
}

func (it *IntegerOnce) IsLessThanEqualZero() bool {
	return it.Value() <= 0
}

func (it *IntegerOnce) IsAbove(i int) bool {
	return it.Value() > i
}

func (it *IntegerOnce) IsAboveEqual(i int) bool {
	return it.Value() >= i
}

func (it *IntegerOnce) IsLessThan(i int) bool {
	return it.Value() < i
}

func (it *IntegerOnce) IsLessThanEqual(i int) bool {
	return it.Value() <= i
}

func (it *IntegerOnce) IsInvalidIndex() bool {
	return it.Value() < 0
}

func (it *IntegerOnce) IsValidIndex() bool {
	return it.Value() >= 0
}

func (it *IntegerOnce) IsNegative() bool {
	return it.Value() < 0
}

func (it *IntegerOnce) IsPositive() bool {
	return it.Value() > 0
}

func (it *IntegerOnce) String() string {
	return strconv.Itoa(it.Value())
}

func (it *IntegerOnce) Serialize() ([]byte, error) {
	value := it.Value()

	return json.Marshal(value)
}
