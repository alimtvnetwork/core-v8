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

package simplewraptests

import (
	"testing"

	"github.com/alimtvnetwork/core/simplewrap"
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
)

func Test_When_DoubleQuoteWrapElements_SkipQuoteOnPresent_Should_Only_Have_SingleDoubleQuotation_NotDuplicates(t *testing.T) {
	// Arrange
	testCases := []string{
		"some-elem",
		"alim-elem",
		"\"has-quote\"",
		"",
		"\"",
		"\"first",
		"last\"",
		"'",
		"simple",
	}
	expectation := []string{
		"\"some-elem\"",
		"\"alim-elem\"",
		"\"has-quote\"",
		"\"\"",
		"\"\"",
		"\"first\"",
		"\"last\"",
		"\"'\"",
		"\"simple\"",
	}

	// Act
	actual := simplewrap.
		DoubleQuoteWrapElements(
			true,
			testCases...,
		)

	// Assert
	convey.Convey(
		"Wrap strings with double quote, if exists already then skip adding", t, func() {
			convey.So(actual, should.Equal, expectation)
		},
	)
}

func Test_When_DoubleQuoteWrapElements_SkipQuoteOnPresent_Disabled_Should_Have_DuplicateDoubleQuotations(t *testing.T) {
	// Arrange
	testCases := []string{
		"some-elem",
		"alim-elem",
		"\"has-quote\"",
		"",
		"\"",
		"\"first",
		"last\"",
		"'",
		"simple",
	}
	expectation := []string{
		"\"some-elem\"",
		"\"alim-elem\"",
		"\"\"has-quote\"\"",
		"\"\"",
		"\"\"\"",
		"\"\"first\"",
		"\"last\"\"",
		"\"'\"",
		"\"simple\"",
	}

	// Act
	actual := simplewrap.
		DoubleQuoteWrapElements(
			false,
			testCases...,
		)

	// Assert
	convey.Convey(
		"Wrap strings with double quote, if exists already then skip adding", t, func() {
			convey.So(actual, should.Equal, expectation)
		},
	)
}
