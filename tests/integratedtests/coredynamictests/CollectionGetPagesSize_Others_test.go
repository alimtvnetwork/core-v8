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

package coredynamictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

func Test_AnyCollection_GetPagesSize_Verification(t *testing.T) {
	for caseIndex, testCase := range anyCollectionGetPagesSizeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.GetAsInt("count")
		eachPageSize, _ := input.GetAsInt("eachPageSize")

		collection := coredynamic.NewAnyCollection(count)
		for i := 0; i < count; i++ {
			collection.Add(i)
		}

		// Act
		result := fmt.Sprintf("%v", collection.GetPagesSize(eachPageSize))

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_DynamicCollection_GetPagesSize_Verification(t *testing.T) {
	for caseIndex, testCase := range dynamicCollectionGetPagesSizeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.GetAsInt("count")
		eachPageSize, _ := input.GetAsInt("eachPageSize")

		collection := coredynamic.NewDynamicCollection(count)
		for i := 0; i < count; i++ {
			collection.Add(coredynamic.NewDynamicValid(i))
		}

		// Act
		result := fmt.Sprintf("%v", collection.GetPagesSize(eachPageSize))

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_KeyValCollection_GetPagesSize_Verification(t *testing.T) {
	for caseIndex, testCase := range keyValCollectionGetPagesSizeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		count, _ := input.GetAsInt("count")
		eachPageSize, _ := input.GetAsInt("eachPageSize")

		collection := coredynamic.NewKeyValCollection(count)
		for i := 0; i < count; i++ {
			collection.Add(coredynamic.KeyVal{
				Key:   fmt.Sprintf("key-%d", i),
				Value: i,
			})
		}

		// Act
		result := fmt.Sprintf("%v", collection.GetPagesSize(eachPageSize))

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}
