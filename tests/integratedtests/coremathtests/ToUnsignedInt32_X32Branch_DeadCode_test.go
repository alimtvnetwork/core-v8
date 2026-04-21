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

package coremathtests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/coremath"
	"github.com/alimtvnetwork/core-v8/osconsts"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage4 — integerOutOfRange.ToUnsignedInt32 x32 branch
//
// Line 21-23: osconsts.IsX32Architecture == true path
// On 64-bit platforms this branch is normally unreachable.
// We temporarily flip the var to exercise it.
// ══════════════════════════════════════════════════════════════════════════════

func Test_ToUnsignedInt32_X32Branch_InRange(t *testing.T) {
	// Arrange
	original := osconsts.IsX32Architecture
	osconsts.IsX32Architecture = true
	defer func() { osconsts.IsX32Architecture = original }()

	// Act
	result := coremath.IsOutOfRange.Integer.ToUnsignedInt32(100)

	// Assert
	convey.Convey("ToUnsignedInt32 returns false for in-range value on x32", t, func() {
		convey.So(result, convey.ShouldBeFalse)
	})
}

func Test_ToUnsignedInt32_X32Branch_Negative(t *testing.T) {
	// Arrange
	original := osconsts.IsX32Architecture
	osconsts.IsX32Architecture = true
	defer func() { osconsts.IsX32Architecture = original }()

	// Act
	result := coremath.IsOutOfRange.Integer.ToUnsignedInt32(-1)

	// Assert
	convey.Convey("ToUnsignedInt32 returns true for negative value on x32", t, func() {
		convey.So(result, convey.ShouldBeTrue)
	})
}
