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

// All returns an iter.Seq2[int, T] over all nodes with index.
//
//	for i, value := range linkedList.All() { ... }
func (it *LinkedList[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		current := it.head

		for current != nil {
			if !yield(i, current.Element) {
				return
			}

			current = current.next
			i++
		}
	}
}

// Values returns an iter.Seq[T] over all node values.
//
//	for value := range linkedList.Values() { ... }
func (it *LinkedList[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		current := it.head

		for current != nil {
			if !yield(current.Element) {
				return
			}

			current = current.next
		}
	}
}
