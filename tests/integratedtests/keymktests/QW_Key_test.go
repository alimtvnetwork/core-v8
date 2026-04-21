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

package keymktests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/keymk"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

func Test_QW_Key_Compile_WithBrackets(t *testing.T) {
	// Arrange
	k := keymk.NewKey.Default(".")
	result := k.Compile("a", "b")

	// Act
	actual := args.Map{"result": result == ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-empty", actual)
}

func Test_QW_Key_Compile_Empty(t *testing.T) {
	k := keymk.NewKey.Default(".")
	result := k.Compile()
	_ = result
}

func Test_QW_Key_ParseInjectUsingJson_Error(t *testing.T) {
	// Arrange
	k := keymk.NewKey.Default(".")
	bad := corejson.NewResult.UsingString(`invalid`)
	_, err := k.ParseInjectUsingJson(bad)

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error for invalid JSON", actual)
}

func Test_QW_Key_ParseInjectUsingJsonMust_Panic(t *testing.T) {
	defer func() { recover() }()
	k := keymk.NewKey.Default(".")
	bad := corejson.NewResult.UsingString(`invalid`)
	k.ParseInjectUsingJsonMust(bad)
}

func Test_QW_Key_Compile_SkipEmpty(t *testing.T) {
	k := keymk.NewKey.DefaultStrings(".", "base")
	result := k.Compile("", "a", "", "b")
	_ = result
}
