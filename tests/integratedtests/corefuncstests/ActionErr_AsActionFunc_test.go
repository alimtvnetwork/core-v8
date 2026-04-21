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

package corefuncstests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core-v8/corefuncs"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ── ActionReturnsErrorFuncWrapper — AsActionFunc ──

func Test_ActionErr_AsActionFunc(t *testing.T) {
	// Arrange
	w := corefuncs.New.ActionErr("test", func() error { return nil })

	// Act — should not panic
	w.AsActionFunc()()

	// Assert
	actual := args.Map{"called": true}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "ActionErr AsActionFunc -- no panic", actual)
}

func Test_ActionErr_AsActionReturnsErrorFunc_Success(t *testing.T) {
	// Arrange
	w := corefuncs.New.ActionErr("test", func() error { return nil })

	// Act
	err := w.AsActionReturnsErrorFunc()()

	// Assert
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "ActionErr success -- nil", actual)
}

// ── InOutErrFuncWrapper — AsActionFunc / AsActionReturnsErrorFunc ──

func Test_LegacyInOutErr_AsActionFunc(t *testing.T) {
	// Arrange
	w := corefuncs.New.LegacyInOutErr("test", func(input any) (any, error) {
		return nil, nil
	})

	// Act
	w.AsActionFunc("input")()

	// Assert
	actual := args.Map{"called": true}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "LegacyInOutErr returns error -- AsActionFunc", actual)
}

func Test_LegacyInOutErr_AsActionReturnsErrorFunc_Fail(t *testing.T) {
	// Arrange
	w := corefuncs.New.LegacyInOutErr("test", func(input any) (any, error) {
		return nil, errors.New("fail")
	})

	// Act
	err := w.AsActionReturnsErrorFunc("input")()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "LegacyInOutErr returns error -- failure", actual)
}

func Test_LegacyInOutErr_AsActionReturnsErrorFunc_Success(t *testing.T) {
	// Arrange
	w := corefuncs.New.LegacyInOutErr("test", func(input any) (any, error) {
		return "ok", nil
	})

	// Act
	err := w.AsActionReturnsErrorFunc("input")()

	// Assert
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "LegacyInOutErr returns error -- success", actual)
}

// ── ResultDelegatingFuncWrapper — AsActionFunc / AsActionReturnsErrorFunc ──

func Test_LegacyResultDelegating_AsActionFunc(t *testing.T) {
	// Arrange
	w := corefuncs.New.LegacyResultDelegating("test", func(target any) error {
		return nil
	})

	// Act
	w.AsActionFunc("target")()

	// Assert
	actual := args.Map{"called": true}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "LegacyResultDelegating returns correct value -- AsActionFunc", actual)
}

func Test_LegacyResultDelegating_AsActionReturnsErrorFunc_Fail(t *testing.T) {
	// Arrange
	w := corefuncs.New.LegacyResultDelegating("test", func(target any) error {
		return errors.New("fail")
	})

	// Act
	err := w.AsActionReturnsErrorFunc("target")()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "LegacyResultDelegating returns correct value -- failure", actual)
}

func Test_LegacyResultDelegating_AsActionReturnsErrorFunc_Success(t *testing.T) {
	// Arrange
	w := corefuncs.New.LegacyResultDelegating("test", func(target any) error {
		return nil
	})

	// Act
	err := w.AsActionReturnsErrorFunc("target")()

	// Assert
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "LegacyResultDelegating returns correct value -- success", actual)
}

// ── Generic wrappers — AsActionFunc / AsActionReturnsErrorFunc ──

func Test_InOutErrWrapperOf_AsActionFunc(t *testing.T) {
	// Arrange
	w := corefuncs.NewInOutErrWrapper[string, int]("test", func(s string) (int, error) {
		return 0, nil
	})

	// Act
	w.AsActionFunc("x")()

	// Assert
	actual := args.Map{"called": true}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "InOutErrWrapperOf returns error -- AsActionFunc", actual)
}

func Test_InOutErrWrapperOf_AsActionReturnsErrorFunc_Fail(t *testing.T) {
	// Arrange
	w := corefuncs.NewInOutErrWrapper[string, int]("test", func(s string) (int, error) {
		return 0, errors.New("fail")
	})

	// Act
	err := w.AsActionReturnsErrorFunc("x")()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "InOutErrWrapperOf returns error -- failure", actual)
}

func Test_InOutFuncWrapperOf_AsActionFunc(t *testing.T) {
	// Arrange
	w := corefuncs.NewInOutWrapper[string, int]("test", func(s string) int { return 0 })

	// Act
	w.AsActionFunc("x")()

	// Assert
	actual := args.Map{"called": true}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "InOutFuncWrapperOf returns correct value -- AsActionFunc", actual)
}

func Test_InOutFuncWrapperOf_AsActionReturnsErrorFunc(t *testing.T) {
	// Arrange
	w := corefuncs.NewInOutWrapper[string, int]("test", func(s string) int { return 0 })

	// Act
	err := w.AsActionReturnsErrorFunc("x")()

	// Assert
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "InOutFuncWrapperOf returns nil -- returns nil", actual)
}

func Test_InActionErrWrapperOf_AsActionFunc(t *testing.T) {
	// Arrange
	w := corefuncs.NewInActionErrWrapper[string]("test", func(s string) error { return nil })

	// Act
	w.AsActionFunc("x")()

	// Assert
	actual := args.Map{"called": true}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "InActionErrWrapperOf returns error -- AsActionFunc", actual)
}

func Test_InActionErrWrapperOf_AsActionReturnsErrorFunc_Fail(t *testing.T) {
	// Arrange
	w := corefuncs.NewInActionErrWrapper[string]("test", func(s string) error {
		return errors.New("fail")
	})

	// Act
	err := w.AsActionReturnsErrorFunc("x")()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "InActionErrWrapperOf returns error -- failure", actual)
}

func Test_InActionErrWrapperOf_AsActionReturnsErrorFunc_Success(t *testing.T) {
	// Arrange
	w := corefuncs.NewInActionErrWrapper[string]("test", func(s string) error { return nil })

	// Act
	err := w.AsActionReturnsErrorFunc("x")()

	// Assert
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "InActionErrWrapperOf returns error -- success", actual)
}

func Test_ResultDelegatingWrapperOf_AsActionFunc(t *testing.T) {
	// Arrange
	w := corefuncs.NewResultDelegatingWrapper[*string]("test", func(t *string) error { return nil })

	// Act
	var s string
	w.AsActionFunc(&s)()

	// Assert
	actual := args.Map{"called": true}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "ResultDelegatingWrapperOf returns correct value -- AsActionFunc", actual)
}

func Test_ResultDelegatingWrapperOf_AsActionReturnsErrorFunc_Fail(t *testing.T) {
	// Arrange
	w := corefuncs.NewResultDelegatingWrapper[*string]("test", func(t *string) error {
		return errors.New("fail")
	})

	// Act
	var s string
	err := w.AsActionReturnsErrorFunc(&s)()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "ResultDelegatingWrapperOf returns correct value -- failure", actual)
}

func Test_SerializeWrapperOf_AsActionReturnsErrorFunc_Fail(t *testing.T) {
	// Arrange
	w := corefuncs.NewSerializeWrapper[string]("test", func(s string) ([]byte, error) {
		return nil, errors.New("fail")
	})

	// Act
	err := w.AsActionReturnsErrorFunc("x")()

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	expected.ShouldBeEqual(t, 0, "SerializeWrapperOf returns correct value -- failure", actual)
}

func Test_SerializeWrapperOf_AsActionReturnsErrorFunc_Success(t *testing.T) {
	// Arrange
	w := corefuncs.NewSerializeWrapper[string]("test", func(s string) ([]byte, error) {
		return []byte(s), nil
	})

	// Act
	err := w.AsActionReturnsErrorFunc("x")()

	// Assert
	actual := args.Map{"isNil": err == nil}
	expected := args.Map{"isNil": true}
	expected.ShouldBeEqual(t, 0, "SerializeWrapperOf returns correct value -- success", actual)
}

// ── NamedActionFuncWrapper — AsActionFunc ──

func Test_NamedAction_AsActionFunc(t *testing.T) {
	// Arrange
	var called bool
	w := corefuncs.New.NamedAction("test", func(name string) { called = true })

	// Act
	w.AsActionFunc()()

	// Assert
	actual := args.Map{"called": called}
	expected := args.Map{"called": true}
	expected.ShouldBeEqual(t, 0, "NamedAction returns correct value -- AsActionFunc", actual)
}
