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

package coreversion

import (
	"github.com/alimtvnetwork/core/corecmp"
	"github.com/alimtvnetwork/core/corecomparator"
)

func Compare(
	left,
	right *Version,
) corecomparator.Compare {
	compare, isApplicable :=
		hasDeductUsingNilNess(left, right)

	if isApplicable {
		return compare
	}

	majorVersionCompare := corecmp.Integer(
		left.VersionMajor,
		right.VersionMajor,
	)

	if majorVersionCompare.IsNotEqualLogically() {
		return majorVersionCompare
	}

	// proceed only on equal
	minorVersionCompare := corecmp.Integer(
		left.VersionMinor,
		right.VersionMinor,
	)

	if minorVersionCompare.IsNotEqualLogically() {
		return minorVersionCompare
	}

	patchVersionCompare := corecmp.Integer(
		left.VersionPatch,
		right.VersionPatch,
	)

	if patchVersionCompare.IsNotEqualLogically() {
		return patchVersionCompare
	}

	buildVersionCompare := corecmp.Integer(
		left.VersionBuild,
		right.VersionBuild,
	)

	if buildVersionCompare.IsNotEqualLogically() {
		return buildVersionCompare
	}

	return corecomparator.Equal
}

// CompareVersionString
//
// See New.Prefix for more details
func CompareVersionString(
	leftVersion,
	rightVersion string,
) corecomparator.Compare {
	left := New.DefaultPtr(leftVersion)
	right := New.DefaultPtr(rightVersion)

	return Compare(left, right)
}

// IsExpectedVersion
//
// See New.Prefix for more details
func IsExpectedVersion(
	expectedCompare corecomparator.Compare,
	leftVersion,
	rightVersion string,
) bool {
	cmp := CompareVersionString(
		leftVersion, rightVersion,
	)

	return cmp.IsCompareEqualLogically(expectedCompare)
}

// IsAtLeast
//
//	returns true if left version is equal or greater than the right
func IsAtLeast(
	leftGreaterOrEqual,
	rightVersion string,
) bool {
	cmp := CompareVersionString(
		leftGreaterOrEqual, rightVersion,
	)

	return cmp.IsLeftGreaterEqualLogically()
}

// IsLower
//
//	returns true if left version is less than the right version
func IsLower(
	leftGreaterOrEqual,
	rightVersion string,
) bool {
	cmp := CompareVersionString(
		leftGreaterOrEqual, rightVersion,
	)

	return cmp.IsLeftLess()
}

// IsLowerOrEqual
//
//	returns true if left version is less or equal than the right version
func IsLowerOrEqual(
	leftGreaterOrEqual,
	rightVersion string,
) bool {
	cmp := CompareVersionString(
		leftGreaterOrEqual, rightVersion,
	)

	return cmp.IsLeftLessEqualLogically()
}
