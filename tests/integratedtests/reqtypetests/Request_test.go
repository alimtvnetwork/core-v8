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

package reqtypetests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/reqtype"
)

func Test_Request_Identity_Verification(t *testing.T) {
	for caseIndex, testCase := range requestIdentityTestCases {
		// Arrange
		input := testCase.ArrangeInput.(reqtype.Request)

		// Act
		actual := args.Map{
			"name":      input.Name(),
			"isValid":   fmt.Sprintf("%v", input.IsValid()),
			"isInvalid": fmt.Sprintf("%v", input.IsInvalid()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Request_LogicalGroups_Verification(t *testing.T) {
	for caseIndex, testCase := range requestLogicalGroupTestCases {
		// Arrange
		input := testCase.ArrangeInput.(reqtype.Request)

		// Act
		actual := args.Map{
			"isCreateLogically": fmt.Sprintf("%v", input.IsCreateLogically()),
			"isDropLogically":   fmt.Sprintf("%v", input.IsDropLogically()),
			"isCrudOnly":        fmt.Sprintf("%v", input.IsCrudOnlyLogically()),
			"isReadOrEdit":      fmt.Sprintf("%v", input.IsReadOrEditLogically()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Request_HttpMethods_Verification(t *testing.T) {
	for caseIndex, testCase := range requestHttpTestCases {
		// Arrange
		input := testCase.ArrangeInput.(reqtype.Request)

		// Act
		actual := args.Map{
			"isGet":    fmt.Sprintf("%v", input.IsGetHttp()),
			"isPost":   fmt.Sprintf("%v", input.IsPostHttp()),
			"isPut":    fmt.Sprintf("%v", input.IsPutHttp()),
			"isDelete": fmt.Sprintf("%v", input.IsDeleteHttp()),
			"isPatch":  fmt.Sprintf("%v", input.IsPatchHttp()),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
