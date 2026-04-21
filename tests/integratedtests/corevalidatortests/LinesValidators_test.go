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
	"strings"
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/corestr"
	"github.com/alimtvnetwork/core-v8/corevalidator"
	"github.com/alimtvnetwork/core-v8/enums/stringcompareas"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ==========================================
// LinesValidators — collection basics
// ==========================================

func Test_LinesValidators_Count(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{})

	// Act
	actual := args.Map{"result": lv.Count() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_LinesValidators_LastIndex(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(2)
	lv.Add(corevalidator.LineValidator{})
	lv.Add(corevalidator.LineValidator{})

	// Act
	actual := args.Map{"result": lv.LastIndex() != 1}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 1", actual)
}

func Test_LinesValidators_Adds_FromLinesValidators(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(3)
	lv.Adds(
		corevalidator.LineValidator{},
		corevalidator.LineValidator{},
		corevalidator.LineValidator{},
	)

	// Act
	actual := args.Map{"result": lv.Length() != 3}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 3", actual)
}

func Test_LinesValidators_String_FromLinesValidators(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: 0},
		TextValidator: corevalidator.TextValidator{
			Search:   "test",
			SearchAs: stringcompareas.Equal,
		},
	})
	s := lv.String()

	// Act
	actual := args.Map{"result": s == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "String should not be empty", actual)
}

// ==========================================
// LinesValidators.IsMatch (with contents)
// ==========================================

func Test_LinesValidators_IsMatch_Empty_FromLinesValidators(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(0)
	items := []corestr.TextWithLineNumber{
		{Text: "hello", LineNumber: 0},
	}

	// Act
	actual := args.Map{"result": lv.IsMatch(false, true, items...)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "empty validators should match", actual)
}

func Test_LinesValidators_IsMatch_NoContentsSkip_FromLinesValidators(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})

	// Act
	actual := args.Map{"result": lv.IsMatch(true, true)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "no contents with skip should match", actual)
}

func Test_LinesValidators_IsMatch_NoContentsNoSkip_FromLinesValidators(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})

	// Act
	actual := args.Map{"result": lv.IsMatch(false, true)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "no contents without skip should not match", actual)
}

func Test_LinesValidators_IsMatch_AllMatch(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	items := []corestr.TextWithLineNumber{
		{Text: "ok", LineNumber: 0},
		{Text: "ok", LineNumber: 1},
	}

	// Act
	actual := args.Map{"result": lv.IsMatch(false, true, items...)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "all matching should return true", actual)
}

func Test_LinesValidators_IsMatch_OneFails(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	items := []corestr.TextWithLineNumber{
		{Text: "ok", LineNumber: 0},
		{Text: "nope", LineNumber: 1},
	}

	// Act
	actual := args.Map{"result": lv.IsMatch(false, true, items...)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "one failing should return false", actual)
}

// ==========================================
// LinesValidators.VerifyFirstDefaultLineNumberError
// ==========================================

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_Empty_FromLinesValidators(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(0)
	params := &corevalidator.Parameter{CaseIndex: 0}
	err := lv.VerifyFirstDefaultLineNumberError(params)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_SkipEmpty(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:                  0,
		IsSkipCompareOnActualEmpty: true,
	}
	err := lv.VerifyFirstDefaultLineNumberError(params)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "skip empty should return nil:", actual)
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_NoSkipEmpty(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:                  0,
		IsSkipCompareOnActualEmpty: false,
	}
	err := lv.VerifyFirstDefaultLineNumberError(params)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty contents without skip should return error", actual)
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_Pass(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "ok", LineNumber: 0},
	}
	err := lv.VerifyFirstDefaultLineNumberError(params, items...)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "match should pass:", actual)
}

func Test_LinesValidators_VerifyFirstDefaultLineNumberError_Fail(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "bad", LineNumber: 0},
	}
	err := lv.VerifyFirstDefaultLineNumberError(params, items...)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mismatch should return error", actual)
}

// ==========================================
// LinesValidators.AllVerifyError
// ==========================================

func Test_LinesValidators_AllVerifyError_Empty_FromLinesValidators(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(0)
	params := &corevalidator.Parameter{CaseIndex: 0}
	err := lv.AllVerifyError(params)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty should return nil", actual)
}

func Test_LinesValidators_AllVerifyError_SkipEmpty(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:                  0,
		IsSkipCompareOnActualEmpty: true,
	}
	err := lv.AllVerifyError(params)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "skip empty should return nil:", actual)
}

func Test_LinesValidators_AllVerifyError_NoSkipEmpty(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		TextValidator: corevalidator.TextValidator{
			Search:   "x",
			SearchAs: stringcompareas.Equal,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:                  0,
		IsSkipCompareOnActualEmpty: false,
	}
	err := lv.AllVerifyError(params)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "empty contents without skip should return error", actual)
}

func Test_LinesValidators_AllVerifyError_Pass(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "ok", LineNumber: 0},
	}
	err := lv.AllVerifyError(params, items...)

	// Act
	actual := args.Map{"result": err != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "match should pass:", actual)
}

func Test_LinesValidators_AllVerifyError_Fail(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(1)
	lv.Add(corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	})
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "bad", LineNumber: 0},
	}
	err := lv.AllVerifyError(params, items...)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "mismatch should return error", actual)
}

func Test_LineValidator_AllVerifyError_CollectsMultipleErrors(t *testing.T) {
	// Arrange: validator expects "ok", but all 3 inputs are wrong
	lv := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "bad1", LineNumber: 0},
		{Text: "bad2", LineNumber: 1},
		{Text: "bad3", LineNumber: 2},
	}

	// Act
	err := lv.AllVerifyError(params, items...)

	// Assert: error should be non-nil and contain all 3 failures
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllVerifyError should return error when all items fail", actual)

	errMsg := err.Error()
	for _, exp := range []string{"bad1", "bad2", "bad3"} {
		actualCheck := args.Map{"result": strings.Contains(errMsg, exp)}
		expectedCheck := args.Map{"result": true}
		expectedCheck.ShouldBeEqual(t, 0, "AllVerifyError should collect all errors, missing '"+exp+"'", actualCheck)
	}
}

func Test_LineValidator_AllVerifyError_FirstFailOthersPass(t *testing.T) {
	// Arrange: validator expects "ok", first fails, rest pass
	lv := corevalidator.LineValidator{
		LineNumber: corevalidator.LineNumber{LineNumber: -1},
		TextValidator: corevalidator.TextValidator{
			Search:    "ok",
			SearchAs:  stringcompareas.Equal,
			Condition: corevalidator.DefaultDisabledCoreCondition,
		},
	}
	params := &corevalidator.Parameter{
		CaseIndex:       0,
		IsCaseSensitive: true,
	}
	items := []corestr.TextWithLineNumber{
		{Text: "bad", LineNumber: 0},
		{Text: "ok", LineNumber: 1},
		{Text: "ok", LineNumber: 2},
	}

	// Act
	err := lv.AllVerifyError(params, items...)

	// Assert: should still report the one failure
	actual := args.Map{"result": err == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "AllVerifyError should return error when any item fails", actual)

	errMsg := err.Error()
	actual = args.Map{"result": strings.Contains(errMsg, "bad")}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "error should mention the failed content 'bad', got:\n", actual)
}

// ==========================================
// LinesValidators.AsBasicSliceContractsBinder
// ==========================================

func Test_LinesValidators_AsBasicSliceContractsBinder_FromLinesValidators(t *testing.T) {
	// Arrange
	lv := corevalidator.NewLinesValidators(1)
	binder := lv.AsBasicSliceContractsBinder()

	// Act
	actual := args.Map{"result": binder == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be nil", actual)
}
