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

// ═══════════════════════════════════════════════════════════════
// Pre-built type aliases — ready to use without specifying T
// ═══════════════════════════════════════════════════════════════

// --- Primitive collection types ---

// StringCollection is a type-safe collection of strings.
type StringCollection = Collection[string]

// IntCollection is a type-safe collection of ints.
type IntCollection = Collection[int]

// Int64Collection is a type-safe collection of int64s.
type Int64Collection = Collection[int64]

// ByteCollection is a type-safe collection of bytes.
type ByteCollection = Collection[byte]

// BoolCollection is a type-safe collection of bools.
type BoolCollection = Collection[bool]

// Float64Collection is a type-safe collection of float64s.
type Float64Collection = Collection[float64]

// Float32Collection is a type-safe collection of float32s.
type Float32Collection = Collection[float32]

// ByteSliceCollection is a type-safe collection of byte slices.
type ByteSliceCollection = Collection[[]byte]

// --- Map collection types ---

// StringMapCollection is a collection of map[string]string.
type StringMapCollection = Collection[map[string]string]

// AnyMapCollection is a collection of map[string]any.
type AnyMapCollection = Collection[map[string]any]

// IntMapCollection is a collection of map[string]int.
type IntMapCollection = Collection[map[string]int]

// ═══════════════════════════════════════════════════════════════
// Factory shortcuts — create common collections without specifying T
// ═══════════════════════════════════════════════════════════════

// NewStringCollection creates a new string collection with the given capacity.
func NewStringCollection(capacity int) *StringCollection {
	return NewCollection[string](capacity)
}

// EmptyStringCollection creates an empty string collection.
func EmptyStringCollection() *StringCollection {
	return EmptyCollection[string]()
}

// NewIntCollection creates a new int collection with the given capacity.
func NewIntCollection(capacity int) *IntCollection {
	return NewCollection[int](capacity)
}

// EmptyIntCollection creates an empty int collection.
func EmptyIntCollection() *IntCollection {
	return EmptyCollection[int]()
}

// NewInt64Collection creates a new int64 collection with the given capacity.
func NewInt64Collection(capacity int) *Int64Collection {
	return NewCollection[int64](capacity)
}

// NewByteCollection creates a new byte collection with the given capacity.
func NewByteCollection(capacity int) *ByteCollection {
	return NewCollection[byte](capacity)
}

// NewBoolCollection creates a new bool collection with the given capacity.
func NewBoolCollection(capacity int) *BoolCollection {
	return NewCollection[bool](capacity)
}

// NewFloat64Collection creates a new float64 collection with the given capacity.
func NewFloat64Collection(capacity int) *Float64Collection {
	return NewCollection[float64](capacity)
}

// NewAnyMapCollection creates a new map[string]any collection with the given capacity.
func NewAnyMapCollection(capacity int) *AnyMapCollection {
	return NewCollection[map[string]any](capacity)
}

// NewStringMapCollection creates a new map[string]string collection with the given capacity.
func NewStringMapCollection(capacity int) *StringMapCollection {
	return NewCollection[map[string]string](capacity)
}
