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
	"github.com/alimtvnetwork/core/coretests/args")

// ══════════════════════════════════════════════════════════════════════════════
// SimpleSlice — Segment 4a: Add, Insert, Accessors, Contains, Index, Length
// ══════════════════════════════════════════════════════════════════════════════

func Test_SS_Add_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Add_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.Add("a").Add("b")

		// Act
		actual := args.Map{
			"len": s.Length(),
			"first": s.First(),
		}

		// Assert
		expected := args.Map{
			"len": 2,
			"first": "a",
		}
		expected.ShouldBeEqual(t, 0, "Add -- 2 items", actual)
	})
}

func Test_SS_AddSplit_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AddSplit_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.AddSplit("a,b,c", ",")

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "AddSplit -- 3 items from CSV", actual)
	})
}

func Test_SS_AddIf_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AddIf_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.AddIf(true, "yes").AddIf(false, "no")

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "AddIf -- only true added", actual)
	})
}

func Test_SS_Adds_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Adds_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.Adds("a", "b", "c")

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 3}
		expected.ShouldBeEqual(t, 0, "Adds -- 3 items", actual)
	})
}

func Test_SS_Adds_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Adds_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.Adds()

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Adds empty -- no change", actual)
	})
}

func Test_SS_Append_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Append_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.Append("a", "b")

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 2}
		expected.ShouldBeEqual(t, 0, "Append -- 2 items", actual)
	})
}

func Test_SS_Append_Empty_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_Append_Empty_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.Append()

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "Append empty -- no change", actual)
	})
}

func Test_SS_AppendFmt_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AppendFmt_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.AppendFmt("hello %s", "world")

		// Act
		actual := args.Map{"val": s.First()}

		// Assert
		expected := args.Map{"val": "hello world"}
		expected.ShouldBeEqual(t, 0, "AppendFmt -- formatted", actual)
	})
}

func Test_SS_AppendFmt_EmptySkip_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AppendFmt_EmptySkip_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.AppendFmt("")

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendFmt empty -- skipped", actual)
	})
}

func Test_SS_AppendFmtIf_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AppendFmtIf_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.AppendFmtIf(true, "v=%d", 42)
		s.AppendFmtIf(false, "skip=%d", 99)

		// Act
		actual := args.Map{
			"len": s.Length(),
			"val": s.First(),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"val": "v=42",
		}
		expected.ShouldBeEqual(t, 0, "AppendFmtIf -- only true", actual)
	})
}

func Test_SS_AppendFmtIf_EmptySkip_FromSeg4(t *testing.T) {
	safeTest(t, "Test_SS_AppendFmtIf_EmptySkip_FromSeg4", func() {
		// Arrange
		s := corestr.SimpleSlice{}
		s.AppendFmtIf(true, "")

		// Act
		actual := args.Map{"len": s.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "AppendFmtIf empty format -- skipped", actual)
	})
}

