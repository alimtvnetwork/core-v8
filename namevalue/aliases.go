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

package namevalue

// StringAny is the backward-compatible alias for Instance[string, any].
// Use this where the old non-generic Instance was used.
type StringAny = Instance[string, any]

// StringString represents a name-value pair where both are strings.
type StringString = Instance[string, string]

// StringInt represents a name-value pair with a string key and int value.
type StringInt = Instance[string, int]

// StringMapAny represents a name-value pair with a string key and map[string]any value.
type StringMapAny = Instance[string, map[string]any]

// StringMapString represents a name-value pair with a string key and map[string]string value.
type StringMapString = Instance[string, map[string]string]

// StringStringCollection is a Collection of StringString items.
type StringStringCollection = Collection[string, string]

// StringIntCollection is a Collection of StringInt items.
type StringIntCollection = Collection[string, int]

// StringMapAnyCollection is a Collection of StringMapAny items.
type StringMapAnyCollection = Collection[string, map[string]any]

// StringMapStringCollection is a Collection of StringMapString items.
type StringMapStringCollection = Collection[string, map[string]string]
