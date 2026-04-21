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
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ── KeyVal.CastKeyVal nil receiver (line 134) ──

func Test_KeyVal_CastKeyVal_NilReceiver_I29(t *testing.T) {
	// Arrange
	var kv *coredynamic.KeyVal

	// Act
	var k string
	var v int
	err := kv.CastKeyVal(&k, &v)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "KeyVal CastKeyVal nil receiver", expected)
}

// ── KeyVal.JsonParseSelfInject (line 300) ──

func Test_KeyVal_JsonParseSelfInject_I29(t *testing.T) {
	// Arrange
	kv := &coredynamic.KeyVal{Key: "a", Value: "b"}
	jsonResult := corejson.New(kv)

	// Act
	err := kv.JsonParseSelfInject(&jsonResult)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "KeyVal JsonParseSelfInject", expected)
}

// ── KeyValCollection operations ──

func Test_KeyValCollection_ParseInjectUsingJson_Valid_I29(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(3)
	kvc.Add(coredynamic.KeyVal{Key: "k1", Value: "v1"})
	jsonResult := corejson.New(kvc)

	// Act
	result, err := kvc.ParseInjectUsingJson(&jsonResult)

	// Assert
	actual := args.Map{
		"notNil":   result != nil,
		"hasError": err != nil,
	}
	expected := args.Map{
		"notNil":   true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "KeyValCollection ParseInjectUsingJson valid", expected)
}

func Test_KeyValCollection_JsonParseSelfInject_I29(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "x", Value: "y"})
	jsonResult := corejson.New(kvc)

	// Act
	err := kvc.JsonParseSelfInject(&jsonResult)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	actual.ShouldBeEqual(t, 1, "KeyValCollection JsonParseSelfInject", expected)
}

func Test_KeyValCollection_Serialize_Valid_I29(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: "b"})

	// Act
	bytes, err := kvc.Serialize()

	// Assert
	actual := args.Map{
		"hasBytes": len(bytes) > 0,
		"hasError": err != nil,
	}
	expected := args.Map{
		"hasBytes": true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "KeyValCollection Serialize valid", expected)
}

func Test_KeyValCollection_JsonString_Valid_I29(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "a", Value: "b"})

	// Act
	jsonStr, err := kvc.JsonString()

	// Assert
	actual := args.Map{
		"hasContent": len(jsonStr) > 0,
		"hasError":   err != nil,
	}
	expected := args.Map{
		"hasContent": true,
		"hasError":   false,
	}
	actual.ShouldBeEqual(t, 1, "KeyValCollection JsonString valid", expected)
}

func Test_KeyValCollection_ParseInjectUsingJsonMust_Valid_I29(t *testing.T) {
	// Arrange
	kvc := coredynamic.NewKeyValCollection(2)
	kvc.Add(coredynamic.KeyVal{Key: "k", Value: "v"})
	jsonResult := corejson.New(kvc)

	// Act
	result := kvc.ParseInjectUsingJsonMust(&jsonResult)

	// Assert
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	actual.ShouldBeEqual(t, 1, "KeyValCollection ParseInjectUsingJsonMust valid", expected)
}
