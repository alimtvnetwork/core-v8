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

// GroupBy groups items by a key function, returning a map of key to Collection[T].
// The key type K must be comparable for use as a map key.
//
// Usage:
//
//	groups := coredynamic.GroupBy(users, func(u User) string { return u.Department })
func GroupBy[T any, K comparable](
	source *Collection[T],
	keyFunc func(T) K,
) map[K]*Collection[T] {
	if source == nil || source.IsEmpty() {
		return map[K]*Collection[T]{}
	}

	result := make(map[K]*Collection[T])

	for _, item := range source.items {
		key := keyFunc(item)
		col, exists := result[key]
		isNewGroup := !exists

		if isNewGroup {
			col = NewCollection[T](0)
			result[key] = col
		}
		col.items = append(col.items, item)
	}
	return result
}

// GroupByLock is the mutex-protected variant of GroupBy.
func GroupByLock[T any, K comparable](
	source *Collection[T],
	keyFunc func(T) K,
) map[K]*Collection[T] {
	source.Lock()
	defer source.Unlock()
	return GroupBy(source, keyFunc)
}

// GroupByCount returns the count of items per group key.
//
// Usage:
//
//	counts := coredynamic.GroupByCount(words, func(w string) string { return w })
func GroupByCount[T any, K comparable](
	source *Collection[T],
	keyFunc func(T) K,
) map[K]int {
	if source == nil || source.IsEmpty() {
		return map[K]int{}
	}

	result := make(map[K]int)

	for _, item := range source.items {
		result[keyFunc(item)]++
	}
	return result
}
