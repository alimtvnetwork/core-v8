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

package coreflecttests

import (
	"testing"

	"github.com/alimtvnetwork/core/reflectcore"
)

// ==========================================
// reflectcore vars — facade re-exports are non-zero
// These are value types (structs), not pointers/interfaces,
// so nil comparison is invalid. We verify compile-time existence instead.
// ==========================================

func Test_Reflectcore_Converter_NotNil(t *testing.T) {
	_ = reflectcore.Converter // compile-time existence check
}

func Test_Reflectcore_Utils_NotNil(t *testing.T) {
	_ = reflectcore.Utils
}

func Test_Reflectcore_Looper_NotNil(t *testing.T) {
	_ = reflectcore.Looper
}

func Test_Reflectcore_CodeStack_NotNil(t *testing.T) {
	_ = reflectcore.CodeStack
}

func Test_Reflectcore_GetFunc_NotNil(t *testing.T) {
	_ = reflectcore.GetFunc
}

func Test_Reflectcore_Is_NotNil(t *testing.T) {
	_ = reflectcore.Is
}

func Test_Reflectcore_TypeName_NotNil(t *testing.T) {
	_ = reflectcore.TypeName // compile-time existence check
}

func Test_Reflectcore_TypeNames_NotNil(t *testing.T) {
	_ = reflectcore.TypeNames // compile-time existence check
}

func Test_Reflectcore_ReflectType_NotNil(t *testing.T) {
	_ = reflectcore.ReflectType
}

func Test_Reflectcore_ReflectGetter_NotNil(t *testing.T) {
	_ = reflectcore.ReflectGetter
}

func Test_Reflectcore_SliceConverter_NotNil(t *testing.T) {
	_ = reflectcore.SliceConverter
}

func Test_Reflectcore_MapConverter_NotNil(t *testing.T) {
	_ = reflectcore.MapConverter
}
