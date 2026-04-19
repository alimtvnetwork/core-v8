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

// =============================================================================
// Comparable constraint functions for Collection[T]
//
// These functions require T to satisfy comparable (==, != operators)
// and provide equality-based searches without custom predicates.
// =============================================================================

// ContainsAll returns true if the collection contains all given items.
func ContainsAll[T comparable](
	source *Collection[T],
	items ...T,
) bool {
	if source == nil {
		return false
	}

	for _, item := range items {
		isMissing := !ContainsItem(source, item)

		if isMissing {
			return false
		}
	}

	return true
}

// ContainsAny returns true if the collection contains any of the given items.
func ContainsAny[T comparable](
	source *Collection[T],
	items ...T,
) bool {
	if source == nil {
		return false
	}

	for _, item := range items {
		if ContainsItem(source, item) {
			return true
		}
	}

	return false
}

// RemoveItem removes the first occurrence of item. Returns true if found.
func RemoveItem[T comparable](
	source *Collection[T],
	item T,
) bool {
	if source == nil {
		return false
	}

	index := IndexOfItem(source, item)
	if index < 0 {
		return false
	}

	return source.RemoveAt(index)
}

// RemoveAllItems removes all occurrences of item. Returns the count removed.
func RemoveAllItems[T comparable](
	source *Collection[T],
	item T,
) int {
	if source == nil {
		return 0
	}

	removed := 0
	newItems := make([]T, 0, source.Length())

	for _, existing := range source.items {
		if existing == item {
			removed++
		} else {
			newItems = append(newItems, existing)
		}
	}

	source.items = newItems

	return removed
}

// ToHashset converts a Collection[T] to a Hashset[T].
// Requires T to be comparable for map key usage.
func ToHashset[T comparable](
	source *Collection[T],
) *Hashset[T] {
	if source == nil {
		return EmptyHashset[T]()
	}

	return HashsetFrom[T](source.items)
}

// DistinctSimpleSlice returns a new SimpleSlice with duplicates removed.
func DistinctSimpleSlice[T comparable](
	source *SimpleSlice[T],
) *SimpleSlice[T] {
	if source == nil {
		return EmptySimpleSlice[T]()
	}

	seen := make(map[T]bool)
	result := EmptySimpleSlice[T]()

	for _, item := range *source {
		if !seen[item] {
			seen[item] = true
			result.Add(item)
		}
	}

	return result
}

// ContainsSimpleSliceItem checks if a comparable item exists in a SimpleSlice.
func ContainsSimpleSliceItem[T comparable](
	source *SimpleSlice[T],
	item T,
) bool {
	if source == nil {
		return false
	}

	for _, existing := range *source {
		if existing == item {
			return true
		}
	}

	return false
}
