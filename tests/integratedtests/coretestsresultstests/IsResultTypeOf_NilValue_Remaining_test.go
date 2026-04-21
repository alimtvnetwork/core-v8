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

package coretestsresultstests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/results"
)

// ── Result.IsResultTypeOf: actualType nil (Value is nil any) ──

func Test_IsResultTypeOf_NilValue_NonNilExpected(t *testing.T) {
	// Arrange
	r := results.ResultAny{Value: nil}

	// Act
	actual := args.Map{"typeOf": r.IsResultTypeOf(42)}

	// Assert
	expected := args.Map{"typeOf": false}
	expected.ShouldBeEqual(t, 0, "IsResultTypeOf returns false -- nil value non-nil expected", actual)
}

func Test_IsResultTypeOf_NilValue_NilExpected_Zero(t *testing.T) {
	// Arrange
	// Value is zero-value int (not nil), expected is nil
	r := results.Result[int]{Value: 0}

	// Act
	actual := args.Map{"typeOf": r.IsResultTypeOf(nil)}

	// Assert
	expected := args.Map{"typeOf": true}
	expected.ShouldBeEqual(t, 0, "IsResultTypeOf returns true -- zero int nil expected", actual)
}

// ── deriveCompareFields: only Error set ──

func Test_ShouldMatchResult_ErrorDerived(t *testing.T) {
	r := results.Result[string]{Value: "x", Error: errors.New("e")}
	exp := results.ResultAny{Error: results.ExpectAnyError}
	// This exercises deriveCompareFields with only Error != nil
	r.ShouldMatchResult(t, 0, "error derived", exp, "panicked", "hasError")
}

// ── FromResultAny: panicked with no AllResults ──

func Test_FromResultAny_Panicked(t *testing.T) {
	// Arrange
	ra := results.ResultAny{Panicked: true, PanicValue: "boom"}
	r := results.FromResultAny[string, int](ra)

	// Act
	actual := args.Map{
		"panicked": r.Panicked,
		"val":      r.Value,
		"val2":     r.Result2,
	}

	// Assert
	expected := args.Map{
		"panicked": true,
		"val":      "",
		"val2":     0,
	}
	expected.ShouldBeEqual(t, 0, "FromResultAny returns zero values -- panicked", actual)
}

// ── FromResultAny: single AllResult ──

func Test_FromResultAny_SingleResult(t *testing.T) {
	// Arrange
	ra := results.ResultAny{AllResults: []any{"hello"}, ReturnCount: 1}
	r := results.FromResultAny[string, int](ra)

	// Act
	actual := args.Map{
		"val": r.Value,
		"val2": r.Result2,
	}

	// Assert
	expected := args.Map{
		"val": "hello",
		"val2": 0,
	}
	expected.ShouldBeEqual(t, 0, "FromResultAny returns first only -- single result", actual)
}

// ── Result.ToMapCompact: panicked ──

func Test_ToMapCompact_Panicked(t *testing.T) {
	// Arrange
	r := results.ResultAny{Panicked: true, PanicValue: "boom"}
	m := r.ToMapCompact()

	// Act
	actual := args.Map{"panicked": m["panicked"]}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "ToMapCompact returns panicked -- panicked result", actual)
}

// ── Invoke with receiver + nil arg that is non-any typed ──

type cov5TypedArgStruct struct{}

func (s *cov5TypedArgStruct) Process(name string, count int) string {
	return name
}

func Test_Invoke_TypedArgs(t *testing.T) {
	// Arrange
	s := &cov5TypedArgStruct{}
	r := results.InvokeWithPanicRecovery((*cov5TypedArgStruct).Process, s, "hello", 5)

	// Act
	actual := args.Map{
		"panicked": r.Panicked,
		"val":      r.Value,
		"count":    r.ReturnCount,
	}

	// Assert
	expected := args.Map{
		"panicked": false,
		"val":      "hello",
		"count":    1,
	}
	expected.ShouldBeEqual(t, 0, "Invoke returns correct value -- typed args", actual)
}

// ── extractErrorFromValue: Ptr that implements error but not nil ──

type cov5CustomErr struct{ msg string }

func (e *cov5CustomErr) Error() string { return e.msg }

type cov5PtrErrReturn struct{}

func (s *cov5PtrErrReturn) ReturnErr() *cov5CustomErr {
	return &cov5CustomErr{msg: "custom"}
}

func Test_ExtractError_PtrImplError(t *testing.T) {
	// Arrange
	s := &cov5PtrErrReturn{}
	r := results.InvokeWithPanicRecovery((*cov5PtrErrReturn).ReturnErr, s)

	// Act
	actual := args.Map{
		"panicked": r.Panicked,
		"hasError": r.HasError(),
	}

	// Assert
	expected := args.Map{
		"panicked": false,
		"hasError": true,
	}
	expected.ShouldBeEqual(t, 0, "extractError returns error -- ptr impl error", actual)
}

// ── ShouldMatchResult: with explicit compareFields including returnCount ──

func Test_ShouldMatchResult_ReturnCountDerived(t *testing.T) {
	r := results.Result[int]{Value: 10, ReturnCount: 2}
	exp := results.ResultAny{ReturnCount: 2}
	r.ShouldMatchResult(t, 0, "returnCount derived", exp)
}

// ── Result.IsResult / IsError combined ──

func Test_Result_IsResult_StringComparison(t *testing.T) {
	// Arrange
	r := results.Result[int]{Value: 0}

	// Act
	actual := args.Map{
		"isZero":   r.IsResult(0),
		"isNotOne": !r.IsResult(1),
	}

	// Assert
	expected := args.Map{
		"isZero":   true,
		"isNotOne": true,
	}
	expected.ShouldBeEqual(t, 0, "IsResult compares via fmt -- zero value", actual)
}
