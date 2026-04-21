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

