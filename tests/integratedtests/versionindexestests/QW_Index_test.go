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

package versionindexestests

import (
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/enums/versionindexes"
)

func Test_QW_Index_JsonParseSelfInject_NilResult(t *testing.T) {
	// Arrange
	idx := versionindexes.Major

	// Act
	err := idx.JsonParseSelfInject(nil)

	// Assert
	actual := args.Map{
		"hasError": err != nil,
	}
	expected := args.Map{
		"hasError": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonParseSelfInject returns error -- nil result", actual)
}

func Test_QW_Index_JsonParseSelfInject_ErrorResult(t *testing.T) {
	// Arrange
	idx := versionindexes.Major
	bad := corejson.NewResult.UsingString(`invalid`)

	// Act
	err := idx.JsonParseSelfInject(bad)

	// Assert
	actual := args.Map{
		"hasError": err != nil,
	}
	expected := args.Map{
		"hasError": true,
	}
	expected.ShouldBeEqual(t, 0, "JsonParseSelfInject returns error -- invalid json", actual)
}
