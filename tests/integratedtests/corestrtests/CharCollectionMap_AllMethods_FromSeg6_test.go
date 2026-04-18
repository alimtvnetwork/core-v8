package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// CharCollectionMap — Segment 6a
// ══════════════════════════════════════════════════════════════════════════════

func Test_CharCollectionMap_IsEmpty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEmpty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)

		// Act
		actual := args.Map{
			"empty": ccm.IsEmpty(),
			"hasItems": ccm.HasItems(),
		}

		// Assert
		expected := args.Map{
			"empty": true,
			"hasItems": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEmpty -- true on empty", actual)
	})
}

func Test_CharCollectionMap_Add_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Add", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple").Add("avocado").Add("banana")

		// Act
		actual := args.Map{
			"len": ccm.Length(),
			"allLen": ccm.AllLengthsSum(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"allLen": 3,
		}
		expected.ShouldBeEqual(t, 0, "Add -- 2 char groups, 3 total items", actual)
	})
}

func Test_CharCollectionMap_AddLock_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.AddLock("apple").AddLock("banana")

		// Act
		actual := args.Map{"len": ccm.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddLock -- 2 char groups", actual)
	})
}

func Test_CharCollectionMap_AddStrings_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddStrings", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.AddStrings("apple", "avocado", "banana")

		// Act
		actual := args.Map{
			"len": ccm.Length(),
			"allLen": ccm.AllLengthsSum(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"allLen": 3,
		}
		expected.ShouldBeEqual(t, 0, "AddStrings -- 2 groups 3 total", actual)
	})
}

func Test_CharCollectionMap_AddStrings_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddStrings_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.AddStrings()

		// Act
		actual := args.Map{"len": ccm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddStrings empty -- no change", actual)
	})
}

func Test_CharCollectionMap_GetChar_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetChar", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)

		// Act
		actual := args.Map{
			"char": ccm.GetChar("abc"),
			"empty": ccm.GetChar(""),
		}

		// Assert
		expected := args.Map{
			"char": byte('a'),
			"empty": byte(0),
		}
		expected.ShouldBeEqual(t, 0, "GetChar -- first char or empty", actual)
	})
}

func Test_CharCollectionMap_Has_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Has", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")

		// Act
		actual := args.Map{
			"has": ccm.Has("apple"),
			"miss": ccm.Has("banana"),
		}

		// Assert
		expected := args.Map{
			"has": true,
			"miss": false,
		}
		expected.ShouldBeEqual(t, 0, "Has -- found and missing", actual)
	})
}

func Test_CharCollectionMap_Has_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Has_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)

		// Act
		actual := args.Map{"has": ccm.Has("apple")}

		// Assert
		expected := args.Map{"has": false}
		expected.ShouldBeEqual(t, 0, "Has empty -- false", actual)
	})
}

func Test_CharCollectionMap_HasWithCollection_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HasWithCollection", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		has, coll := ccm.HasWithCollection("apple")

		// Act
		actual := args.Map{
			"has": has,
			"collNotNil": coll != nil,
		}

		// Assert
		expected := args.Map{
			"has": true,
			"collNotNil": true,
		}
		expected.ShouldBeEqual(t, 0, "HasWithCollection -- found", actual)
	})
}

func Test_CharCollectionMap_HasWithCollection_Miss_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HasWithCollection_Miss", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		has, coll := ccm.HasWithCollection("banana")

		// Act
		actual := args.Map{
			"has": has,
			"collNotNil": coll != nil,
		}

		// Assert
		expected := args.Map{
			"has": false,
			"collNotNil": true,
		}
		expected.ShouldBeEqual(t, 0, "HasWithCollection miss -- not found", actual)
	})
}

func Test_CharCollectionMap_HasWithCollection_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HasWithCollection_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		has, coll := ccm.HasWithCollection("apple")

		// Act
		actual := args.Map{
			"has": has,
			"collNotNil": coll != nil,
		}

		// Assert
		expected := args.Map{
			"has": false,
			"collNotNil": true,
		}
		expected.ShouldBeEqual(t, 0, "HasWithCollection empty -- not found", actual)
	})
}

func Test_CharCollectionMap_HasWithCollectionLock_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HasWithCollectionLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		has, coll := ccm.HasWithCollectionLock("apple")

		// Act
		actual := args.Map{
			"has": has,
			"collNotNil": coll != nil,
		}

		// Assert
		expected := args.Map{
			"has": true,
			"collNotNil": true,
		}
		expected.ShouldBeEqual(t, 0, "HasWithCollectionLock -- found", actual)
	})
}

func Test_CharCollectionMap_HasWithCollectionLock_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HasWithCollectionLock_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		has, coll := ccm.HasWithCollectionLock("apple")

		// Act
		actual := args.Map{
			"has": has,
			"collNotNil": coll != nil,
		}

		// Assert
		expected := args.Map{
			"has": false,
			"collNotNil": true,
		}
		expected.ShouldBeEqual(t, 0, "HasWithCollectionLock empty -- not found", actual)
	})
}

func Test_CharCollectionMap_HasWithCollectionLock_Miss_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HasWithCollectionLock_Miss", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		has, coll := ccm.HasWithCollectionLock("banana")

		// Act
		actual := args.Map{
			"has": has,
			"collNotNil": coll != nil,
		}

		// Assert
		expected := args.Map{
			"has": false,
			"collNotNil": true,
		}
		expected.ShouldBeEqual(t, 0, "HasWithCollectionLock miss -- not found", actual)
	})
}

func Test_CharCollectionMap_LengthOf_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_LengthOf", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple").Add("avocado")

		// Act
		actual := args.Map{
			"lenA": ccm.LengthOf(byte('a')),
			"lenZ": ccm.LengthOf(byte('z')),
		}

		// Assert
		expected := args.Map{
			"lenA": 2,
			"lenZ": 0,
		}
		expected.ShouldBeEqual(t, 0, "LengthOf -- 2 for 'a', 0 for 'z'", actual)
	})
}

func Test_CharCollectionMap_LengthOf_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_LengthOf_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)

		// Act
		actual := args.Map{"len": ccm.LengthOf(byte('a'))}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LengthOf empty -- 0", actual)
	})
}

func Test_CharCollectionMap_LengthOfLock_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_LengthOfLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")

		// Act
		actual := args.Map{
			"len": ccm.LengthOfLock(byte('a')),
			"miss": ccm.LengthOfLock(byte('z')),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"miss": 0,
		}
		expected.ShouldBeEqual(t, 0, "LengthOfLock -- found and missing", actual)
	})
}

func Test_CharCollectionMap_LengthOfLock_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_LengthOfLock_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)

		// Act
		actual := args.Map{"len": ccm.LengthOfLock(byte('a'))}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "LengthOfLock empty -- 0", actual)
	})
}

func Test_CharCollectionMap_LengthOfCollectionFromFirstChar_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_LengthOfCollectionFromFirstChar", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple").Add("avocado")

		// Act
		actual := args.Map{
			"len": ccm.LengthOfCollectionFromFirstChar("abc"),
			"miss": ccm.LengthOfCollectionFromFirstChar("xyz"),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"miss": 0,
		}
		expected.ShouldBeEqual(t, 0, "LengthOfCollectionFromFirstChar -- 2 and 0", actual)
	})
}

func Test_CharCollectionMap_AllLengthsSum_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AllLengthsSum", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple").Add("avocado").Add("banana")

		// Act
		actual := args.Map{"sum": ccm.AllLengthsSum()}

		// Assert
		expected := args.Map{"sum": 3}
		expected.ShouldBeEqual(t, 0, "AllLengthsSum -- 3", actual)
	})
}

func Test_CharCollectionMap_AllLengthsSumLock_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AllLengthsSumLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple").Add("banana")

		// Act
		actual := args.Map{"sum": ccm.AllLengthsSumLock()}

		// Assert
		expected := args.Map{"sum": 2}
		expected.ShouldBeEqual(t, 0, "AllLengthsSumLock -- 2", actual)
	})
}

func Test_CharCollectionMap_IsEmptyLock_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEmptyLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)

		// Act
		actual := args.Map{"empty": ccm.IsEmptyLock()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "IsEmptyLock -- true", actual)
	})
}

func Test_CharCollectionMap_LengthLock_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_LengthLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")

		// Act
		actual := args.Map{"len": ccm.LengthLock()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "LengthLock -- 1", actual)
	})
}

func Test_CharCollectionMap_GetMap_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetMap", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")

		// Act
		actual := args.Map{"notNil": ccm.GetMap() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetMap -- not nil", actual)
	})
}

func Test_CharCollectionMap_GetCopyMapLock_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetCopyMapLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")

		// Act
		actual := args.Map{"len": len(ccm.GetCopyMapLock())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "GetCopyMapLock -- 1", actual)
	})
}

func Test_CharCollectionMap_GetCopyMapLock_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetCopyMapLock_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)

		// Act
		actual := args.Map{"len": len(ccm.GetCopyMapLock())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "GetCopyMapLock empty -- 0", actual)
	})
}

func Test_CharCollectionMap_List_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_List", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple").Add("banana")

		// Act
		actual := args.Map{"len": len(ccm.List())}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "List -- 2 items", actual)
	})
}

func Test_CharCollectionMap_List_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_List_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)

		// Act
		actual := args.Map{"len": len(ccm.List())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "List empty -- 0", actual)
	})
}

func Test_CharCollectionMap_ListLock_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_ListLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")

		// Act
		actual := args.Map{"len": len(ccm.ListLock())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ListLock -- 1", actual)
	})
}

func Test_CharCollectionMap_SortedListAsc_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_SortedListAsc", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("banana").Add("apple")
		sorted := ccm.SortedListAsc()

		// Act
		actual := args.Map{"first": sorted[0]}

		// Assert
		expected := args.Map{"first": "apple"}
		expected.ShouldBeEqual(t, 0, "SortedListAsc -- sorted", actual)
	})
}

func Test_CharCollectionMap_SortedListAsc_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_SortedListAsc_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)

		// Act
		actual := args.Map{"len": len(ccm.SortedListAsc())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "SortedListAsc empty -- 0", actual)
	})
}

func Test_CharCollectionMap_String_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_String", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")

		// Act
		actual := args.Map{"nonEmpty": ccm.String() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "String -- non-empty", actual)
	})
}

func Test_CharCollectionMap_StringLock_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_StringLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")

		// Act
		actual := args.Map{"nonEmpty": ccm.StringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "StringLock -- non-empty", actual)
	})
}

func Test_CharCollectionMap_SummaryString_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_SummaryString", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")

		// Act
		actual := args.Map{"nonEmpty": ccm.SummaryString() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryString -- non-empty", actual)
	})
}

func Test_CharCollectionMap_SummaryStringLock_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_SummaryStringLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")

		// Act
		actual := args.Map{"nonEmpty": ccm.SummaryStringLock() != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "SummaryStringLock -- non-empty", actual)
	})
}

func Test_CharCollectionMap_Print_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Print", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		ccm.Print(false) // skip print
		ccm.Print(true)  // actual print
	})
}

func Test_CharCollectionMap_PrintLock_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_PrintLock", func() {
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		ccm.PrintLock(false)
		ccm.PrintLock(true)
	})
}

// ── IsEquals ────────────────────────────────────────────────────────────────

func Test_CharCollectionMap_IsEquals_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEquals", func() {
		// Arrange
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm1.Add("apple")
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm2.Add("apple")

		// Act
		actual := args.Map{
			"eq":   ccm1.IsEquals(ccm2),
			"same": ccm1.IsEquals(ccm1),
			"nil":  ccm1.IsEquals(nil),
		}

		// Assert
		expected := args.Map{
			"eq": true,
			"same": true,
			"nil": false,
		}
		expected.ShouldBeEqual(t, 0, "IsEquals -- various", actual)
	})
}

func Test_CharCollectionMap_IsEquals_BothEmpty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEquals_BothEmpty", func() {
		// Arrange
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(0, 4)

		// Act
		actual := args.Map{"eq": ccm1.IsEquals(ccm2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEquals both empty -- true", actual)
	})
}

func Test_CharCollectionMap_IsEquals_OneEmpty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEquals_OneEmpty", func() {
		// Arrange
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm1.Add("apple")
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(0, 4)

		// Act
		actual := args.Map{"eq": ccm1.IsEquals(ccm2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals one empty -- false", actual)
	})
}

func Test_CharCollectionMap_IsEquals_DiffLen_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEquals_DiffLen", func() {
		// Arrange
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm1.Add("apple")
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm2.Add("apple").Add("banana")

		// Act
		actual := args.Map{"eq": ccm1.IsEquals(ccm2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff len -- false", actual)
	})
}

func Test_CharCollectionMap_IsEquals_DiffItems_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEquals_DiffItems", func() {
		// Arrange
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm1.Add("apple")
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm2.Add("avocado")

		// Act
		actual := args.Map{"eq": ccm1.IsEquals(ccm2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals diff items -- false", actual)
	})
}

func Test_CharCollectionMap_IsEquals_MissingKey_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEquals_MissingKey", func() {
		// Arrange
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm1.Add("apple")
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm2.Add("banana")

		// Act
		actual := args.Map{"eq": ccm1.IsEquals(ccm2)}

		// Assert
		expected := args.Map{"eq": false}
		expected.ShouldBeEqual(t, 0, "IsEquals missing key -- false", actual)
	})
}

func Test_CharCollectionMap_IsEqualsLock_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEqualsLock", func() {
		// Arrange
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm1.Add("apple")
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm2.Add("apple")

		// Act
		actual := args.Map{"eq": ccm1.IsEqualsLock(ccm2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsLock -- true", actual)
	})
}

func Test_CharCollectionMap_IsEqualsCaseSensitive_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEqualsCaseSensitive", func() {
		// Arrange
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm1.Add("apple")
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm2.Add("apple")

		// Act
		actual := args.Map{"eq": ccm1.IsEqualsCaseSensitive(true, ccm2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsCaseSensitive -- true", actual)
	})
}

func Test_CharCollectionMap_IsEqualsCaseSensitiveLock_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_IsEqualsCaseSensitiveLock", func() {
		// Arrange
		ccm1 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm1.Add("apple")
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm2.Add("apple")

		// Act
		actual := args.Map{"eq": ccm1.IsEqualsCaseSensitiveLock(true, ccm2)}

		// Assert
		expected := args.Map{"eq": true}
		expected.ShouldBeEqual(t, 0, "IsEqualsCaseSensitiveLock -- true", actual)
	})
}

// ── AddSameStartingCharItems ────────────────────────────────────────────────

func Test_CharCollectionMap_AddSameStartingCharItems_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameStartingCharItems", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.AddSameStartingCharItems(byte('a'), []string{"apple", "avocado"}, false)

		// Act
		actual := args.Map{"len": ccm.AllLengthsSum()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddSameStartingCharItems -- 2 items", actual)
	})
}

func Test_CharCollectionMap_AddSameStartingCharItems_Existing_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameStartingCharItems_Existing", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		ccm.AddSameStartingCharItems(byte('a'), []string{"avocado"}, false)

		// Act
		actual := args.Map{"len": ccm.AllLengthsSum()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddSameStartingCharItems existing -- 2 items", actual)
	})
}

func Test_CharCollectionMap_AddSameStartingCharItems_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameStartingCharItems_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.AddSameStartingCharItems(byte('a'), []string{}, false)

		// Act
		actual := args.Map{"len": ccm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddSameStartingCharItems empty -- no change", actual)
	})
}

// ── AddHashmapsValues / AddHashmapsKeysValuesBoth ───────────────────────────

func Test_CharCollectionMap_AddHashmapsValues_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddHashmapsValues", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		hm := corestr.New.Hashmap.Cap(2)
		hm.Set("k", "apple")
		ccm.AddHashmapsValues(hm)

		// Act
		actual := args.Map{"has": ccm.Has("apple")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddHashmapsValues -- added values", actual)
	})
}

func Test_CharCollectionMap_AddHashmapsValues_Nil_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddHashmapsValues_Nil", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.AddHashmapsValues(nil)

		// Act
		actual := args.Map{"len": ccm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashmapsValues nil -- no change", actual)
	})
}

func Test_CharCollectionMap_AddHashmapsKeysValuesBoth_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddHashmapsKeysValuesBoth", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		hm := corestr.New.Hashmap.Cap(2)
		hm.Set("key", "val")
		ccm.AddHashmapsKeysValuesBoth(hm)

		// Act
		actual := args.Map{"has": ccm.AllLengthsSum() > 0}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysValuesBoth -- added", actual)
	})
}

func Test_CharCollectionMap_AddHashmapsKeysValuesBoth_Nil_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddHashmapsKeysValuesBoth_Nil", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.AddHashmapsKeysValuesBoth(nil)

		// Act
		actual := args.Map{"len": ccm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysValuesBoth nil -- no change", actual)
	})
}

func Test_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddHashmapsKeysOrValuesBothUsingFilter", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		hm := corestr.New.Hashmap.Cap(2)
		hm.Set("key", "val")
		ccm.AddHashmapsKeysOrValuesBothUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) {
				return pair.Value, true, false
			},
			hm,
		)

		// Act
		actual := args.Map{"has": ccm.Has("val")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysOrValuesBothUsingFilter -- added", actual)
	})
}

func Test_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter_Nil_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddHashmapsKeysOrValuesBothUsingFilter_Nil", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.AddHashmapsKeysOrValuesBothUsingFilter(nil, nil)

		// Act
		actual := args.Map{"len": ccm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysOrValuesBothUsingFilter nil -- no change", actual)
	})
}

func Test_CharCollectionMap_AddHashmapsKeysOrValuesBothUsingFilter_Break_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddHashmapsKeysOrValuesBothUsingFilter_Break", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		hm := corestr.New.Hashmap.Cap(4)
		hm.Set("a", "1")
		hm.Set("b", "2")
		ccm.AddHashmapsKeysOrValuesBothUsingFilter(
			func(pair corestr.KeyValuePair) (string, bool, bool) {
				return pair.Value, true, true
			},
			hm,
		)

		// Act
		actual := args.Map{"hasItems": ccm.AllLengthsSum() > 0}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "AddHashmapsKeysOrValuesBothUsingFilter break -- stops", actual)
	})
}

// ── AddCollectionItems / AddCharHashsetMap ──────────────────────────────────

func Test_CharCollectionMap_AddCollectionItems_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddCollectionItems", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		c := corestr.New.Collection.Strings([]string{"apple", "banana"})
		ccm.AddCollectionItems(c)

		// Act
		actual := args.Map{"sum": ccm.AllLengthsSum()}

		// Assert
		expected := args.Map{"sum": 2}
		expected.ShouldBeEqual(t, 0, "AddCollectionItems -- 2 items", actual)
	})
}

func Test_CharCollectionMap_AddCollectionItems_Nil_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddCollectionItems_Nil", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.AddCollectionItems(nil)

		// Act
		actual := args.Map{"len": ccm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCollectionItems nil -- no change", actual)
	})
}

func Test_CharCollectionMap_AddCharHashsetMap_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddCharHashsetMap", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		chm := corestr.New.CharHashsetMap.Cap(4, 4)
		chm.Add("apple")
		ccm.AddCharHashsetMap(chm)

		// Act
		actual := args.Map{"has": ccm.Has("apple")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddCharHashsetMap -- added", actual)
	})
}

func Test_CharCollectionMap_AddCharHashsetMap_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddCharHashsetMap_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		chm := corestr.New.CharHashsetMap.Cap(0, 4)
		ccm.AddCharHashsetMap(chm)

		// Act
		actual := args.Map{"len": ccm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddCharHashsetMap empty -- no change", actual)
	})
}

// ── GetCollection / GetCollectionLock ────────────────────────────────────────

func Test_CharCollectionMap_GetCollection_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetCollection", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		coll := ccm.GetCollection("a", false)

		// Act
		actual := args.Map{"notNil": coll != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetCollection -- found", actual)
	})
}

func Test_CharCollectionMap_GetCollection_AddNew_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetCollection_AddNew", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		coll := ccm.GetCollection("z", true)

		// Act
		actual := args.Map{"notNil": coll != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetCollection addNew -- created", actual)
	})
}

func Test_CharCollectionMap_GetCollection_Miss_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetCollection_Miss", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		coll := ccm.GetCollection("z", false)

		// Act
		actual := args.Map{"nil": coll == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "GetCollection miss -- nil", actual)
	})
}

func Test_CharCollectionMap_GetCollectionLock_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetCollectionLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		coll := ccm.GetCollectionLock("a", false)

		// Act
		actual := args.Map{"notNil": coll != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetCollectionLock -- found", actual)
	})
}

func Test_CharCollectionMap_GetCollectionByChar_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetCollectionByChar", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")

		// Act
		actual := args.Map{"notNil": ccm.GetCollectionByChar(byte('a')) != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "GetCollectionByChar -- found", actual)
	})
}

// ── AddSameCharsCollection ──────────────────────────────────────────────────

func Test_CharCollectionMap_AddSameCharsCollection_Existing_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameCharsCollection_Existing", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		c := corestr.New.Collection.Strings([]string{"avocado"})
		result := ccm.AddSameCharsCollection("a", c)

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection existing -- merged", actual)
	})
}

func Test_CharCollectionMap_AddSameCharsCollection_NilCollection_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameCharsCollection_NilCollection", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		result := ccm.AddSameCharsCollection("a", nil)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection nil -- existing returned", actual)
	})
}

func Test_CharCollectionMap_AddSameCharsCollection_NewChar_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameCharsCollection_NewChar", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		c := corestr.New.Collection.Strings([]string{"banana"})
		result := ccm.AddSameCharsCollection("b", c)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection new char -- added", actual)
	})
}

func Test_CharCollectionMap_AddSameCharsCollection_NewCharNilColl_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameCharsCollection_NewCharNilColl", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		result := ccm.AddSameCharsCollection("z", nil)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollection new char nil coll -- created empty", actual)
	})
}

func Test_CharCollectionMap_AddSameCharsCollectionLock_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameCharsCollectionLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		c := corestr.New.Collection.Strings([]string{"avocado"})
		result := ccm.AddSameCharsCollectionLock("a", c)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock -- merged", actual)
	})
}

func Test_CharCollectionMap_AddSameCharsCollectionLock_NewCharNil_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameCharsCollectionLock_NewCharNil", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		result := ccm.AddSameCharsCollectionLock("z", nil)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock new nil -- created", actual)
	})
}

func Test_CharCollectionMap_AddSameCharsCollectionLock_NewChar_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameCharsCollectionLock_NewChar", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		c := corestr.New.Collection.Strings([]string{"banana"})
		result := ccm.AddSameCharsCollectionLock("b", c)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock new char -- added", actual)
	})
}

func Test_CharCollectionMap_AddSameCharsCollectionLock_ExistingNil_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddSameCharsCollectionLock_ExistingNil", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		result := ccm.AddSameCharsCollectionLock("a", nil)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "AddSameCharsCollectionLock existing nil -- returned existing", actual)
	})
}

// ── Hashset conversions ─────────────────────────────────────────────────────

func Test_CharCollectionMap_HashsetByChar_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetByChar", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		hs := ccm.HashsetByChar(byte('a'))

		// Act
		actual := args.Map{
			"notNil": hs != nil,
			"has": hs.Has("apple"),
		}

		// Assert
		expected := args.Map{
			"notNil": true,
			"has": true,
		}
		expected.ShouldBeEqual(t, 0, "HashsetByChar -- found", actual)
	})
}

func Test_CharCollectionMap_HashsetByChar_Miss_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetByChar_Miss", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)

		// Act
		actual := args.Map{"nil": ccm.HashsetByChar(byte('z')) == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByChar miss -- nil", actual)
	})
}

func Test_CharCollectionMap_HashsetByCharLock_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetByCharLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		hs := ccm.HashsetByCharLock(byte('a'))

		// Act
		actual := args.Map{"has": hs.Has("apple")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HashsetByCharLock -- found", actual)
	})
}

func Test_CharCollectionMap_HashsetByCharLock_Miss_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetByCharLock_Miss", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		hs := ccm.HashsetByCharLock(byte('z'))

		// Act
		actual := args.Map{"empty": hs.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetByCharLock miss -- empty", actual)
	})
}

func Test_CharCollectionMap_HashsetByStringFirstChar_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetByStringFirstChar", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		hs := ccm.HashsetByStringFirstChar("avocado")

		// Act
		actual := args.Map{"notNil": hs != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByStringFirstChar -- found", actual)
	})
}

func Test_CharCollectionMap_HashsetByStringFirstCharLock_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetByStringFirstCharLock", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		hs := ccm.HashsetByStringFirstCharLock("avocado")

		// Act
		actual := args.Map{"notNil": hs != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "HashsetByStringFirstCharLock -- found", actual)
	})
}

func Test_CharCollectionMap_HashsetsCollection_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetsCollection", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple").Add("banana")
		hsc := ccm.HashsetsCollection()

		// Act
		actual := args.Map{"len": hsc.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection -- 2 hashsets", actual)
	})
}

func Test_CharCollectionMap_HashsetsCollection_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetsCollection_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		hsc := ccm.HashsetsCollection()

		// Act
		actual := args.Map{"empty": hsc.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollection empty -- empty", actual)
	})
}

func Test_CharCollectionMap_HashsetsCollectionByChars_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetsCollectionByChars", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple").Add("banana")
		hsc := ccm.HashsetsCollectionByChars(byte('a'))

		// Act
		actual := args.Map{"len": hsc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByChars -- 1 hashset", actual)
	})
}

func Test_CharCollectionMap_HashsetsCollectionByChars_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetsCollectionByChars_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		hsc := ccm.HashsetsCollectionByChars(byte('a'))

		// Act
		actual := args.Map{"empty": hsc.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByChars empty -- empty", actual)
	})
}

func Test_CharCollectionMap_HashsetsCollectionByStringFirstChar_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetsCollectionByStringFirstChar", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		hsc := ccm.HashsetsCollectionByStringFirstChar("avocado")

		// Act
		actual := args.Map{"len": hsc.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByStringFirstChar -- 1", actual)
	})
}

func Test_CharCollectionMap_HashsetsCollectionByStringFirstChar_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_HashsetsCollectionByStringFirstChar_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		hsc := ccm.HashsetsCollectionByStringFirstChar("a")

		// Act
		actual := args.Map{"empty": hsc.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashsetsCollectionByStringFirstChar empty -- empty", actual)
	})
}

// ── GetCharsGroups ──────────────────────────────────────────────────────────

func Test_CharCollectionMap_GetCharsGroups_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetCharsGroups", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		result := ccm.GetCharsGroups([]string{"apple", "banana"})

		// Act
		actual := args.Map{"len": result.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "GetCharsGroups -- 2 groups", actual)
	})
}

func Test_CharCollectionMap_GetCharsGroups_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_GetCharsGroups_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		result := ccm.GetCharsGroups([]string{})

		// Act
		actual := args.Map{"same": result == ccm}

		// Assert
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "GetCharsGroups empty -- returns self", actual)
	})
}

// ── Resize / AddLength ──────────────────────────────────────────────────────

func Test_CharCollectionMap_Resize_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Resize", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(2, 4)
		ccm.Add("apple")
		ccm.Resize(10)

		// Act
		actual := args.Map{"has": ccm.Has("apple")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "Resize -- preserved items", actual)
	})
}

func Test_CharCollectionMap_Resize_SmallerThanLen_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Resize_SmallerThanLen", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple").Add("banana")
		ccm.Resize(1)

		// Act
		actual := args.Map{"len": ccm.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Resize smaller -- no change", actual)
	})
}

func Test_CharCollectionMap_AddLength_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddLength", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(2, 4)
		ccm.Add("apple")
		ccm.AddLength(10, 20)

		// Act
		actual := args.Map{"has": ccm.Has("apple")}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "AddLength -- preserved", actual)
	})
}

func Test_CharCollectionMap_AddLength_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_AddLength_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(2, 4)
		ccm.AddLength()

		// Act
		actual := args.Map{"len": ccm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AddLength empty -- no change", actual)
	})
}

// ── JSON ────────────────────────────────────────────────────────────────────

func Test_CharCollectionMap_Json_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Json", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		j := ccm.Json()

		// Act
		actual := args.Map{"noErr": !j.HasError()}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "Json -- no error", actual)
	})
}

func Test_CharCollectionMap_MarshalJSON_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_MarshalJSON", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		b, err := ccm.MarshalJSON()

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

func Test_CharCollectionMap_UnmarshalJSON_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_UnmarshalJSON", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		b, _ := ccm.MarshalJSON()
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		err := ccm2.UnmarshalJSON(b)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON -- success", actual)
	})
}

func Test_CharCollectionMap_UnmarshalJSON_Invalid_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_UnmarshalJSON_Invalid", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		err := ccm.UnmarshalJSON([]byte(`invalid`))

		// Act
		actual := args.Map{"hasErr": err != nil}

		// Assert
		expected := args.Map{"hasErr": true}
		expected.ShouldBeEqual(t, 0, "UnmarshalJSON invalid -- error", actual)
	})
}

func Test_CharCollectionMap_JsonModel_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_JsonModel", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")

		// Act
		actual := args.Map{"notNil": ccm.JsonModel() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModel -- non-nil", actual)
	})
}

func Test_CharCollectionMap_JsonModelAny_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_JsonModelAny", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)

		// Act
		actual := args.Map{"notNil": ccm.JsonModelAny() != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "JsonModelAny -- non-nil", actual)
	})
}

func Test_CharCollectionMap_InterfaceCasts_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_InterfaceCasts", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)

		// Act
		actual := args.Map{
			"jsoner":   ccm.AsJsoner() != nil,
			"binder":   ccm.AsJsonContractsBinder() != nil,
			"injector": ccm.AsJsonParseSelfInjector() != nil,
			"marsh":    ccm.AsJsonMarshaller() != nil,
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

func Test_CharCollectionMap_ParseInjectUsingJson_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_ParseInjectUsingJson", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		jr := ccm.JsonPtr()
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		_, err := ccm2.ParseInjectUsingJson(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJson -- success", actual)
	})
}

func Test_CharCollectionMap_ParseInjectUsingJsonMust_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_ParseInjectUsingJsonMust", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		jr := ccm.JsonPtr()
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		result := ccm2.ParseInjectUsingJsonMust(jr)

		// Act
		actual := args.Map{"notNil": result != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "ParseInjectUsingJsonMust -- success", actual)
	})
}

func Test_CharCollectionMap_JsonParseSelfInject_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_JsonParseSelfInject", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		jr := ccm.JsonPtr()
		ccm2 := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		err := ccm2.JsonParseSelfInject(jr)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "JsonParseSelfInject -- success", actual)
	})
}

// ── Clear / Dispose ─────────────────────────────────────────────────────────

func Test_CharCollectionMap_Clear_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Clear", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		ccm.Clear()

		// Act
		actual := args.Map{"len": ccm.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_CharCollectionMap_Clear_Empty_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Clear_Empty", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(0, 4)
		result := ccm.Clear()

		// Act
		actual := args.Map{"same": result == ccm}

		// Assert
		expected := args.Map{"same": true}
		expected.ShouldBeEqual(t, 0, "Clear empty -- returns self", actual)
	})
}

func Test_CharCollectionMap_Dispose_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Dispose", func() {
		// Arrange
		ccm := corestr.New.CharCollectionMap.CapSelfCap(4, 4)
		ccm.Add("apple")
		ccm.Dispose()

		// Act
		actual := args.Map{"empty": ccm.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "Dispose -- cleaned up", actual)
	})
}

func Test_CharCollectionMap_Dispose_Nil_FromSeg6(t *testing.T) {
	safeTest(t, "Test_Seg6_CCM_Dispose_Nil", func() {
		var ccm *corestr.CharCollectionMap
		ccm.Dispose() // should not panic
	})
}
