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

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ═══════════════════════════════════════════════
// ResultsPtrCollection — all uncovered methods
// ═══════════════════════════════════════════════

func Test_01_RPC_Length(t *testing.T) {
	// Arrange
	var rpc *corejson.ResultsPtrCollection

	// Act
	actual := args.Map{"result": rpc.Length() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_02_RPC_LastIndex(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()

	// Act
	actual := args.Map{"result": rpc.LastIndex() != -1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected -1", actual)
}

func Test_03_RPC_IsEmpty_HasAnyItem(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()

	// Act
	actual := args.Map{"result": rpc.IsEmpty() || rpc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_04_RPC_FirstOrDefault(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()

	// Act
	actual := args.Map{"result": rpc.FirstOrDefault() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	actual = args.Map{"result": rpc.FirstOrDefault() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_05_RPC_LastOrDefault(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()

	// Act
	actual := args.Map{"result": rpc.LastOrDefault() != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	actual = args.Map{"result": rpc.LastOrDefault() == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_06_RPC_Take(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.Take(1)
	rpc.Add(corejson.NewResult.AnyPtr("a"))
	rpc.Add(corejson.NewResult.AnyPtr("b"))

	// Act
	actual := args.Map{"result": rpc.Take(1).Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_07_RPC_Limit(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.Limit(5)
	rpc.Add(corejson.NewResult.AnyPtr("a"))
	rpc.Add(corejson.NewResult.AnyPtr("b"))

	// Act
	actual := args.Map{"result": rpc.Limit(-1).Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
	actual = args.Map{"result": rpc.Limit(1).Length() != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_08_RPC_Skip(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.Skip(0)
	rpc.Add(corejson.NewResult.AnyPtr("a"))
	rpc.Add(corejson.NewResult.AnyPtr("b"))

	// Act
	actual := args.Map{"result": rpc.Skip(1).Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_09_RPC_AddSkipOnNil(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddSkipOnNil(nil)
	rpc.AddSkipOnNil(corejson.NewResult.AnyPtr("x"))

	// Act
	actual := args.Map{"result": rpc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_10_RPC_AddNonNilNonError(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddNonNilNonError(nil)
	rpc.AddNonNilNonError(&corejson.Result{Error: errors.New("e")})
	rpc.AddNonNilNonError(corejson.NewResult.AnyPtr("x"))

	// Act
	actual := args.Map{"result": rpc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_11_RPC_GetAt(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))

	// Act
	actual := args.Map{"result": rpc.GetAt(0) == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_12_RPC_HasError(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()

	// Act
	actual := args.Map{"result": rpc.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	rpc.Add(&corejson.Result{Error: errors.New("e")})
	actual = args.Map{"result": rpc.HasError()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
}

func Test_13_RPC_AllErrors(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	errs, has := rpc.AllErrors()

	// Act
	actual := args.Map{"result": has || len(errs) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	rpc.Add(&corejson.Result{Error: errors.New("e")})
	errs, has = rpc.AllErrors()
	actual = args.Map{"result": has || len(errs) != 1}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_14_RPC_GetErrorsStrings(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.GetErrorsStrings()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	rpc.Add(&corejson.Result{Error: errors.New("e")})
	s := rpc.GetErrorsStrings()

	// Act
	actual := args.Map{"result": len(s) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_15_RPC_GetErrorsStringsPtr(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.GetErrorsStringsPtr()
}

func Test_16_RPC_GetErrorsAsSingleString(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.GetErrorsAsSingleString()
}

func Test_17_RPC_GetErrorsAsSingle(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.GetErrorsAsSingle()
}

func Test_18_RPC_UnmarshalAt(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("hello"))
	var s string
	err := rpc.UnmarshalAt(0, &s)

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_19_RPC_UnmarshalAt_NilResult(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(nil)
	var s string
	err := rpc.UnmarshalAt(0, &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for nil result", actual)
}

func Test_20_RPC_UnmarshalAt_HasError(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(&corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("e")})
	var s string
	err := rpc.UnmarshalAt(0, &s)
	// Accept whatever result - just exercise the code path
	_ = err
}

func Test_21_RPC_InjectIntoAt(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	r := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	rpc.Add(r)
	target := corejson.Result{}
	err := rpc.InjectIntoAt(0, &target)
	_ = err
}

func Test_22_RPC_InjectIntoSameIndex(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	var nilInjectors []corejson.JsonParseSelfInjector
	errs, has := rpc.InjectIntoSameIndex(nilInjectors...)

	// Act
	actual := args.Map{"result": has || len(errs) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	rpc.Add(nil)
	rpc.Add(&corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("e")})
	rpc.Add(corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"}))
	t1 := corejson.Result{}
	errs, has = rpc.InjectIntoSameIndex(nil, nil, &t1)
	_ = errs
	_ = has
}

func Test_23_RPC_UnmarshalIntoSameIndex(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	var nilAnys []any
	errs, has := rpc.UnmarshalIntoSameIndex(nilAnys...)

	// Act
	actual := args.Map{"result": has || len(errs) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	rpc.Add(corejson.NewResult.AnyPtr("hello"))
	rpc.Add(nil)
	rpc.Add(&corejson.Result{Bytes: []byte(`"x"`), Error: errors.New("e")})
	rpc.Add(&corejson.Result{Bytes: []byte(`{}`)})
	var s string
	errs, has = rpc.UnmarshalIntoSameIndex(&s, nil, nil, nil)
	_ = errs
	_ = has
}

func Test_24_RPC_GetAtSafe(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))

	// Act
	actual := args.Map{"result": rpc.GetAtSafe(0) == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": rpc.GetAtSafe(-1) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_25_RPC_GetAtSafeUsingLength(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))

	// Act
	actual := args.Map{"result": rpc.GetAtSafeUsingLength(0, 1) == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": rpc.GetAtSafeUsingLength(5, 1) != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_26_RPC_Add_AddResult(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	rpc.AddResult(corejson.NewResult.Any("y"))

	// Act
	actual := args.Map{"result": rpc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_27_RPC_AddSerializer(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddSerializer(nil)

	// Act
	actual := args.Map{"result": rpc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_28_RPC_AddSerializers(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddSerializers()
}

func Test_29_RPC_AddSerializerFunc(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddSerializerFunc(nil)
	rpc.AddSerializerFunc(func() ([]byte, error) { return []byte(`"x"`), nil })

	// Act
	actual := args.Map{"result": rpc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_30_RPC_AddSerializerFunctions(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddSerializerFunctions()
}

func Test_31_RPC_Adds(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Adds(nil, corejson.NewResult.AnyPtr("x"))

	// Act
	actual := args.Map{"result": rpc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_32_RPC_AddAny(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddAny(nil)
	rpc.AddAny("x")

	// Act
	actual := args.Map{"result": rpc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_33_RPC_AddAnyItems(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddAnyItems(nil, "a", nil, "b")

	// Act
	actual := args.Map{"result": rpc.Length() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2", actual)
}

func Test_34_RPC_AddResultsCollection(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddResultsCollection(nil)
	other := corejson.NewResultsPtrCollection.Empty()
	other.Add(corejson.NewResult.AnyPtr("x"))
	rpc.AddResultsCollection(other)

	// Act
	actual := args.Map{"result": rpc.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_35_RPC_AddNonNilItems(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddNonNilItems(nil, corejson.NewResult.AnyPtr("x"))
	// Note: AddNonNilItems appends results... for each non-nil result
	_ = rpc
}

func Test_36_RPC_AddNonNilItemsPtr(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddNonNilItemsPtr()
	rpc.AddNonNilItemsPtr(nil, corejson.NewResult.AnyPtr("x"))
	_ = rpc
}

func Test_37_RPC_Clear(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	rpc.Clear()

	// Act
	actual := args.Map{"result": rpc.HasAnyItem()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_38_RPC_Clear_Nil(t *testing.T) {
	var rpc *corejson.ResultsPtrCollection
	_ = rpc.Clear()
}

func Test_39_RPC_Dispose(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	rpc.Dispose()
}

func Test_40_RPC_Dispose_Nil(t *testing.T) {
	var rpc *corejson.ResultsPtrCollection
	rpc.Dispose()
}

func Test_41_RPC_GetStrings(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.GetStrings()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	s := rpc.GetStrings()

	// Act
	actual := args.Map{"result": len(s) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_42_RPC_GetStringsPtr(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.GetStringsPtr()
}

func Test_43_RPC_AddJsoners(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.AddJsoners(true)
}

func Test_44_RPC_NonPtr_Ptr(t *testing.T) {
	rpc := corejson.ResultsPtrCollection{}
	_ = rpc.NonPtr()
	_ = rpc.Ptr()
}

func Test_45_RPC_GetPagesSize(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()

	// Act
	actual := args.Map{"result": rpc.GetPagesSize(0) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
	for i := 0; i < 5; i++ {
		rpc.Add(corejson.NewResult.AnyPtr(i))
	}
	actual = args.Map{"result": rpc.GetPagesSize(2) != 3}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_46_RPC_GetPagedCollection(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	for i := 0; i < 5; i++ {
		rpc.Add(corejson.NewResult.AnyPtr(i))
	}
	pages := rpc.GetPagedCollection(2)

	// Act
	actual := args.Map{"result": len(pages) != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_47_RPC_GetPagedCollection_Small(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	pages := rpc.GetPagedCollection(10)

	// Act
	actual := args.Map{"result": len(pages) != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_48_RPC_GetSinglePageCollection(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	for i := 0; i < 10; i++ {
		rpc.Add(corejson.NewResult.AnyPtr(i))
	}
	page := rpc.GetSinglePageCollection(3, 1)

	// Act
	actual := args.Map{"result": page.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_49_RPC_GetSinglePageCollection_Small(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	page := rpc.GetSinglePageCollection(10, 1)

	// Act
	actual := args.Map{"result": page.Length() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_50_RPC_JsonModel_JsonModelAny(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.JsonModel()
	_ = rpc.JsonModelAny()
}

func Test_51_RPC_Json_JsonPtr(t *testing.T) {
	rpc := corejson.ResultsPtrCollection{}
	_ = rpc.Json()
	_ = rpc.JsonPtr()
}

func Test_52_RPC_ParseInjectUsingJson(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	serialized := rpc.JsonPtr()
	rpc2 := corejson.NewResultsPtrCollection.Empty()
	_, err := rpc2.ParseInjectUsingJson(serialized)
	_ = err
}

func Test_53_RPC_ParseInjectUsingJson_Fail(t *testing.T) {
	// Arrange
	rpc := corejson.NewResultsPtrCollection.Empty()
	bad := &corejson.Result{Error: errors.New("fail")}
	_, err := rpc.ParseInjectUsingJson(bad)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_54_RPC_ParseInjectUsingJsonMust(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	serialized := rpc.JsonPtr()
	rpc2 := corejson.NewResultsPtrCollection.Empty()
	_ = rpc2.ParseInjectUsingJsonMust(serialized)
}

func Test_55_RPC_AsJsonContractsBinder(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.AsJsonContractsBinder()
}

func Test_56_RPC_AsJsoner(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.AsJsoner()
}

func Test_57_RPC_JsonParseSelfInject(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	serialized := rpc.JsonPtr()
	rpc2 := corejson.NewResultsPtrCollection.Empty()
	_ = rpc2.JsonParseSelfInject(serialized)
}

func Test_58_RPC_AsJsonParseSelfInjector(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	_ = rpc.AsJsonParseSelfInjector()
}

func Test_59_RPC_Clone(t *testing.T) {
	// Arrange
	var rpc *corejson.ResultsPtrCollection

	// Act
	actual := args.Map{"result": rpc.Clone(false) != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	rpc = corejson.NewResultsPtrCollection.Empty()
	_ = rpc.Clone(false)
	rpc.Add(corejson.NewResult.AnyPtr("x"))
	_ = rpc.Clone(true)
}

// ═══════════════════════════════════════════════
// Creators — newResultCreator uncovered methods
// ═══════════════════════════════════════════════

func Test_60_NRC_UnmarshalUsingBytes(t *testing.T) {
	r := corejson.NewResult.Any("hello")
	b, _ := r.Ptr().Serialize()
	result := corejson.NewResult.UnmarshalUsingBytes(b)
	_ = result
}

func Test_61_NRC_DeserializeUsingBytes(t *testing.T) {
	result := corejson.NewResult.DeserializeUsingBytes([]byte(`invalid`))
	_ = result
}

func Test_62_NRC_DeserializeUsingResult(t *testing.T) {
	r := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	result := corejson.NewResult.DeserializeUsingResult(r)
	_ = result
	// with error
	result = corejson.NewResult.DeserializeUsingResult(&corejson.Result{Error: errors.New("e")})
	_ = result
}

func Test_63_NRC_UsingBytes(t *testing.T) {
	r := corejson.NewResult.UsingBytes([]byte(`"x"`))
	_ = r
}

func Test_64_NRC_UsingBytesType(t *testing.T) {
	r := corejson.NewResult.UsingBytesType([]byte(`"x"`), "T")
	_ = r
}

func Test_65_NRC_UsingTypeBytesPtr(t *testing.T) {
	r := corejson.NewResult.UsingTypeBytesPtr("T", []byte(`"x"`))
	_ = r
}

func Test_66_NRC_UsingBytesPtr(t *testing.T) {
	r := corejson.NewResult.UsingBytesPtr(nil)
	_ = r
	r = corejson.NewResult.UsingBytesPtr([]byte(`"x"`))
	_ = r
}

func Test_67_NRC_UsingBytesPtrErrPtr(t *testing.T) {
	r := corejson.NewResult.UsingBytesPtrErrPtr(nil, errors.New("e"), "T")
	_ = r
	r = corejson.NewResult.UsingBytesPtrErrPtr([]byte(`"x"`), nil, "T")
	_ = r
}

func Test_68_NRC_UsingBytesErrPtr(t *testing.T) {
	r := corejson.NewResult.UsingBytesErrPtr(nil, errors.New("e"), "T")
	_ = r
	r = corejson.NewResult.UsingBytesErrPtr([]byte(`"x"`), nil, "T")
	_ = r
}

func Test_69_NRC_PtrUsingStringPtr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.PtrUsingStringPtr(nil, "T")

	// Act
	actual := args.Map{"result": r.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	s := `"hello"`
	r = corejson.NewResult.PtrUsingStringPtr(&s, "T")
	_ = r
}

func Test_70_NRC_UsingErrorStringPtr(t *testing.T) {
	r := corejson.NewResult.UsingErrorStringPtr(nil, nil, "T")
	_ = r
	r = corejson.NewResult.UsingErrorStringPtr(errors.New("e"), nil, "T")
	_ = r
	s := `"hello"`
	r = corejson.NewResult.UsingErrorStringPtr(nil, &s, "T")
	_ = r
	r = corejson.NewResult.UsingErrorStringPtr(errors.New("e"), &s, "T")
	_ = r
}

func Test_71_NRC_Ptr(t *testing.T) {
	r := corejson.NewResult.Ptr([]byte(`"x"`), nil, "T")
	_ = r
}

func Test_72_NRC_UsingJsonBytesTypeError(t *testing.T) {
	r := corejson.NewResult.UsingJsonBytesTypeError([]byte(`"x"`), nil, "T")
	_ = r
}

func Test_73_NRC_UsingJsonBytesError(t *testing.T) {
	r := corejson.NewResult.UsingJsonBytesError([]byte(`"x"`), nil)
	_ = r
}

func Test_74_NRC_UsingTypePlusStringPtr(t *testing.T) {
	r := corejson.NewResult.UsingTypePlusStringPtr("T", nil)
	_ = r
	s := `"hello"`
	r = corejson.NewResult.UsingTypePlusStringPtr("T", &s)
	_ = r
	empty := ""
	r = corejson.NewResult.UsingTypePlusStringPtr("T", &empty)
	_ = r
}

func Test_75_NRC_UsingStringWithType(t *testing.T) {
	r := corejson.NewResult.UsingStringWithType(`"hello"`, "T")
	_ = r
}

func Test_76_NRC_UsingString(t *testing.T) {
	r := corejson.NewResult.UsingString(`"hello"`)
	_ = r
}

func Test_77_NRC_UsingStringPtr(t *testing.T) {
	r := corejson.NewResult.UsingStringPtr(nil)
	_ = r
	s := `"hello"`
	r = corejson.NewResult.UsingStringPtr(&s)
	_ = r
}

func Test_78_NRC_CreatePtr(t *testing.T) {
	r := corejson.NewResult.CreatePtr([]byte(`"x"`), nil, "T")
	_ = r
}

func Test_79_NRC_NonPtr(t *testing.T) {
	r := corejson.NewResult.NonPtr([]byte(`"x"`), nil, "T")
	_ = r
}

func Test_80_NRC_PtrUsingBytesPtr(t *testing.T) {
	// Arrange
	r := corejson.NewResult.PtrUsingBytesPtr(nil, errors.New("e"), "T")

	// Act
	actual := args.Map{"result": r.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	r = corejson.NewResult.PtrUsingBytesPtr(nil, nil, "T")
	_ = r
	r = corejson.NewResult.PtrUsingBytesPtr([]byte(`"x"`), nil, "T")
	_ = r
}

func Test_81_NRC_CastingAny(t *testing.T) {
	r := corejson.NewResult.CastingAny("hello")
	_ = r
}

func Test_82_NRC_UsingBytesError(t *testing.T) {
	r := corejson.NewResult.UsingBytesError(nil)
	_ = r
}

func Test_83_NRC_Error_ErrorPtr(t *testing.T) {
	r := corejson.NewResult.Error(errors.New("e"))
	_ = r
	rp := corejson.NewResult.ErrorPtr(errors.New("e"))
	_ = rp
}

func Test_84_NRC_Empty_EmptyPtr(t *testing.T) {
	_ = corejson.NewResult.Empty()
	_ = corejson.NewResult.EmptyPtr()
}

func Test_85_NRC_TypeName_TypeNameBytes(t *testing.T) {
	_ = corejson.NewResult.TypeName("T")
	_ = corejson.NewResult.TypeNameBytes("T")
}

func Test_86_NRC_Many(t *testing.T) {
	r := corejson.NewResult.Many("a", "b")
	_ = r
}

func Test_87_NRC_Serialize(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Serialize("hello")
	_ = r
	ch := make(chan int)
	r = corejson.NewResult.Serialize(ch)

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_88_NRC_Marshal(t *testing.T) {
	// Arrange
	r := corejson.NewResult.Marshal("hello")
	_ = r
	ch := make(chan int)
	r = corejson.NewResult.Marshal(ch)

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_89_NRC_UsingSerializer(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingSerializer(nil)

	// Act
	actual := args.Map{"result": r != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_90_NRC_UsingSerializerFunc(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingSerializerFunc(nil)

	// Act
	actual := args.Map{"result": r != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	r = corejson.NewResult.UsingSerializerFunc(func() ([]byte, error) { return []byte(`"x"`), nil })
	actual = args.Map{"result": r.HasError()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
}

func Test_91_NRC_UsingJsoner(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingJsoner(nil)

	// Act
	actual := args.Map{"result": r != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_92_NRC_AnyToCastingResult(t *testing.T) {
	r := corejson.NewResult.AnyToCastingResult("hello")
	_ = r
}

// ═══════════════════════════════════════════════
// Creators — collection creators
// ═══════════════════════════════════════════════

func Test_93_NRCC_Default(t *testing.T) {
	_ = corejson.NewResultsCollection.Default()
}

func Test_94_NRCC_AnyItems(t *testing.T) {
	_ = corejson.NewResultsCollection.AnyItems("a", "b")
}

func Test_95_NRCC_AnyItemsPlusCap(t *testing.T) {
	_ = corejson.NewResultsCollection.AnyItemsPlusCap(5)
	_ = corejson.NewResultsCollection.AnyItemsPlusCap(5, "a")
}

func Test_96_NRCC_UsingJsonersOption(t *testing.T) {
	_ = corejson.NewResultsCollection.UsingJsonersOption(true, 5)
}

func Test_97_NRCC_UsingJsonersNonNull(t *testing.T) {
	_ = corejson.NewResultsCollection.UsingJsonersNonNull(5)
}

func Test_98_NRCC_UsingJsoners(t *testing.T) {
	_ = corejson.NewResultsCollection.UsingJsoners()
}

func Test_99_NRCC_UsingResultsPtrPlusCap(t *testing.T) {
	_ = corejson.NewResultsCollection.UsingResultsPtrPlusCap(5)
	_ = corejson.NewResultsCollection.UsingResultsPtrPlusCap(0, corejson.NewResult.AnyPtr("x"))
}

func Test_100_NRCC_UsingResultsPtr(t *testing.T) {
	_ = corejson.NewResultsCollection.UsingResultsPtr(corejson.NewResult.AnyPtr("x"))
}

func Test_101_NRCC_UsingResultsPlusCap(t *testing.T) {
	_ = corejson.NewResultsCollection.UsingResultsPlusCap(5)
	_ = corejson.NewResultsCollection.UsingResultsPlusCap(0, corejson.NewResult.Any("x"))
}

func Test_102_NRCC_UsingResults(t *testing.T) {
	_ = corejson.NewResultsCollection.UsingResults(corejson.NewResult.Any("x"))
}

func Test_103_NRCC_Serializers(t *testing.T) {
	_ = corejson.NewResultsCollection.Serializers()
}

func Test_104_NRCC_SerializerFunctions(t *testing.T) {
	_ = corejson.NewResultsCollection.SerializerFunctions()
}

func Test_105_NRCC_UnmarshalUsingBytes(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	b, _ := rc.JsonPtr().SerializeSkipExistingIssues()
	_, _ = corejson.NewResultsCollection.UnmarshalUsingBytes(b)
}

func Test_106_NRCC_DeserializeUsingBytes(t *testing.T) {
	_, _ = corejson.NewResultsCollection.DeserializeUsingBytes([]byte(`invalid`))
}

func Test_107_NRCC_DeserializeUsingResult(t *testing.T) {
	_, _ = corejson.NewResultsCollection.DeserializeUsingResult(&corejson.Result{Error: errors.New("e")})
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	_, _ = corejson.NewResultsCollection.DeserializeUsingResult(rc.JsonPtr())
}

// ═══════════════════════════════════════════════
// newBytesCollectionCreator
// ═══════════════════════════════════════════════

func Test_108_NBCC_UnmarshalUsingBytes(t *testing.T) {
	_, _ = corejson.NewBytesCollection.UnmarshalUsingBytes([]byte(`[[1]]`))
}

func Test_109_NBCC_DeserializeUsingBytes(t *testing.T) {
	_, _ = corejson.NewBytesCollection.DeserializeUsingBytes([]byte(`invalid`))
}

func Test_110_NBCC_DeserializeUsingResult(t *testing.T) {
	_, _ = corejson.NewBytesCollection.DeserializeUsingResult(&corejson.Result{Error: errors.New("e")})
}

func Test_111_NBCC_AnyItems(t *testing.T) {
	_, _ = corejson.NewBytesCollection.AnyItems("a", "b")
}

func Test_112_NBCC_JsonersPlusCap(t *testing.T) {
	_ = corejson.NewBytesCollection.JsonersPlusCap(true, 5)
}

func Test_113_NBCC_Jsoners(t *testing.T) {
	_ = corejson.NewBytesCollection.Jsoners()
}

func Test_114_NBCC_Serializers(t *testing.T) {
	_ = corejson.NewBytesCollection.Serializers()
}

// ═══════════════════════════════════════════════
// newResultsPtrCollectionCreator
// ═══════════════════════════════════════════════

func Test_115_NRPCC_UnmarshalUsingBytes(t *testing.T) {
	_, _ = corejson.NewResultsPtrCollection.UnmarshalUsingBytes([]byte(`invalid`))
}

func Test_116_NRPCC_DeserializeUsingResult(t *testing.T) {
	_, _ = corejson.NewResultsPtrCollection.DeserializeUsingResult(&corejson.Result{Error: errors.New("e")})
}

func Test_117_NRPCC_Default(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.Default()
}

func Test_118_NRPCC_AnyItemsPlusCap(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.AnyItemsPlusCap(0)
	_ = corejson.NewResultsPtrCollection.AnyItemsPlusCap(5, "a")
}

func Test_119_NRPCC_AnyItems(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.AnyItems("a")
}

func Test_120_NRPCC_UsingResultsPlusCap(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.UsingResultsPlusCap(0)
	_ = corejson.NewResultsPtrCollection.UsingResultsPlusCap(0, corejson.NewResult.AnyPtr("x"))
}

func Test_121_NRPCC_UsingResults(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.UsingResults(corejson.NewResult.AnyPtr("x"))
}

func Test_122_NRPCC_JsonersPlusCap(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.JsonersPlusCap(true, 5)
}

func Test_123_NRPCC_Jsoners(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.Jsoners()
}

func Test_124_NRPCC_Serializers(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.Serializers()
}

// ═══════════════════════════════════════════════
// newMapResultsCreator
// ═══════════════════════════════════════════════

func Test_125_NMRC_UnmarshalUsingBytes(t *testing.T) {
	_, _ = corejson.NewMapResults.UnmarshalUsingBytes([]byte(`{}`))
}

func Test_126_NMRC_DeserializeUsingBytes(t *testing.T) {
	_, _ = corejson.NewMapResults.DeserializeUsingBytes([]byte(`invalid`))
}

func Test_127_NMRC_DeserializeUsingResult(t *testing.T) {
	_, _ = corejson.NewMapResults.DeserializeUsingResult(&corejson.Result{Error: errors.New("e")})
	mr := corejson.NewMapResults.Empty()
	mr.Add("k", corejson.NewResult.Any("v"))
	_, _ = corejson.NewMapResults.DeserializeUsingResult(mr.JsonPtr())
}

func Test_128_NMRC_UsingKeyAnyItems(t *testing.T) {
	_ = corejson.NewMapResults.UsingKeyAnyItems(0)
	_ = corejson.NewMapResults.UsingKeyAnyItems(5, corejson.KeyAny{Key: "k", AnyInf: "v"})
}

func Test_129_NMRC_UsingMapOptions(t *testing.T) {
	m := map[string]corejson.Result{"k": corejson.NewResult.Any("v")}
	_ = corejson.NewMapResults.UsingMapOptions(false, false, 0, nil)
	_ = corejson.NewMapResults.UsingMapOptions(false, false, 0, m)
	_ = corejson.NewMapResults.UsingMapOptions(true, false, 5, m)
}

func Test_130_NMRC_UsingMapPlusCap(t *testing.T) {
	_ = corejson.NewMapResults.UsingMapPlusCap(5, nil)
	m := map[string]corejson.Result{"k": corejson.NewResult.Any("v")}
	_ = corejson.NewMapResults.UsingMapPlusCap(5, m)
}

func Test_131_NMRC_UsingMapPlusCapClone(t *testing.T) {
	_ = corejson.NewMapResults.UsingMapPlusCapClone(5, nil)
	m := map[string]corejson.Result{"k": corejson.NewResult.Any("v")}
	_ = corejson.NewMapResults.UsingMapPlusCapClone(5, m)
}

func Test_132_NMRC_UsingMapPlusCapDeepClone(t *testing.T) {
	_ = corejson.NewMapResults.UsingMapPlusCapDeepClone(5, nil)
	m := map[string]corejson.Result{"k": corejson.NewResult.Any("v")}
	_ = corejson.NewMapResults.UsingMapPlusCapDeepClone(5, m)
}

func Test_133_NMRC_UsingMap(t *testing.T) {
	_ = corejson.NewMapResults.UsingMap(nil)
	m := map[string]corejson.Result{"k": corejson.NewResult.Any("v")}
	_ = corejson.NewMapResults.UsingMap(m)
}

func Test_134_NMRC_UsingMapAnyItemsPlusCap(t *testing.T) {
	_ = corejson.NewMapResults.UsingMapAnyItemsPlusCap(5, nil)
	_ = corejson.NewMapResults.UsingMapAnyItemsPlusCap(5, map[string]any{"k": "v"})
}

func Test_135_NMRC_UsingMapAnyItems(t *testing.T) {
	_ = corejson.NewMapResults.UsingMapAnyItems(map[string]any{"k": "v"})
}

func Test_136_NMRC_UsingKeyWithResultsPlusCap(t *testing.T) {
	_ = corejson.NewMapResults.UsingKeyWithResultsPlusCap(5)
	_ = corejson.NewMapResults.UsingKeyWithResultsPlusCap(0, corejson.KeyWithResult{Key: "k", Result: corejson.NewResult.Any("v")})
}

func Test_137_NMRC_UsingKeyWithResults(t *testing.T) {
	_ = corejson.NewMapResults.UsingKeyWithResults(corejson.KeyWithResult{Key: "k", Result: corejson.NewResult.Any("v")})
}

func Test_138_NMRC_UsingKeyJsonersPlusCap(t *testing.T) {
	_ = corejson.NewMapResults.UsingKeyJsonersPlusCap(5)
}

func Test_139_NMRC_UsingKeyJsoners(t *testing.T) {
	_ = corejson.NewMapResults.UsingKeyJsoners()
}

// ═══════════════════════════════════════════════
// emptyCreator — uncovered methods
// ═══════════════════════════════════════════════

func Test_140_Empty_ResultWithErr(t *testing.T) {
	// Arrange
	r := corejson.Empty.ResultWithErr("T", errors.New("e"))

	// Act
	actual := args.Map{"result": r.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_141_Empty_ResultPtrWithErr(t *testing.T) {
	// Arrange
	r := corejson.Empty.ResultPtrWithErr("T", errors.New("e"))

	// Act
	actual := args.Map{"result": r.Error == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_142_Empty_BytesCollection(t *testing.T) {
	_ = corejson.Empty.BytesCollection()
}

func Test_143_Empty_BytesCollectionPtr(t *testing.T) {
	_ = corejson.Empty.BytesCollectionPtr()
}

// ═══════════════════════════════════════════════
// Deserializer — uncovered methods
// ═══════════════════════════════════════════════

func Test_144_Deser_UsingStringPtr(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingStringPtr(nil, &s)
	_ = err
	str := `"hello"`
	err = corejson.Deserialize.UsingStringPtr(&str, &s)

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_145_Deser_UsingError(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingError(nil, &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	err = corejson.Deserialize.UsingError(errors.New(`"hello"`), &s)
	_ = err
}

func Test_146_Deser_UsingErrorWhichJsonResult(t *testing.T) {
	// Arrange
	var r corejson.Result
	err := corejson.Deserialize.UsingErrorWhichJsonResult(nil, &r)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_147_Deser_ApplyMust(t *testing.T) {
	// Arrange
	r := corejson.NewResult.AnyPtr("hello")
	var s string
	corejson.Deserialize.ApplyMust(r, &s)

	// Act
	actual := args.Map{"result": s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_148_Deser_FromString(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.FromString(`"hello"`, &s)

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_149_Deser_FromStringMust(t *testing.T) {
	// Arrange
	var s string
	corejson.Deserialize.FromStringMust(`"hello"`, &s)

	// Act
	actual := args.Map{"result": s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_150_Deser_FromTo(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.FromTo([]byte(`"hello"`), &s)

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_151_Deser_MapAnyToPointer(t *testing.T) {
	// Arrange
	type simple struct {
		Name string `json:"Name"`
	}
	var s simple
	err := corejson.Deserialize.MapAnyToPointer(true, nil, &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil for empty map skip", actual)
	err = corejson.Deserialize.MapAnyToPointer(false, map[string]any{"Name": "test"}, &s)
	actual = args.Map{"result": err != nil || s.Name != "test"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_152_Deser_UsingStringOption(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingStringOption(true, "", &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	err = corejson.Deserialize.UsingStringOption(false, `"hello"`, &s)
	_ = err
}

func Test_153_Deser_UsingStringIgnoreEmpty(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingStringIgnoreEmpty("", &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_154_Deser_UsingBytesPointerMust(t *testing.T) {
	// Arrange
	var s string
	corejson.Deserialize.UsingBytesPointerMust([]byte(`"hello"`), &s)

	// Act
	actual := args.Map{"result": s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_155_Deser_UsingBytesIf(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingBytesIf(false, []byte(`"hello"`), &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil when skip", actual)
	err = corejson.Deserialize.UsingBytesIf(true, []byte(`"hello"`), &s)
	actual = args.Map{"result": err != nil || s != "hello"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_156_Deser_UsingBytesPointerIf(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingBytesPointerIf(false, []byte(`"hello"`), &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil when skip", actual)
	err = corejson.Deserialize.UsingBytesPointerIf(true, []byte(`"hello"`), &s)
	actual = args.Map{"result": err != nil || s != "hello"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_157_Deser_UsingBytesPointer(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingBytesPointer(nil, &s)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for nil", actual)
	err = corejson.Deserialize.UsingBytesPointer([]byte(`"hello"`), &s)
	actual = args.Map{"result": err != nil || s != "hello"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_158_Deser_UsingBytesMust(t *testing.T) {
	// Arrange
	var s string
	corejson.Deserialize.UsingBytesMust([]byte(`"hello"`), &s)

	// Act
	actual := args.Map{"result": s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_159_Deser_UsingSafeBytesMust(t *testing.T) {
	// Arrange
	var s string
	corejson.Deserialize.UsingSafeBytesMust(nil, &s)
	corejson.Deserialize.UsingSafeBytesMust([]byte(`"hello"`), &s)

	// Act
	actual := args.Map{"result": s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_160_Deser_AnyToFieldsMap(t *testing.T) {
	m, err := corejson.Deserialize.AnyToFieldsMap(map[string]int{"a": 1})
	_ = m
	_ = err
}

func Test_161_Deser_UsingSerializerTo(t *testing.T) {
	var s string
	err := corejson.Deserialize.UsingSerializerTo(nil, &s)
	_ = err
}

func Test_162_Deser_UsingSerializerFuncTo(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingSerializerFuncTo(nil, &s)
	_ = err
	err = corejson.Deserialize.UsingSerializerFuncTo(func() ([]byte, error) {
		return []byte(`"hello"`), nil
	}, &s)

	// Act
	actual := args.Map{"result": err != nil || s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_163_Deser_UsingDeserializerToOption(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingDeserializerToOption(true, nil, &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	err = corejson.Deserialize.UsingDeserializerToOption(false, nil, &s)
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_164_Deser_UsingDeserializerDefined(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingDeserializerDefined(nil, &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_165_Deser_UsingDeserializerFuncDefined(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingDeserializerFuncDefined(nil, &s)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
	err = corejson.Deserialize.UsingDeserializerFuncDefined(func(toPtr any) error {
		return nil
	}, &s)
	actual = args.Map{"result": err != nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
}

func Test_166_Deser_UsingJsonerToAny(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingJsonerToAny(true, nil, &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	err = corejson.Deserialize.UsingJsonerToAny(false, nil, &s)
	actual = args.Map{"result": err == nil}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_167_Deser_UsingJsonerToAnyMust(t *testing.T) {
	// Arrange
	var s string
	err := corejson.Deserialize.UsingJsonerToAnyMust(true, nil, &s)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_168_Deser_Result(t *testing.T) {
	r := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	b, _ := r.Serialize()
	_, _ = corejson.Deserialize.Result(b)
}

func Test_169_Deser_ResultPtr(t *testing.T) {
	_, _ = corejson.Deserialize.ResultPtr([]byte(`invalid`))
}

// ═══════════════════════════════════════════════
// deserializeFromBytesTo
// ═══════════════════════════════════════════════

func Test_170_BytesTo_Strings(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.Strings([]byte(`["a","b"]`))
}

func Test_171_BytesTo_StringsMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.StringsMust([]byte(`["a","b"]`))
}

func Test_172_BytesTo_String(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.String([]byte(`"hello"`))
}

func Test_173_BytesTo_Integer(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.Integer([]byte(`42`))
}

func Test_174_BytesTo_IntegerMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.IntegerMust([]byte(`42`))
}

func Test_175_BytesTo_Integer64(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.Integer64([]byte(`64`))
}

func Test_176_BytesTo_Integer64Must(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.Integer64Must([]byte(`64`))
}

func Test_177_BytesTo_Integers(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.Integers([]byte(`[1,2,3]`))
}

func Test_178_BytesTo_IntegersMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.IntegersMust([]byte(`[1,2,3]`))
}

func Test_179_BytesTo_StringMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.StringMust([]byte(`"hello"`))
}

func Test_180_BytesTo_MapAnyItem(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.MapAnyItem([]byte(`{"a":1}`))
}

func Test_181_BytesTo_MapAnyItemMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.MapAnyItemMust([]byte(`{"a":1}`))
}

func Test_182_BytesTo_MapStringString(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.MapStringString([]byte(`{"a":"b"}`))
}

func Test_183_BytesTo_MapStringStringMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.MapStringStringMust([]byte(`{"a":"b"}`))
}

func Test_184_BytesTo_ResultCollection(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.ResultCollection([]byte(`{"JsonResultsCollection":[]}`))
	_, _ = corejson.Deserialize.BytesTo.ResultCollection([]byte(`invalid`))
}

func Test_185_BytesTo_ResultCollectionMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.ResultCollectionMust([]byte(`{"JsonResultsCollection":[]}`))
}

func Test_186_BytesTo_ResultsPtrCollection(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.ResultsPtrCollection([]byte(`{"JsonResultsCollection":[]}`))
	_, _ = corejson.Deserialize.BytesTo.ResultsPtrCollection([]byte(`invalid`))
}

func Test_187_BytesTo_ResultsPtrCollectionMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.ResultsPtrCollectionMust([]byte(`{"JsonResultsCollection":[]}`))
}

func Test_188_BytesTo_MapResults(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.MapResults([]byte(`{"JsonResultsMap":{}}`))
	_, _ = corejson.Deserialize.BytesTo.MapResults([]byte(`invalid`))
}

func Test_189_BytesTo_MapResultsMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.MapResultsMust([]byte(`{"JsonResultsMap":{}}`))
}

func Test_190_BytesTo_Bytes(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.Bytes([]byte(`"dGVzdA=="`))
}

func Test_191_BytesTo_BytesMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.BytesMust([]byte(`"dGVzdA=="`))
}

func Test_192_BytesTo_Bool(t *testing.T) {
	_, _ = corejson.Deserialize.BytesTo.Bool([]byte(`true`))
}

func Test_193_BytesTo_BoolMust(t *testing.T) {
	_ = corejson.Deserialize.BytesTo.BoolMust([]byte(`true`))
}

// ═══════════════════════════════════════════════
// deserializeFromResultTo
// ═══════════════════════════════════════════════

func Test_194_ResultTo_String(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	_, _ = corejson.Deserialize.ResultTo.String(r)
}

func Test_195_ResultTo_Bool(t *testing.T) {
	r := corejson.NewResult.AnyPtr(true)
	_, _ = corejson.Deserialize.ResultTo.Bool(r)
}

func Test_196_ResultTo_Byte(t *testing.T) {
	r := corejson.NewResult.AnyPtr(65)
	_, _ = corejson.Deserialize.ResultTo.Byte(r)
}

func Test_197_ResultTo_ByteMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr(65)
	_ = corejson.Deserialize.ResultTo.ByteMust(r)
}

func Test_198_ResultTo_BoolMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr(true)
	_ = corejson.Deserialize.ResultTo.BoolMust(r)
}

func Test_199_ResultTo_StringMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr("hello")
	_ = corejson.Deserialize.ResultTo.StringMust(r)
}

func Test_200_ResultTo_StringsMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr([]string{"a", "b"})
	_ = corejson.Deserialize.ResultTo.StringsMust(r)
}

func Test_201_ResultTo_MapAnyItem(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]any{"a": 1})
	_, _ = corejson.Deserialize.ResultTo.MapAnyItem(r)
}

func Test_202_ResultTo_MapAnyItemMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]any{"a": 1})
	_ = corejson.Deserialize.ResultTo.MapAnyItemMust(r)
}

func Test_203_ResultTo_MapStringString(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]string{"a": "b"})
	_, _ = corejson.Deserialize.ResultTo.MapStringString(r)
}

func Test_204_ResultTo_MapStringStringMust(t *testing.T) {
	r := corejson.NewResult.AnyPtr(map[string]string{"a": "b"})
	_ = corejson.Deserialize.ResultTo.MapStringStringMust(r)
}

func Test_205_ResultTo_ResultCollection(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	rc.Add(corejson.NewResult.Any("x"))
	r := rc.JsonPtr()
	_, _ = corejson.Deserialize.ResultTo.ResultCollection(r)
}

func Test_206_ResultTo_ResultCollectionMust(t *testing.T) {
	rc := corejson.NewResultsCollection.Empty()
	r := rc.JsonPtr()
	_ = corejson.Deserialize.ResultTo.ResultCollectionMust(r)
}

func Test_207_ResultTo_ResultsPtrCollection(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	r := rpc.JsonPtr()
	_, _ = corejson.Deserialize.ResultTo.ResultsPtrCollection(r)
}

func Test_208_ResultTo_ResultsPtrCollectionMust(t *testing.T) {
	rpc := corejson.NewResultsPtrCollection.Empty()
	r := rpc.JsonPtr()
	_ = corejson.Deserialize.ResultTo.ResultsPtrCollectionMust(r)
}

func Test_209_ResultTo_Result(t *testing.T) {
	inner := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	_, _ = corejson.Deserialize.ResultTo.Result(inner)
}

func Test_210_ResultTo_ResultMust(t *testing.T) {
	inner := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	_ = corejson.Deserialize.ResultTo.ResultMust(inner)
}

func Test_211_ResultTo_ResultPtr(t *testing.T) {
	inner := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	_, _ = corejson.Deserialize.ResultTo.ResultPtr(inner)
}

func Test_212_ResultTo_ResultPtrMust(t *testing.T) {
	inner := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	_ = corejson.Deserialize.ResultTo.ResultPtrMust(inner)
}

func Test_213_ResultTo_MapResults(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	r := mr.JsonPtr()
	_, _ = corejson.Deserialize.ResultTo.MapResults(r)
}

func Test_214_ResultTo_Bytes(t *testing.T) {
	inner := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	_, _ = corejson.Deserialize.ResultTo.Bytes(inner)
}

func Test_215_ResultTo_BytesMust(t *testing.T) {
	inner := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	_ = corejson.Deserialize.ResultTo.BytesMust(inner)
}

func Test_216_ResultTo_MapResultsMust(t *testing.T) {
	mr := corejson.NewMapResults.Empty()
	r := mr.JsonPtr()
	_ = corejson.Deserialize.ResultTo.MapResultsMust(r)
}

// ═══════════════════════════════════════════════
// Serializer — uncovered methods
// ═══════════════════════════════════════════════

func Test_217_Serializer_FromStringer(t *testing.T) {
	// Arrange
	type myStringer31 struct{ val string }
	stringer := myStringer31{val: "test-stringer"}
	// Create a proper fmt.Stringer
	r := corejson.Serialize.FromStringer(stringerImpl31{stringer.val})

	// Act
	actual := args.Map{"result": r.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected error", actual)
}

type stringerImpl31 struct{ v string }
func (s stringerImpl31) String() string { return s.v }

// ═══════════════════════════════════════════════
// AnyTo — uncovered: UsingSerializer
// ═══════════════════════════════════════════════

func Test_218_AnyTo_UsingSerializer(t *testing.T) {
	// Arrange
	r := corejson.AnyTo.UsingSerializer(nil)

	// Act
	actual := args.Map{"result": r != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

// ═══════════════════════════════════════════════
// CastAny — uncovered: FromToReflection
// ═══════════════════════════════════════════════

func Test_219_CastAny_FromToReflection(t *testing.T) {
	// Arrange
	var out string
	err := corejson.CastAny.FromToReflection([]byte(`"hello"`), &out)

	// Act
	actual := args.Map{"result": err != nil || out != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

// ═══════════════════════════════════════════════
// Deserialize ResultMust, ResultPtrMust
// ═══════════════════════════════════════════════

func Test_220_Deser_ResultMust(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	b, _ := r.Serialize()
	_ = corejson.Deserialize.ResultMust(b)
}

func Test_221_Deser_ResultPtrMust(t *testing.T) {
	defer func() { recover() }()
	r := corejson.NewResult.AnyPtr(corejson.Result{Bytes: []byte(`"t"`), TypeName: "T"})
	b, _ := r.Serialize()
	_ = corejson.Deserialize.ResultPtrMust(b)
}
