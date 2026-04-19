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

package core

func EmptyAnysPtr() *[]any {
	return EmptySlicePtr[any]()
}

func EmptyFloat32Ptr() *[]float32 {
	return EmptySlicePtr[float32]()
}

func EmptyFloat64Ptr() *[]float64 {
	return EmptySlicePtr[float64]()
}

func EmptyBoolsPtr() *[]bool {
	return EmptySlicePtr[bool]()
}

func EmptyIntsPtr() *[]int {
	return EmptySlicePtr[int]()
}

func EmptyBytePtr() []byte {
	return []byte{}
}

func EmptyStringsMapPtr() *map[string]string {
	return EmptyMapPtr[string, string]()
}

func EmptyStringToIntMapPtr() *map[string]int {
	return EmptyMapPtr[string, int]()
}

func EmptyStringsPtr() *[]string {
	return EmptySlicePtr[string]()
}

func EmptyPointerStringsPtr() *[]*string {
	return EmptySlicePtr[*string]()
}

func StringsPtrByLength(length int) *[]string {
	return SlicePtrByLength[string](length)
}

func StringsPtrByCapacity(length, cap int) *[]string {
	return SlicePtrByCapacity[string](length, cap)
}

func PointerStringsPtrByCapacity(length, cap int) *[]*string {
	return SlicePtrByCapacity[*string](length, cap)
}
