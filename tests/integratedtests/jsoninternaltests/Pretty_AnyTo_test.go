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

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/internal/jsoninternal"
)

func Test_Pretty_AnyTo_String(t *testing.T) {
	// Arrange
	item := map[string]string{"key": "value"}

	// Act
	str, err := jsoninternal.Pretty.AnyTo.String(item)
	safeStr := jsoninternal.Pretty.AnyTo.SafeString(item)
	nilStr := jsoninternal.Pretty.AnyTo.SafeString(nil)

	actual := args.Map{
		"noError":     err == nil,
		"strNotEmpty": str != "",
		"safeNotEmpty": safeStr != "",
		"nilEmpty":    nilStr,
	}
	expected := args.Map{
		"noError":     true,
		"strNotEmpty": true,
		"safeNotEmpty": true,
		"nilEmpty":    nilStr,
	}
	expected.ShouldBeEqual(t, 0, "Pretty_AnyTo_String returns correct value -- with args", actual)
}

func Test_Pretty_AnyTo_PrettyString(t *testing.T) {
	// Arrange
	item := map[string]string{"key": "value"}

	// Act
	pretty, err := jsoninternal.Pretty.AnyTo.PrettyString("", item)
	prettyIndent, err2 := jsoninternal.Pretty.AnyTo.PrettyStringIndent("", "  ", item)
	safePretty := jsoninternal.Pretty.AnyTo.SafePrettyString("", item)
	defaultPretty := jsoninternal.Pretty.AnyTo.PrettyStringDefault(item)
	nilPretty := jsoninternal.Pretty.AnyTo.PrettyStringDefault(nil)

	actual := args.Map{
		"noErr":            err == nil,
		"noErr2":           err2 == nil,
		"prettyNotEmpty":   pretty != "",
		"indentNotEmpty":   prettyIndent != "",
		"safeNotEmpty":     safePretty != "",
		"defaultNotEmpty":  defaultPretty != "",
		"nilEmpty":         nilPretty,
	}
	expected := args.Map{
		"noErr":            true,
		"noErr2":           true,
		"prettyNotEmpty":   true,
		"indentNotEmpty":   true,
		"safeNotEmpty":     true,
		"defaultNotEmpty":  true,
		"nilEmpty":         nilPretty,
	}
	expected.ShouldBeEqual(t, 0, "Pretty_AnyTo_PrettyString returns correct value -- with args", actual)
}

func Test_Pretty_Bytes(t *testing.T) {
	// Arrange
	jsonBytes := []byte(`{"key":"value"}`)

	// Act
	safe := jsoninternal.Pretty.Bytes.Safe("", jsonBytes)
	safeDefault := jsoninternal.Pretty.Bytes.SafeDefault(jsonBytes)
	prefix, err := jsoninternal.Pretty.Bytes.Prefix("", jsonBytes)
	indent, err2 := jsoninternal.Pretty.Bytes.Indent("", "  ", jsonBytes)

	actual := args.Map{
		"safeNotEmpty":    safe != "",
		"safeDefNotEmpty": safeDefault != "",
		"noErr":           err == nil,
		"noErr2":          err2 == nil,
		"prefixNotEmpty":  prefix != "",
		"indentNotEmpty":  indent != "",
	}
	expected := args.Map{
		"safeNotEmpty":    true,
		"safeDefNotEmpty": true,
		"noErr":           true,
		"noErr2":          true,
		"prefixNotEmpty":  true,
		"indentNotEmpty":  true,
	}
	expected.ShouldBeEqual(t, 0, "Pretty_Bytes returns correct value -- with args", actual)
}

func Test_Pretty_String(t *testing.T) {
	// Arrange
	jsonStr := `{"key":"value"}`

	// Act
	safe := jsoninternal.Pretty.String.Safe("", jsonStr)
	safeDefault := jsoninternal.Pretty.String.SafeDefault(jsonStr)
	prefix, err := jsoninternal.Pretty.String.Prefix("", jsonStr)
	indent, err2 := jsoninternal.Pretty.String.Indent("", "  ", jsonStr)

	actual := args.Map{
		"safeNotEmpty":    safe != "",
		"safeDefNotEmpty": safeDefault != "",
		"noErr":           err == nil,
		"noErr2":          err2 == nil,
		"prefixNotEmpty":  prefix != "",
		"indentNotEmpty":  indent != "",
	}
	expected := args.Map{
		"safeNotEmpty":    true,
		"safeDefNotEmpty": true,
		"noErr":           true,
		"noErr2":          true,
		"prefixNotEmpty":  true,
		"indentNotEmpty":  true,
	}
	expected.ShouldBeEqual(t, 0, "Pretty_String returns correct value -- with args", actual)
}

func Test_StringJsonConverter(t *testing.T) {
	// Arrange
	item := map[string]string{"key": "value"}

	// Act
	defStr, err := jsoninternal.String.Default(item)
	safeStr := jsoninternal.String.SafeDefault(item)
	prettyStr, err2 := jsoninternal.String.Pretty(item)
	strVal := jsoninternal.String.StringValue("hello")

	actual := args.Map{
		"noErr":            err == nil,
		"noErr2":           err2 == nil,
		"defNotEmpty":      defStr != "",
		"safeNotEmpty":     safeStr != "",
		"prettyNotEmpty":   prettyStr != "",
		"strValNotEmpty":   len(strVal) > 0,
	}
	expected := args.Map{
		"noErr":            true,
		"noErr2":           true,
		"defNotEmpty":      true,
		"safeNotEmpty":     true,
		"prettyNotEmpty":   true,
		"strValNotEmpty":   true,
	}
	expected.ShouldBeEqual(t, 0, "StringJsonConverter returns correct value -- with args", actual)
}
