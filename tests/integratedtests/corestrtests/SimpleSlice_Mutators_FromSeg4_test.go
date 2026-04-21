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
// SimpleSlice — Segment 4a: Add, Insert, Accessors, Contains, Index, Length
// ══════════════════════════════════════════════════════════════════════════════

func Test_SS_InsertAt_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_InsertAt_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "c"}
		s.InsertAt(1, "b")

		// Act
		actual := args.Map{
			"len": s.Length(),
			"mid": s[1],
		}

		// Assert
		expected := args.Map{
			"len": 3,
			"mid": "b",
		}
		expected.ShouldBeEqual(t, 0, "InsertAt -- inserted in middle", actual)
	})
}

func Test_SS_InsertAt_OutOfBounds_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_InsertAt_OutOfBounds_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		s.InsertAt(-1, "x")
		s.InsertAt(99, "y")

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "InsertAt out of bounds -- no change", actual)
	})
}

func Test_SS_Sort_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Sort_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"c", "a", "b"}
		s.Sort()

		// Act
		actual := args.Map{"first": s.First()}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "Sort -- ascending", actual)
	})
}

func Test_SS_Reverse_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Reverse_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b", "c", "d"}
		s.Reverse()

		// Act
		actual := args.Map{
			"first": s.First(),
			"last": s.Last(),
		}

		// Assert
		expected := args.Map{
			"first": "d",
			"last": "a",
		}
		expected.ShouldBeEqual(t, 0, "Reverse -- reversed", actual)
	})
}

func Test_SS_Reverse_Two_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Reverse_Two_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}
		s.Reverse()

		// Act
		actual := args.Map{"first": s.First()}

		// Assert
		expected := args.Map{"first": "b"}
		expected.ShouldBeEqual(t, 0, "Reverse 2 -- swapped", actual)
	})
}

func Test_SS_Reverse_Single_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Reverse_Single_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		s.Reverse()

		// Act
		actual := args.Map{"first": s.First()}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "Reverse 1 -- unchanged", actual)
	})
}

func Test_SS_Clone_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Clone_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}
		c := s.Clone(true)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Clone deep -- same items", actual)
	})
}

func Test_SS_ClonePtr_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_ClonePtr_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a"}
		c := s.ClonePtr(true)

		// Act
		actual := args.Map{"len": c.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "ClonePtr -- same items", actual)
	})
}

func Test_SS_ClonePtr_Nil_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_ClonePtr_Nil_FromSeg4", func() {
		// Arrange
		var s *corestr.SimpleSlice

		// Act
		actual := args.Map{"nil": s.ClonePtr(true) == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "ClonePtr nil -- returns nil", actual)
	})
}

func Test_SS_Clear_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Clear_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{"a", "b"}
		s.Clear()

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Clear -- emptied", actual)
	})
}

func Test_SS_Clear_Nil_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Clear_Nil_FromSeg4", func() {
		// Arrange
		var s *corestr.SimpleSlice

		// Act
		actual := args.Map{"nil": s.Clear() == nil}

		// Assert
		expected := args.Map{"nil": true}
		expected.ShouldBeEqual(t, 0, "Clear nil -- returns nil", actual)
	})
}

