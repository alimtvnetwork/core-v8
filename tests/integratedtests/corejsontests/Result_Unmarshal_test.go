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

package corejsontests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

func Test_Result_Unmarshal_Valid_FromResultUnmarshal(t *testing.T) {
	tc := resultUnmarshalValidTestCase

	// Arrange
	src := exampleStruct{Name: "Alice", Age: 30}
	jsonResult := corejson.NewPtr(src)
	target := &exampleStruct{}

	// Act
	err := jsonResult.Unmarshal(target)

	actual := args.Map{
		"error":            fmt.Sprintf("%v", err),
		"deserializedName": target.Name,
		"deserializedAge":  fmt.Sprintf("%v", target.Age),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Result_Unmarshal_NilReceiver_ResultUnmarshal(t *testing.T) {
	tc := resultUnmarshalNilTestCase

	// Arrange
	var nilResult *corejson.Result
	target := &exampleStruct{}

	// Act
	err := nilResult.Unmarshal(target)

	actual := args.Map{
		"hasError":          err != nil,
		"errorContainsNull": strings.Contains(err.Error(), "null"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Result_Unmarshal_InvalidBytes(t *testing.T) {
	tc := resultUnmarshalInvalidTestCase

	// Arrange
	result := corejson.NewResult.UsingBytesTypePtr([]byte(`{invalid-json`), "TestType")
	target := &exampleStruct{}

	// Act
	err := result.Unmarshal(target)

	actual := args.Map{
		"hasError":               err != nil,
		"errorContainsUnmarshal": strings.Contains(err.Error(), "unmarshal"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Result_Unmarshal_ExistingError(t *testing.T) {
	tc := resultUnmarshalExistingErrorTestCase

	// Arrange
	ch := make(chan int)
	result := corejson.NewPtr(ch)
	target := &exampleStruct{}

	// Act
	err := result.Unmarshal(target)

	actual := args.Map{
		"hasError":               err != nil,
		"errorContainsUnmarshal": strings.Contains(err.Error(), "unmarshal"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
