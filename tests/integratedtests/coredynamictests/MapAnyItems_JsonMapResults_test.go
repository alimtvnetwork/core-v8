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

package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── MapAnyItems.JsonMapResults (line 903) ──

func Test_MapAnyItems_JsonMapResults_Valid_I29(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(3)
	m.Add("key1", "val1")
	m.Add("key2", 42)

	// Act
	results, err := m.JsonMapResults()

	// Assert
	actual := args.Map{
		"notNil":   results != nil,
		"hasError": err != nil,
	}
	expected := args.Map{
		"notNil":   true,
		"hasError": false,
	}
	actual.ShouldBeEqual(t, 1, "MapAnyItems JsonMapResults valid", expected)
}

// ── MapAnyItems.GetUsingUnmarshallAt type mismatch (line 350) ──

func Test_MapAnyItems_GetUsingUnmarshallAt_TypeMismatch_I29(t *testing.T) {
	// Arrange
	m := coredynamic.NewMapAnyItems(3)
	m.Add("key1", "stringValue")

	var target int

	// Act
	err := m.GetUsingUnmarshallAt("key1", &target)

	// Assert
	actual := args.Map{"hasError": err != nil}
	expected := args.Map{"hasError": true}
	actual.ShouldBeEqual(t, 1, "MapAnyItems GetUsingUnmarshallAt type mismatch", expected)
}
