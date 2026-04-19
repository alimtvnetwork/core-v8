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

func Test_DoubleQuoteWrapElementsWithIndexes_Verification(t *testing.T) {
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
		"\"some-elem[0]\"",
		"\"alim-elem[1]\"",
		"\"\"has-quote\"[2]\"",
		"\"[3]\"",
		"\"\"[4]\"",
		"\"\"first[5]\"",
		"\"last\"[6]\"",
		"\"'[7]\"",
		"\"simple[8]\"",
	}

	// Act
	actual := simplewrap.
		DoubleQuoteWrapElementsWithIndexes(
			testCases...,
		)

	// Assert
	convey.Convey(
		"Wrap strings with double quote with indexes - "+
			"doesn't verify existing double quote, "+
			"and possible duplicate double quote possible", t, func() {
			convey.So(actual, should.Equal, expectation)
		},
	)
}
