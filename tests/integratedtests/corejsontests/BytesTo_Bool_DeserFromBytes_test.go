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

package corejsontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── deserializeFromBytesTo: Bool, BoolMust ──

func Test_BytesTo_Bool_Valid(t *testing.T) {
	// Arrange
	jsonBytes, _ := corejson.Serialize.Raw(true)

	// Act
	result, err := corejson.Deserialize.BytesTo.Bool(jsonBytes)

	// Assert
	actual := args.Map{
		"val": result,
		"noErr": err == nil,
	}
	expected := args.Map{
		"val": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesTo.Bool returns true -- valid", actual)
}

func Test_BytesTo_BoolMust_Valid(t *testing.T) {
	// Arrange
	jsonBytes, _ := corejson.Serialize.Raw(false)

	// Act
	result := corejson.Deserialize.BytesTo.BoolMust(jsonBytes)

	// Assert
	actual := args.Map{"val": result}
	expected := args.Map{"val": false}
	expected.ShouldBeEqual(t, 0, "BytesTo.BoolMust returns false -- valid", actual)
}

// ── deserializeFromBytesTo: Integer64, Integer64Must ──

func Test_BytesTo_Integer64_Valid(t *testing.T) {
	// Arrange
	jsonBytes, _ := corejson.Serialize.Raw(int64(9999999999))

	// Act
	result, err := corejson.Deserialize.BytesTo.Integer64(jsonBytes)

	// Assert
	actual := args.Map{
		"val": result,
		"noErr": err == nil,
	}
	expected := args.Map{
		"val": int64(9999999999),
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesTo.Integer64 returns value -- valid", actual)
}

func Test_BytesTo_Integer64Must_Valid(t *testing.T) {
	// Arrange
	jsonBytes, _ := corejson.Serialize.Raw(int64(42))

	// Act
	result := corejson.Deserialize.BytesTo.Integer64Must(jsonBytes)

	// Assert
	actual := args.Map{"val": result}
	expected := args.Map{"val": int64(42)}
	expected.ShouldBeEqual(t, 0, "BytesTo.Integer64Must returns 42 -- valid", actual)
}

// ── deserializeFromBytesTo: MapStringString, MapStringStringMust ──

func Test_BytesTo_MapStringString_Valid(t *testing.T) {
	// Arrange
	m := map[string]string{"a": "1", "b": "2"}
	jsonBytes, _ := corejson.Serialize.Raw(m)

	// Act
	result, err := corejson.Deserialize.BytesTo.MapStringString(jsonBytes)

	// Assert
	actual := args.Map{
		"len": len(result),
		"noErr": err == nil,
		"a": result["a"],
	}
	expected := args.Map{
		"len": 2,
		"noErr": true,
		"a": "1",
	}
	expected.ShouldBeEqual(t, 0, "BytesTo.MapStringString returns map -- valid", actual)
}

func Test_BytesTo_MapStringStringMust_Valid(t *testing.T) {
	// Arrange
	m := map[string]string{"x": "y"}
	jsonBytes, _ := corejson.Serialize.Raw(m)

	// Act
	result := corejson.Deserialize.BytesTo.MapStringStringMust(jsonBytes)

	// Assert
	actual := args.Map{"x": result["x"]}
	expected := args.Map{"x": "y"}
	expected.ShouldBeEqual(t, 0, "BytesTo.MapStringStringMust returns map -- valid", actual)
}

// ── deserializeFromBytesTo: Bytes, BytesMust ──

func Test_BytesTo_Bytes_Valid(t *testing.T) {
	// Arrange
	original := []byte("hello")
	jsonBytes, _ := corejson.Serialize.Raw(original)

	// Act
	result, err := corejson.Deserialize.BytesTo.Bytes(jsonBytes)

	// Assert
	actual := args.Map{
		"hasContent": len(result) > 0,
		"noErr": err == nil,
	}
	expected := args.Map{
		"hasContent": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesTo.Bytes returns bytes -- valid", actual)
}

func Test_BytesTo_BytesMust_Valid(t *testing.T) {
	// Arrange
	original := []byte("world")
	jsonBytes, _ := corejson.Serialize.Raw(original)

	// Act
	result := corejson.Deserialize.BytesTo.BytesMust(jsonBytes)

	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	expected.ShouldBeEqual(t, 0, "BytesTo.BytesMust returns bytes -- valid", actual)
}

// ── deserializeFromBytesTo: ResultsPtrCollection, ResultsPtrCollectionMust ──

func Test_BytesTo_ResultsPtrCollection_Valid(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.AnyItems("a", "b")
	jsonBytes, _ := corejson.Serialize.Raw(coll)

	// Act
	result, err := corejson.Deserialize.BytesTo.ResultsPtrCollection(jsonBytes)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"hasItems": result != nil && result.Length() > 0,
	}
	expected := args.Map{
		"noErr": true,
		"hasItems": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesTo.ResultsPtrCollection returns coll -- valid", actual)
}

func Test_BytesTo_ResultsPtrCollectionMust_Valid(t *testing.T) {
	// Arrange
	coll := corejson.NewResultsPtrCollection.AnyItems("x")
	jsonBytes, _ := corejson.Serialize.Raw(coll)

	// Act
	result := corejson.Deserialize.BytesTo.ResultsPtrCollectionMust(jsonBytes)

	// Assert
	actual := args.Map{"hasItems": result != nil && result.Length() > 0}
	expected := args.Map{"hasItems": true}
	expected.ShouldBeEqual(t, 0, "BytesTo.ResultsPtrCollectionMust returns coll -- valid", actual)
}

// ── deserializeFromBytesTo: MapResults, MapResultsMust ──

func Test_BytesTo_MapResults_Valid(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.UsingMapAnyItems(map[string]any{
		"key1": "val1",
	})
	jsonBytes, _ := corejson.Serialize.Raw(mr)

	// Act
	result, err := corejson.Deserialize.BytesTo.MapResults(jsonBytes)

	// Assert
	actual := args.Map{
		"noErr": err == nil,
		"notNil": result != nil,
	}
	expected := args.Map{
		"noErr": true,
		"notNil": true,
	}
	expected.ShouldBeEqual(t, 0, "BytesTo.MapResults returns map -- valid", actual)
}

func Test_BytesTo_MapResultsMust_Valid(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.UsingMapAnyItems(map[string]any{
		"k": "v",
	})
	jsonBytes, _ := corejson.Serialize.Raw(mr)

	// Act
	result := corejson.Deserialize.BytesTo.MapResultsMust(jsonBytes)

	// Assert
	actual := args.Map{"notNil": result != nil}
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "BytesTo.MapResultsMust returns map -- valid", actual)
}
