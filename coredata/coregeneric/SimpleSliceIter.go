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

package coregeneric

import "iter"

// All returns an iter.Seq2[int, T] over all items with index.
//
//	for i, item := range simpleSlice.All() { ... }
func (it *SimpleSlice[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, item := range *it {
			if !yield(i, item) {
				return
			}
		}
	}
}

// Values returns an iter.Seq[T] over all items.
//
//	for item := range simpleSlice.Values() { ... }
func (it *SimpleSlice[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range *it {
			if !yield(item) {
				return
			}
		}
	}
}
