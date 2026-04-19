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

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// CollectionsOfCollection — Segment 6c
// ══════════════════════════════════════════════════════════════════════════════

func Test_CollectionOfCollections_IsEmpty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_IsEmpty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 0)

		// Act
		actual := args.Map{
			"empty": coc.IsEmpty(),
			"hasItems": coc.HasItems(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"hasItems": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty -- true on empty", actual)
	})
}

func Test_CollectionOfCollections_Add_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_Add", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		coc.Add(c)

		// Act
		actual := args.Map{
			"len": coc.Length(),
			"allLen": coc.AllIndividualItemsLength(),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"allLen": 2,
		}
		expected.ShouldBeEqual(t, 0, "Add -- 1 collection 2 items", actual)
	})
}

func Test_CollectionOfCollections_Add_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_Add_Empty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Cap(0)
		coc.Add(c)

		// Act
		actual := args.Map{"len": coc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Add empty collection -- skipped", actual)
	})
}

func Test_CollectionOfCollections_Adds_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_Adds", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c1 := *corestr.New.Collection.Strings([]string{"a"})
		c2 := *corestr.New.Collection.Strings([]string{"b"})
		coc.Adds(c1, c2)

		// Act
		actual := args.Map{"len": coc.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Adds -- 2 collections", actual)
	})
}

func Test_CollectionOfCollections_Adds_Nil_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_Adds_Nil", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		coc.Adds(nil...)

		// Act
		actual := args.Map{"len": coc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Adds nil -- no change", actual)
	})
}

func Test_CollectionOfCollections_AddCollections_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_AddCollections", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := *corestr.New.Collection.Strings([]string{"a"})
		coc.AddCollections(c)

		// Act
		actual := args.Map{"len": coc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddCollections -- 1 collection", actual)
	})
}

func Test_CollectionOfCollections_AddCollections_Nil_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_AddCollections_Nil", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		coc.AddCollections(nil...)

		// Act
		actual := args.Map{"len": coc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollections nil -- no change", actual)
	})
}

func Test_CollectionOfCollections_AddStrings_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_AddStrings", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		coc.AddStrings(false, []string{"a", "b"})

		// Act
		actual := args.Map{"len": coc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddStrings -- 1 collection", actual)
	})
}

func Test_CollectionOfCollections_AddStrings_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_AddStrings_Empty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		coc.AddStrings(false, []string{})

		// Act
		actual := args.Map{"len": coc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStrings empty -- no change", actual)
	})
}

func Test_CollectionOfCollections_AddsStringsOfStrings_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_AddsStringsOfStrings", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 4)
		coc.AddsStringsOfStrings(false, []string{"a"}, []string{"b"})

		// Act
		actual := args.Map{"len": coc.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsStringsOfStrings -- 2 collections", actual)
	})
}

func Test_CollectionOfCollections_AddsStringsOfStrings_Nil_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_AddsStringsOfStrings_Nil", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		coc.AddsStringsOfStrings(false, nil...)

		// Act
		actual := args.Map{"len": coc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsStringsOfStrings nil -- no change", actual)
	})
}

func Test_CollectionOfCollections_Items_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_Items", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)

		// Act
		actual := args.Map{"len": len(coc.Items())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Items -- 1 collection", actual)
	})
}

func Test_CollectionOfCollections_List_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_List", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 4)
		c1 := corestr.New.Collection.Strings([]string{"a", "b"})
		c2 := corestr.New.Collection.Strings([]string{"c"})
		coc.Add(c1).Add(c2)

		// Act
		actual := args.Map{"len": len(coc.List(0))}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "List -- 3 items total", actual)
	})
}

func Test_CollectionOfCollections_List_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_List_Empty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 0)

		// Act
		actual := args.Map{"len": len(coc.List(0))}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "List empty -- 0", actual)
	})
}

func Test_CollectionOfCollections_ToCollection_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_ToCollection", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a", "b"})
		coc.Add(c)
		result := coc.ToCollection()

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "ToCollection -- 2 items", actual)
	})
}

func Test_CollectionOfCollections_AllIndividualItemsLength_WithEmpty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_AllIndividualItemsLength_WithEmpty", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 4)
		c1 := corestr.New.Collection.Strings([]string{"a"})
		c2 := corestr.New.Collection.Cap(0)
		coc.Add(c1).Add(c2)

		// Act
		actual := args.Map{"len": coc.AllIndividualItemsLength()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AllIndividualItemsLength -- skips empty", actual)
	})
}

func Test_CollectionOfCollections_String_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_String", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)

		// Act
		actual := args.Map{"nonEmpty": coc.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

// ── JSON ────────────────────────────────────────────────────────────────────

func Test_CollectionOfCollections_Json_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_Json", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)
		j := coc.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_CollectionOfCollections_MarshalJSON_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_MarshalJSON", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)
		b, err := coc.MarshalJSON()

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

func Test_CollectionOfCollections_UnmarshalJSON_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_UnmarshalJSON", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)
		b, _ := coc.MarshalJSON()
		coc2 := corestr.New.CollectionsOfCollection.LenCap(0, 0)
		err := coc2.UnmarshalJSON(b)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- success", actual)
	})
}

func Test_CollectionOfCollections_UnmarshalJSON_Invalid_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_UnmarshalJSON_Invalid", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 0)
		err := coc.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
	})
}

func Test_CollectionOfCollections_JsonModel_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_JsonModel", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)

		// Act
		actual := args.Map{"len": len(coc.JsonModel().Items)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "JsonModel -- 1 item", actual)
	})
}

func Test_CollectionOfCollections_JsonModelAny_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_JsonModelAny", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)

		// Act
		actual := args.Map{"notNil": coc.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_CollectionOfCollections_InterfaceCasts_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_InterfaceCasts", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)

		// Act
		actual := args.Map{
			"jsoner":   coc.AsJsoner() != nil,
			"binder":   coc.AsJsonContractsBinder() != nil,
			"injector": coc.AsJsonParseSelfInjector() != nil,
			"marsh":    coc.AsJsonMarshaller() != nil,
		}

		// Assert
		expected := args.Map{
			"jsoner": true,
			"binder": true,
			"injector": true,
			"marsh": true,
		}
		expected.ShouldBeEqual(t, 0, "Interface casts -- all non-nil", actual)
	})
}

func Test_CollectionOfCollections_ParseInjectUsingJson_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_ParseInjectUsingJson", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)
		jr := coc.JsonPtr()
		coc2 := corestr.New.CollectionsOfCollection.LenCap(0, 0)
		_, err := coc2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- success", actual)
	})
}

func Test_CollectionOfCollections_ParseInjectUsingJsonMust_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_ParseInjectUsingJsonMust", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)
		jr := coc.JsonPtr()
		coc2 := corestr.New.CollectionsOfCollection.LenCap(0, 0)
		result := coc2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_CollectionOfCollections_JsonParseSelfInject_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_COC_JsonParseSelfInject", func() {
		// Arrange
		coc := corestr.New.CollectionsOfCollection.LenCap(0, 2)
		c := corestr.New.Collection.Strings([]string{"a"})
		coc.Add(c)
		jr := coc.JsonPtr()
		coc2 := corestr.New.CollectionsOfCollection.LenCap(0, 0)
		err := coc2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyValueCollection — Segment 6d
// ══════════════════════════════════════════════════════════════════════════════

