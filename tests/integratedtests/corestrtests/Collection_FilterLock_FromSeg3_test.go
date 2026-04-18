package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Segment 3: Filter variants, Hashset, Sort, Contains, Join, CSV,
//   GetAllExcept, JSON, Clear/Dispose, Resize, remaining methods
// ══════════════════════════════════════════════════════════════════════════════

// ── FilterLock ───────────────────────────────────────────────────────────────

func Test_Collection_FilterLock_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_FilterLock_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "bb", "ccc")
		filtered := c.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, len(s) > 1, false
		})

		// Act
		actual := args.Map{"len": len(filtered)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "FilterLock -- keeps len>1", actual)
	})
}

func Test_Collection_FilterLock_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_FilterLock_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		filtered := c.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"len": len(filtered)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "FilterLock empty -- returns empty", actual)
	})
}

func Test_Collection_FilterLock_Break_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_FilterLock_Break_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")
		filtered := c.FilterLock(func(s string, i int) (string, bool, bool) {
			return s, true, i == 0
		})

		// Act
		actual := args.Map{"len": len(filtered)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilterLock break -- stops after first", actual)
	})
}

func Test_Collection_FilteredCollectionLock_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_FilteredCollectionLock_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("x", "yy")
		fc := c.FilteredCollectionLock(func(s string, i int) (string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"len": fc.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "FilteredCollectionLock -- returns new collection", actual)
	})
}

