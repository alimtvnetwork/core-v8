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

// newCollectionCreator aggregates per-type collection creators.
//
// Simple types use newGenericCollectionCreator[T] directly.
// Types with extra methods (LenCap, Create) embed the generic creator
// and extend it.
//
// Usage:
//
//	coredynamic.New.Collection.String.Cap(10)
//	coredynamic.New.Collection.Int.Empty()
//	coredynamic.New.Collection.Any.From(items)
type newCollectionCreator struct {
	Any       newAnyCollectionCreator
	String    newStringCollectionCreator
	Int       newIntCollectionCreator
	Int64     newInt64CollectionCreator
	Byte      newByteCollectionCreator
	ByteSlice newGenericCollectionCreator[[]byte]
	Bool      newGenericCollectionCreator[bool]
	Float32   newGenericCollectionCreator[float32]
	Float64   newGenericCollectionCreator[float64]
	AnyMap    newGenericCollectionCreator[map[string]any]
	StringMap newGenericCollectionCreator[map[string]string]
	IntMap    newGenericCollectionCreator[map[string]int]
}
