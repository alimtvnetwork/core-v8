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

package ostypetests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/ostype"
)

// Test_Cov4_Group_UnmarshallEnumToValue covers
// ostype/Group.go L191-193: UnmarshallEnumToValue.
func Test_Group_UnmarshallEnumToValue(t *testing.T) {
	// Arrange
	g := ostype.UnixGroup
	validBytes := []byte("1")

	// Act
	val, err := g.UnmarshallEnumToValue(validBytes)

	// Assert
	actual := args.Map{
		"value": val,
		"hasError": err != nil,
	}
	expected := args.Map{
		"value": byte(ostype.UnixGroup),
		"hasError": false,
	}
	expected.ShouldBeEqual(t, 0, "Group returns correct value -- UnmarshallEnumToValue", actual)
}
