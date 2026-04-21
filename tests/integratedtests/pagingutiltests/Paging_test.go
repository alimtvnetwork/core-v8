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

package pagingutiltests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/pagingutil"
)

func Test_GetPagesSize_Verification(t *testing.T) {
	for caseIndex, testCase := range getPagesSizeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		eachPageSize, _ := input.GetAsInt("eachPageSize")
		totalLength, _ := input.GetAsInt("totalLength")

		// Act
		result := pagingutil.GetPagesSize(eachPageSize, totalLength)

		// Assert
		actual := args.Map{
			"pagesSize": result,
		}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_GetPagingInfo_Verification(t *testing.T) {
	for caseIndex, testCase := range getPagingInfoTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		length, _ := input.GetAsInt("length")
		pageIndex, _ := input.GetAsInt("pageIndex")
		eachPageSize, _ := input.GetAsInt("eachPageSize")

		// Act
		info := pagingutil.GetPagingInfo(pagingutil.PagingRequest{
			Length:       length,
			PageIndex:    pageIndex,
			EachPageSize: eachPageSize,
		})

		// Assert
		actual := args.Map{
			"pageIndex":        info.PageIndex,
			"skipItems":        info.SkipItems,
			"endingLength":     info.EndingLength,
			"isPagingPossible": info.IsPagingPossible,
		}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
