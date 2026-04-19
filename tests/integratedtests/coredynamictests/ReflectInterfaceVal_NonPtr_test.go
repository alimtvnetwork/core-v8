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

package coredynamictests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── ReflectInterfaceVal non-ptr (line 20) ──

func Test_ReflectInterfaceVal_NonPtr_I29(t *testing.T) {
	// Arrange / Act
	result := coredynamic.ReflectInterfaceVal(42)

	// Assert
	actual := args.Map{"result": result}
	expected := args.Map{"result": 42}
	actual.ShouldBeEqual(t, 1, "ReflectInterfaceVal non-ptr", expected)
}

// ── ReflectSetFromTo: []byte → struct (line 159-167) ──

func Test_ReflectSetFromTo_BytesToStruct_I29(t *testing.T) {
	// Arrange
	type sample struct {
		Name string `json:"name"`
	}
	data, _ := json.Marshal(sample{Name: "test"})
	var target sample

	// Act
	err := coredynamic.ReflectSetFromTo(data, &target)

	// Assert
	actual := args.Map{
		"hasError": err != nil,
		"name":     target.Name,
	}
	expected := args.Map{
		"hasError": false,
		"name":     "test",
	}
	actual.ShouldBeEqual(t, 1, "ReflectSetFromTo bytes to struct", expected)
}

// ── ReflectSetFromTo: struct → *[]byte (line 174-180) ──

func Test_ReflectSetFromTo_StructToBytes_I29(t *testing.T) {
	// Arrange
	type sample struct {
		Name string `json:"name"`
	}
	from := sample{Name: "test"}
	var target []byte

	// Act
	err := coredynamic.ReflectSetFromTo(from, &target)

	// Assert
	actual := args.Map{
		"hasError":   err != nil,
		"hasContent": len(target) > 0,
	}
	expected := args.Map{
		"hasError":   false,
		"hasContent": true,
	}
	actual.ShouldBeEqual(t, 1, "ReflectSetFromTo struct to bytes", expected)
}

// ── SafeZeroSet non-pointer (line 18) ──

func Test_SafeZeroSet_NonPointer_I29(t *testing.T) {
	// Arrange
	val := reflect.ValueOf(42)

	// Act — should not panic on non-pointer
	coredynamic.SafeZeroSet(val)

	// Assert
	actual := args.Map{"noPanic": true}
	expected := args.Map{"noPanic": true}
	actual.ShouldBeEqual(t, 1, "SafeZeroSet non-pointer", expected)
}
