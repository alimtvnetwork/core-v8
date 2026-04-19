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

package errcore

import (
	"errors"
	"testing"
)

// ══════════════════════════════════════════════════════════════════════════════
// Consolidated internal tests for unexported symbols in errcore.
// These tests MUST remain in the source package because they access
// unexported functions/fields: getReferenceMessage, typesNamesString,
// stackTraceEnhance, expectedString.
//
// Source: Coverage06_Formatters_test.go, Coverage09_Iteration14_test.go,
//         Coverage10_Iteration7_test.go
// ══════════════════════════════════════════════════════════════════════════════

// ── getReferenceMessage, typesNamesString (unexported functions) ──

func TestInternal_GetReferenceMessage_Nil(t *testing.T) {
	if getReferenceMessage(nil) != "" {
		t.Fatal("expected empty")
	}
}

func TestInternal_GetReferenceMessage_EmptyString(t *testing.T) {
	if getReferenceMessage("") != "" {
		t.Fatal("expected empty")
	}
}

func TestInternal_GetReferenceMessage_WithRef(t *testing.T) {
	if getReferenceMessage("ref") == "" {
		t.Fatal("expected non-empty")
	}
}

func TestInternal_TypesNamesString(t *testing.T) {
	if typesNamesString("a", 1) == "" {
		t.Fatal("expected non-empty")
	}
}

// ── stackTraceEnhance (unexported struct) ──

func TestInternal_StackTraceEnhance_MsgErrorSkip_AlreadyTraced(t *testing.T) {
	ste := stackTraceEnhance{}
	msg := "some error\nStack-Trace: already"
	result := ste.MsgErrorSkip(0, msg, errors.New("wrapped"))
	if result == "" {
		t.Fatal("expected non-empty result")
	}
}

func TestInternal_StackTraceEnhance_MsgErrorSkip_NilErr(t *testing.T) {
	ste := stackTraceEnhance{}
	result := ste.MsgErrorSkip(0, "msg", nil)
	if result != "" {
		t.Fatal("expected empty for nil error")
	}
}

func TestInternal_StackTraceEnhance_MsgErrorToErrSkip_NilErr(t *testing.T) {
	ste := stackTraceEnhance{}
	err := ste.MsgErrorToErrSkip(0, "msg", nil)
	if err != nil {
		t.Fatal("expected nil for nil error")
	}
}

func TestInternal_StackTraceEnhance_FmtSkip_Empty(t *testing.T) {
	ste := stackTraceEnhance{}
	err := ste.FmtSkip(0, "")
	if err != nil {
		t.Fatal("expected nil for empty format")
	}
}

// ── ExpectationMessageDef.expectedString (unexported field) ──

func TestInternal_ExpectationMessageDef_ExpectedSafeString_Cached(t *testing.T) {
	cached := "pre-cached"
	def := ExpectationMessageDef{
		Expected:       "test",
		expectedString: &cached,
	}
	result := def.ExpectedSafeString()
	if result != "pre-cached" {
		t.Fatal("expected pre-cached")
	}
}
