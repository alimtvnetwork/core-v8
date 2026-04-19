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

package iserrortests

import (
	"errors"
	"os/exec"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/iserror"
)

// ── Defined / Empty additional coverage ──

func Test_Defined_Error(t *testing.T) {
	// Act
	actual := args.Map{"result": iserror.Defined(errors.New("x"))}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Defined error -- true", actual)
}

func Test_Defined_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": iserror.Defined(nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Defined nil -- false", actual)
}

func Test_Empty_Error(t *testing.T) {
	// Act
	actual := args.Map{"result": iserror.Empty(errors.New("x"))}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Empty error -- false", actual)
}

func Test_Empty_Nil(t *testing.T) {
	// Act
	actual := args.Map{"result": iserror.Empty(nil)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Empty nil -- true", actual)
}

// ── ExitError with actual ExitError ──

func Test_ExitError_RealExitError(t *testing.T) {
	// Arrange
	exitErr := &exec.ExitError{}

	// Act
	actual := args.Map{"result": iserror.ExitError(exitErr)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "ExitError real ExitError -- true", actual)
}

// ── Equal with same error message different instances ──

func Test_Equal_SameMessage(t *testing.T) {
	// Arrange
	e1 := errors.New("same msg")
	e2 := errors.New("same msg")

	// Act
	actual := args.Map{"result": iserror.Equal(e1, e2)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Equal same message different instances -- true", actual)
}

// ── NotEqual both nil ──

func Test_NotEqual_BothNil(t *testing.T) {
	// Act
	actual := args.Map{"result": iserror.NotEqual(nil, nil)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotEqual both nil -- false", actual)
}

// ── NotEqualString same ──

func Test_NotEqualString_EmptyStrings(t *testing.T) {
	// Act
	actual := args.Map{"result": iserror.NotEqualString("", "")}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "NotEqualString empty -- false", actual)
}
