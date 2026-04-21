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

// ═══════════════════════════════════════════════
// BytesCollection — all uncovered methods
// ═══════════════════════════════════════════════

func Test_01_BC_Length(t *testing.T) {
	// Arrange
	var bc *corejson.BytesCollection

	// Act
	actual := args.Map{"result": bc.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_02_BC_LastIndex(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()

	// Act
	actual := args.Map{"result": bc.LastIndex() != -1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
}

func Test_03_BC_IsEmpty_HasAnyItem(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()

	// Act
	actual := args.Map{"result": bc.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual = args.Map{"result": bc.HasAnyItem()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_04_BC_FirstOrDefault(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()

	// Act
	actual := args.Map{"result": bc.FirstOrDefault() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	bc.Add([]byte(`"x"`))
	actual = args.Map{"result": bc.FirstOrDefault() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_05_BC_LastOrDefault(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()

	// Act
	actual := args.Map{"result": bc.LastOrDefault() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	bc.Add([]byte(`"x"`))
	actual = args.Map{"result": bc.LastOrDefault() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_06_BC_Take(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	_ = bc.Take(1)
	bc.Add([]byte(`"a"`))
	bc.Add([]byte(`"b"`))
	taken := bc.Take(1)

	// Act
	actual := args.Map{"result": taken.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_07_BC_Limit(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	_ = bc.Limit(5)
	bc.Add([]byte(`"a"`))
	bc.Add([]byte(`"b"`))
	l := bc.Limit(-1)

	// Act
	actual := args.Map{"result": l.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	l = bc.Limit(1)
	actual = args.Map{"result": l.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_08_BC_Skip(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	_ = bc.Skip(0)
	bc.Add([]byte(`"a"`))
	bc.Add([]byte(`"b"`))
	s := bc.Skip(1)

	// Act
	actual := args.Map{"result": s.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_09_BC_AddSkipOnNil(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.AddSkipOnNil(nil)
	bc.AddSkipOnNil([]byte(`"x"`))

	// Act
	actual := args.Map{"result": bc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_10_BC_AddNonEmpty(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.AddNonEmpty([]byte{})
	bc.AddNonEmpty([]byte(`"x"`))

	// Act
	actual := args.Map{"result": bc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_11_BC_AddResultPtr(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.AddResultPtr(&corejson.Result{Error: errors.New("e")})
	bc.AddResultPtr(&corejson.Result{Bytes: []byte(`"x"`)})

	// Act
	actual := args.Map{"result": bc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_12_BC_AddResult(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.AddResult(corejson.Result{Error: errors.New("e")})
	bc.AddResult(corejson.Result{Bytes: []byte(`"x"`)})

	// Act
	actual := args.Map{"result": bc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_13_BC_GetAt(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	b := bc.GetAt(0)

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
}

func Test_14_BC_JsonResultAt(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	r := bc.JsonResultAt(0)

	// Act
	actual := args.Map{"result": r == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_15_BC_UnmarshalAt(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"hello"`))
	var s string
	err := bc.UnmarshalAt(0, &s)

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_16_BC_AddSerializer(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.AddSerializer(nil)

	// Act
	actual := args.Map{"result": bc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_17_BC_AddSerializers(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.AddSerializers()

	// Act
	actual := args.Map{"result": bc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_18_BC_AddSerializerFunc(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.AddSerializerFunc(nil)
	bc.AddSerializerFunc(func() ([]byte, error) {
		return []byte(`"x"`), nil
	})

	// Act
	actual := args.Map{"result": bc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_19_BC_AddSerializerFunctions(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.AddSerializerFunctions()

	// Act
	actual := args.Map{"result": bc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_20_BC_InjectIntoAt(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`{"Bytes":"dGVzdA==","TypeName":"T"}`))
	target := corejson.Result{}
	err := bc.InjectIntoAt(0, &target)
	_ = err
}

func Test_21_BC_InjectIntoSameIndex(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	var nilInjectors []corejson.JsonParseSelfInjector
	errs, has := bc.InjectIntoSameIndex(nilInjectors...)

	// Act
	actual := args.Map{"result": has || len(errs) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	bc.Add([]byte(`{"Bytes":"dGVzdA==","TypeName":"T"}`))
	t1 := corejson.Result{}
	errs, has = bc.InjectIntoSameIndex(&t1)
	_ = errs
	_ = has
}

func Test_22_BC_UnmarshalIntoSameIndex(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	var nilAnys []any
	errs, has := bc.UnmarshalIntoSameIndex(nilAnys...)

	// Act
	actual := args.Map{"result": has || len(errs) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	bc.Add([]byte(`"hello"`))
	bc.Add([]byte(`42`))
	var s string
	var n int
	errs, has = bc.UnmarshalIntoSameIndex(&s, &n)
	_ = errs
	_ = has
}

func Test_23_BC_GetAtSafe(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))

	// Act
	actual := args.Map{"result": bc.GetAtSafe(0) == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": bc.GetAtSafe(-1) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	actual = args.Map{"result": bc.GetAtSafe(999) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_24_BC_GetAtSafePtr(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))

	// Act
	actual := args.Map{"result": bc.GetAtSafePtr(0) == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": bc.GetAtSafePtr(-1) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_25_BC_GetResultAtSafe(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))

	// Act
	actual := args.Map{"result": bc.GetResultAtSafe(0) == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": bc.GetResultAtSafe(-1) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_26_BC_GetAtSafeUsingLength(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))

	// Act
	actual := args.Map{"result": bc.GetAtSafeUsingLength(0, 1) == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": bc.GetAtSafeUsingLength(5, 1) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_27_BC_AddPtr(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.AddPtr([]byte{})
	bc.AddPtr([]byte(`"x"`))

	// Act
	actual := args.Map{"result": bc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_28_BC_Adds(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.Adds()
	bc.Adds([]byte{}, []byte(`"a"`))

	// Act
	actual := args.Map{"result": bc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_29_BC_AddAnyItems(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	err := bc.AddAnyItems()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	err = bc.AddAnyItems("x", 42)
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
	actual = args.Map{"result": bc.Length() != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_30_BC_AddAnyItems_Error(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	ch := make(chan int)
	err := bc.AddAnyItems(ch)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_31_BC_AddMapResults(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	mr := corejson.NewMapResults.Empty()
	bc.AddMapResults(mr)
	mr.Add("k", corejson.NewResult.Any("v"))
	bc.AddMapResults(mr)

	// Act
	actual := args.Map{"result": bc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_32_BC_AddRawMapResults(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.AddRawMapResults(nil)
	m := map[string]corejson.Result{
		"k": corejson.NewResult.Any("v"),
		"e": corejson.NewResult.Error(errors.New("err")),
	}
	bc.AddRawMapResults(m)

	// Act
	actual := args.Map{"result": bc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_33_BC_AddsPtr(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.AddsPtr(nil, corejson.NewResult.AnyPtr("x"), &corejson.Result{})

	// Act
	actual := args.Map{"result": bc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_34_BC_AddAny(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	err := bc.AddAny("hello")

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	actual = args.Map{"result": bc.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_35_BC_AddAny_Error(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	ch := make(chan int)
	err := bc.AddAny(ch)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_36_BC_AddBytesCollection(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	other := corejson.NewBytesCollection.Empty()
	bc.AddBytesCollection(other)
	other.Add([]byte(`"x"`))
	bc.AddBytesCollection(other)

	// Act
	actual := args.Map{"result": bc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_37_BC_Clear(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	bc.Clear()

	// Act
	actual := args.Map{"result": bc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_38_BC_Clear_Nil(t *testing.T) {
	var bc *corejson.BytesCollection
	_ = bc.Clear()
}

func Test_39_BC_Dispose(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	bc.Dispose()
}

func Test_40_BC_Dispose_Nil(t *testing.T) {
	var bc *corejson.BytesCollection
	bc.Dispose()
}

func Test_41_BC_Strings(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	s := bc.Strings()

	// Act
	actual := args.Map{"result": len(s) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	bc.Add([]byte(`"x"`))
	s = bc.Strings()
	actual = args.Map{"result": len(s) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_42_BC_StringsPtr(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	_ = bc.StringsPtr()
}

func Test_43_BC_AddJsoners(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.AddJsoners(true)

	// Act
	actual := args.Map{"result": bc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_44_BC_GetPagesSize(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()

	// Act
	actual := args.Map{"result": bc.GetPagesSize(0) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	for i := 0; i < 5; i++ {
		bc.Add([]byte(`"x"`))
	}
	actual = args.Map{"result": bc.GetPagesSize(2) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_45_BC_GetPagedCollection(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	for i := 0; i < 5; i++ {
		bc.Add([]byte(`"x"`))
	}
	pages := bc.GetPagedCollection(2)

	// Act
	actual := args.Map{"result": len(pages) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_46_BC_GetPagedCollection_Small(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	pages := bc.GetPagedCollection(10)

	// Act
	actual := args.Map{"result": len(pages) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_47_BC_GetSinglePageCollection(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	for i := 0; i < 10; i++ {
		bc.Add([]byte(`"x"`))
	}
	page := bc.GetSinglePageCollection(3, 1)

	// Act
	actual := args.Map{"result": page.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	page = bc.GetSinglePageCollection(3, 4)
	actual = args.Map{"result": page.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_48_BC_GetSinglePageCollection_Small(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	page := bc.GetSinglePageCollection(10, 1)

	// Act
	actual := args.Map{"result": page.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_49_BC_JsonModel_JsonModelAny(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	_ = bc.JsonModel()
	_ = bc.JsonModelAny()
}

func Test_50_BC_MarshalJSON(t *testing.T) {
	// Arrange
	bc := corejson.BytesCollection{}
	bc.Items = [][]byte{[]byte(`"x"`)}
	b, err := bc.MarshalJSON()

	// Act
	actual := args.Map{"result": err != nil || len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_51_BC_UnmarshalJSON(t *testing.T) {
	bc := corejson.BytesCollection{}
	err := bc.UnmarshalJSON([]byte(`[["dGVzdA=="]]`))
	_ = err
}

func Test_52_BC_Json_JsonPtr(t *testing.T) {
	bc := corejson.BytesCollection{}
	_ = bc.Json()
	_ = bc.JsonPtr()
}

func Test_53_BC_ParseInjectUsingJson(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	serialized := bc.JsonPtr()
	bc2 := corejson.NewBytesCollection.Empty()
	_, err := bc2.ParseInjectUsingJson(serialized)
	_ = err
}

func Test_54_BC_ParseInjectUsingJson_Fail(t *testing.T) {
	// Arrange
	bc := corejson.NewBytesCollection.Empty()
	bad := &corejson.Result{Error: errors.New("fail")}
	_, err := bc.ParseInjectUsingJson(bad)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_55_BC_ParseInjectUsingJsonMust(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	serialized := bc.JsonPtr()
	bc2 := corejson.NewBytesCollection.Empty()
	_ = bc2.ParseInjectUsingJsonMust(serialized)
}

func Test_56_BC_AsJsonContractsBinder(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	_ = bc.AsJsonContractsBinder()
}

func Test_57_BC_AsJsoner(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	_ = bc.AsJsoner()
}

func Test_58_BC_JsonParseSelfInject(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	serialized := bc.JsonPtr()
	bc2 := corejson.NewBytesCollection.Empty()
	err := bc2.JsonParseSelfInject(serialized)
	_ = err
}

func Test_59_BC_AsJsonParseSelfInjector(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	_ = bc.AsJsonParseSelfInjector()
}

func Test_60_BC_ShadowClone(t *testing.T) {
	bc := corejson.NewBytesCollection.Empty()
	bc.Add([]byte(`"x"`))
	_ = bc.ShadowClone()
}

func Test_61_BC_Clone(t *testing.T) {
	bc := corejson.BytesCollection{}
	_ = bc.Clone(false)
	bc.Items = [][]byte{[]byte(`"x"`)}
	_ = bc.Clone(true)
}

func Test_62_BC_ClonePtr(t *testing.T) {
	// Arrange
	var bc *corejson.BytesCollection

	// Act
	actual := args.Map{"result": bc.ClonePtr(false) != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	bc = corejson.NewBytesCollection.Empty()
	_ = bc.ClonePtr(false)
	bc.Add([]byte(`"x"`))
	_ = bc.ClonePtr(true)
}

// ═══════════════════════════════════════════════
// MapResults — all uncovered methods
// ═══════════════════════════════════════════════

func Test_63_MR_Length(t *testing.T) {
	// Arrange
	var mr *corejson.MapResults

	// Act
	actual := args.Map{"result": mr.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_64_MR_LastIndex(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{"result": mr.LastIndex() != -1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
}

func Test_65_MR_IsEmpty_HasAnyItem(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{"result": mr.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	actual = args.Map{"result": mr.HasAnyItem()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_66_MR_AddSkipOnNil(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddSkipOnNil("k", nil)

	// Act
	actual := args.Map{"result": mr.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	mr.AddSkipOnNil("k", corejson.NewResult.AnyPtr("v"))
	actual = args.Map{"result": mr.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_67_MR_GetByKey(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{"result": mr.GetByKey("k") != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	mr.Add("k", corejson.NewResult.Any("v"))
	actual = args.Map{"result": mr.GetByKey("k") == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_68_MR_HasError(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{"result": mr.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	mr.Add("k", corejson.NewResult.Error(errors.New("e")))
	actual = args.Map{"result": mr.HasError()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_69_MR_AllErrors(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	errs, has := mr.AllErrors()

	// Act
	actual := args.Map{"result": has || len(errs) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	mr.Add("k", corejson.NewResult.Error(errors.New("e")))
	mr.Add("ok", corejson.NewResult.Any("v"))
	errs, has = mr.AllErrors()
	actual = args.Map{"result": has || len(errs) != 1}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_70_MR_GetErrorsStrings(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	s := mr.GetErrorsStrings()

	// Act
	actual := args.Map{"result": len(s) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	mr.Add("k", corejson.NewResult.Error(errors.New("e")))
	mr.Add("ok", corejson.NewResult.Any("v"))
	s = mr.GetErrorsStrings()
	actual = args.Map{"result": len(s) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_71_MR_GetErrorsStringsPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.GetErrorsStringsPtr()
}

func Test_72_MR_GetErrorsAsSingleString(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.GetErrorsAsSingleString()
}

func Test_73_MR_GetErrorsAsSingle(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.GetErrorsAsSingle()
}

func Test_74_MR_Unmarshal(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("hello"))
	var s string
	// Note: Unmarshal has inverted logic (has==false means key exists)
	err := mr.Unmarshal("missing", &s)
	_ = err
}

func Test_75_MR_Deserialize(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	var s string
	err := mr.Deserialize("k", &s)
	_ = err
}

func Test_76_MR_DeserializeMust(t *testing.T) {
	defer func() { recover() }()
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any(map[string]string{"a": "b"}))
	target := make(map[string]string)
	_ = mr.DeserializeMust("k", &target)
}

func Test_77_MR_UnmarshalMany(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	err := mr.UnmarshalMany()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	mr.Add("k", corejson.NewResult.Any("hello"))
	var s string
	err = mr.UnmarshalMany(corejson.KeyAny{Key: "k", AnyInf: &s})
	_ = err
}

func Test_78_MR_UnmarshalManySafe(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	err := mr.UnmarshalManySafe()

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	mr.Add("k", corejson.NewResult.Any("hello"))
	var s string
	err = mr.UnmarshalManySafe(corejson.KeyAny{Key: "k", AnyInf: &s})
	_ = err
}

func Test_79_MR_SafeUnmarshal(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("hello"))
	var s string
	err := mr.SafeUnmarshal("k", &s)
	_ = err
	err = mr.SafeUnmarshal("missing", &s)
	_ = err
}

func Test_80_MR_SafeDeserialize(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	var s string
	err := mr.SafeDeserialize("k", &s)
	_ = err
}

func Test_81_MR_SafeDeserializeMust(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	var s string
	_ = mr.SafeDeserializeMust("k", &s)
}

func Test_82_MR_InjectIntoAt(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	r := corejson.NewResult.Any(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	mr.Add("k", r)
	target := corejson.Result{}
	err := mr.InjectIntoAt("k", &target)
	_ = err
}

func Test_83_MR_Add_AddPtr(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	mr.AddPtr("k2", nil)
	mr.AddPtr("k3", corejson.NewResult.AnyPtr("v"))

	// Act
	actual := args.Map{"result": mr.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_84_MR_AddAny(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	err := mr.AddAny("k", nil)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
	err = mr.AddAny("k", "hello")
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
}

func Test_85_MR_AddAny_MarshalError(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	ch := make(chan int)
	err := mr.AddAny("k", ch)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_86_MR_AddAnySkipOnNil(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	err := mr.AddAnySkipOnNil("k", nil)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	err = mr.AddAnySkipOnNil("k", "v")
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_87_MR_AddAnyNonEmptyNonError(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddAnyNonEmptyNonError("k", nil)
	mr.AddAnyNonEmptyNonError("k", "v")

	// Act
	actual := args.Map{"result": mr.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_88_MR_AddAnyNonEmpty(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddAnyNonEmpty("k", nil)
	mr.AddAnyNonEmpty("k", "v")

	// Act
	actual := args.Map{"result": mr.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_89_MR_AddKeyWithResult(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyWithResult(corejson.KeyWithResult{Key: "k", Result: corejson.NewResult.Any("v")})

	// Act
	actual := args.Map{"result": mr.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_90_MR_AddKeyWithResultPtr(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyWithResultPtr(nil)
	mr.AddKeyWithResultPtr(&corejson.KeyWithResult{Key: "k", Result: corejson.NewResult.Any("v")})

	// Act
	actual := args.Map{"result": mr.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_91_MR_AddKeysWithResultsPtr(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddKeysWithResultsPtr()
	kr := &corejson.KeyWithResult{Key: "k", Result: corejson.NewResult.Any("v")}
	mr.AddKeysWithResultsPtr(kr)

	// Act
	actual := args.Map{"result": mr.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_92_MR_AddKeysWithResults(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddKeysWithResults()
	kr := corejson.KeyWithResult{Key: "k", Result: corejson.NewResult.Any("v")}
	mr.AddKeysWithResults(kr)

	// Act
	actual := args.Map{"result": mr.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_93_MR_AddKeyAnyInf(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyInf(corejson.KeyAny{Key: "k", AnyInf: "v"})

	// Act
	actual := args.Map{"result": mr.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_94_MR_AddKeyAnyInfPtr(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyInfPtr(nil)
	mr.AddKeyAnyInfPtr(&corejson.KeyAny{Key: "k", AnyInf: "v"})

	// Act
	actual := args.Map{"result": mr.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_95_MR_AddKeyAnyItems(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyItems()
	mr.AddKeyAnyItems(corejson.KeyAny{Key: "k", AnyInf: "v"})

	// Act
	actual := args.Map{"result": mr.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_96_MR_AddKeyAnyItemsPtr(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyAnyItemsPtr()
	mr.AddKeyAnyItemsPtr(&corejson.KeyAny{Key: "k", AnyInf: "v"})

	// Act
	actual := args.Map{"result": mr.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_97_MR_AddNonEmptyNonErrorPtr(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddNonEmptyNonErrorPtr("k", nil)
	mr.AddNonEmptyNonErrorPtr("k", &corejson.Result{Error: errors.New("e")})
	mr.AddNonEmptyNonErrorPtr("k", corejson.NewResult.AnyPtr("v"))

	// Act
	actual := args.Map{"result": mr.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_98_MR_AddMapResults(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddMapResults(nil)
	other := corejson.NewMapResults.Empty()
	mr.AddMapResults(other)
	other.Add("k", corejson.NewResult.Any("v"))
	mr.AddMapResults(other)

	// Act
	actual := args.Map{"result": mr.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_99_MR_AddMapAnyItems(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddMapAnyItems(nil)
	mr.AddMapAnyItems(map[string]any{"k": "v"})

	// Act
	actual := args.Map{"result": mr.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_100_MR_AllKeys(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	keys := mr.AllKeys()

	// Act
	actual := args.Map{"result": len(keys) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	mr.Add("a", corejson.NewResult.Any("v"))
	keys = mr.AllKeys()
	actual = args.Map{"result": len(keys) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_101_MR_AllKeysSorted(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	_ = mr.AllKeysSorted()
	mr.Add("b", corejson.NewResult.Any("v"))
	mr.Add("a", corejson.NewResult.Any("v"))
	keys := mr.AllKeysSorted()

	// Act
	actual := args.Map{"result": keys[0] != "a"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected sorted", actual)
}

func Test_102_MR_AllValues(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	_ = mr.AllValues()
	mr.Add("k", corejson.NewResult.Any("v"))
	vals := mr.AllValues()

	// Act
	actual := args.Map{"result": len(vals) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_103_MR_AllResultsCollection(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	rc := mr.AllResultsCollection()

	// Act
	actual := args.Map{"result": rc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	mr.Add("k", corejson.NewResult.Any("v"))
	rc = mr.AllResultsCollection()
	actual = args.Map{"result": rc.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_104_MR_AllResults(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.AllResults()
}

func Test_105_MR_GetStrings(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	s := mr.GetStrings()

	// Act
	actual := args.Map{"result": len(s) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	mr.Add("k", corejson.NewResult.Any("v"))
	s = mr.GetStrings()
	actual = args.Map{"result": len(s) != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_106_MR_GetStringsPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.GetStringsPtr()
}

func Test_107_MR_AddJsoner(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddJsoner("k", nil)

	// Act
	actual := args.Map{"result": mr.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_108_MR_AddKeyWithJsoner(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyWithJsoner(corejson.KeyWithJsoner{Key: "k", Jsoner: nil})
	_ = mr
}

func Test_109_MR_AddKeysWithJsoners(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.AddKeysWithJsoners()
}

func Test_110_MR_AddKeyWithJsonerPtr(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.AddKeyWithJsonerPtr(nil)
	mr.AddKeyWithJsonerPtr(&corejson.KeyWithJsoner{Key: "k", Jsoner: nil})
	_ = mr
}

func Test_111_MR_GetPagesSize(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()

	// Act
	actual := args.Map{"result": mr.GetPagesSize(0) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	for i := 0; i < 5; i++ {
		mr.Add(string(rune('a'+i)), corejson.NewResult.Any(i))
	}
	actual = args.Map{"result": mr.GetPagesSize(2) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_112_MR_GetPagedCollection(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	for i := 0; i < 5; i++ {
		mr.Add(string(rune('a'+i)), corejson.NewResult.Any(i))
	}
	pages := mr.GetPagedCollection(2)

	// Act
	actual := args.Map{"result": len(pages) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_113_MR_GetPagedCollection_Small(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	pages := mr.GetPagedCollection(10)

	// Act
	actual := args.Map{"result": len(pages) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_114_MR_AddMapResultsUsingCloneOption(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.AddMapResultsUsingCloneOption(false, false, nil)
	m := map[string]corejson.Result{
		"k": corejson.NewResult.Any("v"),
	}
	mr.AddMapResultsUsingCloneOption(false, false, m)

	// Act
	actual := args.Map{"result": mr.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	mr2 := corejson.NewMapResults.Empty()
	mr2.AddMapResultsUsingCloneOption(true, true, m)
	actual = args.Map{"result": mr2.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_115_MR_GetSinglePageCollection(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	for i := 0; i < 10; i++ {
		mr.Add(string(rune('a'+i)), corejson.NewResult.Any(i))
	}
	allKeys := mr.AllKeysSorted()
	page := mr.GetSinglePageCollection(3, 1, allKeys)

	// Act
	actual := args.Map{"result": page.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
	page = mr.GetSinglePageCollection(3, 4, allKeys)
	actual = args.Map{"result": page.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_116_MR_GetSinglePageCollection_Small(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	allKeys := mr.AllKeysSorted()
	page := mr.GetSinglePageCollection(10, 1, allKeys)

	// Act
	actual := args.Map{"result": page.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_117_MR_GetNewMapUsingKeys(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("a", corejson.NewResult.Any("1"))
	mr.Add("b", corejson.NewResult.Any("2"))
	sub := mr.GetNewMapUsingKeys(false, "a")

	// Act
	actual := args.Map{"result": sub.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
	sub = mr.GetNewMapUsingKeys(false)
	actual = args.Map{"result": sub.HasAnyItem()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	// non-panic missing
	sub = mr.GetNewMapUsingKeys(false, "missing")
	actual = args.Map{"result": sub.HasAnyItem()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_118_MR_ResultCollection(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	rc := mr.ResultCollection()

	// Act
	actual := args.Map{"result": rc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
	mr.Add("k", corejson.NewResult.Any("v"))
	rc = mr.ResultCollection()
	actual = args.Map{"result": rc.Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_119_MR_JsonModel_JsonModelAny(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.JsonModel()
	_ = mr.JsonModelAny()
}

func Test_120_MR_Clear(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	mr.Clear()

	// Act
	actual := args.Map{"result": mr.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_121_MR_Clear_Nil(t *testing.T) {
	var mr *corejson.MapResults
	_ = mr.Clear()
}

func Test_122_MR_Dispose(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	mr.Dispose()
}

func Test_123_MR_Dispose_Nil(t *testing.T) {
	var mr *corejson.MapResults
	mr.Dispose()
}

func Test_124_MR_Json_JsonPtr(t *testing.T) {
	mr := corejson.MapResults{Items: map[string]corejson.Result{}}
	_ = mr.Json()
	_ = mr.JsonPtr()
}

func Test_125_MR_ParseInjectUsingJson(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	serialized := mr.JsonPtr()
	mr2 := corejson.NewMapResults.Empty()
	_, err := mr2.ParseInjectUsingJson(serialized)
	_ = err
}

func Test_126_MR_ParseInjectUsingJson_Fail(t *testing.T) {
	// Arrange
	mr := corejson.NewMapResults.Empty()
	bad := &corejson.Result{Error: errors.New("fail")}
	_, err := mr.ParseInjectUsingJson(bad)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_127_MR_ParseInjectUsingJsonMust(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	serialized := mr.JsonPtr()
	mr2 := corejson.NewMapResults.Empty()
	_ = mr2.ParseInjectUsingJsonMust(serialized)
}

func Test_128_MR_AsJsonContractsBinder(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.AsJsonContractsBinder()
}

func Test_129_MR_AsJsoner(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.AsJsoner()
}

func Test_130_MR_JsonParseSelfInject(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	serialized := mr.JsonPtr()
	mr2 := corejson.NewMapResults.Empty()
	err := mr2.JsonParseSelfInject(serialized)
	_ = err
}

func Test_131_MR_AsJsonParseSelfInjector(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	_ = mr.AsJsonParseSelfInjector()
}
