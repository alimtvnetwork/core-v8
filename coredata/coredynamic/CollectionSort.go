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

import (
	"cmp"
	"sort"
)

// --- Sort with custom comparator (works for any T) ---

// SortFunc sorts the collection in place using the provided less function.
func (it *Collection[T]) SortFunc(less func(a, b T) bool) *Collection[T] {
	if it.Length() <= 1 {
		return it
	}

	sort.Slice(it.items, func(i, j int) bool {
		return less(it.items[i], it.items[j])
	})
	return it
}

// SortFuncLock sorts the collection in place with mutex protection.
func (it *Collection[T]) SortFuncLock(less func(a, b T) bool) *Collection[T] {
	it.Lock()
	defer it.Unlock()

	return it.SortFunc(less)
}

// SortedFunc returns a new sorted collection without mutating the original.
func (it *Collection[T]) SortedFunc(less func(a, b T) bool) *Collection[T] {
	cloned := it.Clone()
	return cloned.SortFunc(less)
}

// --- Package-level functions for Ordered types ---

// SortAsc sorts a collection of ordered types in ascending order (mutates).
func SortAsc[T cmp.Ordered](col *Collection[T]) *Collection[T] {
	return col.SortFunc(func(a, b T) bool {
		return a < b
	})
}

// SortDesc sorts a collection of ordered types in descending order (mutates).
func SortDesc[T cmp.Ordered](col *Collection[T]) *Collection[T] {
	return col.SortFunc(func(a, b T) bool {
		return a > b
	})
}

// SortAscLock sorts ascending with mutex protection.
func SortAscLock[T cmp.Ordered](col *Collection[T]) *Collection[T] {
	return col.SortFuncLock(func(a, b T) bool {
		return a < b
	})
}

// SortDescLock sorts descending with mutex protection.
func SortDescLock[T cmp.Ordered](col *Collection[T]) *Collection[T] {
	return col.SortFuncLock(func(a, b T) bool {
		return a > b
	})
}

// SortedAsc returns a new ascending-sorted collection without mutating the original.
func SortedAsc[T cmp.Ordered](col *Collection[T]) *Collection[T] {
	return col.SortedFunc(func(a, b T) bool {
		return a < b
	})
}

// SortedDesc returns a new descending-sorted collection without mutating the original.
func SortedDesc[T cmp.Ordered](col *Collection[T]) *Collection[T] {
	return col.SortedFunc(func(a, b T) bool {
		return a > b
	})
}

// IsSorted returns true if the collection is sorted according to the less function.
func (it *Collection[T]) IsSorted(less func(a, b T) bool) bool {
	if it.Length() <= 1 {
		return true
	}

	return sort.SliceIsSorted(it.items, func(i, j int) bool {
		return less(it.items[i], it.items[j])
	})
}

// IsSortedAsc returns true if an ordered collection is sorted ascending.
func IsSortedAsc[T cmp.Ordered](col *Collection[T]) bool {
	return col.IsSorted(func(a, b T) bool {
		return a < b
	})
}

// IsSortedDesc returns true if an ordered collection is sorted descending.
func IsSortedDesc[T cmp.Ordered](col *Collection[T]) bool {
	return col.IsSorted(func(a, b T) bool {
		return a > b
	})
}
