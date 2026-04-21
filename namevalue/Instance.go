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

package namevalue

import (
	"encoding/json"
	"fmt"

	"github.com/alimtvnetwork/core-v8/constants"
)

type Instance[K comparable, V any] struct {
	Name  K
	Value V
}

func (it *Instance[K, V]) IsNull() bool {
	return it == nil
}

func (it *Instance[K, V]) String() string {
	if it.IsNull() {
		return constants.EmptyString
	}

	return fmt.Sprintf(
		"%v = %v",
		it.Name,
		it.Value)
}

func (it *Instance[K, V]) JsonString() string {
	if it.IsNull() {
		return constants.EmptyString
	}

	rawBytes, err := json.Marshal(it)

	if err != nil || rawBytes == nil {
		return constants.EmptyString
	}

	return string(rawBytes)
}

func (it *Instance[K, V]) Dispose() {
	if it == nil {
		return
	}

	var zeroK K
	var zeroV V

	it.Name = zeroK
	it.Value = zeroV
}
