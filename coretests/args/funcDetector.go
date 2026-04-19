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

// funcDetector provides utility methods for detecting and extracting
// FuncWrapAny instances from various input types.
type funcDetector struct{}

// GetFuncWrap extracts a FuncWrapAny from various input types.
// Supports Map, *FuncWrapAny, FuncWrapGetter, ArgsMapper,
// and falls back to creating a new FuncWrapAny via reflection.
func (it funcDetector) GetFuncWrap(i any) *FuncWrapAny {
	switch v := i.(type) {
	case Map:
		return v.FuncWrap()
	case *FuncWrapAny:
		return v
	case FuncWrapGetter:
		return v.FuncWrap()
	case ArgsMapper:
		return v.FuncWrap()
	default:
		return NewFuncWrap.Default(i)
	}
}
