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

func Test_KeyValueCollection_IsEmpty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_IsEmpty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{
			"empty": kvc.IsEmpty(),
			"hasAny": kvc.HasAnyItem(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"hasAny": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty -- true on empty", actual)
	})
}

func Test_KeyValueCollection_Add_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Add", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1").Add("b", "2")

		// Act
		actual := args.Map{
			"len": kvc.Length(),
			"count": kvc.Count(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"count": 2,
		}
		expected.ShouldBeEqual(t, 0, "Add -- 2 pairs", actual)
	})
}

func Test_KeyValueCollection_AddIf_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddIf", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddIf(true, "a", "1").AddIf(false, "b", "2")

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddIf -- only true", actual)
	})
}

func Test_KeyValueCollection_Adds_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Adds", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Adds(
			corestr.KeyValuePair{Key: "a", Value: "1"},
			corestr.KeyValuePair{Key: "b", Value: "2"},
		)

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Adds -- 2 pairs", actual)
	})
}

func Test_KeyValueCollection_Adds_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Adds_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Adds()

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Adds empty -- no change", actual)
	})
}

func Test_KeyValueCollection_AddMap_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddMap", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddMap(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddMap -- 1 pair", actual)
	})
}

func Test_KeyValueCollection_AddMap_Nil_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddMap_Nil", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddMap(nil)

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddMap nil -- no change", actual)
	})
}

func Test_KeyValueCollection_AddHashsetMap_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddHashsetMap", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashsetMap(map[string]bool{"a": true})

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddHashsetMap -- 1 pair", actual)
	})
}

func Test_KeyValueCollection_AddHashsetMap_Nil_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddHashsetMap_Nil", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashsetMap(nil)

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashsetMap nil -- no change", actual)
	})
}

func Test_KeyValueCollection_AddHashset_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddHashset", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		hs := corestr.New.Hashset.Strings([]string{"a"})
		kvc.AddHashset(hs)

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddHashset -- 1 pair", actual)
	})
}

func Test_KeyValueCollection_AddHashset_Nil_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddHashset_Nil", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddHashset(nil)

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashset nil -- no change", actual)
	})
}

func Test_KeyValueCollection_AddsHashmap_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddsHashmap", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		hm := corestr.New.Hashmap.Cap(2)
		hm.Set("a", "1")
		kvc.AddsHashmap(hm)

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddsHashmap -- 1 pair", actual)
	})
}

func Test_KeyValueCollection_AddsHashmap_Nil_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddsHashmap_Nil", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddsHashmap(nil)

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsHashmap nil -- no change", actual)
	})
}

func Test_KeyValueCollection_AddsHashmaps_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddsHashmaps", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		hm1 := corestr.New.Hashmap.Cap(2)
		hm1.Set("a", "1")
		hm2 := corestr.New.Hashmap.Cap(2)
		hm2.Set("b", "2")
		kvc.AddsHashmaps(hm1, hm2)

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddsHashmaps -- 2 pairs", actual)
	})
}

func Test_KeyValueCollection_AddsHashmaps_Nil_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddsHashmaps_Nil", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddsHashmaps(nil)

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddsHashmaps nil -- no change", actual)
	})
}

func Test_KeyValueCollection_AddStringBySplit_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddStringBySplit", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddStringBySplit("=", "key=value")
		val, found := kvc.Get("key")

		// Act
		actual := args.Map{
			"found": found,
			"val": val,
		}

		// Assert
		expected := args.Map{
			"found": true,
			"val": "value",
		}
		expected.ShouldBeEqual(t, 0, "AddStringBySplit -- key=value", actual)
	})
}

func Test_KeyValueCollection_AddStringBySplitTrim_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AddStringBySplitTrim", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.AddStringBySplitTrim("=", " key = value ")
		val, found := kvc.Get("key")

		// Act
		actual := args.Map{
			"found": found,
			"val": val,
		}

		// Assert
		expected := args.Map{
			"found": true,
			"val": "value",
		}
		expected.ShouldBeEqual(t, 0, "AddStringBySplitTrim -- trimmed", actual)
	})
}

// ── Accessors ───────────────────────────────────────────────────────────────

func Test_KeyValueCollection_First_Last_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_First_Last", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1").Add("b", "2")

		// Act
		actual := args.Map{
			"first": kvc.First().Key,
			"last": kvc.Last().Key,
			"lastIdx": kvc.LastIndex(),
		}

		// Assert
		expected := args.Map{
			"first": "a",
			"last": "b",
			"lastIdx": 1,
		}
		expected.ShouldBeEqual(t, 0, "First/Last -- correct", actual)
	})
}

func Test_KeyValueCollection_FirstOrDefault_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_FirstOrDefault_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"nil": kvc.FirstOrDefault() == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "FirstOrDefault empty -- nil", actual)
	})
}

func Test_KeyValueCollection_LastOrDefault_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_LastOrDefault_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"nil": kvc.LastOrDefault() == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "LastOrDefault empty -- nil", actual)
	})
}

func Test_KeyValueCollection_HasKey_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_HasKey", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{
			"has": kvc.HasKey("a"),
			"miss": kvc.HasKey("z"),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "HasKey -- found and missing", actual)
	})
}

func Test_KeyValueCollection_IsContains_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_IsContains", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{
			"has": kvc.IsContains("a"),
			"miss": kvc.IsContains("z"),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "IsContains -- found and missing", actual)
	})
}

func Test_KeyValueCollection_HasIndex_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_HasIndex", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{
			"has": kvc.HasIndex(0),
			"miss": kvc.HasIndex(5),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "HasIndex -- found and missing", actual)
	})
}

func Test_KeyValueCollection_Get_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Get", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		val, found := kvc.Get("a")
		val2, found2 := kvc.Get("z")

		// Act
		actual := args.Map{
			"val": val,
			"found": found,
			"val2": val2,
			"found2": found2,
		}

		// Assert
		expected := args.Map{
			"val": "1",
			"found": true,
			"val2": "",
			"found2": false,
		}
		expected.ShouldBeEqual(t, 0, "Get -- found and missing", actual)
	})
}

func Test_KeyValueCollection_SafeValueAt_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_SafeValueAt", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{
			"val": kvc.SafeValueAt(0),
			"miss": kvc.SafeValueAt(5),
		}

		// Assert
		expected := args.Map{
			"val": "1",
			"miss": "",
		}
		expected.ShouldBeEqual(t, 0, "SafeValueAt -- found and out of bounds", actual)
	})
}

func Test_KeyValueCollection_SafeValueAt_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_SafeValueAt_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"val": kvc.SafeValueAt(0)}

		// Assert
		expected := args.Map{"val": ""}
		expected.ShouldBeEqual(t, 0, "SafeValueAt empty -- empty string", actual)
	})
}

func Test_KeyValueCollection_SafeValuesAtIndexes_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_SafeValuesAtIndexes", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1").Add("b", "2")
		result := kvc.SafeValuesAtIndexes(0, 1, 5)

		// Act
		actual := args.Map{
			"len": len(result),
			"first": result[0],
			"last": result[2],
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"first": "1",
			"last": "",
		}
		expected.ShouldBeEqual(t, 0, "SafeValuesAtIndexes -- values", actual)
	})
}

func Test_KeyValueCollection_SafeValuesAtIndexes_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_SafeValuesAtIndexes_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		result := kvc.SafeValuesAtIndexes()

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SafeValuesAtIndexes empty -- 0", actual)
	})
}

func Test_KeyValueCollection_Find_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Find", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1").Add("b", "2")
		result := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, kv.Key == "a", false
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Find -- 1 match", actual)
	})
}

func Test_KeyValueCollection_Find_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Find_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		result := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, true, false
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Find empty -- 0", actual)
	})
}

func Test_KeyValueCollection_Find_Break_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Find_Break", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1").Add("b", "2").Add("c", "3")
		result := kvc.Find(func(i int, kv corestr.KeyValuePair) (corestr.KeyValuePair, bool, bool) {
			return kv, true, true
		})

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Find break -- 1 item", actual)
	})
}

// ── Keys / Values ───────────────────────────────────────────────────────────

func Test_KeyValueCollection_AllKeys_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AllKeys", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1").Add("b", "2")

		// Act
		actual := args.Map{"len": len(kvc.AllKeys())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AllKeys -- 2 keys", actual)
	})
}

func Test_KeyValueCollection_AllKeys_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AllKeys_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"len": len(kvc.AllKeys())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllKeys empty -- 0", actual)
	})
}

func Test_KeyValueCollection_AllKeysSorted_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AllKeysSorted", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("b", "2").Add("a", "1")
		sorted := kvc.AllKeysSorted()

		// Act
		actual := args.Map{"first": sorted[0]}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "AllKeysSorted -- sorted", actual)
	})
}

func Test_KeyValueCollection_AllValues_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AllValues", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1").Add("b", "2")

		// Act
		actual := args.Map{"len": len(kvc.AllValues())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AllValues -- 2 values", actual)
	})
}

func Test_KeyValueCollection_AllValues_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_AllValues_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"len": len(kvc.AllValues())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AllValues empty -- 0", actual)
	})
}

// ── String / Join ───────────────────────────────────────────────────────────

func Test_KeyValueCollection_Strings_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Strings", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"len": len(kvc.Strings())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Strings -- 1 string", actual)
	})
}

func Test_KeyValueCollection_Strings_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Strings_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"len": len(kvc.Strings())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Strings empty -- 0", actual)
	})
}

func Test_KeyValueCollection_StringsUsingFormat_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_StringsUsingFormat", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		result := kvc.StringsUsingFormat("%v=%v")

		// Act
		actual := args.Map{"val": result[0]}

		// Assert
		expected := args.Map{"val": "a=1"}
		expected.ShouldBeEqual(t, 0, "StringsUsingFormat -- formatted", actual)
	})
}

func Test_KeyValueCollection_StringsUsingFormat_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_StringsUsingFormat_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		result := kvc.StringsUsingFormat("%v=%v")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "StringsUsingFormat empty -- 0", actual)
	})
}

func Test_KeyValueCollection_String_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_String", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"nonEmpty": kvc.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_KeyValueCollection_Compile_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Compile", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"nonEmpty": kvc.Compile() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Compile -- delegates to String", actual)
	})
}

func Test_KeyValueCollection_Join_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Join", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"nonEmpty": kvc.Join(",") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "Join -- non-empty", actual)
	})
}

func Test_KeyValueCollection_JoinKeys_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_JoinKeys", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1").Add("b", "2")

		// Act
		actual := args.Map{"nonEmpty": kvc.JoinKeys(",") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinKeys -- non-empty", actual)
	})
}

func Test_KeyValueCollection_JoinValues_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_JoinValues", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1").Add("b", "2")

		// Act
		actual := args.Map{"nonEmpty": kvc.JoinValues(",") != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "JoinValues -- non-empty", actual)
	})
}

// ── Hashmap / Map ───────────────────────────────────────────────────────────

func Test_KeyValueCollection_Hashmap_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Hashmap", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		hm := kvc.Hashmap()

		// Act
		actual := args.Map{"len": hm.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Hashmap -- 1 item", actual)
	})
}

func Test_KeyValueCollection_Hashmap_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Hashmap_Empty", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		hm := kvc.Hashmap()

		// Act
		actual := args.Map{"empty": hm.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Hashmap empty -- empty", actual)
	})
}

func Test_KeyValueCollection_Map_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Map", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"len": len(kvc.Map())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Map -- 1 item", actual)
	})
}

// ── JSON ────────────────────────────────────────────────────────────────────

func Test_KeyValueCollection_Json_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Json", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		j := kvc.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_KeyValueCollection_MarshalJSON_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_MarshalJSON", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		b, err := kvc.MarshalJSON()

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

func Test_KeyValueCollection_UnmarshalJSON_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_UnmarshalJSON", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		b, _ := kvc.MarshalJSON()
		kvc2 := &corestr.KeyValueCollection{}
		err := kvc2.UnmarshalJSON(b)

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": kvc2.Length(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 1,
		}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- success", actual)
	})
}

func Test_KeyValueCollection_UnmarshalJSON_EmptyArray_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_UnmarshalJSON_EmptyArray", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		err := kvc.UnmarshalJSON([]byte(`[]`))

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"len": kvc.Length(),
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"len": 0,
		}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON empty array -- empty", actual)
	})
}

func Test_KeyValueCollection_Serialize_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Serialize", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		b, err := kvc.Serialize()

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

func Test_KeyValueCollection_SerializeMust_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_SerializeMust", func() {
		// Arrange
		kvc := corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		b := kvc.SerializeMust()

		// Act
		actual := args.Map{"hasBytes": len(b) > 0}

		// Assert
		expected := args.Map{"hasBytes": true}
		expected.ShouldBeEqual(t, 0, "SerializeMust -- success", actual)
	})
}

func Test_KeyValueCollection_Deserialize_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Deserialize", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		var dest interface{}
		err := kvc.Deserialize(&dest)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Deserialize -- success", actual)
	})
}

func Test_KeyValueCollection_JsonModel_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_JsonModel", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")

		// Act
		actual := args.Map{"len": len(kvc.JsonModel())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "JsonModel -- 1 item", actual)
	})
}

func Test_KeyValueCollection_JsonModelAny_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_JsonModelAny", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{"notNil": kvc.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_KeyValueCollection_InterfaceCasts_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_InterfaceCasts", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}

		// Act
		actual := args.Map{
			"jsoner":   kvc.AsJsoner() != nil,
			"binder":   kvc.AsJsonContractsBinder() != nil,
			"injector": kvc.AsJsonParseSelfInjector() != nil,
		}

		// Assert
		expected := args.Map{
			"jsoner": true,
			"binder": true,
			"injector": true,
		}
		expected.ShouldBeEqual(t, 0, "Interface casts -- all non-nil", actual)
	})
}

func Test_KeyValueCollection_ParseInjectUsingJson_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_ParseInjectUsingJson", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		jr := kvc.JsonPtr()
		kvc2 := &corestr.KeyValueCollection{}
		_, err := kvc2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- success", actual)
	})
}

func Test_KeyValueCollection_JsonParseSelfInject_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_JsonParseSelfInject", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		jr := kvc.JsonPtr()
		kvc2 := &corestr.KeyValueCollection{}
		err := kvc2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

// ── Clear / Dispose ─────────────────────────────────────────────────────────

func Test_KeyValueCollection_Clear_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Clear", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Clear()

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_KeyValueCollection_Dispose_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_KVC_Dispose", func() {
		// Arrange
		kvc := &corestr.KeyValueCollection{}
		kvc.Add("a", "1")
		kvc.Dispose()

		// Act
		actual := args.Map{"len": kvc.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Dispose -- cleaned up", actual)
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// KeyAnyValuePair — Segment 6e
// ══════════════════════════════════════════════════════════════════════════════

