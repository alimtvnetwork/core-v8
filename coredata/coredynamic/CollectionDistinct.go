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

package coredynamic

// --- Package-level deduplication for comparable types ---

// Distinct returns a new collection with duplicates removed, preserving order.
func Distinct[T comparable](col *Collection[T]) *Collection[T] {
	if col.IsEmpty() {
		return EmptyCollection[T]()
	}

	seen := make(map[T]bool, col.Length())
	result := NewCollection[T](col.Length())

	for _, item := range col.items {
		if !seen[item] {
			seen[item] = true
			result.items = append(result.items, item)
		}
	}
	return result
}

// Unique is an alias for Distinct.
func Unique[T comparable](col *Collection[T]) *Collection[T] {
	return Distinct(col)
}

// DistinctLock returns a new deduplicated collection with mutex protection.
func DistinctLock[T comparable](col *Collection[T]) *Collection[T] {
	col.Lock()
	defer col.Unlock()
	return Distinct(col)
}

// DistinctCount returns the number of unique items.
func DistinctCount[T comparable](col *Collection[T]) int {
	if col.IsEmpty() {
		return 0
	}

	seen := make(map[T]bool, col.Length())

	for _, item := range col.items {
		seen[item] = true
	}
	return len(seen)
}

// IsDistinct returns true if the collection has no duplicates.
func IsDistinct[T comparable](col *Collection[T]) bool {
	return col.Length() == DistinctCount(col)
}
