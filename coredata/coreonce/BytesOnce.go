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
)

type BytesOnce struct {
	innerData       []byte
	initializerFunc func() []byte
	isInitialized   bool
}

func NewBytesOnce(initializerFunc func() []byte) BytesOnce {
	return BytesOnce{
		initializerFunc: initializerFunc,
	}
}

func NewBytesOncePtr(initializerFunc func() []byte) *BytesOnce {
	return &BytesOnce{
		initializerFunc: initializerFunc,
	}
}

func (it *BytesOnce) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.Value())
}

func (it *BytesOnce) UnmarshalJSON(data []byte) error {
	it.isInitialized = true

	return json.Unmarshal(data, &it.innerData)
}

func (it *BytesOnce) Value() []byte {
	if it.isInitialized {
		return it.innerData
	}

	if it.initializerFunc == nil {
		it.isInitialized = true

		return it.innerData
	}

	it.innerData = it.initializerFunc()
	it.isInitialized = true

	return it.innerData
}

func (it *BytesOnce) Execute() []byte {
	return it.Value()
}

// IsEmpty returns true if zero
func (it *BytesOnce) IsEmpty() bool {
	return it == nil || it.initializerFunc == nil || len(it.Value()) == 0
}

func (it *BytesOnce) String() string {
	return string(it.Value())
}

func (it *BytesOnce) Length() int {
	if it == nil || it.initializerFunc == nil {
		return 0
	}

	return len(it.Value())
}

func (it *BytesOnce) Serialize() ([]byte, error) {
	value := it.Value()

	return json.Marshal(value)
}
