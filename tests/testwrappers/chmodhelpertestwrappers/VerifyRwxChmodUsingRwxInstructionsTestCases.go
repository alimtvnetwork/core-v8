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

import "github.com/alimtvnetwork/core/chmodhelper/chmodins"

var VerifyRwxChmodUsingRwxInstructionsTestCases = []VerifyRwxChmodUsingRwxInstructionsWrapper{
	{
		Header: "rwx - missing paths",
		RwxInstruction: chmodins.RwxInstruction{
			RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
				Owner: "rwx",
				Group: "rwx",
				Other: "---",
			},
			Condition: chmodins.Condition{
				IsSkipOnInvalid:   false,
				IsContinueOnError: false,
				IsRecursive:       false,
			},
		},
		Locations:            SimpleLocations,
		ExpectedErrorMessage: "Path missing or having other access issues! Ref(s) { \"[/tmp/core/test-cases-3s /tmp/core/test-cases-3x]\" }",
	},
	{
		Header: "rwx - expectation failed",
		RwxInstruction: chmodins.RwxInstruction{
			RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
				Owner: "rwx",
				Group: "r-x",
				Other: "---",
			},
			Condition: chmodins.Condition{
				IsSkipOnInvalid:   true,
				IsContinueOnError: true,
				IsRecursive:       false,
			},
		},
		Locations: SimpleLocations,
		ExpectedErrorMessage: "Path:/tmp/core/test-cases-2 - " +
			"Expect [\"rwxr-x---\"] != [\"rwxr-xr--\"] Actual\n" +
			"Path:/tmp/core/test-cases-3 - " +
			"Expect [\"rwxr-x---\"] != [\"rwxr-xr--\"] Actual",
	},
	{
		Header: "Recursive not supported",
		RwxInstruction: chmodins.RwxInstruction{
			RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
				Owner: "rwx",
				Group: "r-x",
				Other: "---",
			},
			Condition: chmodins.Condition{
				IsSkipOnInvalid:   true,
				IsContinueOnError: true,
				IsRecursive:       true,
			},
		},
		Locations: SimpleLocations,
		ExpectedErrorMessage: "Not Supported: Feature or method is not supported yet. " +
			"IsRecursive is not supported for Verify chmod. Ref(s) { " +
			"\"[" +
			"/tmp/core/test-cases-2 " +
			"/tmp/core/test-cases-3s " +
			"/tmp/core/test-cases-3x " +
			"/tmp/core/test-cases-3]\" }",
	},
	{
		Header: "Missing paths + Expectation failed",
		RwxInstruction: chmodins.RwxInstruction{
			RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
				Owner: "rwx",
				Group: "r-x",
				Other: "---",
			},
			Condition: chmodins.Condition{
				IsSkipOnInvalid:   false,
				IsContinueOnError: true,
				IsRecursive:       false,
			},
		},
		Locations: SimpleLocations,
		ExpectedErrorMessage: "Path missing or having other access issues! Ref(s) { " +
			"\"[/tmp/core/test-cases-3s /tmp/core/test-cases-3x]\" " +
			"}\n" +
			"Path:/tmp/core/test-cases-2 - " +
			"Expect [\"rwxr-x---\"] != [\"rwxr-xr--\"] Actual\n" +
			"Path:/tmp/core/test-cases-3 - " +
			"Expect [\"rwxr-x---\"] != [\"rwxr-xr--\"] Actual",
	},
	{
		Header: "Expectation and missing paths, isContinue false so will fail for missing paths only",
		RwxInstruction: chmodins.RwxInstruction{
			RwxOwnerGroupOther: chmodins.RwxOwnerGroupOther{
				Owner: "rwx",
				Group: "r-x",
				Other: "---",
			},
			Condition: chmodins.Condition{
				IsSkipOnInvalid:   false,
				IsContinueOnError: false,
				IsRecursive:       false,
			},
		},
		Locations:            SimpleLocations,
		ExpectedErrorMessage: "\"[/tmp/core/test-cases-3s /tmp/core/test-cases-3x]\" Path Ref(s) access having issues! missing or other { }",
	},
}
