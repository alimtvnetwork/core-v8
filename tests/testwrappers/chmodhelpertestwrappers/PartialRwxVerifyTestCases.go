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

package chmodhelpertestwrappers

type PartialRwxVerify struct {
	Header                                string
	PartialRwxInput1, FullRwxVerifyInput2 string
	IsMatchesExpectation                  bool
}

var PartialRwxVerifyTestCases = []PartialRwxVerify{
	{
		Header:               "Same input returns true.",
		PartialRwxInput1:     "-rwx-*-r*x",
		FullRwxVerifyInput2:  "-rwx-*-r*x",
		IsMatchesExpectation: true,
	},
	{
		Header: "Same [-rwx---r*x] comparing " +
			"with [-rwx-*-r*x] returns false.",
		PartialRwxInput1:     "-rwx---r*x",
		FullRwxVerifyInput2:  "-rwx-*-r*x",
		IsMatchesExpectation: false,
	},
	{
		Header: "Same [-rwx-*-r*x] comparing with " +
			"[-rwx-w-r*x] returns true.",
		PartialRwxInput1:     "-rwx-*-r*x",
		FullRwxVerifyInput2:  "-rwx-w-r*x",
		IsMatchesExpectation: true,
	},
	{
		Header: "Same [-rwx-*-] or [-rwx-*-***] (not givens ones are wildcard) " +
			"comparing with [-rwx-w--*x] returns true.",
		PartialRwxInput1:     "-rwx-*-",
		FullRwxVerifyInput2:  "-rwx-w--*x",
		IsMatchesExpectation: true,
	},
	{
		Header: "Same [-rwxr*-] or [-rwxr*-***] (not givens ones are wildcard) " +
			"comparing with [-rwx-w--*x] returns false.",
		PartialRwxInput1:     "-rwxr*-",
		FullRwxVerifyInput2:  "-rwx-w--*x",
		IsMatchesExpectation: false,
	},
}
