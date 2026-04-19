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

package reflectinternal

import (
	"testing"
)

func TestTypeName(t *testing.T) {
	s := TypeName("hello")
	if s != "string" {
		t.Fatal("expected string")
	}

	s2 := TypeName(42)
	if s2 != "int" {
		t.Fatal("expected int")
	}

	s3 := TypeName(nil)
	if s3 != "" {
		t.Fatal("expected empty for nil")
	}
}

func TestTypeNames(t *testing.T) {
	names := TypeNames(true, "a", 1)
	if len(names) != 2 {
		t.Fatal("expected 2")
	}
	if names[0] != "string" {
		t.Fatal("expected string")
	}

	names2 := TypeNames(false, "a", 1)
	if len(names2) != 2 {
		t.Fatal("expected 2")
	}
}

func TestTypeNamesString(t *testing.T) {
	s := TypeNamesString(true, "a", 1)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestTypeNamesReferenceString(t *testing.T) {
	s := TypeNamesReferenceString(true, "a", 1)
	if s == "" {
		t.Fatal("expected non-empty")
	}
}

func TestTypeNameToValidVariableName(t *testing.T) {
	// simple
	s := TypeNameToValidVariableName("string")
	if s != "string" {
		t.Fatal("expected string, got:", s)
	}

	// empty
	s2 := TypeNameToValidVariableName("")
	if s2 != "" {
		t.Fatal("expected empty")
	}

	// with dot (complex)
	s3 := TypeNameToValidVariableName("*pkg.Type")
	_ = s3

	// slice type
	s4 := TypeNameToValidVariableName("[]string")
	_ = s4
}
