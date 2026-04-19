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

// newHashmapCreator aggregates per-type hashmap creators.
//
// Usage:
//
//	coregeneric.New.Hashmap.StringString.Cap(10)
//	coregeneric.New.Hashmap.StringInt.Empty()
//	coregeneric.New.Hashmap.StringAny.From(existingMap)
type newHashmapCreator struct {
	StringString  typedHashmapCreator[string, string]
	StringInt     typedHashmapCreator[string, int]
	StringInt64   typedHashmapCreator[string, int64]
	StringFloat64 typedHashmapCreator[string, float64]
	StringBool    typedHashmapCreator[string, bool]
	StringAny     typedHashmapCreator[string, any]
	IntString     typedHashmapCreator[int, string]
	IntInt        typedHashmapCreator[int, int]
	IntAny        typedHashmapCreator[int, any]
}
