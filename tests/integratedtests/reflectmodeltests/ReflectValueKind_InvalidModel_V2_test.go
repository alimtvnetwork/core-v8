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

package reflectmodeltests

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/reflectcore/reflectmodel"
)

func Test_ReflectValueKind_InvalidModel(t *testing.T) {
	// Arrange
	m := reflectmodel.InvalidReflectValueKindModel("test error")

	// Act
	actual := args.Map{
		"isInvalid": m.IsInvalid(), "hasError": m.HasError(),
		"isEmptyErr": m.IsEmptyError(), "typeName": m.TypeName(),
	}

	// Assert
	expected := args.Map{
		"isInvalid": true, "hasError": true,
		"isEmptyErr": false, "typeName": "",
	}
	expected.ShouldBeEqual(t, 0, "InvalidModel returns error -- with args", actual)
}

func Test_ReflectValueKind_NilMethods(t *testing.T) {
	// Arrange
	var m *reflectmodel.ReflectValueKind

	// Act
	actual := args.Map{
		"isInvalid": m.IsInvalid(), "hasError": m.HasError(),
		"isEmptyErr": m.IsEmptyError(), "actualNil": m.ActualInstance() == nil,
		"pkgPath": m.PkgPath(), "pointerRv": m.PointerRv() == nil,
		"typeName": m.TypeName(), "pointerInf": m.PointerInterface() == nil,
	}

	// Assert
	expected := args.Map{
		"isInvalid": true, "hasError": false,
		"isEmptyErr": true, "actualNil": true,
		"pkgPath": "", "pointerRv": true,
		"typeName": "", "pointerInf": true,
	}
	expected.ShouldBeEqual(t, 0, "NilMethods returns nil -- with args", actual)
}

func Test_ReflectValueKind_ValidModel(t *testing.T) {
	// Arrange
	rv := reflect.ValueOf(42)
	m := &reflectmodel.ReflectValueKind{IsValid: true, FinalReflectVal: rv, Kind: rv.Kind()}

	// Act
	actual := args.Map{
		"isInvalid": m.IsInvalid(), "hasError": m.HasError(),
		"typeName": m.TypeName() != "", "pkgPath": m.PkgPath(),
		"actualVal": m.ActualInstance(), "ptrNotNil": m.PointerRv() != nil,
		"ptrInfNotNil": m.PointerInterface() != nil,
	}

	// Assert
	expected := args.Map{
		"isInvalid": false, "hasError": false,
		"typeName": true, "pkgPath": "",
		"actualVal": 42, "ptrNotNil": true,
		"ptrInfNotNil": true,
	}
	expected.ShouldBeEqual(t, 0, "ValidModel returns non-empty -- with args", actual)
}

func Test_ReflectValueKind_InvalidNotValid(t *testing.T) {
	// Arrange
	m := &reflectmodel.ReflectValueKind{IsValid: false, FinalReflectVal: reflect.ValueOf(42)}

	// Act
	actual := args.Map{"ptrRvNotNil": m.PointerRv() != nil}

	// Assert
	expected := args.Map{"ptrRvNotNil": true}
	expected.ShouldBeEqual(t, 0, "InvalidNotValid returns error -- with args", actual)
}

func Test_ReflectValueKind_WithError(t *testing.T) {
	// Arrange
	m := &reflectmodel.ReflectValueKind{
		IsValid: true, Error: errors.New("some error"),
		FinalReflectVal: reflect.ValueOf(42), Kind: reflect.Int,
	}

	// Act
	actual := args.Map{
		"isInvalid": m.IsInvalid(),
		"hasError": m.HasError(),
	}

	// Assert
	expected := args.Map{
		"isInvalid": true,
		"hasError": true,
	}
	expected.ShouldBeEqual(t, 0, "WithError returns error -- with args", actual)
}

func Test_ReflectValue_Fields_FromReflectValueKindInva(t *testing.T) {
	// Arrange
	rv := reflectmodel.ReflectValue{
		TypeName: "MyType", FieldsNames: []string{"A", "B"},
		MethodsNames: []string{"M1"}, RawData: 42,
	}

	// Act
	actual := args.Map{
		"typeName": rv.TypeName, "fieldsLen": len(rv.FieldsNames),
		"methodsLen": len(rv.MethodsNames), "rawData": rv.RawData,
	}

	// Assert
	expected := args.Map{
		"typeName": "MyType", "fieldsLen": 2,
		"methodsLen": 1, "rawData": 42,
	}
	expected.ShouldBeEqual(t, 0, "ReflectValue returns correct value -- with args", actual)
}
