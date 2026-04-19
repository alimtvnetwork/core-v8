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

// typedHashsetCreator provides factory methods for Hashset[T].
type typedHashsetCreator[T comparable] struct{}

// Empty creates a zero-capacity Hashset[T].
func (it typedHashsetCreator[T]) Empty() *Hashset[T] {
	return EmptyHashset[T]()
}

// Cap creates a Hashset[T] with pre-allocated capacity.
func (it typedHashsetCreator[T]) Cap(capacity int) *Hashset[T] {
	return NewHashset[T](capacity)
}

// From creates a Hashset[T] from a slice of items.
func (it typedHashsetCreator[T]) From(items []T) *Hashset[T] {
	return HashsetFrom[T](items)
}

// Items creates a Hashset[T] from variadic arguments.
func (it typedHashsetCreator[T]) Items(items ...T) *Hashset[T] {
	return HashsetFrom[T](items)
}

// UsingMap creates a Hashset[T] from an existing map.
func (it typedHashsetCreator[T]) UsingMap(itemsMap map[T]bool) *Hashset[T] {
	return HashsetFromMap[T](itemsMap)
}
