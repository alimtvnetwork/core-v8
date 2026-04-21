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

// Test_Cov3_JsonString_ZeroInfo exercises JsonString on a zero-value Info.
func Test_JsonString_ZeroInfo(t *testing.T) {
	// Arrange
	info := coretaskinfo.Info{}

	// Act
	actual := info.JsonString()

	// Assert — zero-value Info is not nil, just verify it doesn't panic
	_ = actual
}

// Test_Cov3_MapWithPayloadAsAny_SerializeError tests the HasError branch in MapWithPayloadAsAny.
func Test_MapWithPayloadAsAny_SerializeError(t *testing.T) {
	// Arrange
	info := &coretaskinfo.Info{
		RootName:    "test",
		Description: "desc",
	}

	// Act — pass a channel which cannot be JSON-marshalled
	ch := make(chan int)
	result := info.MapWithPayloadAsAny(ch)

	// Assert — should have a serializing error field
	_, hasPayloadsErr := result["Payloads.SerializingErr"]
	actual := args.Map{"result": hasPayloadsErr}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "MapWithPayloadAsAny with unmarshal-able payload should have error key, got keys:", actual)
}
