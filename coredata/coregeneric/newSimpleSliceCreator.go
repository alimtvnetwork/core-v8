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

// newSimpleSliceCreator aggregates per-type simple slice creators.
//
// Usage:
//
//	coregeneric.New.SimpleSlice.String.Cap(10)
//	coregeneric.New.SimpleSlice.Int.Items(1, 2, 3)
//	coregeneric.New.SimpleSlice.Float64.From(existingSlice)
type newSimpleSliceCreator struct {
	String  typedSimpleSliceCreator[string]
	Int     typedSimpleSliceCreator[int]
	Int8    typedSimpleSliceCreator[int8]
	Int16   typedSimpleSliceCreator[int16]
	Int32   typedSimpleSliceCreator[int32]
	Int64   typedSimpleSliceCreator[int64]
	Uint    typedSimpleSliceCreator[uint]
	Uint8   typedSimpleSliceCreator[uint8]
	Uint16  typedSimpleSliceCreator[uint16]
	Uint32  typedSimpleSliceCreator[uint32]
	Uint64  typedSimpleSliceCreator[uint64]
	Float32 typedSimpleSliceCreator[float32]
	Float64 typedSimpleSliceCreator[float64]
	Byte    typedSimpleSliceCreator[byte]
	Bool    typedSimpleSliceCreator[bool]
	Any     typedSimpleSliceCreator[any]
}

// newSimpleSliceCreatorNote:
// SimpleSlice uses [T any] constraint because it wraps a raw []T.
// For ordered operations (Sort, Min, Max), use the package-level
// SortSimpleSlice[T cmp.Ordered]() functions.
