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

// ── Migrated from Creators_test.go ──

func Test_EmptyCreator(t *testing.T) {
	_ = corejson.Empty.Result()
	_ = corejson.Empty.ResultPtr()
	_ = corejson.Empty.ResultWithErr("t", errors.New("e"))
	_ = corejson.Empty.ResultPtrWithErr("t", errors.New("e"))
	_ = corejson.Empty.BytesCollection()
	_ = corejson.Empty.BytesCollectionPtr()
	_ = corejson.Empty.ResultsCollection()
	_ = corejson.Empty.ResultsPtrCollection()
	_ = corejson.Empty.MapResults()
}

func Test_NewResultCreator(t *testing.T) {
	_ = corejson.NewResult.UsingBytes([]byte(`"x"`))
	_ = corejson.NewResult.UsingBytesType([]byte(`"x"`), "string")
	_ = corejson.NewResult.UsingBytesTypePtr([]byte(`"x"`), "string")
	_ = corejson.NewResult.UsingTypeBytesPtr("string", []byte(`"x"`))
	_ = corejson.NewResult.UsingBytesPtr([]byte(`"x"`))
	_ = corejson.NewResult.UsingBytesPtr(nil)
	_ = corejson.NewResult.UsingBytesPtrErrPtr(nil, errors.New("e"), "t")
	_ = corejson.NewResult.UsingBytesPtrErrPtr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.UsingBytesErrPtr([]byte{}, nil, "t")
	_ = corejson.NewResult.UsingBytesErrPtr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.Ptr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.UsingJsonBytesTypeError([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.UsingJsonBytesError([]byte(`"x"`), nil)
	_ = corejson.NewResult.UsingTypePlusString("t", `"x"`)
	_ = corejson.NewResult.UsingStringWithType(`"x"`, "t")
	_ = corejson.NewResult.UsingString(`"x"`)
	_ = corejson.NewResult.CreatePtr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.NonPtr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.Create([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.PtrUsingBytesPtr(nil, errors.New("e"), "t")
	_ = corejson.NewResult.PtrUsingBytesPtr(nil, nil, "t")
	_ = corejson.NewResult.PtrUsingBytesPtr([]byte(`"x"`), nil, "t")
	_ = corejson.NewResult.Error(errors.New("e"))
	_ = corejson.NewResult.ErrorPtr(errors.New("e"))
	_ = corejson.NewResult.Empty()
	_ = corejson.NewResult.EmptyPtr()
	_ = corejson.NewResult.TypeName("t")
	_ = corejson.NewResult.TypeNameBytes("t")
	_ = corejson.NewResult.Many("a", "b")
	_ = corejson.NewResult.Any("x")
	_ = corejson.NewResult.AnyPtr("x")
	_ = corejson.NewResult.Serialize("x")
	_ = corejson.NewResult.Marshal("x")
	_ = corejson.NewResult.CastingAny("x")
	_ = corejson.NewResult.AnyToCastingResult("x")
}

func Test_NewResultCreator_UsingStringPtr(t *testing.T) {
	s := `"hello"`
	_ = corejson.NewResult.UsingStringPtr(&s)
	_ = corejson.NewResult.UsingStringPtr(nil)
	empty := ""
	_ = corejson.NewResult.UsingStringPtr(&empty)
}

func Test_NewResultCreator_UsingTypePlusStringPtr(t *testing.T) {
	s := `"x"`
	_ = corejson.NewResult.UsingTypePlusStringPtr("t", &s)
	_ = corejson.NewResult.UsingTypePlusStringPtr("t", nil)
}

func Test_NewResultCreator_PtrUsingStringPtr(t *testing.T) {
	s := `"x"`
	_ = corejson.NewResult.PtrUsingStringPtr(&s, "t")
	_ = corejson.NewResult.PtrUsingStringPtr(nil, "t")
}

func Test_NewResultCreator_UsingErrorStringPtr(t *testing.T) {
	s := `"x"`
	_ = corejson.NewResult.UsingErrorStringPtr(nil, &s, "t")
	_ = corejson.NewResult.UsingErrorStringPtr(errors.New("e"), &s, "t")
	_ = corejson.NewResult.UsingErrorStringPtr(errors.New("e"), nil, "t")
}

func Test_NewResultCreator_UsingSerializer(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingSerializer(nil)

	// Act
	actual := args.Map{"result": r != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_NewResultCreator_UsingSerializerFunc(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingSerializerFunc(nil)

	// Act
	actual := args.Map{"result": r != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
	r2 := corejson.NewResult.UsingSerializerFunc(func() ([]byte, error) {
		return []byte(`"x"`), nil
	})
	actual = args.Map{"result": r2 == nil || r2.HasError()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewResultCreator_UsingJsoner(t *testing.T) {
	// Arrange
	r := corejson.NewResult.UsingJsoner(nil)

	// Act
	actual := args.Map{"result": r != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected nil", actual)
}

func Test_NewResultCreator_Deserialize(t *testing.T) {
	orig := corejson.NewResult.Any("hello")
	b, _ := orig.Serialize()
	_ = corejson.NewResult.DeserializeUsingBytes(b)
}

func Test_NewResultCreator_DeserializeUsingResult(t *testing.T) {
	orig := corejson.NewResult.Any("hello")
	b, _ := orig.Serialize()
	jr := corejson.NewResult.UsingBytes(b)
	_ = corejson.NewResult.DeserializeUsingResult(jr.Ptr())
}

func Test_NewResultsCollectionCreator(t *testing.T) {
	_ = corejson.NewResultsCollection.Empty()
	_ = corejson.NewResultsCollection.Default()
	_ = corejson.NewResultsCollection.UsingCap(5)
	_ = corejson.NewResultsCollection.AnyItems("a", "b")
	_ = corejson.NewResultsCollection.AnyItemsPlusCap(2, "a")
	_ = corejson.NewResultsCollection.AnyItemsPlusCap(2)
	_ = corejson.NewResultsCollection.UsingResults(corejson.NewResult.Any("a"))
	_ = corejson.NewResultsCollection.UsingResultsPtr(corejson.NewResult.AnyPtr("a"))
	_ = corejson.NewResultsCollection.UsingResultsPlusCap(2, corejson.NewResult.Any("a"))
	_ = corejson.NewResultsCollection.UsingResultsPtrPlusCap(2, corejson.NewResult.AnyPtr("a"))
}

func Test_NewResultsPtrCollectionCreator(t *testing.T) {
	_ = corejson.NewResultsPtrCollection.Empty()
	_ = corejson.NewResultsPtrCollection.Default()
	_ = corejson.NewResultsPtrCollection.UsingCap(5)
	_ = corejson.NewResultsPtrCollection.AnyItems("a", "b")
	_ = corejson.NewResultsPtrCollection.AnyItemsPlusCap(2, "a")
	_ = corejson.NewResultsPtrCollection.UsingResults(corejson.NewResult.AnyPtr("a"))
	_ = corejson.NewResultsPtrCollection.UsingResultsPlusCap(2, corejson.NewResult.AnyPtr("a"))
}

func Test_NewBytesCollectionCreator(t *testing.T) {
	_ = corejson.NewBytesCollection.Empty()
	_ = corejson.NewBytesCollection.UsingCap(5)
	_, _ = corejson.NewBytesCollection.AnyItems("a", "b")
}

func Test_NewMapResultsCreator(t *testing.T) {
	_ = corejson.NewMapResults.Empty()
	_ = corejson.NewMapResults.UsingCap(5)
	_ = corejson.NewMapResults.UsingMap(map[string]corejson.Result{})
	_ = corejson.NewMapResults.UsingMap(map[string]corejson.Result{"k": corejson.NewResult.Any("v")})
	_ = corejson.NewMapResults.UsingMapPlusCap(2, map[string]corejson.Result{})
	_ = corejson.NewMapResults.UsingMapPlusCap(2, map[string]corejson.Result{"k": corejson.NewResult.Any("v")})
	_ = corejson.NewMapResults.UsingMapPlusCapClone(2, map[string]corejson.Result{"k": corejson.NewResult.Any("v")})
	_ = corejson.NewMapResults.UsingMapPlusCapDeepClone(2, map[string]corejson.Result{"k": corejson.NewResult.Any("v")})
	_ = corejson.NewMapResults.UsingMapAnyItems(map[string]any{"k": "v"})
	_ = corejson.NewMapResults.UsingMapAnyItemsPlusCap(2, map[string]any{})
	_ = corejson.NewMapResults.UsingKeyWithResults(corejson.KeyWithResult{Key: "k", Result: corejson.NewResult.Any("v")})
	_ = corejson.NewMapResults.UsingKeyWithResultsPlusCap(2, corejson.KeyWithResult{Key: "k", Result: corejson.NewResult.Any("v")})
	_ = corejson.NewMapResults.UsingKeyAnyItems(2, corejson.KeyAny{Key: "k", AnyInf: "v"})
}

func Test_BytesCloneIf_Func(t *testing.T) {
	_ = corejson.BytesCloneIf(true, []byte("hello"))
	_ = corejson.BytesCloneIf(false, []byte("hello"))
	_ = corejson.BytesCloneIf(true, []byte{})
}

func Test_BytesDeepClone_Func(t *testing.T) {
	// Arrange
	b := corejson.BytesDeepClone([]byte("hello"))

	// Act
	actual := args.Map{"result": len(b) == 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected bytes", actual)
	b2 := corejson.BytesDeepClone([]byte{})
	actual = args.Map{"result": len(b2) != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_BytesToString_Func(t *testing.T) {
	// Arrange
	s := corejson.BytesToString([]byte("hello"))

	// Act
	actual := args.Map{"result": s != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
	s2 := corejson.BytesToString([]byte{})
	actual = args.Map{"result": s2 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_BytesToPrettyString_Func(t *testing.T) {
	// Arrange
	s := corejson.BytesToPrettyString([]byte(`{"a":1}`))

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	s2 := corejson.BytesToPrettyString([]byte{})
	actual = args.Map{"result": s2 != ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty", actual)
}

func Test_JsonString_Func(t *testing.T) {
	// Arrange
	s, err := corejson.JsonString("hello")

	// Act
	actual := args.Map{"result": err != nil || s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "unexpected", actual)
}

func Test_JsonStringOrErrMsg_Func(t *testing.T) {
	// Arrange
	s := corejson.JsonStringOrErrMsg("hello")

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	ch := make(chan int)
	s2 := corejson.JsonStringOrErrMsg(ch)
	actual = args.Map{"result": s2 == ""}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error message", actual)
}

func Test_StaticJsonError(t *testing.T) {
	// Act
	actual := args.Map{"result": corejson.StaticJsonError == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
}

func Test_NewResult_Various(t *testing.T) {
	_ = corejson.NewResult.UnmarshalUsingBytes([]byte(`{}`))
	_ = corejson.NewResult.DeserializeUsingBytes([]byte(`{}`))
	_ = corejson.NewResult.CastingAny("hello")
	_ = corejson.NewResult.AnyToCastingResult("hello")
}
