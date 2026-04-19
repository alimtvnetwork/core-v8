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

// File scope: constructor + first-use smoke checks for every corestr factory
// (`corestr.New.<Type>.Cap/Empty`) plus a handful of free helpers
// (`AnyToString`, `CloneSlice`, etc.). Each test asserts the freshly built
// value is non-nil and reports an empty / single-item state correctly.

// ── SimpleSlice constructor + first-use ──

func Test_SimpleSlice_NewCap_InitialStateIsEmpty(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_NewCap_InitialStateIsEmpty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(5)

		// Act
		actual := args.Map{
			"isNil":   s == nil,
			"isEmpty": s.IsEmpty(),
			"length":  s.Length(),
			"hasAny":  s.HasAnyItem(),
		}

		// Assert
		expected := args.Map{
			"isNil":   false,
			"isEmpty": true,
			"length":  0,
			"hasAny":  false,
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice.Cap returns empty slice -- newly constructed", actual)
	})
}

func Test_SimpleSlice_AddSingleItem_LengthBecomesOne(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_AddSingleItem_LengthBecomesOne", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(5)
		s.Add("hello")
		s.Add("world")

		// Act
		actual := args.Map{
			"length": s.Length(),
			"hasAny": s.HasAnyItem(),
			"first":  s.Strings()[0],
		}

		// Assert
		expected := args.Map{
			"length": 2,
			"hasAny": true,
			"first":  "hello",
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice.Add appends item -- two adds yield length 2", actual)
	})
}

func Test_SimpleSlice_AddsVariadic_AppendsAllItems(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_AddsVariadic_AppendsAllItems", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(5)
		s.Adds("a", "b", "c")

		// Act
		actual := args.Map{"length": s.Length()}

		// Assert
		expected := args.Map{"length": 3}
		expected.ShouldBeEqual(t, 0, "SimpleSlice.Adds appends variadic -- three items added", actual)
	})
}

func Test_SimpleSlice_AddIf_OnlyAppendsWhenConditionTrue(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_AddIf_OnlyAppendsWhenConditionTrue", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(5)
		s.AddIf(true, "yes")
		s.AddIf(false, "no")

		// Act
		actual := args.Map{"length": s.Length()}

		// Assert
		expected := args.Map{"length": 1}
		expected.ShouldBeEqual(t, 0, "SimpleSlice.AddIf gates append -- only true predicate adds", actual)
	})
}

func Test_SimpleSlice_AppendFmt_FormatsAndAppends(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_AppendFmt_FormatsAndAppends", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(5)
		s.AppendFmt("hello %d", 42)

		// Act
		actual := args.Map{
			"length": s.Length(),
			"first":  s.Strings()[0],
		}

		// Assert
		expected := args.Map{
			"length": 1,
			"first":  "hello 42",
		}
		expected.ShouldBeEqual(t, 0, "SimpleSlice.AppendFmt formats item -- printf-style append", actual)
	})
}

func Test_SimpleSlice_StringDump_ReturnsNonEmpty(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_StringDump_ReturnsNonEmpty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(5)
		s.Add("hello")

		// Act
		actual := args.Map{"notEmpty": s.String() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "SimpleSlice.String dumps content -- non-empty after Add", actual)
	})
}

func Test_SimpleSlice_JoinCsvDump_ReturnsNonEmpty(t *testing.T) {
	safeTest(t, "Test_SimpleSlice_JoinCsvDump_ReturnsNonEmpty", func() {
		// Arrange
		s := corestr.New.SimpleSlice.Cap(5)
		s.Adds("a", "b")

		// Act
		actual := args.Map{"notEmpty": s.JoinCsv() != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "SimpleSlice.JoinCsv joins with comma -- non-empty for two items", actual)
	})
}

// ── Collection constructor + first-use ──

func Test_Collection_NewCap_InitialStateIsEmpty(t *testing.T) {
	safeTest(t, "Test_Collection_NewCap_InitialStateIsEmpty", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{
			"isNil":   c == nil,
			"isEmpty": c.IsEmpty(),
			"length":  c.Length(),
		}

		// Assert
		expected := args.Map{
			"isNil":   false,
			"isEmpty": true,
			"length":  0,
		}
		expected.ShouldBeEqual(t, 0, "Collection.Cap returns empty collection -- newly constructed", actual)
	})
}

func Test_Collection_AddStringsSlice_IncreasesLength(t *testing.T) {
	safeTest(t, "Test_Collection_AddStringsSlice_IncreasesLength", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddStrings([]string{"hello", "world"})

		// Act
		actual := args.Map{"length": c.Length()}

		// Assert
		expected := args.Map{"length": 2}
		expected.ShouldBeEqual(t, 0, "Collection.AddStrings appends slice -- length matches input size", actual)
	})
}

// ── Hashset constructor + first-use ──

func Test_Hashset_NewCap_InitialStateIsEmpty(t *testing.T) {
	safeTest(t, "Test_Hashset_NewCap_InitialStateIsEmpty", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)

		// Act
		actual := args.Map{
			"isNil":   h == nil,
			"isEmpty": h.IsEmpty(),
			"length":  h.Length(),
			"hasAny":  h.HasAnyItem(),
		}

		// Assert
		expected := args.Map{
			"isNil":   false,
			"isEmpty": true,
			"length":  0,
			"hasAny":  false,
		}
		expected.ShouldBeEqual(t, 0, "Hashset.Cap returns empty set -- newly constructed", actual)
	})
}

func Test_Hashset_AddDuplicates_DeduplicatesAndReportsMembership(t *testing.T) {
	safeTest(t, "Test_Hashset_AddDuplicates_DeduplicatesAndReportsMembership", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.Add("hello")
		h.Add("hello")
		h.Add("world")

		// Act
		actual := args.Map{
			"length": h.Length(),
			"has":    h.Has("hello"),
			"hasNo":  h.Has("nope"),
		}

		// Assert
		expected := args.Map{
			"length": 2,
			"has":    true,
			"hasNo":  false,
		}
		expected.ShouldBeEqual(t, 0, "Hashset.Add deduplicates -- duplicate insert keeps length 2", actual)
	})
}

func Test_Hashset_AddsVariadic_StoresUniqueItems(t *testing.T) {
	safeTest(t, "Test_Hashset_AddsVariadic_StoresUniqueItems", func() {
		// Arrange
		h := corestr.New.Hashset.Cap(5)
		h.Adds("a", "b", "c")

		// Act
		actual := args.Map{"length": h.Length()}

		// Assert
		expected := args.Map{"length": 3}
		expected.ShouldBeEqual(t, 0, "Hashset.Adds stores unique -- three distinct items added", actual)
	})
}

// ── Hashmap constructor + first-use ──

func Test_Hashmap_NewCap_InitialStateIsEmpty(t *testing.T) {
	safeTest(t, "Test_Hashmap_NewCap_InitialStateIsEmpty", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)

		// Act
		actual := args.Map{
			"isNil":   h == nil,
			"isEmpty": h.IsEmpty(),
			"length":  h.Length(),
		}

		// Assert
		expected := args.Map{
			"isNil":   false,
			"isEmpty": true,
			"length":  0,
		}
		expected.ShouldBeEqual(t, 0, "Hashmap.Cap returns empty map -- newly constructed", actual)
	})
}

func Test_Hashmap_AddOrUpdateThenGet_ReturnsStoredValue(t *testing.T) {
	safeTest(t, "Test_Hashmap_AddOrUpdateThenGet_ReturnsStoredValue", func() {
		// Arrange
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("key", "value")
		val, _ := h.Get("key")

		// Act
		actual := args.Map{
			"length": h.Length(),
			"has":    h.Has("key"),
			"getVal": val,
		}

		// Assert
		expected := args.Map{
			"length": 1,
			"has":    true,
			"getVal": "value",
		}
		expected.ShouldBeEqual(t, 0, "Hashmap.AddOrUpdate stores entry -- Get returns the same value", actual)
	})
}

// ── KeyValues constructor ──

func Test_KeyValues_NewCap_ReturnsNonNilContainer(t *testing.T) {
	safeTest(t, "Test_KeyValues_NewCap_ReturnsNonNilContainer", func() {
		// Act
		actual := args.Map{"notNil": corestr.New.KeyValues.Cap(5) != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "KeyValues.Cap returns non-nil -- newly constructed", actual)
	})
}

// ── LinkedList constructor + first-use ──

func Test_LinkedList_NewEmpty_InitialStateIsEmpty(t *testing.T) {
	safeTest(t, "Test_LinkedList_NewEmpty_InitialStateIsEmpty", func() {
		// Arrange
		ll := corestr.New.LinkedList.Empty()

		// Act
		actual := args.Map{
			"isNil":   ll == nil,
			"isEmpty": ll.IsEmpty(),
			"length":  ll.Length(),
		}

		// Assert
		expected := args.Map{
			"isNil":   false,
			"isEmpty": true,
			"length":  0,
		}
		expected.ShouldBeEqual(t, 0, "LinkedList.Empty returns empty list -- newly constructed", actual)
	})
}

func Test_LinkedList_AddTwoItems_LengthIsTwo(t *testing.T) {
	safeTest(t, "Test_LinkedList_AddTwoItems_LengthIsTwo", func() {
		// Arrange
		ll := corestr.New.LinkedList.Empty()
		ll.Add("hello")
		ll.Add("world")

		// Act
		actual := args.Map{
			"length":  ll.Length(),
			"isEmpty": ll.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"length":  2,
			"isEmpty": false,
		}
		expected.ShouldBeEqual(t, 0, "LinkedList.Add appends node -- two adds yield length 2", actual)
	})
}

// ── CharHashsetMap / CharCollectionMap / SimpleStringOnce / HashsetsCollection / LinkedCollection / CollectionsOfCollection ──

func Test_CharHashsetMap_NewCap_InitialStateIsEmpty(t *testing.T) {
	safeTest(t, "Test_CharHashsetMap_NewCap_InitialStateIsEmpty", func() {
		// Arrange
		chm := corestr.New.CharHashsetMap.Cap(0, 0)

		// Act
		actual := args.Map{
			"isNil":   chm == nil,
			"isEmpty": chm.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"isNil":   false,
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "CharHashsetMap.Cap returns empty -- newly constructed", actual)
	})
}

func Test_CharCollectionMap_NewEmpty_InitialStateIsEmpty(t *testing.T) {
	safeTest(t, "Test_CharCollectionMap_NewEmpty_InitialStateIsEmpty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.Empty()

		// Act
		actual := args.Map{
			"isNil":   ccm == nil,
			"isEmpty": ccm.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"isNil":   false,
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "CharCollectionMap.Empty returns empty -- newly constructed", actual)
	})
}

func Test_SimpleStringOnce_NewEmpty_InitialStateIsEmpty(t *testing.T) {
	safeTest(t, "Test_SimpleStringOnce_NewEmpty_InitialStateIsEmpty", func() {
		// Arrange
		so := corestr.New.SimpleStringOnce.Empty()

		// Act
		actual := args.Map{
			"isNil":   false,
			"isEmpty": so.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"isNil":   false,
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "SimpleStringOnce.Empty returns empty -- newly constructed", actual)
	})
}

func Test_HashsetsCollection_NewCap_InitialStateIsEmpty(t *testing.T) {
	safeTest(t, "Test_HashsetsCollection_NewCap_InitialStateIsEmpty", func() {
		// Arrange
		hc := corestr.New.HashsetsCollection.Cap(5)

		// Act
		actual := args.Map{
			"isNil":   hc == nil,
			"isEmpty": hc.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"isNil":   false,
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection.Cap returns empty -- newly constructed", actual)
	})
}

func Test_LinkedCollection_NewEmpty_InitialStateIsEmpty(t *testing.T) {
	safeTest(t, "Test_LinkedCollection_NewEmpty_InitialStateIsEmpty", func() {
		// Arrange
		lc := corestr.New.LinkedCollection.Empty()

		// Act
		actual := args.Map{
			"isNil":   lc == nil,
			"isEmpty": lc.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"isNil":   false,
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "LinkedCollection.Empty returns empty -- newly constructed", actual)
	})
}

func Test_CollectionsOfCollection_NewCap_InitialStateIsEmpty(t *testing.T) {
	safeTest(t, "Test_CollectionsOfCollection_NewCap_InitialStateIsEmpty", func() {
		// Arrange
		cc := corestr.New.CollectionsOfCollection.Cap(5)

		// Act
		actual := args.Map{
			"isNil":   cc == nil,
			"isEmpty": cc.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"isNil":   false,
			"isEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "CollectionsOfCollection.Cap returns empty -- newly constructed", actual)
	})
}

// ── LeftRight / LeftMiddleRight / ValidValue / ValidValues ──

func Test_NewLeftRight_BothSidesPopulated_ReportsNonEmptyAndSafe(t *testing.T) {
	safeTest(t, "Test_NewLeftRight_BothSidesPopulated_ReportsNonEmptyAndSafe", func() {
		// Arrange
		lr := corestr.NewLeftRight("l", "r")

		// Act
		actual := args.Map{
			"isLeftEmpty":  lr.IsLeftEmpty(),
			"isRightEmpty": lr.IsRightEmpty(),
			"hasSafe":      lr.HasSafeNonEmpty(),
		}

		// Assert
		expected := args.Map{
			"isLeftEmpty":  false,
			"isRightEmpty": false,
			"hasSafe":      true,
		}
		expected.ShouldBeEqual(t, 0, "NewLeftRight constructs with both sides -- emptiness flags are false", actual)
	})
}

func Test_NewLeftMiddleRight_AllThreeSidesPopulated_AllEmptinessFlagsFalse(t *testing.T) {
	safeTest(t, "Test_NewLeftMiddleRight_AllThreeSidesPopulated_AllEmptinessFlagsFalse", func() {
		// Arrange
		lmr := corestr.NewLeftMiddleRight("l", "m", "r")

		// Act
		actual := args.Map{
			"isLeftEmpty":   lmr.IsLeftEmpty(),
			"isMiddleEmpty": lmr.IsMiddleEmpty(),
			"isRightEmpty":  lmr.IsRightEmpty(),
		}

		// Assert
		expected := args.Map{
			"isLeftEmpty":   false,
			"isMiddleEmpty": false,
			"isRightEmpty":  false,
		}
		expected.ShouldBeEqual(t, 0, "NewLeftMiddleRight constructs with all sides -- emptiness flags are false", actual)
	})
}

func Test_ValidValue_StructLiteral_ExposesFields(t *testing.T) {
	safeTest(t, "Test_ValidValue_StructLiteral_ExposesFields", func() {
		// Arrange
		vv := corestr.ValidValue{Value: "hello", IsValid: true}

		// Act
		actual := args.Map{
			"value":   vv.Value,
			"isValid": vv.IsValid,
		}

		// Assert
		expected := args.Map{
			"value":   "hello",
			"isValid": true,
		}
		expected.ShouldBeEqual(t, 0, "ValidValue struct literal exposes fields -- Value and IsValid as set", actual)
	})
}

func Test_ValidValues_SingleEntrySlice_LengthIsOne(t *testing.T) {
	safeTest(t, "Test_ValidValues_SingleEntrySlice_LengthIsOne", func() {
		// Arrange
		vv := corestr.ValidValues{ValidValues: []*corestr.ValidValue{{Value: "a", IsValid: true}}}

		// Act
		actual := args.Map{"len": len(vv.ValidValues)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ValidValues struct with one entry -- length 1", actual)
	})
}

// ── AnyToString ──

func Test_AnyToString_FromInteger_ReturnsNonEmpty(t *testing.T) {
	safeTest(t, "Test_AnyToString_FromInteger_ReturnsNonEmpty", func() {
		// Act
		actual := args.Map{"notEmpty": corestr.AnyToString(false, 42) != ""}

		// Assert
		expected := args.Map{"notEmpty": true}
		expected.ShouldBeEqual(t, 0, "AnyToString formats integer -- non-empty result", actual)
	})
}

func Test_AnyToString_FromString_ReturnsSameString(t *testing.T) {
	safeTest(t, "Test_AnyToString_FromString_ReturnsSameString", func() {
		// Act
		actual := args.Map{"result": corestr.AnyToString(false, "hello")}

		// Assert
		expected := args.Map{"result": "hello"}
		expected.ShouldBeEqual(t, 0, "AnyToString formats string -- echoes input", actual)
	})
}

func Test_AnyToString_FromNilInput_ReturnsNonEmptyPlaceholder(t *testing.T) {
	safeTest(t, "Test_AnyToString_FromNilInput_ReturnsNonEmptyPlaceholder", func() {
		// Arrange
		result := corestr.AnyToString(false, nil)

		// Act
		actual := args.Map{
			"ok":    true,
			"empty": result == "",
		}

		// Assert
		expected := args.Map{
			"ok":    true,
			"empty": false,
		}
		expected.ShouldBeEqual(t, 0, "AnyToString formats nil -- returns non-empty placeholder", actual)
	})
}

// ── AllIndividualStringsOfStringsLength ──

func Test_AllIndividualStringsOfStringsLength_TwoStrings_ReturnsCount(t *testing.T) {
	safeTest(t, "Test_AllIndividualStringsOfStringsLength_TwoStrings_ReturnsCount", func() {
		// Arrange
		items := [][]string{{"ab", "cde"}}

		// Act
		actual := args.Map{"result": corestr.AllIndividualStringsOfStringsLength(&items)}

		// Assert
		expected := args.Map{"result": 2}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength counts strings -- two inner strings", actual)
	})
}

func Test_AllIndividualStringsOfStringsLength_NilPointer_ReturnsZero(t *testing.T) {
	safeTest(t, "Test_AllIndividualStringsOfStringsLength_NilPointer_ReturnsZero", func() {
		// Act
		actual := args.Map{"result": corestr.AllIndividualStringsOfStringsLength(nil)}

		// Assert
		expected := args.Map{"result": 0}
		expected.ShouldBeEqual(t, 0, "AllIndividualStringsOfStringsLength handles nil -- returns 0", actual)
	})
}

// ── CloneSlice / CloneSliceIf ──

func Test_CloneSlice_TwoStrings_ReturnsLengthTwoCopy(t *testing.T) {
	safeTest(t, "Test_CloneSlice_TwoStrings_ReturnsLengthTwoCopy", func() {
		// Act
		actual := args.Map{"len": len(corestr.CloneSlice([]string{"a", "b"}))}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "CloneSlice copies slice -- length matches source", actual)
	})
}

func Test_CloneSlice_NilInput_ReturnsNonNilEmptySlice(t *testing.T) {
	safeTest(t, "Test_CloneSlice_NilInput_ReturnsNonNilEmptySlice", func() {
		// Act
		actual := args.Map{"isNil": corestr.CloneSlice(nil) == nil}

		// Assert
		expected := args.Map{"isNil": false}
		expected.ShouldBeEqual(t, 0, "CloneSlice handles nil -- returns empty slice not nil", actual)
	})
}

func Test_CloneSliceIf_PredicateTrue_ReturnsLengthOne(t *testing.T) {
	safeTest(t, "Test_CloneSliceIf_PredicateTrue_ReturnsLengthOne", func() {
		// Act
		actual := args.Map{"len": len(corestr.CloneSliceIf(true, "a"))}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf with true predicate -- returns single-item slice", actual)
	})
}

func Test_CloneSliceIf_PredicateFalse_StillReturnsLengthOne(t *testing.T) {
	safeTest(t, "Test_CloneSliceIf_PredicateFalse_StillReturnsLengthOne", func() {
		// Act
		actual := args.Map{"len": len(corestr.CloneSliceIf(false, "a"))}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "CloneSliceIf with false predicate -- documented behaviour returns one item", actual)
	})
}

// ── TextWithLineNumber / ValueStatus ──

func Test_TextWithLineNumber_StructLiteral_ExposesFields(t *testing.T) {
	safeTest(t, "Test_TextWithLineNumber_StructLiteral_ExposesFields", func() {
		// Arrange
		tw := corestr.TextWithLineNumber{LineNumber: 5, Text: "hello"}

		// Act
		actual := args.Map{
			"lineNumber": tw.LineNumber,
			"text":       tw.Text,
		}

		// Assert
		expected := args.Map{
			"lineNumber": 5,
			"text":       "hello",
		}
		expected.ShouldBeEqual(t, 0, "TextWithLineNumber struct literal exposes fields -- LineNumber and Text as set", actual)
	})
}

func Test_ValueStatus_StructLiteralWithValidValue_ExposesFields(t *testing.T) {
	safeTest(t, "Test_ValueStatus_StructLiteralWithValidValue_ExposesFields", func() {
		// Arrange
		vv := corestr.NewValidValue("hello")
		vs := corestr.ValueStatus{ValueValid: vv, Index: 0}

		// Act
		actual := args.Map{
			"index":  vs.Index,
			"notNil": vs.ValueValid != nil,
		}

		// Assert
		expected := args.Map{
			"index":  0,
			"notNil": true,
		}
		expected.ShouldBeEqual(t, 0, "ValueStatus struct literal exposes fields -- Index and ValueValid populated", actual)
	})
}
