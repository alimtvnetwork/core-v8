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

import "github.com/alimtvnetwork/core/constants"

func BytePtr(val byte) *byte {
	return &val
}

// BytePtrToSimple if nil then 0
func BytePtrToSimple(val *byte) byte {
	if val == nil {
		return constants.Zero
	}

	return *val
}

// BytePtrToSimpleDef if nil then 0
func BytePtrToSimpleDef(val *byte, defVal byte) byte {
	if val == nil {
		return defVal
	}

	return *val
}

// BytePtrToDefPtr if nil then 0
func BytePtrToDefPtr(val *byte, defVal byte) *byte {
	if val == nil {
		return &defVal
	}

	return val
}

// BytePtrDefValFunc if nil then executes returns defValFunc result as pointer
func BytePtrDefValFunc(val *byte, defValFunc func() byte) *byte {
	if val == nil {
		result := defValFunc()

		return &result
	}

	return val
}
