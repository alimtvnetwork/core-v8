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
)

type ByteOnce struct {
	innerData       byte
	initializerFunc func() byte
	isInitialized   bool
}

func NewByteOnce(initializerFunc func() byte) ByteOnce {
	return ByteOnce{
		initializerFunc: initializerFunc,
	}
}

func NewByteOncePtr(initializerFunc func() byte) *ByteOnce {
	return &ByteOnce{
		initializerFunc: initializerFunc,
	}
}

func (it *ByteOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.Value())
}

func (it *ByteOnce) UnmarshalJSON(data []byte) error {
	it.isInitialized = true

	return json.Unmarshal(data, &it.innerData)
}

func (it *ByteOnce) Value() byte {
	if it.isInitialized {
		return it.innerData
	}

	it.innerData = it.initializerFunc()
	it.isInitialized = true

	return it.innerData
}

func (it *ByteOnce) Execute() byte {
	return it.Value()
}

func (it *ByteOnce) Int() int {
	return int(it.Value())
}

// IsEmpty returns true if zero
func (it *ByteOnce) IsEmpty() bool {
	return it.Value() == 0
}

func (it *ByteOnce) IsZero() bool {
	return it.Value() == 0
}

func (it *ByteOnce) IsNegative() bool {
	return it.Value() < 0
}

func (it *ByteOnce) IsPositive() bool {
	return it.Value() > 0
}

func (it *ByteOnce) String() string {
	return strconv.Itoa(int(it.Value()))
}

func (it *ByteOnce) Serialize() ([]byte, error) {
	value := it.Value()

	return json.Marshal(value)
}
