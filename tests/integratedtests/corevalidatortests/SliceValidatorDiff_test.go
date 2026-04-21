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

package corevalidatortests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core-v8/corevalidator"
	"github.com/alimtvnetwork/core-v8/enums/stringcompareas"
	"github.com/alimtvnetwork/core-v8/errcore"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ==========================================
// SliceValidator — AllVerifyError with diff
// ==========================================

func Test_SliceValidator_AllVerifyError_MultiLineMismatch_WithDiff(t *testing.T) {
	// Arrange: 5 lines, 2 mismatches at lines 1 and 3
	actualLines := []string{"alpha", "bravo-wrong", "charlie", "delta-wrong", "echo"}
	expectedLines := []string{"alpha", "bravo", "charlie", "delta", "echo"}

	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   actualLines,
		ExpectedLines: expectedLines,
	}

	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "Multi-line mismatch with diff output",
		IsCaseSensitive:    true,
		IsAttachUserInputs: true,
	}

	// Act
	err := v.AllVerifyError(params)

	// Assert: must fail
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for multi-line mismatch", actual)

	// Print line-by-line diff for diagnostics
	errcore.PrintDiffOnMismatch(0, params.Header, actualLines, expectedLines)

	errMsg := err.Error()
	actual = args.Map{"result": strings.Contains(errMsg, "bravo")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "error should mention 'bravo' mismatch, got:\n", actual)
	actual = args.Map{"result": strings.Contains(errMsg, "delta")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "error should mention 'delta' mismatch, got:\n", actual)
}

func Test_SliceValidator_AllVerifyError_ExtraActualLines_WithDiff(t *testing.T) {
	// Act
	actualLines := []string{"line1", "line2", "line3", "extra-line"}

	// Assert
	expectedLines := []string{"line1", "line2", "line3"}

	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   actualLines,
		ExpectedLines: expectedLines,
	}

	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "Extra actual lines diff",
		IsCaseSensitive:    true,
		IsAttachUserInputs: true,
	}

	err := v.AllVerifyError(params)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for length mismatch", actual)

	// Print diff showing extra line
	errcore.PrintDiffOnMismatch(0, params.Header, actualLines, expectedLines)
	summary := errcore.SliceDiffSummary(actualLines, expectedLines)
	actual2 := args.Map{"hasSummary": len(summary) > 0}
	expected2 := args.Map{"hasSummary": true}
	expected2.ShouldBeEqual(t, 0, "diff summary should be non-empty", actual2)
}

func Test_SliceValidator_AllVerifyError_MissingActualLines_WithDiff(t *testing.T) {
	// Act
	actualLines := []string{"line1"}

	// Assert
	expectedLines := []string{"line1", "line2", "line3"}

	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   actualLines,
		ExpectedLines: expectedLines,
	}

	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "Missing actual lines diff",
		IsCaseSensitive:    true,
		IsAttachUserInputs: true,
	}

	err := v.AllVerifyError(params)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for missing actual lines", actual)

	errcore.PrintDiffOnMismatch(0, params.Header, actualLines, expectedLines)
}

// ==========================================
// SliceValidator — VerifyFirstError with diff
// ==========================================

func Test_SliceValidator_VerifyFirstError_StopsAtFirst_WithDiff(t *testing.T) {
	// Act
	actualLines := []string{"a", "WRONG1", "WRONG2"}

	// Assert
	expectedLines := []string{"a", "b", "c"}

	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   actualLines,
		ExpectedLines: expectedLines,
	}

	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "VerifyFirst stops at first mismatch",
		IsCaseSensitive:    true,
		IsAttachUserInputs: true,
	}

	err := v.VerifyFirstError(params)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	errcore.PrintDiffOnMismatch(0, params.Header, actualLines, expectedLines)

	// VerifyFirst should mention line 1 mismatch
	errMsg := err.Error()
	actual = args.Map{"result": strings.Contains(errMsg, "WRONG1")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should mention first mismatch 'WRONG1', got:\n", actual)
}

// ==========================================
// SliceValidator — AllVerifyErrorTestCase with diff
// ==========================================

func Test_SliceValidator_AllVerifyErrorTestCase_WithDiff(t *testing.T) {
	// Act
	actualLines := []string{"hello", "world-different"}

	// Assert
	expectedLines := []string{"hello", "world"}

	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   actualLines,
		ExpectedLines: expectedLines,
	}

	err := v.AllVerifyErrorTestCase(0, "TestCase with diff", true)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)

	// Also print our enhanced diff
	errcore.PrintDiffOnMismatch(0, "TestCase with diff", actualLines, expectedLines)
}

// ==========================================
// SliceValidator — Contains with multiple mismatches
// ==========================================

func Test_SliceValidator_AllVerifyError_Contains_MultiMismatch(t *testing.T) {
	// Act
	actualLines := []string{
		"path/to/file.go:10",
		"some other text",
		"path/to/other.go:20",
	}

	// Assert
	expectedLines := []string{
		"file.go",
		"expected-missing",
		"other.go",
	}

	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Contains,
		ActualLines:   actualLines,
		ExpectedLines: expectedLines,
	}

	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "Contains multi-mismatch",
		IsCaseSensitive:    true,
		IsAttachUserInputs: true,
	}

	err := v.AllVerifyError(params)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for line 1 mismatch", actual)

	errcore.PrintDiffOnMismatch(0, params.Header, actualLines, expectedLines)

	errMsg := err.Error()
	actual = args.Map{"result": strings.Contains(errMsg, "expected-missing")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "error should reference missing substring, got:\n", actual)
}

// ==========================================
// SliceValidator — Trim + diff
// ==========================================

func Test_SliceValidator_AllVerifyError_Trim_WithDiff(t *testing.T) {
	// Act
	actualLines := []string{"  hello  ", "  world  "}

	// Assert
	expectedLines := []string{"hello", "universe"}

	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultTrimCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   actualLines,
		ExpectedLines: expectedLines,
	}

	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "Trim with diff - line 1 mismatch",
		IsCaseSensitive: true,
	}

	err := v.AllVerifyError(params)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error: world != universe after trim", actual)

	errcore.PrintDiffOnMismatch(0, params.Header, actualLines, expectedLines)
}

// ==========================================
// SliceValidator — Glob pattern with diff
// ==========================================

func Test_SliceValidator_AllVerifyError_Glob_WithDiff(t *testing.T) {
	// Act
	actualLines := []string{
		"build-20260303/result.json",
		"build-20260303/output.txt",
		"build-20260303/data.csv",
	}

	// Assert
	expectedLines := []string{
		"build-*/result.json",
		"build-*/output.txt",
		"build-*/WRONG.csv",
	}

	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Glob,
		ActualLines:   actualLines,
		ExpectedLines: expectedLines,
	}

	params := &corevalidator.Parameter{
		CaseIndex:          0,
		Header:             "Glob pattern - line 2 mismatch",
		IsCaseSensitive:    true,
		IsAttachUserInputs: true,
	}

	err := v.AllVerifyError(params)
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error: data.csv doesn't match WRONG.csv glob", actual)

	errcore.PrintDiffOnMismatch(0, params.Header, actualLines, expectedLines)
}

// ==========================================
// SliceValidator — AllVerifyErrorExceptLast with diff
// ==========================================

func Test_SliceValidator_AllVerifyErrorExceptLast_WithDiff(t *testing.T) {
	// Act
	actualLines := []string{"a", "b", "INTENTIONALLY-DIFFERENT"}

	// Assert
	expectedLines := []string{"a", "b", "c"}

	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   actualLines,
		ExpectedLines: expectedLines,
	}

	params := &corevalidator.Parameter{
		CaseIndex:       0,
		Header:          "ExceptLast should skip last line",
		IsCaseSensitive: true,
	}

	err := v.AllVerifyErrorExceptLast(params)
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": false}
	expected.ShouldBeEqual(t, 0, "AllVerifyErrorExceptLast passes -- skipping last line", actual)
}

// ==========================================
// SliceValidator — Dispose then verify
// ==========================================

func Test_SliceValidator_Dispose_ThenAllVerifyError(t *testing.T) {
	// Arrange
	v := corevalidator.SliceValidator{
		Condition:     corevalidator.DefaultDisabledCoreCondition,
		CompareAs:     stringcompareas.Equal,
		ActualLines:   []string{"a"},
		ExpectedLines: []string{"a"},
	}

	v.Dispose()

	params := &corevalidator.Parameter{CaseIndex: 0}
	err := v.AllVerifyError(params)

	// After dispose, both are nil, so nil receiver-like behavior

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "disposed validator with nil lines should not error:", actual)
}

// ==========================================
// errcore.LineDiff utility direct tests
// ==========================================

func Test_LineDiff_BothEmpty(t *testing.T) {
	// Arrange
	diffs := errcore.LineDiff([]string{}, []string{})

	// Act
	actual := args.Map{"result": len(diffs) != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "both empty should produce 0 diffs", actual)
}

func Test_LineDiff_ExactMatch(t *testing.T) {
	// Act
	actualLines := []string{"a", "b", "c"}

	// Assert
	expectedLines := []string{"a", "b", "c"}
	diffs := errcore.LineDiff(actualLines, expectedLines)

	for i, d := range diffs {
		actual := args.Map{"result": d.Status != "  "}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "line should match, got status", actual)
		actual = args.Map{"result": d.LineNumber != i}
		expected = args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "line number should be", actual)
	}
}

func Test_LineDiff_Mismatches(t *testing.T) {
	// Act
	actualSlice := []string{"a", "WRONG", "c"}

	// Assert
	expectedSlice := []string{"a", "b", "c"}
	diffs := errcore.LineDiff(actualSlice, expectedSlice)

	actual := args.Map{"result": diffs[0].Status != "  "}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "line 0 should match", actual)
	actual = args.Map{"result": diffs[1].Status != "!!"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "line 1 should be mismatch", actual)
	actual = args.Map{"result": diffs[1].LineNumber != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mismatch line number should be 1", actual)
	actual = args.Map{"result": diffs[2].Status != "  "}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "line 2 should match", actual)
}

func Test_LineDiff_ExtraActual(t *testing.T) {
	// Act
	actualSlice := []string{"a", "b", "extra"}

	// Assert
	expectedSlice := []string{"a", "b"}
	diffs := errcore.LineDiff(actualSlice, expectedSlice)

	actual := args.Map{"result": len(diffs) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 diffs", actual)
	actual = args.Map{"result": diffs[2].Status != "+"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "extra line should have '+' status", actual)
	actual = args.Map{"result": diffs[2].LineNumber != 2}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "extra line number should be 2", actual)
}

func Test_LineDiff_MissingActual(t *testing.T) {
	// Act
	actualSlice := []string{"a"}

	// Assert
	expectedSlice := []string{"a", "b", "c"}
	diffs := errcore.LineDiff(actualSlice, expectedSlice)

	actual := args.Map{"result": len(diffs) != 3}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3 diffs", actual)
	actual = args.Map{"result": diffs[1].Status != "-"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "missing line should have '-' status", actual)
	actual = args.Map{"result": diffs[2].Status != "-"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "missing line should have '-' status", actual)
}

func Test_LineDiffToString_ContainsLineNumbers(t *testing.T) {
	// Act
	actualLines := []string{"a", "WRONG"}

	// Assert
	expectedLines := []string{"a", "b"}

	result := errcore.LineDiffToString(0, "test header", actualLines, expectedLines)

	actualCheck := args.Map{"result": strings.Contains(result, "Line")}
	expectedCheck := args.Map{"result": true}
	expectedCheck.ShouldBeEqual(t, 0, "diff output should contain 'Line' labels", actualCheck)
	actualCheck = args.Map{"result": strings.Contains(result, "MISMATCH")}
	expectedCheck = args.Map{"result": true}
	expectedCheck.ShouldBeEqual(t, 0, "diff output should contain 'MISMATCH' for differing lines", actualCheck)
	actualCheck = args.Map{"result": strings.Contains(result, "test header")}
	expectedCheck = args.Map{"result": true}
	expectedCheck.ShouldBeEqual(t, 0, "diff output should contain the header", actualCheck)
	actualCheck = args.Map{"result": strings.Contains(result, "Case 0")}
	expectedCheck = args.Map{"result": true}
	expectedCheck.ShouldBeEqual(t, 0, "diff output should contain the case index", actualCheck)

	// Print for visual inspection during test runs
	fmt.Print(result)
}

func Test_HasAnyMismatchOnLines_True(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"b"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "different content should be mismatch", actual)
}

func Test_HasAnyMismatchOnLines_DifferentLength(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.HasAnyMismatchOnLines([]string{"a"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "different length should be mismatch", actual)
}

func Test_HasAnyMismatchOnLines_False(t *testing.T) {
	// Act
	actual := args.Map{"result": errcore.HasAnyMismatchOnLines([]string{"a", "b"}, []string{"a", "b"})}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "same content should not be mismatch", actual)
}

func Test_SliceDiffSummary_AllMatch(t *testing.T) {
	// Arrange
	result := errcore.SliceDiffSummary([]string{"a", "b"}, []string{"a", "b"})

	// Act
	actual := args.Map{"result": result != "all lines match"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 'all lines match'", actual)
}

func Test_SliceDiffSummary_HasMismatches(t *testing.T) {
	// Arrange
	result := errcore.SliceDiffSummary(
		[]string{"a", "WRONG", "c"},
		[]string{"a", "b", "c"},
	)

	// Act
	actual := args.Map{"result": strings.Contains(result, "1 mismatches")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "summary should show mismatch count", actual)
	actual = args.Map{"result": strings.Contains(result, "line 1")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "summary should show line number", actual)
}
