package corestrtests

import (
	"errors"
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// AllIndividualStringsOfStringsLength
// ══════════════════════════════════════════════════════════════════════════════

func Test_Collection_BasicOps_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_BasicOps_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a").Add("b").Add("c")

		// Act
		actual := args.Map{
			"len":      c.Length(),
			"count":    c.Count(),
			"hasAny":   c.HasAnyItem(),
			"hasItems": c.HasItems(),
			"isEmpty":  c.IsEmpty(),
			"lastIdx":  c.LastIndex(),
			"hasIdx0":  c.HasIndex(0),
			"hasIdx5":  c.HasIndex(5),
		}

		// Assert
		expected := args.Map{
			"len":      3,
			"count":    3,
			"hasAny":   true,
			"hasItems": true,
			"isEmpty":  false,
			"lastIdx":  2,
			"hasIdx0":  true,
			"hasIdx5":  false,
		}
		expected.ShouldBeEqual(t, 0, "Collection basic ops -- 3 items", actual)
	})
}

func Test_Collection_NilLength_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_NilLength_FromSeg1", func() {
		// Arrange
		var c *corestr.Collection

		// Act
		actual := args.Map{
			"len": c.Length(),
			"empty": c.IsEmpty(),
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"empty": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection nil receiver -- length 0", actual)
	})
}

func Test_Collection_Capacity_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_Capacity_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)

		// Act
		actual := args.Map{"cap": c.Capacity() >= 10}

		// Assert
		expected := args.Map{"cap": true}
		expected.ShouldBeEqual(t, 0, "Collection Capacity -- at least 10", actual)
	})
}

func Test_Collection_RemoveAt_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_RemoveAt_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a").Add("b").Add("c")
		ok := c.RemoveAt(1)

		// Act
		actual := args.Map{
			"ok": ok,
			"len": c.Length(),
			"first": c.ListStrings()[0],
		}

		// Assert
		expected := args.Map{
			"ok": true,
			"len": 2,
			"first": "a",
		}
		expected.ShouldBeEqual(t, 0, "Collection RemoveAt -- middle removed", actual)
	})
}

func Test_Collection_RemoveAt_OutOfBounds_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_RemoveAt_OutOfBounds_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a")

		// Act
		actual := args.Map{
			"neg": c.RemoveAt(-1),
			"over": c.RemoveAt(5),
		}

		// Assert
		expected := args.Map{
			"neg": false,
			"over": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection RemoveAt -- out of bounds false", actual)
	})
}

func Test_Collection_ListStringsPtr_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_ListStringsPtr_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a")

		// Act
		actual := args.Map{"len": len(c.ListStringsPtr())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection ListStringsPtr -- returns items", actual)
	})
}

func Test_Collection_AddNonEmpty_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_AddNonEmpty_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddNonEmpty("a").AddNonEmpty("").AddNonEmpty("b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddNonEmpty -- skips empty", actual)
	})
}

func Test_Collection_AddNonEmptyWhitespace_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_AddNonEmptyWhitespace_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddNonEmptyWhitespace("a").AddNonEmptyWhitespace("   ").AddNonEmptyWhitespace("b")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddNonEmptyWhitespace -- skips whitespace", actual)
	})
}

func Test_Collection_AddError_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_AddError_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddError(errors.New("err1")).AddError(nil).AddError(errors.New("err2"))

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddError -- skips nil error", actual)
	})
}

func Test_Collection_AsError_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_AsError_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("err1").Add("err2")
		err := c.AsError("; ")

		// Act
		actual := args.Map{"notNil": err != nil}

		// Assert
		expected := args.Map{"notNil": true}
		expected.ShouldBeEqual(t, 0, "Collection AsError -- non-nil error", actual)
	})
}

func Test_Collection_AsErrorEmpty_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_AsErrorEmpty_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)

		// Act
		actual := args.Map{
			"nil": c.AsError("; ") == nil,
			"defNil": c.AsDefaultError() == nil,
		}

		// Assert
		expected := args.Map{
			"nil": true,
			"defNil": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection AsError -- nil when empty", actual)
	})
}

func Test_Collection_AddIf_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_AddIf_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddIf(true, "yes").AddIf(false, "no")

		// Act
		actual := args.Map{
			"len": c.Length(),
			"first": c.ListStrings()[0],
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"first": "yes",
		}
		expected.ShouldBeEqual(t, 0, "Collection AddIf -- only true added", actual)
	})
}

func Test_Collection_AddIfMany_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_AddIfMany_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddIfMany(true, "a", "b").AddIfMany(false, "c", "d")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddIfMany -- only true batch added", actual)
	})
}

func Test_Collection_AddFunc_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_AddFunc_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddFunc(func() string { return "hello" })

		// Act
		actual := args.Map{
			"len": c.Length(),
			"val": c.ListStrings()[0],
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"val": "hello",
		}
		expected.ShouldBeEqual(t, 0, "Collection AddFunc -- adds func result", actual)
	})
}

func Test_Collection_AddFuncErr_Success_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_AddFuncErr_Success_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddFuncErr(
			func() (string, error) { return "ok", nil },

		// Assert
			func(err error) { actual := args.Map{"errCalled": true}; expected := args.Map{"errCalled": false}; expected.ShouldBeEqual(t, 0, "error handler should not be called", actual) },
		)

		// Act
		actual := args.Map{"len": c.Length()}
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection AddFuncErr -- success path", actual)
	})
}

func Test_Collection_AddFuncErr_Error_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_AddFuncErr_Error_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		called := false
		c.AddFuncErr(
			func() (string, error) { return "", errors.New("fail") },
			func(err error) { called = true },
		)

		// Act
		actual := args.Map{
			"len": c.Length(),
			"called": called,
		}

		// Assert
		expected := args.Map{
			"len": 0,
			"called": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection AddFuncErr -- error path", actual)
	})
}

func Test_Collection_EachItemSplitBy_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_EachItemSplitBy_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a,b").Add("c,d")
		result := c.EachItemSplitBy(",")

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 4}
		expected.ShouldBeEqual(t, 0, "Collection EachItemSplitBy -- 4 items", actual)
	})
}

func Test_Collection_Adds_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_Adds_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b", "c")

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection Adds -- 3 items", actual)
	})
}

func Test_Collection_AddStrings_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_AddStrings_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddStrings([]string{"a", "b"})

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddStrings -- 2 items", actual)
	})
}

func Test_Collection_AddCollection_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_AddCollection_FromSeg1", func() {
		// Arrange
		c1 := corestr.New.Collection.Cap(5)
		c1.Add("a")
		c2 := corestr.New.Collection.Cap(5)
		c2.Add("b").Add("c")
		c1.AddCollection(c2)

		// Act
		actual := args.Map{"len": c1.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Collection AddCollection -- merged", actual)
	})
}

func Test_Collection_AddCollectionEmpty_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_AddCollectionEmpty_FromSeg1", func() {
		// Arrange
		c1 := corestr.New.Collection.Cap(5)
		c1.Add("a")
		c2 := corestr.New.Collection.Cap(5)
		c1.AddCollection(c2)

		// Act
		actual := args.Map{"len": c1.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection AddCollection empty -- no change", actual)
	})
}

func Test_Collection_AddCollections_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_AddCollections_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c1 := corestr.New.Collection.Cap(5)
		c1.Add("a")
		c2 := corestr.New.Collection.Cap(5)
		c2.Add("b")
		empty := corestr.New.Collection.Cap(5)
		c.AddCollections(c1, empty, c2)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddCollections -- skips empty", actual)
	})
}

func Test_Collection_LockMethods_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_LockMethods_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddLock("a")
		c.AddsLock("b", "c")

		// Act
		actual := args.Map{
			"len": c.LengthLock(),
			"emptyLock": c.IsEmptyLock(),
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"emptyLock": false,
		}
		expected.ShouldBeEqual(t, 0, "Collection lock methods -- 3 items", actual)
	})
}

func Test_Collection_IsEquals_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_IsEquals_FromSeg1", func() {
		// Arrange
		c1 := corestr.New.Collection.Cap(5)
		c1.Adds("a", "b")
		c2 := corestr.New.Collection.Cap(5)
		c2.Adds("a", "b")
		c3 := corestr.New.Collection.Cap(5)
		c3.Adds("a", "c")

		// Act
		actual := args.Map{
			"eq":   c1.IsEquals(c2),
			"neq":  c1.IsEquals(c3),
		}

		// Assert
		expected := args.Map{
			"eq":   true,
			"neq":  false,
		}
		expected.ShouldBeEqual(t, 0, "Collection IsEquals -- equal and not equal", actual)
	})
}

func Test_Collection_IsEqualsInsensitive_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_IsEqualsInsensitive_FromSeg1", func() {
		// Arrange
		c1 := corestr.New.Collection.Cap(5)
		c1.Adds("Hello", "World")
		c2 := corestr.New.Collection.Cap(5)
		c2.Adds("hello", "world")

		// Act
		actual := args.Map{
			"sensitive":   c1.IsEqualsWithSensitive(true, c2),
			"insensitive": c1.IsEqualsWithSensitive(false, c2),
		}

		// Assert
		expected := args.Map{
			"sensitive":   false,
			"insensitive": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection IsEqualsWithSensitive -- case comparison", actual)
	})
}

func Test_Collection_IsEqualsEdge_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_IsEqualsEdge_FromSeg1", func() {
		// Arrange
		var nilC *corestr.Collection
		emptyC := corestr.New.Collection.Cap(0)
		c := corestr.New.Collection.Cap(5)
		c.Add("a")

		// Act
		actual := args.Map{
			"bothNil":   nilC.IsEquals(nil),
			"nilVsNon":  nilC.IsEquals(c),
			"emptyBoth": emptyC.IsEquals(corestr.New.Collection.Cap(0)),
			"diffLen":   c.IsEquals(emptyC),
		}

		// Assert
		expected := args.Map{
			"bothNil":   true,
			"nilVsNon":  false,
			"emptyBoth": true,
			"diffLen":   false,
		}
		expected.ShouldBeEqual(t, 0, "Collection IsEquals edge -- nil and empty", actual)
	})
}

func Test_Collection_ToError_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_ToError_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("err1").Add("err2")
		err := c.ToError("; ")
		defErr := c.ToDefaultError()

		// Act
		actual := args.Map{
			"notNil": err != nil,
			"defNotNil": defErr != nil,
		}

		// Assert
		expected := args.Map{
			"notNil": true,
			"defNotNil": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection ToError -- non-nil", actual)
	})
}

func Test_Collection_ConcatNew_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_ConcatNew_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")
		c2 := c.ConcatNew(0, "c", "d")

		// Act
		actual := args.Map{
			"origLen": c.Length(),
			"newLen": c2.Length(),
		}

		// Assert
		expected := args.Map{
			"origLen": 2,
			"newLen": 4,
		}
		expected.ShouldBeEqual(t, 0, "Collection ConcatNew -- new collection with all items", actual)
	})
}

func Test_Collection_ConcatNewEmpty_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_ConcatNewEmpty_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Adds("a", "b")
		c2 := c.ConcatNew(0)

		// Act
		actual := args.Map{"newLen": c2.Length()}

		// Assert
		expected := args.Map{"newLen": 2}
		expected.ShouldBeEqual(t, 0, "Collection ConcatNew -- empty adds returns copy", actual)
	})
}

func Test_Collection_JsonString_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_JsonString_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.Add("a")
		s := c.JsonString()
		s2 := c.JsonStringMust()
		s3 := c.StringJSON()

		// Act
		actual := args.Map{
			"nonEmpty": s != "",
			"eq": s == s2,
			"eq2": s == s3,
		}

		// Assert
		expected := args.Map{
			"nonEmpty": true,
			"eq": true,
			"eq2": true,
		}
		expected.ShouldBeEqual(t, 0, "Collection JsonString -- all variants match", actual)
	})
}

func Test_Collection_AddHashmapsValues_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_AddHashmapsValues_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("k1", "v1")
		h.AddOrUpdate("k2", "v2")
		c.AddHashmapsValues(h)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddHashmapsValues -- 2 values added", actual)
	})
}

func Test_Collection_AddHashmapsValuesNil_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_AddHashmapsValuesNil_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddHashmapsValues(nil, nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Collection AddHashmapsValues nil -- no items", actual)
	})
}

func Test_Collection_AddHashmapsKeys_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_AddHashmapsKeys_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		h := corestr.New.Hashmap.Cap(5)
		h.AddOrUpdate("k1", "v1")
		h.AddOrUpdate("k2", "v2")
		c.AddHashmapsKeys(h)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Collection AddHashmapsKeys -- 2 keys added", actual)
	})
}

func Test_Collection_AddHashmapsKeysNil_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_AddHashmapsKeysNil_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c.AddHashmapsKeys(nil, nil)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Collection AddHashmapsKeys nil -- no items", actual)
	})
}

func Test_Collection_AddPointerCollectionsLock_FromSeg1(t *testing.T) {
	safeTest(t, "Test_Collection_AddPointerCollectionsLock_FromSeg1", func() {
		// Arrange
		c := corestr.New.Collection.Cap(5)
		c2 := corestr.New.Collection.Cap(5)
		c2.Add("a")
		c.AddPointerCollectionsLock(c2)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "Collection AddPointerCollectionsLock -- 1 item", actual)
	})
}

