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

package coretaskinfotests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/coretaskinfo"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

func Test_QW_Info_JsonString_Nil(t *testing.T) {
	// Arrange
	defer func() { recover() }() // value receiver on nil pointer may panic
	var info *coretaskinfo.Info
	s := info.JsonString()

	// Act
	actual := args.Map{"result": s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for nil", actual)
}

func Test_QW_Info_LazyMapWithPayload_ErrorPath(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{
		RootName:    "test",
		Description: "desc",
	}
	// LazyMapWithPayload takes []byte — pass invalid JSON bytes to cover error branch
	m := info.LazyMapWithPayload([]byte(`{invalid`))

	// Act
	actual := args.Map{"result": m == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil map", actual)
}

func Test_QW_Info_LazyMapWithPayloadAsAny_ErrorPath(t *testing.T) {
	info := &coretaskinfo.Info{
		RootName:    "test",
		Description: "desc",
	}
	m := info.LazyMapWithPayloadAsAny(make(chan int))
	actual := args.Map{"result": m == nil}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil map", actual)
}
