package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args")

// ══════════════════════════════════════════════════════════════════════════════
// Collection — Segment 3: Filter variants, Hashset, Sort, Contains, Join, CSV,
//   GetAllExcept, JSON, Clear/Dispose, Resize, remaining methods
// ══════════════════════════════════════════════════════════════════════════════

// ── FilterLock ───────────────────────────────────────────────────────────────

func Test_Collection_FilterPtr_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_FilterPtr_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "bb", "ccc")
		result := c.FilterPtr(func(s *string, i int) (*string, bool, bool) {
			return s, len(*s) > 1, false
		})

		// Act
		actual := args.Map{"len": len(*result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "FilterPtr -- keeps len>1", actual)
	})
}

func Test_Collection_FilterPtr_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_FilterPtr_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		result := c.FilterPtr(func(s *string, i int) (*string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"len": len(*result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "FilterPtr empty -- returns empty", actual)
	})
}

func Test_Collection_FilterPtr_Break_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_FilterPtr_Break_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")
		result := c.FilterPtr(func(s *string, i int) (*string, bool, bool) {
			return s, true, i == 0
		})

		// Act
		actual := args.Map{"len": len(*result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilterPtr break -- stops after first", actual)
	})
}

func Test_Collection_FilterPtrLock_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_FilterPtrLock_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "bb")
		result := c.FilterPtrLock(func(s *string, i int) (*string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"len": len(*result)}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "FilterPtrLock -- returns all", actual)
	})
}

func Test_Collection_FilterPtrLock_Empty_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_FilterPtrLock_Empty_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		result := c.FilterPtrLock(func(s *string, i int) (*string, bool, bool) {
			return s, true, false
		})

		// Act
		actual := args.Map{"len": len(*result)}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "FilterPtrLock empty -- returns empty", actual)
	})
}

func Test_Collection_FilterPtrLock_Break_FromSeg3(t *testing.T) {
	safeTest(t, "Test_Collection_FilterPtrLock_Break_FromSeg3", func() {
		// Arrange
		c := corestr.New.Collection.Cap(10)
		c.Adds("a", "b", "c")
		result := c.FilterPtrLock(func(s *string, i int) (*string, bool, bool) {
			return s, true, i == 0
		})

		// Act
		actual := args.Map{"len": len(*result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "FilterPtrLock break -- stops after first", actual)
	})
}

