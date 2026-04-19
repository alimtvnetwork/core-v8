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

package coreversiontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coreversion"
)

func Test_Creation_Verification(t *testing.T) {
	for caseIndex, testCase := range versionCreationTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]coreversion.Version)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, input := range inputs {
			if input.IsSafeInvalidCheck() {
				actualSlice.AppendFmt(
					defaultInvalidV2CreationFmt,
					i,
					input.String(),
				)
			} else {
				actualSlice.AppendFmt(
					defaultCreationFmt,
					i,
					input.String(),
					input.VersionCompact,
					input.VersionDisplay(),
				)
			}
		}

		finalActLines := actualSlice.Strings()
		finalCase := testCase.AsCaseV1()

		// Assert
		finalCase.ShouldBeTrimEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}

func Test_Creation_UsingString_Verification(t *testing.T) {
	for caseIndex, testCase := range versionCreationUsingStringTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]string)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))
		creatorFunc := coreversion.New.Default

		// Act
		for i, input := range inputs {
			toVersion := creatorFunc(input)

			if toVersion.IsSafeInvalidCheck() {
				actualSlice.AppendFmt(
					defaultInvalidV1CreationFmt,
					i,
					toVersion.String(),
					input,
				)
			} else {
				actualSlice.AppendFmt(
					defaultCreationFmt,
					i,
					toVersion.String(),
					toVersion.VersionCompact,
					toVersion.VersionDisplay(),
				)
			}
		}

		finalActLines := actualSlice.Strings()
		finalCase := testCase.AsCaseV1()

		// Assert
		finalCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
