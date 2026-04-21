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

package typesconv

import "github.com/alimtvnetwork/core-v8/constants"

func FloatPtr(val float32) *float32 {
	return &val
}

// FloatPtrToSimple if nil then 0
func FloatPtrToSimple(val *float32) float32 {
	if val == nil {
		return constants.Zero
	}

	return *val
}

// FloatPtrToSimpleDef if nil then 0
func FloatPtrToSimpleDef(val *float32, defVal float32) float32 {
	if val == nil {
		return defVal
	}

	return *val
}

// FloatPtrToDefPtr if nil then 0
func FloatPtrToDefPtr(val *float32, defVal float32) *float32 {
	if val == nil {
		return &defVal
	}

	return val
}

// FloatPtrDefValFunc if nil then executes returns defValFunc result as pointer
func FloatPtrDefValFunc(val *float32, defValFunc func() float32) *float32 {
	if val == nil {
		result := defValFunc()

		return &result
	}

	return val
}
