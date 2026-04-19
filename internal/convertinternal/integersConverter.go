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

package convertinternal

import "strconv"

type integersConverter struct{}

func (it integersConverter) ToMapBool(
	items ...int,
) (hashMap map[int]bool) {
	if len(items) == 0 {
		return map[int]bool{}
	}

	hashMap = make(map[int]bool, len(items))

	for _, item := range items {
		hashMap[item] = true
	}

	return hashMap
}

func (it integersConverter) Int8ToMapBool(
	items ...int8,
) (hashMap map[int8]bool) {
	if len(items) == 0 {
		return map[int8]bool{}
	}

	hashMap = make(map[int8]bool, len(items))

	for _, item := range items {
		hashMap[item] = true
	}

	return hashMap
}

func (it integersConverter) FromIntegersToMap(inputArray ...int) map[int]bool {
	if len(inputArray) == 0 {
		return map[int]bool{}
	}

	return Map.FromIntegersToMap(inputArray...)
}

func (it integersConverter) IntegersToStrings(intSlice []int) []string {
	stringSlice := make([]string, len(intSlice))
	for index, value := range intSlice {
		stringSlice[index] = strconv.Itoa(value)
	}

	return stringSlice
}
