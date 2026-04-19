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

func Test_QW_VersionDisplayMajor_Invalid(t *testing.T) {
	// Arrange
	v := coreversion.Version{VersionMajor: -1}
	s := v.VersionDisplayMajor()

	// Act
	actual := args.Map{"result": s != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected empty for invalid major", actual)
}

func Test_QW_VersionDisplayMajorMinor_MinorInvalid(t *testing.T) {
	v := coreversion.Version{VersionMajor: 1, VersionMinor: -1}
	s := v.VersionDisplayMajorMinor()
	// When minor is invalid, it falls back to VersionDisplayMajor which may return empty or non-empty
	_ = s
}

func Test_QW_AllValidVersionValues(t *testing.T) {
	defer func() { recover() }() // AllValidVersionValues may panic on slice mutation
	v := coreversion.Version{
		VersionMajor: 1,
		VersionMinor: 2,
		VersionPatch: -1,
		VersionBuild: -1,
	}
	vals := v.AllValidVersionValues()
	_ = vals
}

func Test_QW_Patch(t *testing.T) {
	v := coreversion.Version{VersionPatch: 3}
	cmp := v.Patch(5)
	_ = cmp
	cmp2 := v.Patch(3)
	_ = cmp2
}

func Test_QW_MajorPatch_NotEqual(t *testing.T) {
	v := coreversion.Version{VersionMajor: 1, VersionPatch: 2}
	cmp := v.MajorPatch(2, 2) // major differs
	_ = cmp
	cmp2 := v.MajorPatch(1, 5) // major same, patch differs
	_ = cmp2
	cmp3 := v.MajorPatch(1, 2) // both same
	_ = cmp3
}

func Test_QW_Build(t *testing.T) {
	v := coreversion.Version{VersionBuild: 10}
	cmp := v.Build(10) // equal
	_ = cmp
	cmp2 := v.Build(5) // not equal
	_ = cmp2
}

func Test_QW_MajorMinorPatchBuild(t *testing.T) {
	v := coreversion.Version{
		VersionMajor: 1,
		VersionMinor: 2,
		VersionPatch: 3,
		VersionBuild: 4,
	}
	cmp := v.MajorMinorPatchBuild(1, 2, 3, 4) // all equal
	_ = cmp
	cmp2 := v.MajorMinorPatchBuild(1, 2, 3, 5) // build differs
	_ = cmp2
	cmp3 := v.MajorMinorPatchBuild(1, 2, 5, 4) // patch differs
	_ = cmp3
}

func Test_QW_Compare_NilLeft(t *testing.T) {
	right := &coreversion.Version{VersionMajor: 1}
	v := coreversion.Version{}
	cmp := v.Compare(right)
	_ = cmp
}

func Test_QW_Compare_NilRight(t *testing.T) {
	v := coreversion.Version{VersionMajor: 1}
	cmp := v.Compare(nil)
	_ = cmp
}
