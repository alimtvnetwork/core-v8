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

package coretests

import "github.com/alimtvnetwork/core-v8/internal/msgcreator"

// IsEqualMessage
//
// Gives generic and consistent
// test message using msgformats.IsEqualMessageFormat
func (it getAssert) IsEqualMessage(
	when,
	actual,
	expected any,
) string {
	return msgcreator.Assert.IsEqualMessage(
		when,
		actual,
		expected,
	)
}

// IsNotEqualMessage
//
// Gives generic and consistent
// test message using msgformats.IsNotEqualMessageFormat
func (it getAssert) IsNotEqualMessage(
	when,
	actual,
	expected any,
) string {
	return msgcreator.Assert.IsNotEqualMessage(
		when,
		actual,
		expected,
	)
}

// IsTrueMessage
//
// Gives generic and consistent
// test message using msgformats.IsTrueMessageFormat
func (it getAssert) IsTrueMessage(
	when,
	actual any,
) string {
	return msgcreator.Assert.IsTrueMessage(
		when,
		actual,
	)
}

// IsFalseMessage
//
// Gives generic and consistent
// test message using msgformats.IsFalseMessageFormat
func (it getAssert) IsFalseMessage(
	when,
	actual any,
) string {
	return msgcreator.Assert.IsFalseMessage(
		when,
		actual,
	)
}

// IsNilMessage
//
// Gives generic and consistent
// test message using msgformats.IsNilMessageFormat
func (it getAssert) IsNilMessage(
	when,
	actual any,
) string {
	return msgcreator.Assert.IsNilMessage(
		when,
		actual,
	)
}

// IsNotNilMessage
//
// Gives generic and consistent
// test message using msgformats.IsNotNilMessageFormat
func (it getAssert) IsNotNilMessage(
	when,
	actual any,
) string {
	return msgcreator.Assert.IsNotNilMessage(
		when,
		actual,
	)
}

// ShouldBeMessage
//
// Gives generic and consistent
// test message using msgformats.ShouldBeMessageFormat
func (it getAssert) ShouldBeMessage(
	title string,
	actual,
	expected any,
) string {
	return msgcreator.Assert.ShouldBeMessage(
		title,
		actual,
		expected,
	)
}

// ShouldNotBeMessage
//
// Gives generic and consistent
// test message using msgformats.ShouldNotBeMessageFormat
func (it getAssert) ShouldNotBeMessage(
	title string,
	actual,
	expected any,
) string {
	return msgcreator.Assert.ShouldNotBeMessage(
		title,
		actual,
		expected,
	)
}
