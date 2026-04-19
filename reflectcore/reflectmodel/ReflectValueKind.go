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

package reflectmodel

import (
	"errors"
	"reflect"
	"unsafe"
)

type ReflectValueKind struct {
	IsValid         bool
	FinalReflectVal reflect.Value
	Kind            reflect.Kind
	Error           error
}

func InvalidReflectValueKindModel(err string) *ReflectValueKind {
	return &ReflectValueKind{
		IsValid:         false,
		FinalReflectVal: reflect.ValueOf(nil),
		Kind:            0,
		Error:           errors.New(err),
	}
}

func (it *ReflectValueKind) IsInvalid() bool {
	return it == nil || !it.IsValid || it.HasError()
}

func (it *ReflectValueKind) HasError() bool {
	return it != nil && it.Error != nil
}

func (it *ReflectValueKind) IsEmptyError() bool {
	return it == nil || it.Error == nil
}

func (it *ReflectValueKind) ActualInstance() any {
	if it == nil {
		return nil
	}

	return it.FinalReflectVal.Interface()
}

func (it *ReflectValueKind) PkgPath() string {
	if it == nil || !it.IsValid {
		return ""
	}

	return it.FinalReflectVal.Type().PkgPath()
}

func (it *ReflectValueKind) PointerRv() *reflect.Value {
	if it == nil {
		return nil
	}

	if !it.IsValid {
		return &it.FinalReflectVal
	}

	rv := it.FinalReflectVal

	toInterface := rv.Interface()
	toPointer := &toInterface
	unsafePtr := unsafe.Pointer(&toPointer)

	newRv := reflect.NewAt(rv.Type(), unsafePtr)

	return &newRv
}

func (it *ReflectValueKind) TypeName() string {
	if it == nil || !it.IsValid {
		return ""
	}

	rv := it.FinalReflectVal

	return rv.String()
}

func (it *ReflectValueKind) PointerInterface() any {
	rv := it.PointerRv()

	if rv == nil {
		return nil
	}

	return rv.Interface()
}
