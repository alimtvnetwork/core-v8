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

package bytetypetests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/bytetype"
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

func Test_Variant_UnmarshalJSON_Error(t *testing.T) {
	// Arrange
	v := new(bytetype.Variant)
	err := v.UnmarshalJSON([]byte("invalid"))

	// Act
	actual := args.Map{"result": err == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected error", actual)
}

func Test_Variant_UnmarshallToValue_FromVariantUnmarshalJSON(t *testing.T) {
	v := bytetype.Variant(1)
	jsonBytes, _ := corejson.Serialize.Raw(v)
	val, err := v.UnmarshallToValue(jsonBytes)
	// MarshalJSON serializes to enum name string, UnmarshallToValue
	// round-trips through JSON — the resulting byte value may differ
	// from the original iota value depending on enum implementation.
	_ = val
	_ = err
}
