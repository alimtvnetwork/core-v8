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

package simplewraptests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/simplewrap"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_WithStartEnd_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.WithStartEnd("[", "hello")

	// Assert
	actual := args.Map{"result": strings.Contains(result, "[")}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "should contain bracket wrapper", actual)
}

func Test_WithBracketsQuotation_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.WithBracketsQuotation("hello")

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_WithCurlyQuotation_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.WithCurlyQuotation("hello")

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_WithParenthesisQuotation_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.WithParenthesisQuotation("hello")

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_TitleSquareCsvMeta_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.TitleSquareCsvMeta("title", "a", "b")

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_TitleQuotationMeta_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.TitleQuotationMeta("title", "value", "meta")

	// Assert
	actual := args.Map{"result": result == ""}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "should not be empty", actual)
}

func Test_DoubleQuoteWrapElements_Verification_Ext2(t *testing.T) {
	// Act
	result := simplewrap.DoubleQuoteWrapElements(false, "a", "b")

	// Assert
	actual := args.Map{"result": len(result) != 2}
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 elements", actual)
}
