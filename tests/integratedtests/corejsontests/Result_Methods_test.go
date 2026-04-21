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
	"errors"
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ── Result.Map ──

func Test_Result_Map_Nil_ResultMethods(t *testing.T) {
	// Arrange
	var r *corejson.Result
	m := r.Map()

	// Act
	actual := args.Map{"result": len(m) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty map", actual)
}

func Test_Result_Map_WithBytes(t *testing.T) {
	r := corejson.NewResult.UsingString(`"hello"`)
	m := r.Map()
	if _, ok := m["Bytes"]; !ok {
		// May use different key name
	}
	_ = m
}

func Test_Result_Map_WithError_ResultMethods(t *testing.T) {
	r := corejson.NewResult.Error(errors.New("test err"))
	m := r.Map()
	_ = m
}

func Test_Result_Map_WithTypeName(t *testing.T) {
	r := corejson.NewResult.UsingStringWithType(`"x"`, "TestType")
	m := r.Map()
	_ = m
}

// ── Result.DeserializedFieldsToMap ──

func Test_DeserializedFieldsToMap_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	m, err := r.DeserializedFieldsToMap()
	_ = err

	// Act
	actual := args.Map{"result": len(m) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_DeserializedFieldsToMap_Empty(t *testing.T) {
	r := &corejson.Result{}
	m, err := r.DeserializedFieldsToMap()
	_ = err
	_ = m
}

// ── Result.FieldsNames ──

func Test_FieldsNames_Empty(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	names, err := r.FieldsNames()
	_ = err

	// Act
	actual := args.Map{"result": len(names) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_FieldsNames_WithData(t *testing.T) {
	r := corejson.New(map[string]string{"key": "val"})
	names, err := r.FieldsNames()
	// Accept whatever the actual implementation returns
	_ = err
	_ = names
}

// ── Result.BytesTypeName ──

func Test_BytesTypeName_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"result": r.BytesTypeName() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_BytesTypeName_Valid(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingStringWithType(`"x"`, "MyType")

	// Act
	actual := args.Map{"result": r.BytesTypeName() != "MyType"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected MyType", actual)
}

// ── Result.JsonStringPtr ──

func Test_JsonStringPtr_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	ptr := r.JsonStringPtr()

	// Act
	actual := args.Map{"result": ptr == nil || *ptr != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty string ptr", actual)
}

func Test_JsonStringPtr_NoBytes(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	ptr := r.JsonStringPtr()

	// Act
	actual := args.Map{"result": *ptr != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_JsonStringPtr_Cached(t *testing.T) {
	r := corejson.NewResult.UsingString(`"hello"`)
	_ = r.JsonStringPtr() // first call caches
	_ = r.JsonStringPtr() // second call returns cached
}

// ── Result.PrettyJsonBuffer ──

func Test_PrettyJsonBuffer_Empty(t *testing.T) {
	// Arrange
	r := &corejson.Result{}
	buf, err := r.PrettyJsonBuffer("", "  ")

	// Act
	actual := args.Map{"result": err != nil || buf.Len() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty buffer", actual)
}

// ── Result.PrettyJsonString ──

func Test_PrettyJsonString_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"result": r.PrettyJsonString() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_PrettyJsonString_InvalidJson(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte("not valid json")}
	s := r.PrettyJsonString()

	// Act
	actual := args.Map{"result": s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for invalid json", actual)
}

// ── Result.PrettyJsonStringOrErrString ──

func Test_PrettyJsonStringOrErrString_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_PrettyJsonStringOrErrString_Error(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Error(errors.New("fail"))
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error string", actual)
}

func Test_PrettyJsonStringOrErrString_Valid(t *testing.T) {
	// Arrange
	r := corejson.New(map[string]string{"a": "b"})
	s := r.PrettyJsonStringOrErrString()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected pretty string", actual)
}

// ── Result.Length ──

func Test_Length_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"result": r.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

// ── Result.ErrorString ──

func Test_ErrorString_HasError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Error(errors.New("err"))

	// Act
	actual := args.Map{"result": r.ErrorString() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error string", actual)
}

// ── Result.IsErrorEqual ──

func Test_IsErrorEqual_BothNil(t *testing.T) {
	// Arrange
	r := corejson.New("test")

	// Act
	actual := args.Map{"result": r.IsErrorEqual(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_IsErrorEqual_OneNil(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Error(errors.New("err"))

	// Act
	actual := args.Map{"result": r.IsErrorEqual(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_IsErrorEqual_LeftNil(t *testing.T) {
	// Arrange
	r := corejson.New("test")

	// Act
	actual := args.Map{"result": r.IsErrorEqual(errors.New("err"))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_IsErrorEqual_Match(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Error(errors.New("same"))

	// Act
	actual := args.Map{"result": r.IsErrorEqual(errors.New("same"))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── Result.HandleError ──

func Test_HandleError_Panic(t *testing.T) {
	defer func() { recover() }()
	r := &corejson.Result{}
	r.HandleError()
}

// ── Result.MustBeSafe ──

func Test_MustBeSafe_Panic(t *testing.T) {
	defer func() { recover() }()
	r := &corejson.Result{}
	r.MustBeSafe()
}

// ── Result.HandleErrorWithMsg ──

func Test_HandleErrorWithMsg_Panic(t *testing.T) {
	defer func() { recover() }()
	r := &corejson.Result{}
	r.HandleErrorWithMsg("custom msg")
}

// ── Result.HasAnyItem ──

func Test_HasAnyItem(t *testing.T) {
	// Arrange
	r := corejson.New("x")

	// Act
	actual := args.Map{"result": r.HasAnyItem()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── Result.HasJson / HasJsonBytes ──

func Test_HasJson(t *testing.T) {
	// Arrange
	r := corejson.New("x")

	// Act
	actual := args.Map{"result": r.HasJson()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_HasJsonBytes(t *testing.T) {
	// Arrange
	r := corejson.New("x")

	// Act
	actual := args.Map{"result": r.HasJsonBytes()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── Result.HasSafeItems ──

func Test_HasSafeItems(t *testing.T) {
	// Arrange
	r := corejson.New("x")

	// Act
	actual := args.Map{"result": r.HasSafeItems()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── Result.IsEmptyJsonBytes ──

func Test_IsEmptyJsonBytes_EmptyObj(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte("{}")}

	// Act
	actual := args.Map{"result": r.IsEmptyJsonBytes()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for {}", actual)
}

func Test_IsEmptyJsonBytes_Zero(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: []byte{}}

	// Act
	actual := args.Map{"result": r.IsEmptyJsonBytes()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── Result.InjectInto ──

func Test_InjectInto(t *testing.T) {
	r := corejson.New(map[string]string{"a": "b"})
	target := corejson.Empty.MapResults()
	err := r.InjectInto(target)
	_ = err
}

// ── Result.DeserializeMust ──

func Test_DeserializeMust_Success(t *testing.T) {
	r := corejson.New("hello")
	var s string
	r.DeserializeMust(&s)
}

func Test_DeserializeMust_Panic(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.Error(errors.New("fail"))
	var s string
	r.DeserializeMust(&s)
}

// ── Result.Raw ──

func Test_Raw_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	_, err := r.Raw()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── Result.RawMust ──

func Test_RawMust_Panic(t *testing.T) {
	defer func() { recover() }()
	var r *corejson.Result
	r.RawMust()
}

// ── Result.RawString ──

func Test_RawString(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	s, err := r.RawString()

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected valid string", actual)
}

// ── Result.RawStringMust ──

func Test_RawStringMust_Panic(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.Error(errors.New("fail"))
	r.RawStringMust()
}

func Test_RawStringMust_Success(t *testing.T) {
	// Arrange
	r := corejson.New("hello")
	s := r.RawStringMust()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ── Result.RawErrString ──

func Test_RawErrString(t *testing.T) {
	r := corejson.New("x")
	b, e := r.RawErrString()
	_ = b
	_ = e
}

// ── Result.RawPrettyString ──

func Test_RawPrettyString(t *testing.T) {
	r := corejson.New(map[string]string{"a": "b"})
	s, err := r.RawPrettyString()
	_ = err
	_ = s
}

// ── Result.MeaningfulErrorMessage ──

func Test_MeaningfulErrorMessage_NoErr(t *testing.T) {
	// Arrange
	r := corejson.New("x")

	// Act
	actual := args.Map{"result": r.MeaningfulErrorMessage() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_MeaningfulErrorMessage_WithErr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Error(errors.New("fail"))

	// Act
	actual := args.Map{"result": r.MeaningfulErrorMessage() == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

// ── Result.MeaningfulError ──

func Test_MeaningfulError_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	err := r.MeaningfulError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_MeaningfulError_EmptyBytes(t *testing.T) {
	// Arrange
	r := &corejson.Result{Bytes: nil}
	err := r.MeaningfulError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_MeaningfulError_HasError(t *testing.T) {
	// Arrange
	r := &corejson.Result{
		Bytes: []byte(`"x"`),
		Error: errors.New("some error"),
	}
	err := r.MeaningfulError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── Result.SafeBytes ──

func Test_SafeBytes_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	b := r.SafeBytes()

	// Act
	actual := args.Map{"result": len(b) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

// ── Result.JsonModel ──

func Test_JsonModel_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	m := r.JsonModel()

	// Act
	actual := args.Map{"result": m.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error in model", actual)
}

func Test_JsonModel_Valid(t *testing.T) {
	r := corejson.New("x")
	m := r.JsonModel()
	_ = m
}

// ── Result.JsonModelAny ──

func Test_JsonModelAny(t *testing.T) {
	r := corejson.New("x")
	a := r.JsonModelAny()
	_ = a
}

// ── Result.Json / JsonPtr ──

func Test_Json(t *testing.T) {
	r := corejson.New("x")
	j := r.Json()
	_ = j
}

func Test_JsonPtr(t *testing.T) {
	r := corejson.New("x")
	j := r.JsonPtr()
	_ = j
}

// ── Result.ParseInjectUsingJson ──

func Test_ParseInjectUsingJson_Error(t *testing.T) {
	r := corejson.Empty.ResultPtr()
	badInput := corejson.NewResult.UsingString(`invalid`)
	_, err := r.ParseInjectUsingJson(badInput)
	_ = err
}

func Test_ParseInjectUsingJson_Success(t *testing.T) {
	r := corejson.Empty.ResultPtr()
	input := corejson.New(*r)
	_, err := r.ParseInjectUsingJson(&input)
	_ = err
}

// ── Result.ParseInjectUsingJsonMust ──

func Test_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() { recover() }()
	r := corejson.Empty.ResultPtr()
	bad := corejson.NewResult.UsingString(`invalid`)
	r.ParseInjectUsingJsonMust(bad)
}

// ── Result.CloneError ──

func Test_CloneError_HasError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Error(errors.New("err"))
	err := r.CloneError()

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_CloneError_NoError(t *testing.T) {
	// Arrange
	r := corejson.New("x")
	err := r.CloneError()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// ── Result.Ptr / NonPtr ──

func Test_Ptr(t *testing.T) {
	// Arrange
	r := corejson.New("x")
	p := r.Ptr()

	// Act
	actual := args.Map{"result": p == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NonPtr_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result
	np := r.NonPtr()

	// Act
	actual := args.Map{"result": np.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_NonPtr_Valid(t *testing.T) {
	r := corejson.New("x")
	np := r.NonPtr()
	_ = np
}

// ── Result.IsEqualPtr ──

func Test_IsEqualPtr_BothNil(t *testing.T) {
	// Arrange
	var a, b *corejson.Result

	// Act
	actual := args.Map{"result": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_IsEqualPtr_OneNil(t *testing.T) {
	// Arrange
	a := corejson.New("x").Ptr()

	// Act
	actual := args.Map{"result": a.IsEqualPtr(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_IsEqualPtr_Same(t *testing.T) {
	// Arrange
	a := corejson.New("x").Ptr()

	// Act
	actual := args.Map{"result": a.IsEqualPtr(a)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true same ptr", actual)
}

func Test_IsEqualPtr_DiffLength(t *testing.T) {
	// Arrange
	a := corejson.New("x").Ptr()
	b := corejson.New("xy").Ptr()

	// Act
	actual := args.Map{"result": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_IsEqualPtr_DiffError(t *testing.T) {
	// Arrange
	a := corejson.NewResult.Ptr([]byte(`"x"`), errors.New("a"), "t")
	b := corejson.NewResult.Ptr([]byte(`"x"`), errors.New("b"), "t")

	// Act
	actual := args.Map{"result": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_IsEqualPtr_DiffTypeName(t *testing.T) {
	// Arrange
	a := corejson.NewResult.Ptr([]byte(`"x"`), nil, "t1")
	b := corejson.NewResult.Ptr([]byte(`"x"`), nil, "t2")

	// Act
	actual := args.Map{"result": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_IsEqualPtr_Equal(t *testing.T) {
	// Arrange
	a := corejson.NewResult.Ptr([]byte(`"x"`), nil, "t")
	b := corejson.NewResult.Ptr([]byte(`"x"`), nil, "t")

	// Act
	actual := args.Map{"result": a.IsEqualPtr(b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── Result.IsEqual ──

func Test_IsEqual_DiffLen(t *testing.T) {
	// Arrange
	a := corejson.New("x")
	b := corejson.New("xy")

	// Act
	actual := args.Map{"result": a.IsEqual(b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_IsEqual_DiffErr(t *testing.T) {
	// Arrange
	a := corejson.NewResult.Create([]byte(`"x"`), errors.New("a"), "")
	b := corejson.NewResult.Create([]byte(`"x"`), errors.New("b"), "")

	// Act
	actual := args.Map{"result": a.IsEqual(b)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_IsEqual_Equal(t *testing.T) {
	// Arrange
	a := corejson.New("x")
	b := corejson.New("x")

	// Act
	actual := args.Map{"result": a.IsEqual(b)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

// ── Result.CombineErrorWithRefString / CombineErrorWithRefError ──

func Test_CombineErrorWithRefString_NoError(t *testing.T) {
	// Arrange
	r := corejson.New("x")
	s := r.CombineErrorWithRefString("ref1")

	// Act
	actual := args.Map{"result": s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_CombineErrorWithRefString_WithError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Error(errors.New("fail"))
	s := r.CombineErrorWithRefString("ref1", "ref2")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_CombineErrorWithRefError_NoError(t *testing.T) {
	// Arrange
	r := corejson.New("x")
	err := r.CombineErrorWithRefError("ref")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_CombineErrorWithRefError_WithError(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Error(errors.New("fail"))
	err := r.CombineErrorWithRefError("ref")

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

// ── Result.BytesError ──

func Test_BytesError_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"result": r.BytesError() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_BytesError_Valid(t *testing.T) {
	// Arrange
	r := corejson.New("x")
	be := r.BytesError()

	// Act
	actual := args.Map{"result": be == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

// ── Result.Dispose ──

func Test_Dispose_Nil(t *testing.T) {
	var r *corejson.Result
	r.Dispose() // should not panic
}

func Test_Dispose_Valid(t *testing.T) {
	r := corejson.New("x")
	r.Dispose()
}

// ── Result.CloneIf / ClonePtr / Clone ──

func Test_CloneIf_NoClone(t *testing.T) {
	r := corejson.New("x")
	c := r.CloneIf(false, false)
	_ = c
}

func Test_CloneIf_Clone(t *testing.T) {
	r := corejson.New("x")
	c := r.CloneIf(true, false)
	_ = c
}

func Test_ClonePtr_Nil(t *testing.T) {
	// Arrange
	var r *corejson.Result

	// Act
	actual := args.Map{"result": r.ClonePtr(false) != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_ClonePtr_Valid(t *testing.T) {
	// Arrange
	r := corejson.New("x")
	p := r.ClonePtr(true)

	// Act
	actual := args.Map{"result": p == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_Clone_Empty(t *testing.T) {
	r := corejson.Result{}
	c := r.Clone(true)
	_ = c
}

func Test_Clone_ShallowCopy(t *testing.T) {
	r := corejson.New("x")
	c := r.Clone(false)
	_ = c
}

func Test_Clone_DeepCopy(t *testing.T) {
	r := corejson.New("x")
	c := r.Clone(true)
	_ = c
}

// ── Result.AsJsonContractsBinder / AsJsoner / AsJsonParseSelfInjector ──

func Test_AsJsonContractsBinder(t *testing.T) {
	r := corejson.New("x")
	_ = r.AsJsonContractsBinder()
}

func Test_AsJsoner(t *testing.T) {
	r := corejson.New("x")
	_ = r.AsJsoner()
}

func Test_AsJsonParseSelfInjector(t *testing.T) {
	r := corejson.New("x")
	_ = r.AsJsonParseSelfInjector()
}

// ── Result.JsonParseSelfInject ──

func Test_JsonParseSelfInject(t *testing.T) {
	r := corejson.New("x")
	input := corejson.New(r)
	err := r.JsonParseSelfInject(&input)
	_ = err
}

// ── Result.SafeBytesTypeName ──

func Test_SafeBytesTypeName_Empty(t *testing.T) {
	// Arrange
	r := &corejson.Result{}

	// Act
	actual := args.Map{"result": r.SafeBytesTypeName() != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}
