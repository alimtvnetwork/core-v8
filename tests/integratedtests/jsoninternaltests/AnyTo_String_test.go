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

package jsoninternaltests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/jsoninternal"
)

// ── AnyTo — String (not pretty) ──

func Test_AnyTo_String(t *testing.T) {
	// Act
	result, err := jsoninternal.Pretty.AnyTo.String(map[string]int{"a": 1})
	emptyResult, _ := jsoninternal.Pretty.AnyTo.String(nil)

	// Assert
	actual := args.Map{
		"notEmpty": result != "",
		"noErr": err == nil,
		"empty": emptyResult,
	}
	expected := args.Map{
		"notEmpty": true,
		"noErr": true,
		"empty": "null",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- String", actual)
}

func Test_AnyTo_SafeString(t *testing.T) {
	// Act
	result := jsoninternal.Pretty.AnyTo.SafeString(map[string]int{"a": 1})
	emptyResult := jsoninternal.Pretty.AnyTo.SafeString(nil)

	// Assert
	actual := args.Map{
		"notEmpty": result != "",
		"empty": emptyResult,
	}
	expected := args.Map{
		"notEmpty": true,
		"empty": "null",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- SafeString", actual)
}

// ── AnyTo — PrettyString ──

func Test_AnyTo_PrettyString(t *testing.T) {
	// Act
	result, err := jsoninternal.Pretty.AnyTo.PrettyString("", map[string]int{"a": 1})
	emptyResult, _ := jsoninternal.Pretty.AnyTo.PrettyString("", nil)

	// Assert
	actual := args.Map{
		"notEmpty": result != "",
		"noErr": err == nil,
		"empty": emptyResult,
	}
	expected := args.Map{
		"notEmpty": true,
		"noErr": true,
		"empty": "null",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- PrettyString", actual)
}

func Test_AnyTo_SafePrettyString(t *testing.T) {
	// Act
	result := jsoninternal.Pretty.AnyTo.SafePrettyString("", map[string]int{"a": 1})
	emptyResult := jsoninternal.Pretty.AnyTo.SafePrettyString("", nil)

	// Assert
	actual := args.Map{
		"notEmpty": result != "",
		"empty": emptyResult,
	}
	expected := args.Map{
		"notEmpty": true,
		"empty": "null",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- SafePrettyString", actual)
}

func Test_AnyTo_PrettyStringDefault(t *testing.T) {
	// Act
	result := jsoninternal.Pretty.AnyTo.PrettyStringDefault(map[string]int{"a": 1})
	emptyResult := jsoninternal.Pretty.AnyTo.PrettyStringDefault(nil)

	// Assert
	actual := args.Map{
		"notEmpty": result != "",
		"empty": emptyResult,
	}
	expected := args.Map{
		"notEmpty": true,
		"empty": "null",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- PrettyStringDefault", actual)
}

// ── AnyTo — PrettyStringIndent ──

func Test_AnyTo_PrettyStringIndent(t *testing.T) {
	// Act
	result, err := jsoninternal.Pretty.AnyTo.PrettyStringIndent("", "\t", map[string]int{"a": 1})
	emptyResult, _ := jsoninternal.Pretty.AnyTo.PrettyStringIndent("", "\t", nil)

	// Assert
	actual := args.Map{
		"notEmpty": result != "",
		"noErr": err == nil,
		"empty": emptyResult,
	}
	expected := args.Map{
		"notEmpty": true,
		"noErr": true,
		"empty": "null",
	}
	expected.ShouldBeEqual(t, 0, "AnyTo returns correct value -- PrettyStringIndent", actual)
}

// ── Bytes — SafeDefault / Indent ──

func Test_Bytes_SafeDefault(t *testing.T) {
	// Act
	result := jsoninternal.Pretty.Bytes.SafeDefault([]byte(`{"a":1}`))

	// Assert
	actual := args.Map{"notEmpty": result != ""}
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Bytes returns correct value -- SafeDefault", actual)
}

func Test_Bytes_Indent(t *testing.T) {
	// Act
	result, err := jsoninternal.Pretty.Bytes.Indent("", "\t", []byte(`{"a":1}`))

	// Assert
	actual := args.Map{
		"notEmpty": result != "",
		"noErr": err == nil,
	}
	expected := args.Map{
		"notEmpty": true,
		"noErr": true,
	}
	expected.ShouldBeEqual(t, 0, "Bytes returns correct value -- Indent", actual)
}
