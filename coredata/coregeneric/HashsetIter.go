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

// All returns an iter.Seq2[T, bool] over all set members.
// The bool value is always true (matching the internal map representation).
//
//	for item, _ := range hashset.All() { ... }
func (it *Hashset[T]) All() iter.Seq2[T, bool] {
	return func(yield func(T, bool) bool) {
		for item, v := range it.items {
			if !yield(item, v) {
				return
			}
		}
	}
}

// Values returns an iter.Seq[T] over all set members.
//
//	for item := range hashset.Values() { ... }
func (it *Hashset[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range it.items {
			if !yield(item) {
				return
			}
		}
	}
}
