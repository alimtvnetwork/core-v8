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

package stringslice

import (
	"reflect"
	"sync"
)

func AnyLinesProcessAsyncUsingProcessor(
	lines any,
	lineProcessor func(index int, lineIn any) (lineOut string),
) []string {
	if lines == nil {
		return []string{}
	}

	reflectType := reflect.TypeOf(lines)
	kind := reflectType.Kind()
	isArrayOrSlice := kind == reflect.Slice ||
		kind == reflect.Array

	isNotSliceOrArray := !isArrayOrSlice

	if isNotSliceOrArray {
		return []string{}
	}

	reflectValue := reflect.ValueOf(lines)
	length := reflectValue.Len()

	if length == 0 {
		return []string{}
	}

	slice := MakeLen(length)
	wg := sync.WaitGroup{}

	wg.Add(length)

	asyncProcessor := func(index int, lineIn any) {
		slice[index] = lineProcessor(index, lineIn)

		wg.Done()
	}

	for i := 0; i < length; i++ {
		itemAt := reflectValue.Index(i)
		go asyncProcessor(i, itemAt.Interface())
	}

	wg.Wait()

	return slice
}
