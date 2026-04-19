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
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// ── FieldProcessor ──

func Test_FieldProcessor_IsFieldType(t *testing.T) {
	// Arrange
	fp := &reflectmodel.FieldProcessor{
		Name: "A", Index: 0,
		FieldType: reflect.TypeOf(42),
	}
	var nilFP *reflectmodel.FieldProcessor

	// Act
	actual := args.Map{
		"isInt":    fp.IsFieldType(reflect.TypeOf(42)),
		"isString": fp.IsFieldType(reflect.TypeOf("")),
		"nil":      nilFP.IsFieldType(reflect.TypeOf(42)),
	}

	// Assert
	expected := args.Map{
		"isInt": true,
		"isString": false,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "FieldProcessor returns correct value -- IsFieldType", actual)
}

func Test_FieldProcessor_IsFieldKind(t *testing.T) {
	// Arrange
	fp := &reflectmodel.FieldProcessor{
		FieldType: reflect.TypeOf(42),
	}
	var nilFP *reflectmodel.FieldProcessor

	// Act
	actual := args.Map{
		"isInt":  fp.IsFieldKind(reflect.Int),
		"isStr":  fp.IsFieldKind(reflect.String),
		"nil":    nilFP.IsFieldKind(reflect.Int),
	}

	// Assert
	expected := args.Map{
		"isInt": true,
		"isStr": false,
		"nil": false,
	}
	expected.ShouldBeEqual(t, 0, "FieldProcessor returns correct value -- IsFieldKind", actual)
}

// ── MethodProcessor extended ──

func Test_MethodProcessor_Nil(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	// Only test methods that are safe on nil receivers

	// Act
	actual := args.Map{
		"hasValid":  mp.HasValidFunc(),
		"isInvalid": mp.IsInvalid(),
		"funcNil":   mp.Func() == nil,
		"isPublic":  mp.IsPublicMethod(),
		"isPrivate": mp.IsPrivateMethod(),
	}

	// Assert
	expected := args.Map{
		"hasValid": false, "isInvalid": true,
		"funcNil": true, "isPublic": false, "isPrivate": false,
	}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns nil -- Nil", actual)
}

func Test_MethodProcessor_IsEqual(t *testing.T) {
	// Arrange
	var mp1, mp2 *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{"nilNil": mp1.IsEqual(mp2)}

	// Assert
	expected := args.Map{"nilNil": true}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns nil -- IsEqual nil", actual)
}

func Test_MethodProcessor_IsNotEqual(t *testing.T) {
	// Arrange
	var mp1, mp2 *reflectmodel.MethodProcessor

	// Act
	actual := args.Map{"notEqual": mp1.IsNotEqual(mp2)}

	// Assert
	expected := args.Map{"notEqual": false}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns correct value -- IsNotEqual", actual)
}

func Test_MethodProcessor_GetOutArgsTypes_Nil(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	result := mp.GetOutArgsTypes()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns nil -- GetOutArgsTypes nil", actual)
}

func Test_MethodProcessor_GetInArgsTypes_Nil(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	result := mp.GetInArgsTypes()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns nil -- GetInArgsTypes nil", actual)
}

func Test_MethodProcessor_GetInArgsTypesNames_Nil(t *testing.T) {
	// Arrange
	var mp *reflectmodel.MethodProcessor
	result := mp.GetInArgsTypesNames()

	// Act
	actual := args.Map{"len": len(result)}

	// Assert
	expected := args.Map{"len": 0}
	expected.ShouldBeEqual(t, 0, "MethodProcessor returns nil -- GetInArgsTypesNames nil", actual)
}

// ── rvUtils is unexported — cannot be tested from external package ──
