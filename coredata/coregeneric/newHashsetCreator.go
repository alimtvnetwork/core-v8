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

// newHashsetCreator aggregates per-type hashset creators.
//
// Usage:
//
//	coregeneric.New.Hashset.String.Cap(10)
//	coregeneric.New.Hashset.Int.From([]int{1, 2, 3})
//	coregeneric.New.Hashset.Float64.Items(1.0, 2.5)
type newHashsetCreator struct {
	String  typedHashsetCreator[string]
	Int     typedHashsetCreator[int]
	Int8    typedHashsetCreator[int8]
	Int16   typedHashsetCreator[int16]
	Int32   typedHashsetCreator[int32]
	Int64   typedHashsetCreator[int64]
	Uint    typedHashsetCreator[uint]
	Uint8   typedHashsetCreator[uint8]
	Uint16  typedHashsetCreator[uint16]
	Uint32  typedHashsetCreator[uint32]
	Uint64  typedHashsetCreator[uint64]
	Float32 typedHashsetCreator[float32]
	Float64 typedHashsetCreator[float64]
	Byte    typedHashsetCreator[byte]
	Bool    typedHashsetCreator[bool]
}
