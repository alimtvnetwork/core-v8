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

package enumimpltests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_EnumByte_MinMax(t *testing.T) {
	for caseIndex, tc := range enumByteTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		enumMap := input["enum-map"].(enumimpl.DynamicMap)
		enumImpl := enumMap.BasicByte("unknown type")

		// Act
		actual := args.Map{
			"min": fmt.Sprintf("%v", enumImpl.Min()),
			"max": fmt.Sprintf("%v", enumImpl.Max()),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_EnumInt8_MinMax(t *testing.T) {
	for caseIndex, tc := range enumInt8TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		enumMap := input["enum-map"].(enumimpl.DynamicMap)
		enumImpl := enumMap.BasicInt8("unknown type")

		// Act
		actual := args.Map{
			"min": fmt.Sprintf("%v", enumImpl.Min()),
			"max": fmt.Sprintf("%v", enumImpl.Max()),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_EnumInt16_MinMax(t *testing.T) {
	for caseIndex, tc := range enumInt16TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		enumMap := input["enum-map"].(enumimpl.DynamicMap)
		enumImpl := enumMap.BasicInt16("unknown type")

		// Act
		actual := args.Map{
			"min": fmt.Sprintf("%v", enumImpl.Min()),
			"max": fmt.Sprintf("%v", enumImpl.Max()),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_EnumInt32_MinMax(t *testing.T) {
	for caseIndex, tc := range enumInt32TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		enumMap := input["enum-map"].(enumimpl.DynamicMap)
		enumImpl := enumMap.BasicInt32("unknown type")

		// Act
		actual := args.Map{
			"min": fmt.Sprintf("%v", enumImpl.Min()),
			"max": fmt.Sprintf("%v", enumImpl.Max()),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_EnumUInt16_MinMax(t *testing.T) {
	for caseIndex, tc := range enumUInt16TestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		enumMap := input["enum-map"].(enumimpl.DynamicMap)
		enumImpl := enumMap.BasicUInt16("unknown type")

		// Act
		actual := args.Map{
			"min": fmt.Sprintf("%v", enumImpl.Min()),
			"max": fmt.Sprintf("%v", enumImpl.Max()),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_EnumString_MinMax(t *testing.T) {
	for caseIndex, tc := range enumStringTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		enumMap := input["enum-map"].(enumimpl.DynamicMap)
		enumImpl := enumMap.BasicString("unknown type")

		// Act
		actual := args.Map{
			"min": fmt.Sprintf("%v", enumImpl.Min()),
			"max": fmt.Sprintf("%v", enumImpl.Max()),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
