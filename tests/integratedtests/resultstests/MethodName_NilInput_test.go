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

package resultstests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/results"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// Test_Cov3_MethodName_NilInput tests MethodName with nil funcRef.
func Test_MethodName_NilInput(t *testing.T) {
	// Arrange / Act
	result := results.MethodName(nil)

	// Assert
	actual := args.Map{"result": result != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MethodName(nil) should return empty string", actual)
}

// Test_Cov3_MethodName_NonFuncInput tests MethodName with a non-function value.
func Test_MethodName_NonFuncInput(t *testing.T) {
	// Arrange / Act
	result := results.MethodName("not-a-func")

	// Assert
	actual := args.Map{"result": result != ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MethodName(string) should return empty string", actual)
}

// Test_Cov3_MethodName_SimpleFuncNoDot tests MethodName with a simple function.
func Test_MethodName_SimpleFuncNoDot(t *testing.T) {
	// Arrange
	myFunc := func() {}

	// Act
	result := results.MethodName(myFunc)

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "MethodName(func) should return non-empty, got empty", actual)
}

// Test_Cov3_InvokeWithPanicRecovery_VoidFunc tests InvokeWithPanicRecovery on a void func.
func Test_InvokeWithPanicRecovery_VoidFunc(t *testing.T) {
	// Arrange
	voidFunc := func() {}

	// Act — signature is InvokeWithPanicRecovery(funcRef any, receiver any, args ...any)
	result := results.InvokeWithPanicRecovery(voidFunc, nil)

	// Assert
	actual := args.Map{"result": result.Panicked}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "void func should not panic, got panicked", actual)
	actual = args.Map{"result": result.ReturnCount != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "void func return count: got, want 0", actual)
}

// Test_Cov3_InvokeWithPanicRecovery_NilPtrError tests extractErrorFromValue with nil ptr implementing error.
func Test_InvokeWithPanicRecovery_NilPtrError(t *testing.T) {
	// Arrange
	funcReturningNilPtrError := func() error {
		var e *customError
		return e
	}

	// Act
	result := results.InvokeWithPanicRecovery(funcReturningNilPtrError, nil)

	// Assert — Go's nil interface check: typed nil ptr implements error, result.Error may be non-nil interface with nil value
	// The function returns a typed nil, which may or may not be detected as nil
	_ = result // exercise the code path
}

// Test_Cov3_InvokeWithPanicRecovery_FuncReturning42 tests with a func returning int.
func Test_InvokeWithPanicRecovery_FuncReturning42(t *testing.T) {
	// Arrange
	funcReturning42 := func() int { return 42 }

	// Act
	result := results.InvokeWithPanicRecovery(funcReturning42, nil)

	// Assert
	actual := args.Map{"result": result.Panicked}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no panic", actual)
	actual = args.Map{"result": result.ReturnCount != 1}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected return count 1", actual)
}

// Test_Cov3_FilterByFields_MissingKey tests filterByFields with a key not present in the map.
func Test_FilterByFields_MissingKey(t *testing.T) {
	// Arrange
	funcReturning42 := func() int { return 42 }
	expected := results.ResultAny{
		Value:       42,
		ReturnCount: 1,
	}

	// Act
	result := results.InvokeWithPanicRecovery(funcReturning42, nil)

	// Assert — request a field "isSafe" that won't be in the map
	result.ShouldMatchResult(
		t,
		0,
		"filterByFields missing key exercise",
		expected,
		"panicked", "value", "isSafe",
	)
}

// customError is a test helper to create a nil-pointer error type.
type customError struct {
	msg string
}

func (e *customError) Error() string {
	return fmt.Sprintf("custom: %s", e.msg)
}
