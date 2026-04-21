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

package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/corestr"
	"github.com/alimtvnetwork/core-v8/coretests/args")

// ══════════════════════════════════════════════════════════════════════════════
// Hashset — Segment 5b
// ══════════════════════════════════════════════════════════════════════════════

func Test_HashsetCollection_IsEmpty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_IsEmpty", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 0)

		// Act
		actual := args.Map{
			"empty": hsc.IsEmpty(),
			"hasItems": hsc.HasItems(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"hasItems": false,
		}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection IsEmpty -- true", actual)
	})
}

func Test_HashsetCollection_Add_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_Add", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)

		// Act
		actual := args.Map{"len": hsc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Add -- 1 hashset", actual)
	})
}

func Test_HashsetCollection_AddNonNil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_AddNonNil", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.AddNonNil(h).AddNonNil(nil)

		// Act
		actual := args.Map{"len": hsc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddNonNil -- skips nil", actual)
	})
}

func Test_HashsetCollection_AddNonEmpty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_AddNonEmpty", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		empty := corestr.New.Hashset.Empty()
		hsc.AddNonEmpty(h).AddNonEmpty(empty)

		// Act
		actual := args.Map{"len": hsc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddNonEmpty -- skips empty", actual)
	})
}

func Test_HashsetCollection_Adds_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_Adds", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 4)
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		h2 := corestr.New.Hashset.Strings([]string{"b"})
		hsc.Adds(h1, h2)

		// Act
		actual := args.Map{"len": hsc.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Adds -- 2 hashsets", actual)
	})
}

func Test_HashsetCollection_Adds_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_Adds_Nil", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		hsc.Adds(nil...)

		// Act
		actual := args.Map{"len": hsc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Adds nil -- no change", actual)
	})
}

func Test_HashsetCollection_Length_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_Length_Nil", func() {
		// Arrange
		var hsc *corestr.HashsetsCollection

		// Act
		actual := args.Map{"len": hsc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Length nil -- 0", actual)
	})
}

func Test_HashsetCollection_LastIndex_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_LastIndex", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)

		// Act
		actual := args.Map{"last": hsc.LastIndex()}

		// Assert
		expected := args.Map{"last": 0}
		expected.ShouldBeEqual(t, 0, "LastIndex -- 0", actual)
	})
}

func Test_HashsetCollection_List_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_List", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)

		// Act
		actual := args.Map{
			"len": len(hsc.List()),
			"ptrLen": len(*hsc.ListPtr()),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"ptrLen": 1,
		}
		expected.ShouldBeEqual(t, 0, "List/ListPtr -- 1 item", actual)
	})
}

func Test_HashsetCollection_StringsList_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_StringsList", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		hsc.Add(h)

		// Act
		actual := args.Map{"len": len(hsc.StringsList())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "StringsList -- 2 strings", actual)
	})
}

func Test_HashsetCollection_StringsList_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_StringsList_Empty", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 0)

		// Act
		actual := args.Map{"len": len(hsc.StringsList())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "StringsList empty -- 0", actual)
	})
}

func Test_HashsetCollection_ListDirectPtr_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_ListDirectPtr", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)

		// Act
		actual := args.Map{"len": len(*hsc.ListDirectPtr())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListDirectPtr -- 1 item", actual)
	})
}

func Test_HashsetCollection_AddHashsetsCollection_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_AddHashsetsCollection", func() {
		// Arrange
		hsc1 := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc1.Add(h)
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 2)
		h2 := corestr.New.Hashset.Strings([]string{"b"})
		hsc2.Add(h2)
		hsc1.AddHashsetsCollection(hsc2)

		// Act
		actual := args.Map{"len": hsc1.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddHashsetsCollection -- merged", actual)
	})
}

func Test_HashsetCollection_AddHashsetsCollection_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_AddHashsetsCollection_Nil", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		hsc.AddHashsetsCollection(nil)

		// Act
		actual := args.Map{"len": hsc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddHashsetsCollection nil -- no change", actual)
	})
}

func Test_HashsetCollection_ConcatNew_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_ConcatNew", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 2)
		h2 := corestr.New.Hashset.Strings([]string{"b"})
		hsc2.Add(h2)
		result := hsc.ConcatNew(hsc2)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ConcatNew -- merged", actual)
	})
}

func Test_HashsetCollection_ConcatNew_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_ConcatNew_Empty", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		result := hsc.ConcatNew()

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ConcatNew empty -- cloned", actual)
	})
}

func Test_HashsetCollection_HasAll_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_HasAll", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a", "b"})
		hsc.Add(h)

		// Act
		actual := args.Map{
			"has": hsc.HasAll("a", "b"),
			"miss": hsc.HasAll("a", "z"),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "HasAll -- found and missing", actual)
	})
}

func Test_HashsetCollection_HasAll_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_HasAll_Empty", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 0)

		// Act
		actual := args.Map{"has": hsc.HasAll("a")}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "HasAll empty -- false", actual)
	})
}

func Test_HashsetCollection_IsEqualPtr_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_IsEqualPtr", func() {
		// Arrange
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		hsc1 := corestr.New.HashsetsCollection.LenCap(0, 2)
		hsc1.Add(h1)
		h2 := corestr.New.Hashset.Strings([]string{"a"})
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 2)
		hsc2.Add(h2)

		// Act
		actual := args.Map{
			"eq":      hsc1.IsEqualPtr(hsc2),
			"same":    hsc1.IsEqualPtr(hsc1),
			"nilBoth": (*corestr.HashsetsCollection)(nil).IsEqualPtr(nil),
			"nilOne":  hsc1.IsEqualPtr(nil),
		}

		// Assert
		expected := args.Map{
			"eq":      true,
			"same":    true,
			"nilBoth": true,
			"nilOne":  false,
		}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr -- various", actual)
	})
}

func Test_HashsetCollection_IsEqualPtr_DiffLen_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_IsEqualPtr_DiffLen", func() {
		// Arrange
		hsc1 := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc1.Add(h)
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 2)

		// Act
		actual := args.Map{"eq": hsc1.IsEqualPtr(hsc2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr diff len -- false", actual)
	})
}

func Test_HashsetCollection_IsEqualPtr_BothEmpty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_IsEqualPtr_BothEmpty", func() {
		// Arrange
		hsc1 := corestr.New.HashsetsCollection.LenCap(0, 0)
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 0)

		// Act
		actual := args.Map{"eq": hsc1.IsEqualPtr(hsc2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr both empty -- true", actual)
	})
}

func Test_HashsetCollection_IsEqualPtr_OneEmpty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_IsEqualPtr_OneEmpty", func() {
		// Arrange
		hsc1 := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc1.Add(h)
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 0)

		// Act
		actual := args.Map{"eq": hsc1.IsEqualPtr(hsc2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr one empty -- false", actual)
	})
}

func Test_HashsetCollection_IsEqualPtr_Different_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_IsEqualPtr_Different", func() {
		// Arrange
		hsc1 := corestr.New.HashsetsCollection.LenCap(0, 2)
		h1 := corestr.New.Hashset.Strings([]string{"a"})
		hsc1.Add(h1)
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 2)
		h2 := corestr.New.Hashset.Strings([]string{"b"})
		hsc2.Add(h2)

		// Act
		actual := args.Map{"eq": hsc1.IsEqualPtr(hsc2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEqualPtr different -- false", actual)
	})
}

func Test_HashsetCollection_IsEqual_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_IsEqual", func() {
		// Arrange
		hsc1 := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc1.Add(h)
		hsc2 := *corestr.New.HashsetsCollection.LenCap(0, 2)
		h2 := corestr.New.Hashset.Strings([]string{"a"})
		hsc2.Add(h2)

		// Act
		actual := args.Map{"eq": hsc1.IsEqual(hsc2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqual -- delegates to IsEqualPtr", actual)
	})
}

func Test_HashsetCollection_String_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_String", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)

		// Act
		actual := args.Map{"nonEmpty": hsc.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_HashsetCollection_String_Empty_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_String_Empty", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 0)

		// Act
		actual := args.Map{"nonEmpty": hsc.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String empty -- has NoElements", actual)
	})
}

func Test_HashsetCollection_Join_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_Join", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)

		// Act
		actual := args.Map{"nonEmpty": hsc.Join(",") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Join -- non-empty", actual)
	})
}

func Test_HashsetCollection_Json_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_Json", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		j := hsc.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_HashsetCollection_MarshalJSON_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_MarshalJSON", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		b, err := hsc.MarshalJSON()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "MarshalJSON -- success", actual)
	})
}

func Test_HashsetCollection_UnmarshalJSON_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_UnmarshalJSON", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		b, _ := hsc.MarshalJSON()
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 0)
		err := hsc2.UnmarshalJSON(b)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- success", actual)
	})
}

func Test_HashsetCollection_UnmarshalJSON_Invalid_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_UnmarshalJSON_Invalid", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 0)
		err := hsc.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
	})
}

func Test_HashsetCollection_Serialize_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_Serialize", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		b, err := hsc.Serialize()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "Serialize -- success", actual)
	})
}

func Test_HashsetCollection_Deserialize_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_Deserialize", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		var dest interface{}
		err := hsc.Deserialize(&dest)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Deserialize -- success", actual)
	})
}

func Test_HashsetCollection_JsonModel_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_JsonModel", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)

		// Act
		actual := args.Map{"notNil": hsc.JsonModel() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModel -- non-nil", actual)
	})
}

func Test_HashsetCollection_JsonModelAny_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_JsonModelAny", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)

		// Act
		actual := args.Map{"notNil": hsc.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_HashsetCollection_ParseInjectUsingJson_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_ParseInjectUsingJson", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		jr := hsc.JsonPtr()
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 0)
		_, err := hsc2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- success", actual)
	})
}

func Test_HashsetCollection_ParseInjectUsingJsonMust_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_ParseInjectUsingJsonMust", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		jr := hsc.JsonPtr()
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 0)
		result := hsc2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_HashsetCollection_InterfaceCasts_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_InterfaceCasts", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)

		// Act
		actual := args.Map{
			"jsoner":   hsc.AsJsoner() != nil,
			"binder":   hsc.AsJsonContractsBinder() != nil,
			"injector": hsc.AsJsonParseSelfInjector() != nil,
			"marsh":    hsc.AsJsonMarshaller() != nil,
		}

		// Assert
		expected := args.Map{
			"jsoner":   true,
			"binder":   true,
			"injector": true,
			"marsh":    true,
		}
		expected.ShouldBeEqual(t, 0, "Interface casts -- all non-nil", actual)
	})
}

func Test_HashsetCollection_JsonParseSelfInject_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSC_JsonParseSelfInject", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		jr := hsc.JsonPtr()
		hsc2 := corestr.New.HashsetsCollection.LenCap(0, 0)
		err := hsc2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

// ── HashsetsCollectionDataModel ─────────────────────────────────────────────

func Test_HashsetCollectionDataModel_DataModel_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSCDM_DataModel", func() {
		// Arrange
		h := corestr.New.Hashset.Strings([]string{"a"})
		dm := &corestr.HashsetsCollectionDataModel{Items: []*corestr.Hashset{h}}
		hsc := corestr.NewHashsetsCollectionUsingDataModel(dm)

		// Act
		actual := args.Map{"len": hsc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "NewHashsetsCollectionUsingDataModel -- 1 item", actual)
	})
}

func Test_HashsetCollectionDataModel_DataModel_Reverse_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HSCDM_DataModel_Reverse", func() {
		// Arrange
		hsc := corestr.New.HashsetsCollection.LenCap(0, 2)
		h := corestr.New.Hashset.Strings([]string{"a"})
		hsc.Add(h)
		dm := corestr.NewHashsetsCollectionDataModelUsing(hsc)

		// Act
		actual := args.Map{"len": len(dm.Items)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "NewHashsetsCollectionDataModelUsing -- 1 item", actual)
	})
}
