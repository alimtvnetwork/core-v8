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

import (
	"reflect"

	"github.com/alimtvnetwork/core-v8/errcore"
	"github.com/alimtvnetwork/core-v8/internal/reflectinternal"
	"github.com/alimtvnetwork/core-v8/iserror"
)

type newFuncWrapCreator struct{}

// Default creates a FuncWrapAny from any value.
// If the value is nil, not a function, or already a FuncWrap,
// it handles each case appropriately.
func (it newFuncWrapCreator) Default(anyFunc any) *FuncWrapAny {
	if reflectinternal.Is.Null(anyFunc) {
		return &FuncWrapAny{
			Func:      anyFunc,
			isInvalid: true,
		}
	}

	switch v := anyFunc.(type) {
	case *FuncWrapAny:
		return v
	case FuncWrapGetter:
		return v.FuncWrap()
	}

	typeOf := reflect.TypeOf(anyFunc)
	kind := typeOf.Kind()

	if kind != reflect.Func {
		return &FuncWrapAny{
			Func:      anyFunc,
			isInvalid: true,
			rvType:    typeOf,
		}
	}

	fullName, nameOnly := reflectinternal.GetFunc.FullNameWithName(anyFunc)

	return &FuncWrapAny{
		Name:      nameOnly,
		FullName:  fullName,
		Func:      anyFunc,
		isInvalid: false,
		rvType:    typeOf,
		rv:        reflect.ValueOf(anyFunc),
	}
}

// Single is an alias for Default.
func (it newFuncWrapCreator) Single(anyFunc any) *FuncWrapAny {
	return it.Default(anyFunc)
}

// Invalid creates an invalid FuncWrapAny.
func (it newFuncWrapCreator) Invalid() *FuncWrapAny {
	return &FuncWrapAny{
		isInvalid: true,
	}
}

// Map creates a FuncMap from multiple function values.
func (it newFuncWrapCreator) Map(anyFunctions ...any) FuncMap {
	if len(anyFunctions) == 0 {
		return map[string]FuncWrapAny{}
	}

	newMap := make(map[string]FuncWrapAny, len(anyFunctions))

	for _, function := range anyFunctions {
		v := it.Default(function)

		if v.IsValid() {
			newMap[v.GetFuncName()] = *v
		}
	}

	return newMap
}

// Many creates a slice of FuncWrapAny pointers from multiple function values.
func (it newFuncWrapCreator) Many(anyFunctions ...any) []*FuncWrapAny {
	if len(anyFunctions) == 0 {
		return []*FuncWrapAny{}
	}

	slice := make([]*FuncWrapAny, len(anyFunctions))

	for i, function := range anyFunctions {
		slice[i] = it.Default(function)
	}

	return slice
}

// MethodToFunc converts a reflect.Method to a FuncWrapAny.
func (it newFuncWrapCreator) MethodToFunc(m *reflect.Method) (*FuncWrapAny, error) {
	if m == nil {
		return it.Invalid(), errcore.CannotBeNilType.ErrorNoRefs(
			"m * method cannot be nil",
		)
	}

	name := m.Name
	fullName := m.PkgPath + name

	return &FuncWrapAny{
		Name:      name,
		FullName:  fullName,
		Func:      m.Func.Interface(),
		isInvalid: false,
		rvType:    m.Func.Type(),
		rv:        m.Func,
	}, nil
}

// StructToMap creates a FuncMap from all public methods of a struct.
func (it newFuncWrapCreator) StructToMap(i any) (FuncMap, error) {
	methods, err := reflectinternal.Looper.MethodsMap(i)

	if iserror.Defined(err) {
		return Empty.FuncMap(), err
	}

	newMap := make(map[string]FuncWrapAny, len(methods))
	var rawErr errcore.RawErrCollection

	for index, method := range methods {
		v, nErr := it.MethodToFunc(method)
		rawErr.Add(nErr)

		if v.IsValid() {
			newMap[index] = *v
		}
	}

	return newMap, rawErr.CompiledErrorWithStackTraces()
}
