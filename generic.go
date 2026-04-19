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

// EmptySlice returns an empty slice of type T.
//
// Usage:
//
//	ints := core.EmptySlice[int]()       // returns []int
//	strs := core.EmptySlice[string]()    // returns []string
func EmptySlice[T any]() []T {
	return make([]T, 0)
}

// SliceByLength returns a zero-valued slice of type T with the given length.
func SliceByLength[T any](length int) []T {
	return make([]T, length)
}

// SliceByCapacity returns a slice of type T with the given length and capacity.
func SliceByCapacity[T any](length, cap int) []T {
	return make([]T, length, cap)
}

func EmptySlicePtr[T any]() *[]T {
	s := make([]T, 0)
	return &s
}

func SlicePtrByLength[T any](length int) *[]T {
	s := make([]T, length)
	return &s
}

func SlicePtrByCapacity[T any](length, cap int) *[]T {
	s := make([]T, length, cap)
	return &s
}

// EmptyMapPtr returns a pointer to an empty map of type map[K]V.
// It replaces EmptyStringsMapPtr, EmptyStringToIntMapPtr and similar per-type functions.
func EmptyMapPtr[K comparable, V any]() *map[K]V {
	m := make(map[K]V)
	return &m
}
