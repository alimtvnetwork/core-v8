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

// All returns an iter.Seq2[int, T] that yields each item with its index.
// Supports Go 1.23+ range-over-func:
//
//	for i, item := range collection.All() { ... }
func (it *Collection[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, item := range it.items {
			if !yield(i, item) {
				return
			}
		}
	}
}

// Values returns an iter.Seq[T] that yields each item without index.
//
//	for item := range collection.Values() { ... }
func (it *Collection[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range it.items {
			if !yield(item) {
				return
			}
		}
	}
}

// Backward returns an iter.Seq2[int, T] that yields items in reverse order.
//
//	for i, item := range collection.Backward() { ... }
func (it *Collection[T]) Backward() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i := len(it.items) - 1; i >= 0; i-- {
			if !yield(i, it.items[i]) {
				return
			}
		}
	}
}
