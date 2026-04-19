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

package args

// emptyCreator provides factory methods for creating empty/invalid
// instances of common arg types.
type emptyCreator struct{}

// Map returns an empty Map.
func (it emptyCreator) Map() Map {
	return map[string]any{}
}

// FuncWrap returns an invalid FuncWrapAny.
func (it emptyCreator) FuncWrap() *FuncWrapAny {
	return &FuncWrapAny{
		isInvalid: true,
	}
}

// FuncMap returns an empty FuncMap.
func (it emptyCreator) FuncMap() FuncMap {
	return map[string]FuncWrapAny{}
}

// Holder returns an empty HolderAny.
func (it emptyCreator) Holder() HolderAny {
	return HolderAny{}
}
