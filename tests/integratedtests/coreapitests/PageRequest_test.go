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

package coreapitests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/coreapi"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

func Test_PageRequest_IsPageSizeEmpty(t *testing.T) {
	for caseIndex, tc := range pageRequestIsPageSizeEmptyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		result := fmt.Sprintf("%v", req.IsPageSizeEmpty())

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_PageRequest_IsPageIndexEmpty(t *testing.T) {
	for caseIndex, tc := range pageRequestIsPageIndexEmptyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		result := fmt.Sprintf("%v", req.IsPageIndexEmpty())

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_PageRequest_HasPageSize(t *testing.T) {
	for caseIndex, tc := range pageRequestHasPageSizeTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		result := fmt.Sprintf("%v", req.HasPageSize())

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_PageRequest_HasPageIndex(t *testing.T) {
	for caseIndex, tc := range pageRequestHasPageIndexTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		result := fmt.Sprintf("%v", req.HasPageIndex())

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_PageRequest_Clone_Nil(t *testing.T) {
	for caseIndex, tc := range pageRequestCloneNilTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		result := fmt.Sprintf("%v", req.Clone() == nil)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_PageRequest_Clone_Fields(t *testing.T) {
	for caseIndex, tc := range pageRequestCloneFieldsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		clone := req.Clone()
		actual := args.Map{
			"pageSize":  fmt.Sprintf("%v", clone.PageSize),
			"pageIndex": fmt.Sprintf("%v", clone.PageIndex),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PageRequest_Clone_Independence(t *testing.T) {
	for caseIndex, tc := range pageRequestCloneIndependenceTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		clone := req.Clone()
		clone.PageSize = 99
		clone.PageIndex = 99
		actual := args.Map{
			"pageSize":  fmt.Sprintf("%v", req.PageSize),
			"pageIndex": fmt.Sprintf("%v", req.PageIndex),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
