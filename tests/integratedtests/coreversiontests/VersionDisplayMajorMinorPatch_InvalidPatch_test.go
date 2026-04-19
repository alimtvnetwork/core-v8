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

	"github.com/alimtvnetwork/core/coreversion"
	"github.com/alimtvnetwork/core/coretests/args"
)

// Test_Cov5_VersionDisplayMajorMinorPatch_InvalidPatch tests the IsPatchInvalid branch.
func Test_VersionDisplayMajorMinorPatch_InvalidPatch(t *testing.T) {
	// Arrange
	v := coreversion.Version{
		VersionMajor: 1,
		VersionMinor: 2,
		VersionPatch: -1,
	}

	// Act
	result := v.VersionDisplayMajorMinorPatch()

	// Assert
	expectedStr := "v1.2"
	actualCheck := args.Map{"result": result != expectedStr}
	expectedCheck := args.Map{"result": false}
	expectedCheck.ShouldBeEqual(t, 0, "VersionDisplayMajorMinorPatch with invalid patch", actualCheck)
}

// Test_Cov5_Major_Compare tests the Major() comparison method.
func Test_Major_Compare(t *testing.T) {
	// Arrange
	v := coreversion.Version{
		VersionMajor: 3,
		VersionMinor: 0,
	}

	// Act
	result := v.Major(3)

	// Assert
	actual := args.Map{"result": result.IsEqual()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Major(3) on version with major=3 should be Equal", actual)
}
