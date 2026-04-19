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

import "fmt"

// LinkedListNode is a generic singly-linked list node.
// It generalizes corestr.LinkedListNode from string to any type T.
type LinkedListNode[T any] struct {
	Element T
	next    *LinkedListNode[T]
}

// HasNext returns true if there is a next node.
func (it *LinkedListNode[T]) HasNext() bool {
	return it.next != nil
}

// Next returns the next node in the chain.
func (it *LinkedListNode[T]) Next() *LinkedListNode[T] {
	return it.next
}

// EndOfChain traverses to the end of the chain and returns the last node and total length.
func (it *LinkedListNode[T]) EndOfChain() (endOfChain *LinkedListNode[T], length int) {
	node := it
	length++

	for node.HasNext() {
		node = node.Next()
		length++
	}

	return node, length
}

// Clone creates a copy of this node (without the chain).
func (it *LinkedListNode[T]) Clone() *LinkedListNode[T] {
	return &LinkedListNode[T]{
		Element: it.Element,
		next:    nil,
	}
}

// ListPtr collects all elements from this node onward into a slice.
func (it *LinkedListNode[T]) ListPtr() *[]T {
	list := make([]T, 0, 100)

	node := it
	list = append(list, node.Element)

	for node.HasNext() {
		node = node.Next()
		list = append(list, node.Element)
	}

	return &list
}

// String returns a string representation of the element.
func (it *LinkedListNode[T]) String() string {
	return fmt.Sprintf("%v", it.Element)
}
