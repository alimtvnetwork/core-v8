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

// typedCollectionCreator provides factory methods for Collection[T].
// A single generic definition covers all primitive type variants.
type typedCollectionCreator[T any] struct{}

// Empty creates a zero-capacity Collection[T].
func (it typedCollectionCreator[T]) Empty() *Collection[T] {
	return EmptyCollection[T]()
}

// Cap creates a Collection[T] with pre-allocated capacity.
func (it typedCollectionCreator[T]) Cap(capacity int) *Collection[T] {
	return NewCollection[T](capacity)
}

// From wraps an existing slice into a Collection[T] (no copy).
func (it typedCollectionCreator[T]) From(items []T) *Collection[T] {
	return CollectionFrom[T](items)
}

// Clone copies items into a new Collection[T].
func (it typedCollectionCreator[T]) Clone(items []T) *Collection[T] {
	return CollectionClone[T](items)
}

// Items creates a Collection[T] from variadic arguments.
func (it typedCollectionCreator[T]) Items(items ...T) *Collection[T] {
	return CollectionFrom[T](items)
}

// LenCap creates a Collection[T] with specific length and capacity.
func (it typedCollectionCreator[T]) LenCap(length, capacity int) *Collection[T] {
	return CollectionLenCap[T](length, capacity)
}
