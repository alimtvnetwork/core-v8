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
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════
// S11a — SimpleSlice.go Lines 1-600 — Add, Query, Join, Wrap
// ══════════════════════════════════════════════════════════════

func Test_SimpleSlice_01_SimpleSlice_Add_FromS11a(t *testing.T) {
	safeTest(t, "Test_01_SimpleSlice_Add", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.Add("a")

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_02_SimpleSlice_AddSplit_FromS11a(t *testing.T) {
	safeTest(t, "Test_02_SimpleSlice_AddSplit", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddSplit("a,b,c", ",")

		// Assert
		actual := args.Map{"result": ss.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_SimpleSlice_03_SimpleSlice_AddIf_FromS11a(t *testing.T) {
	safeTest(t, "Test_03_SimpleSlice_AddIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddIf(true, "yes")
		ss.AddIf(false, "no")

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_04_SimpleSlice_Adds_FromS11a(t *testing.T) {
	safeTest(t, "Test_04_SimpleSlice_Adds", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.Adds("a", "b")

		// Assert
		actual := args.Map{"result": ss.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SimpleSlice_05_SimpleSlice_Adds_Empty_FromS11a(t *testing.T) {
	safeTest(t, "Test_05_SimpleSlice_Adds_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.Adds()

		// Assert
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_SimpleSlice_06_SimpleSlice_Append_FromS11a(t *testing.T) {
	safeTest(t, "Test_06_SimpleSlice_Append", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.Append("a")

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_07_SimpleSlice_Append_Empty_FromS11a(t *testing.T) {
	safeTest(t, "Test_07_SimpleSlice_Append_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.Append()

		// Assert
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_SimpleSlice_08_SimpleSlice_AppendFmt_FromS11a(t *testing.T) {
	safeTest(t, "Test_08_SimpleSlice_AppendFmt", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AppendFmt("hello %s", "world")

		// Assert
		actual := args.Map{"result": ss.Length() != 1 || ss.First() != "hello world"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'hello world'", actual)
	})
}

func Test_SimpleSlice_09_SimpleSlice_AppendFmt_EmptySkip_FromS11a(t *testing.T) {
	safeTest(t, "Test_09_SimpleSlice_AppendFmt_EmptySkip", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AppendFmt("")

		// Assert
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_SimpleSlice_10_SimpleSlice_AppendFmtIf_FromS11a(t *testing.T) {
	safeTest(t, "Test_10_SimpleSlice_AppendFmtIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AppendFmtIf(true, "val=%d", 42)
		ss.AppendFmtIf(false, "skip")

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_11_SimpleSlice_AppendFmtIf_EmptyFormatSkip_FromS11a(t *testing.T) {
	safeTest(t, "Test_11_SimpleSlice_AppendFmtIf_EmptyFormatSkip", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AppendFmtIf(true, "")

		// Assert
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_SimpleSlice_12_SimpleSlice_AddAsTitleValue_FromS11a(t *testing.T) {
	safeTest(t, "Test_12_SimpleSlice_AddAsTitleValue", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddAsTitleValue("Name", "John")

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_13_SimpleSlice_AddAsCurlyTitleWrap_FromS11a(t *testing.T) {
	safeTest(t, "Test_13_SimpleSlice_AddAsCurlyTitleWrap", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddAsCurlyTitleWrap("Key", "Val")

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_14_SimpleSlice_AddAsCurlyTitleWrapIf_FromS11a(t *testing.T) {
	safeTest(t, "Test_14_SimpleSlice_AddAsCurlyTitleWrapIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddAsCurlyTitleWrapIf(true, "K", "V")
		ss.AddAsCurlyTitleWrapIf(false, "K2", "V2")

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_15_SimpleSlice_AddAsTitleValueIf_FromS11a(t *testing.T) {
	safeTest(t, "Test_15_SimpleSlice_AddAsTitleValueIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddAsTitleValueIf(true, "K", "V")
		ss.AddAsTitleValueIf(false, "K2", "V2")

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_16_SimpleSlice_InsertAt_FromS11a(t *testing.T) {
	safeTest(t, "Test_16_SimpleSlice_InsertAt", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "c"})

		// Act
		ss.InsertAt(1, "b")

		// Assert
		actual := args.Map{"result": ss.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_SimpleSlice_17_SimpleSlice_InsertAt_OutOfRange_FromS11a(t *testing.T) {
	safeTest(t, "Test_17_SimpleSlice_InsertAt_OutOfRange", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		ss.InsertAt(-1, "x")
		ss.InsertAt(100, "y")

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1 — out of range skipped", actual)
	})
}

func Test_SimpleSlice_18_SimpleSlice_AddStruct_FromS11a(t *testing.T) {
	safeTest(t, "Test_18_SimpleSlice_AddStruct", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddStruct(true, "hello")
		ss.AddStruct(false, nil) // nil skipped

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_19_SimpleSlice_AddPointer_FromS11a(t *testing.T) {
	safeTest(t, "Test_19_SimpleSlice_AddPointer", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddPointer(true, "hello")
		ss.AddPointer(false, nil) // nil skipped

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_20_SimpleSlice_AddsIf_FromS11a(t *testing.T) {
	safeTest(t, "Test_20_SimpleSlice_AddsIf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddsIf(true, "a", "b")
		ss.AddsIf(false, "c")

		// Assert
		actual := args.Map{"result": ss.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SimpleSlice_21_SimpleSlice_AddError_FromS11a(t *testing.T) {
	safeTest(t, "Test_21_SimpleSlice_AddError", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Cap(5)

		// Act
		ss.AddError(nil)
		ss.AddError(&testErr{})

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_22_SimpleSlice_AsDefaultError_FromS11a(t *testing.T) {
	safeTest(t, "Test_22_SimpleSlice_AsDefaultError", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"err1", "err2"})

		// Act
		err := ss.AsDefaultError()

		// Assert
		actual := args.Map{"result": err == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected error", actual)
	})
}

func Test_SimpleSlice_23_SimpleSlice_AsError_Empty_FromS11a(t *testing.T) {
	safeTest(t, "Test_23_SimpleSlice_AsError_Empty", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()

		// Act
		err := ss.AsError(",")

		// Assert
		actual := args.Map{"result": err != nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected nil", actual)
	})
}

func Test_SimpleSlice_24_SimpleSlice_FirstAndLast_FromS11a(t *testing.T) {
	safeTest(t, "Test_24_SimpleSlice_FirstAndLast", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})

		// Act & Assert
		actual := args.Map{"result": ss.First() != "a"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a", actual)
		actual = args.Map{"result": ss.Last() != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c", actual)
		actual = args.Map{"result": ss.FirstDynamic().(string) != "a"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a dynamic", actual)
		actual = args.Map{"result": ss.LastDynamic().(string) != "c"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected c dynamic", actual)
	})
}

func Test_SimpleSlice_25_SimpleSlice_FirstOrDefault_FromS11a(t *testing.T) {
	safeTest(t, "Test_25_SimpleSlice_FirstOrDefault", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()
		ss2 := corestr.New.SimpleSlice.Strings([]string{"x"})

		// Act & Assert
		actual := args.Map{"result": ss.FirstOrDefault() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": ss.FirstOrDefaultDynamic().(string) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty dynamic", actual)
		actual = args.Map{"result": ss2.FirstOrDefault() != "x"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected x", actual)
	})
}

func Test_SimpleSlice_26_SimpleSlice_LastOrDefault_FromS11a(t *testing.T) {
	safeTest(t, "Test_26_SimpleSlice_LastOrDefault", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()
		ss2 := corestr.New.SimpleSlice.Strings([]string{"x"})

		// Act & Assert
		actual := args.Map{"result": ss.LastOrDefault() != ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
		actual = args.Map{"result": ss.LastOrDefaultDynamic().(string) != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty dynamic", actual)
		actual = args.Map{"result": ss2.LastOrDefault() != "x"}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected x", actual)
	})
}

func Test_SimpleSlice_27_SimpleSlice_Skip_FromS11a(t *testing.T) {
	safeTest(t, "Test_27_SimpleSlice_Skip", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})

		// Act
		skipped := ss.Skip(1)
		skippedAll := ss.Skip(10)

		// Assert
		actual := args.Map{"result": len(skipped) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual = args.Map{"result": len(skippedAll) != 0}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_SimpleSlice_28_SimpleSlice_SkipDynamic_FromS11a(t *testing.T) {
	safeTest(t, "Test_28_SimpleSlice_SkipDynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		result := ss.SkipDynamic(1)

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_SimpleSlice_29_SimpleSlice_Take_FromS11a(t *testing.T) {
	safeTest(t, "Test_29_SimpleSlice_Take", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b", "c"})

		// Act
		taken := ss.Take(2)
		takenAll := ss.Take(10)

		// Assert
		actual := args.Map{"result": len(taken) != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
		actual = args.Map{"result": len(takenAll) != 3}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_SimpleSlice_30_SimpleSlice_TakeDynamic_FromS11a(t *testing.T) {
	safeTest(t, "Test_30_SimpleSlice_TakeDynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		result := ss.TakeDynamic(1)

		// Assert
		actual := args.Map{"result": result == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	})
}

func Test_SimpleSlice_31_SimpleSlice_LimitDynamic_FromS11a(t *testing.T) {
	safeTest(t, "Test_31_SimpleSlice_LimitDynamic", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": ss.LimitDynamic(1) == nil}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
		actual = args.Map{"result": len(ss.Limit(1)) != 1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_32_SimpleSlice_Length_Count_FromS11a(t *testing.T) {
	safeTest(t, "Test_32_SimpleSlice_Length_Count", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": ss.Length() != 1 || ss.Count() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_33_SimpleSlice_Length_Nil_FromS11a(t *testing.T) {
	safeTest(t, "Test_33_SimpleSlice_Length_Nil", func() {
		// Arrange
		var ss *corestr.SimpleSlice

		// Act & Assert
		actual := args.Map{"result": ss.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_SimpleSlice_34_SimpleSlice_CountFunc_FromS11a(t *testing.T) {
	safeTest(t, "Test_34_SimpleSlice_CountFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "bb", "ccc"})

		// Act
		count := ss.CountFunc(func(index int, item string) bool {
			return len(item) > 1
		})

		// Assert
		actual := args.Map{"result": count != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SimpleSlice_35_SimpleSlice_CountFunc_Empty_FromS11a(t *testing.T) {
	safeTest(t, "Test_35_SimpleSlice_CountFunc_Empty", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()

		// Act
		count := ss.CountFunc(func(index int, item string) bool { return true })

		// Assert
		actual := args.Map{"result": count != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_SimpleSlice_36_SimpleSlice_IsEmpty_FromS11a(t *testing.T) {
	safeTest(t, "Test_36_SimpleSlice_IsEmpty", func() {
		// Act & Assert
		actual := args.Map{"result": corestr.Empty.SimpleSlice().IsEmpty()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_SimpleSlice_37_SimpleSlice_IsContains_FromS11a(t *testing.T) {
	safeTest(t, "Test_37_SimpleSlice_IsContains", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": ss.IsContains("a")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ss.IsContains("z")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().IsContains("a")}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

func Test_SimpleSlice_38_SimpleSlice_IsContainsFunc_FromS11a(t *testing.T) {
	safeTest(t, "Test_38_SimpleSlice_IsContainsFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"abc"})

		// Act & Assert
		actual := args.Map{"result": ss.IsContainsFunc("ab", strings.Contains)}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().IsContainsFunc("a", strings.Contains)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false for empty", actual)
	})
}

func Test_SimpleSlice_39_SimpleSlice_IndexOf_FromS11a(t *testing.T) {
	safeTest(t, "Test_39_SimpleSlice_IndexOf", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": ss.IndexOf("b") != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": ss.IndexOf("z") != -1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().IndexOf("a") != -1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1 for empty", actual)
	})
}

func Test_SimpleSlice_40_SimpleSlice_IndexOfFunc_FromS11a(t *testing.T) {
	safeTest(t, "Test_40_SimpleSlice_IndexOfFunc", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"abc", "def"})

		// Act
		idx := ss.IndexOfFunc("de", strings.Contains)

		// Assert
		actual := args.Map{"result": idx != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().IndexOfFunc("a", strings.Contains) != -1}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected -1 for empty", actual)
	})
}

func Test_SimpleSlice_41_SimpleSlice_HasAnyItem_FromS11a(t *testing.T) {
	safeTest(t, "Test_41_SimpleSlice_HasAnyItem", func() {
		// Act & Assert
		actual := args.Map{"result": corestr.New.SimpleSlice.Strings([]string{"a"}).HasAnyItem()}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
	})
}

func Test_SimpleSlice_42_SimpleSlice_LastIndex_HasIndex_FromS11a(t *testing.T) {
	safeTest(t, "Test_42_SimpleSlice_LastIndex_HasIndex", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": ss.LastIndex() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
		actual = args.Map{"result": ss.HasIndex(0) || !ss.HasIndex(1)}
		expected = args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected true", actual)
		actual = args.Map{"result": ss.HasIndex(2) || ss.HasIndex(-1)}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected false", actual)
	})
}

func Test_SimpleSlice_43_SimpleSlice_Strings_List_FromS11a(t *testing.T) {
	safeTest(t, "Test_43_SimpleSlice_Strings_List", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": len(ss.Strings()) != 1 || len(ss.List()) != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_44_SimpleSlice_WrapDoubleQuote_FromS11a(t *testing.T) {
	safeTest(t, "Test_44_SimpleSlice_WrapDoubleQuote", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": ss.WrapDoubleQuote().Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_45_SimpleSlice_WrapSingleQuote_FromS11a(t *testing.T) {
	safeTest(t, "Test_45_SimpleSlice_WrapSingleQuote", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.WrapSingleQuote().Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_46_SimpleSlice_WrapTildaQuote_FromS11a(t *testing.T) {
	safeTest(t, "Test_46_SimpleSlice_WrapTildaQuote", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.WrapTildaQuote().Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_47_SimpleSlice_WrapDoubleQuoteIfMissing_FromS11a(t *testing.T) {
	safeTest(t, "Test_47_SimpleSlice_WrapDoubleQuoteIfMissing", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.WrapDoubleQuoteIfMissing().Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_48_SimpleSlice_WrapSingleQuoteIfMissing_FromS11a(t *testing.T) {
	safeTest(t, "Test_48_SimpleSlice_WrapSingleQuoteIfMissing", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.WrapSingleQuoteIfMissing().Length() != 1}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}

func Test_SimpleSlice_49_SimpleSlice_Transpile_FromS11a(t *testing.T) {
	safeTest(t, "Test_49_SimpleSlice_Transpile", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		result := ss.Transpile(strings.ToUpper)

		// Assert
		actual := args.Map{"result": result.First() != "A"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected A", actual)
	})
}

func Test_SimpleSlice_50_SimpleSlice_Transpile_Empty_FromS11a(t *testing.T) {
	safeTest(t, "Test_50_SimpleSlice_Transpile_Empty", func() {
		// Arrange
		ss := corestr.Empty.SimpleSlice()

		// Act
		result := ss.Transpile(strings.ToUpper)

		// Assert
		actual := args.Map{"result": result.Length() != 0}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 0", actual)
	})
}

func Test_SimpleSlice_51_SimpleSlice_TranspileJoin_FromS11a(t *testing.T) {
	safeTest(t, "Test_51_SimpleSlice_TranspileJoin", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		result := ss.TranspileJoin(strings.ToUpper, ",")

		// Assert
		actual := args.Map{"result": result != "A,B"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'A,B', got ''", actual)
	})
}

func Test_SimpleSlice_52_SimpleSlice_Hashset_FromS11a(t *testing.T) {
	safeTest(t, "Test_52_SimpleSlice_Hashset", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		hs := ss.Hashset()

		// Assert
		actual := args.Map{"result": hs.Length() != 2}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 2", actual)
	})
}

func Test_SimpleSlice_53_SimpleSlice_Join_FromS11a(t *testing.T) {
	safeTest(t, "Test_53_SimpleSlice_Join", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act & Assert
		actual := args.Map{"result": ss.Join(",") != "a,b"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected a,b", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().Join(",") != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_SimpleSlice_54_SimpleSlice_JoinLine_FromS11a(t *testing.T) {
	safeTest(t, "Test_54_SimpleSlice_JoinLine", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act & Assert
		actual := args.Map{"result": ss.JoinLine() == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().JoinLine() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_SimpleSlice_55_SimpleSlice_JoinLineEofLine_FromS11a(t *testing.T) {
	safeTest(t, "Test_55_SimpleSlice_JoinLineEofLine", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		result := ss.JoinLineEofLine()

		// Assert
		actual := args.Map{"result": strings.HasSuffix(result, "\n")}
		expected := args.Map{"result": true}
		expected.ShouldBeEqual(t, 0, "expected newline at end", actual)
		actual = args.Map{"result": corestr.Empty.SimpleSlice().JoinLineEofLine() != ""}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected empty", actual)
	})
}

func Test_SimpleSlice_56_SimpleSlice_JoinLineEofLine_AlreadyHas_FromS11a(t *testing.T) {
	safeTest(t, "Test_56_SimpleSlice_JoinLineEofLine_AlreadyHas", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a\n"})

		// Act
		result := ss.JoinLineEofLine()

		// Assert
		actual := args.Map{"result": result == ""}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_SimpleSlice_57_SimpleSlice_JoinSpace_FromS11a(t *testing.T) {
	safeTest(t, "Test_57_SimpleSlice_JoinSpace", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ss.JoinSpace() != "a b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a b'", actual)
	})
}

func Test_SimpleSlice_58_SimpleSlice_JoinComma_FromS11a(t *testing.T) {
	safeTest(t, "Test_58_SimpleSlice_JoinComma", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a", "b"})

		// Act
		actual := args.Map{"result": ss.JoinComma() != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b'", actual)
	})
}

func Test_SimpleSlice_59_SimpleSlice_JoinCsv_FromS11a(t *testing.T) {
	safeTest(t, "Test_59_SimpleSlice_JoinCsv", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		result := ss.JoinCsv()

		// Act
		actual := args.Map{"result": result == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_SimpleSlice_60_SimpleSlice_JoinCsvLine_FromS11a(t *testing.T) {
	safeTest(t, "Test_60_SimpleSlice_JoinCsvLine", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})

		// Act
		actual := args.Map{"result": ss.JoinCsvLine() == ""}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
	})
}

func Test_SimpleSlice_61_SimpleSlice_EachItemSplitBy_FromS11a(t *testing.T) {
	safeTest(t, "Test_61_SimpleSlice_EachItemSplitBy", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a,b", "c,d"})

		// Act
		result := ss.EachItemSplitBy(",")

		// Assert
		actual := args.Map{"result": result.Length() != 4}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 4", actual)
	})
}

func Test_SimpleSlice_62_SimpleSlice_PrependJoin_FromS11a(t *testing.T) {
	safeTest(t, "Test_62_SimpleSlice_PrependJoin", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"b"})
		result := ss.PrependJoin(",", "a")

		// Act
		actual := args.Map{"result": result != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
	})
}

func Test_SimpleSlice_63_SimpleSlice_AppendJoin_FromS11a(t *testing.T) {
	safeTest(t, "Test_63_SimpleSlice_AppendJoin", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"a"})
		result := ss.AppendJoin(",", "b")

		// Act
		actual := args.Map{"result": result != "a,b"}

		// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 'a,b', got ''", actual)
	})
}

func Test_SimpleSlice_64_SimpleSlice_PrependAppend_FromS11a(t *testing.T) {
	safeTest(t, "Test_64_SimpleSlice_PrependAppend", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"b"})

		// Act
		ss.PrependAppend([]string{"a"}, []string{"c"})

		// Assert
		actual := args.Map{"result": ss.Length() != 3}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 3", actual)
	})
}

func Test_SimpleSlice_65_SimpleSlice_PrependAppend_Empty_FromS11a(t *testing.T) {
	safeTest(t, "Test_65_SimpleSlice_PrependAppend_Empty", func() {
		// Arrange
		ss := corestr.New.SimpleSlice.Strings([]string{"b"})

		// Act
		ss.PrependAppend(nil, nil)

		// Assert
		actual := args.Map{"result": ss.Length() != 1}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected 1", actual)
	})
}
