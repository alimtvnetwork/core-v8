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

import (
	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/errcore"
)

// SimpleGherkins
//
// https://www.guru99.com/gherkin-test-cucumber.html
// Feature:  Login functionality of social networking site Facebook.
// Given:  I am a facebook user.
// When: I enter username as username.
// And I enter the password as the password
// Then I should be redirected to the home page of facebook
// Given -> When -> Then
type SimpleGherkins struct {
	Feature,
	Given,
	When,
	Then,
	Expect,
	Actual string
}

func (it *SimpleGherkins) ToString(testIndex int) string {
	return errcore.GherkinsString(
		testIndex,
		it.Feature,
		it.Given,
		it.When,
		it.Then)
}

func (it *SimpleGherkins) String() string {
	return it.ToString(constants.Zero)
}

func (it *SimpleGherkins) GetWithExpectation(
	testIndex int,
) string {
	return errcore.GherkinsStringWithExpectation(
		testIndex,
		it.Feature,
		it.Given,
		it.When,
		it.Then,
		it.Actual,
		it.Expect)
}

func (it *SimpleGherkins) GetMessageConditional(
	isExpectation bool,
	testIndex int,
) string {
	if isExpectation {
		return it.GetWithExpectation(testIndex)
	}

	return it.ToString(testIndex)
}
