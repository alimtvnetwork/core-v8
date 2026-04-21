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

package regexnewtests

import (
	"regexp"
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/regexnew"
)

// ══════════════════════════════════════════════════════════════════════════════
// LazyRegex — basic operations
// ══════════════════════════════════════════════════════════════════════════════

func Test_LazyRegex_Compile_Valid_FromLazyRegexCompile(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`^cov7\d+$`)
	re, err := lr.Compile()

	// Act
	actual := args.Map{
		"notNil": re != nil,
		"noErr": err == nil,
	}

	// Assert
	expected := args.Map{
		"notNil": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Compile returns non-empty -- valid", actual)
}

func Test_LazyRegex_Compile_Invalid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`[invalid`)
	re, err := lr.Compile()

	// Act
	actual := args.Map{
		"nil": re == nil,
		"hasErr": err != nil,
	}

	// Assert
	expected := args.Map{
		"nil": true,
		"hasErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Compile returns error -- invalid", actual)
}

func Test_LazyRegex_IsMatch_True(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov7m\d+$`)

	// Act
	actual := args.Map{"v": lr.IsMatch("cov7m123")}

	// Assert
	expected := args.Map{"v": true}
	expected.ShouldBeEqual(t, 0, "IsMatch returns correct value -- true", actual)
}

func Test_LazyRegex_IsMatch_False(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov7m\d+$`)

	// Act
	actual := args.Map{"v": lr.IsMatch("notmatch")}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsMatch returns correct value -- false", actual)
}

func Test_LazyRegex_MatchError_Valid_FromLazyRegexCompile(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov7me\d+$`)
	err := lr.MatchError("cov7me123")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError returns error -- valid", actual)
}

func Test_LazyRegex_MatchError_NoMatch_FromLazyRegexCompile(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov7me\d+$`)
	err := lr.MatchError("notmatch")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError returns empty -- no match", actual)
}

func Test_LazyRegex_MatchError_Undefined(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy("")
	err := lr.MatchError("test")

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchError returns error -- undefined", actual)
}

func Test_LazyRegex_MatchUsingFuncError_Undefined(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy("")
	err := lr.MatchUsingFuncError("test", func(re *regexp.Regexp, s string) bool {
		return false
	})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncError returns error -- undefined", actual)
}

func Test_LazyRegex_IsMatch_Undefined(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy("")

	// Act
	actual := args.Map{"v": lr.IsMatch("test")}

	// Assert
	expected := args.Map{"v": false}
	expected.ShouldBeEqual(t, 0, "IsMatch returns correct value -- undefined", actual)
}

func Test_LazyRegex_CompileMust_Valid(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov7cm\d+$`)
	re := lr.CompileMust()

	// Act
	actual := args.Map{"notNil": re != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "CompileMust returns non-empty -- valid", actual)
}

func Test_LazyRegex_CompileMust_Panic(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`[invalid`)
	panicked := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		lr.CompileMust()
	}()

	// Act
	actual := args.Map{"panicked": panicked}

	// Assert
	expected := args.Map{"panicked": true}
	expected.ShouldBeEqual(t, 0, "CompileMust panics -- panic", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MatchUsingCustomizeErrorFuncLock
// ══════════════════════════════════════════════════════════════════════════════

func Test_MatchUsingCustomizeErrorFuncLock_InvalidRegex(t *testing.T) {
	// Arrange
	err := regexnew.MatchUsingCustomizeErrorFuncLock(`[invalid`, "abc",
		func(re *regexp.Regexp, s string) bool { return false },
		nil,
	)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CustomErrLock returns error -- invalid regex", actual)
}

func Test_MatchUsingCustomizeErrorFuncLock_Valid(t *testing.T) {
	// Arrange
	err := regexnew.MatchUsingCustomizeErrorFuncLock(`^cov7cust\d+$`, "cov7cust123",
		func(re *regexp.Regexp, s string) bool { return re.MatchString(s) },
		nil,
	)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "CustomErrLock returns error -- valid", actual)
}

func Test_MatchUsingCustomizeErrorFuncLock_NoMatch(t *testing.T) {
	// Arrange
	err := regexnew.MatchUsingCustomizeErrorFuncLock(`^cov7cust\d+$`, "nomatch",
		func(re *regexp.Regexp, s string) bool { return re.MatchString(s) },
		nil,
	)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "CustomErrLock returns empty -- no match", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// MatchUsingFuncErrorLock
// ══════════════════════════════════════════════════════════════════════════════

func Test_MatchUsingFuncErrorLock_Valid(t *testing.T) {
	// Arrange
	err := regexnew.MatchUsingFuncErrorLock(`^cov7fl\d+$`, "cov7fl123",
		func(re *regexp.Regexp, s string) bool { return re.MatchString(s) },
	)

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncErrorLock returns error -- valid", actual)
}

func Test_MatchUsingFuncErrorLock_Invalid(t *testing.T) {
	// Arrange
	err := regexnew.MatchUsingFuncErrorLock(`[invalid`, "abc",
		func(re *regexp.Regexp, s string) bool { return false },
	)

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncErrorLock returns error -- invalid", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LazyRegex FullString
// ══════════════════════════════════════════════════════════════════════════════

func Test_LazyRegex_FullString_Valid_FromLazyRegexCompile(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov7fs\d+$`)
	_ = lr.CompileMust()
	s := lr.FullString()

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FullString returns non-empty -- valid", actual)
}

func Test_LazyRegex_FullString_Invalid(t *testing.T) {
	// Arrange
	lr := regexnew.New.Lazy(`[invalid`)
	s := lr.FullString()

	// Act
	actual := args.Map{"notEmpty": s != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "FullString returns error -- invalid pattern", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// LazyRegex MatchUsingFuncError with func returning false
// ══════════════════════════════════════════════════════════════════════════════

func Test_LazyRegex_MatchUsingFuncError_FuncReturnsFalse(t *testing.T) {
	// Arrange
	lr := regexnew.New.LazyLock(`^cov7mfe\d+$`)
	err := lr.MatchUsingFuncError("cov7mfe123", func(re *regexp.Regexp, s string) bool {
		return false
	})

	// Act
	actual := args.Map{"hasErr": err != nil}

	// Assert
	expected := args.Map{"hasErr": true}
	expected.ShouldBeEqual(t, 0, "MatchUsingFuncError returns error -- func returns false", actual)
}
