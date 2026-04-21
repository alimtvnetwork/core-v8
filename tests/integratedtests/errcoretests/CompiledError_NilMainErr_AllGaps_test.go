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

package errcoretests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/errcore"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage17 — errcore remaining 11 lines
// ══════════════════════════════════════════════════════════════════════════════

// ── CompiledError: nil mainErr (line 30-32) ──

func Test_CompiledError_NilMainErr(t *testing.T) {
	// Arrange
	var nilErr error

	// Act
	compiled := errcore.CompiledError(nilErr, "extra message")

	// Assert
	actual := args.Map{"isNil": compiled == nil}
	expected := args.Map{"isNil": true}
	actual.ShouldBeEqual(t, 1, "CompiledError nil main error", expected)
}

// ── RawErrCollection.CompiledJsonErrorWithStackTraces (line 237) ──

func Test_RawErrCollection_CompiledJsonErrorWithStackTraces(t *testing.T) {
	// Arrange
	coll := &errcore.RawErrCollection{
		Items: []error{errors.New("err1"), errors.New("err2")},
	}

	// Act
	compiledErr := coll.CompiledJsonErrorWithStackTraces()

	// Assert
	actual := args.Map{"hasError": compiledErr != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "RawErrCollection CompiledJsonErrorWithStackTraces", expected)
}

// ── RawErrCollection.CompiledJsonStringWithStackTraces (line 243-245) ──

func Test_RawErrCollection_CompiledJsonStringWithStackTraces(t *testing.T) {
	// Arrange
	coll := &errcore.RawErrCollection{
		Items: []error{errors.New("err1")},
	}

	// Act
	result := coll.CompiledJsonStringWithStackTraces()

	// Assert
	actual := args.Map{"hasContent": len(result) > 0}
	expected := args.Map{"hasContent": true}
	actual.ShouldBeEqual(t, 1, "RawErrCollection CompiledJsonStringWithStackTraces", expected)
}

// ── RawErrCollection.SerializeWithoutTraces (line 449-455) ──

func Test_RawErrCollection_SerializeWithoutTraces_FromCompiledErrorNilMain(t *testing.T) {
	// Arrange
	coll := &errcore.RawErrCollection{
		Items: []error{errors.New("test error")},
	}

	// Act
	rawBytes, err := coll.SerializeWithoutTraces()

	// Assert
	actual := args.Map{
		"hasBytes": len(rawBytes) > 0,
		"hasError": err != nil,
	}
	expected := args.Map{
		"hasBytes": true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "RawErrCollection SerializeWithoutTraces", expected)
}

// ── RawErrCollection.Serialize (line 458-464) ──

func Test_RawErrCollection_Serialize(t *testing.T) {
	// Arrange
	coll := &errcore.RawErrCollection{
		Items: []error{errors.New("test error")},
	}

	// Act
	rawBytes, err := coll.Serialize()

	// Assert
	actual := args.Map{
		"hasBytes": len(rawBytes) > 0,
		"hasError": err != nil,
	}
	expected := args.Map{
		"hasBytes": true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "RawErrCollection Serialize", expected)
}

// ── RawErrCollection.SerializeMust panic path (line 468-470) ──

func Test_RawErrCollection_SerializeMust(t *testing.T) {
	// Arrange
	coll := &errcore.RawErrCollection{
		Items: []error{errors.New("test")},
	}

	// Act
	rawBytes := coll.SerializeMust()

	// Assert
	actual := args.Map{"hasBytes": len(rawBytes) > 0}
	expected := args.Map{"hasBytes": true}
	actual.ShouldBeEqual(t, 1, "RawErrCollection SerializeMust", expected)
}

// ── stackTraceEnhance: message without stack trace (line 115-121) ──
// This is an internal unexported method — coverage is achieved indirectly
// through ErrorWithCompiledTraceRef and similar public functions.
