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

var VerifyRwxPartialChmodLocationsTestCases = []VerifyRwxPartialChmodLocationsWrapper{
	{
		Header:             "Missing Paths should NOT have error with it's location!",
		Locations:          SimpleLocations,
		IsContinueOnError:  true,
		IsSkipOnInvalid:    true,
		ExpectedPartialRwx: "-rwxrwx",
		ExpectationErrorMessage: "/tmp/core/test-cases-2 - " +
			"Expect [\"rwxrwx***\"] != [\"rwxr-xr--\"] Actual\n" +
			"/tmp/core/test-cases-3 - " +
			"Expect [\"rwxrwx***\"] != [\"rwxr-xr--\"] Actual",
	},
	{
		Header:                  "Missing Paths should NOT have error with it's location and all matches with WhatIsExpected RWX!",
		Locations:               SimpleLocations,
		IsContinueOnError:       true,
		IsSkipOnInvalid:         true,
		ExpectedPartialRwx:      "-rwx",
		ExpectationErrorMessage: "",
	},
	{
		Header:             "Missing Paths should have error with it's location!",
		Locations:          SimpleLocations,
		IsContinueOnError:  true,
		IsSkipOnInvalid:    false,
		ExpectedPartialRwx: "-rwxrwx-",
		ExpectationErrorMessage: "/tmp/core/test-cases-2 - " +
			"Expect [\"rwxrwx-**\"] != [\"rwxr-xr--\"] Actual\n" +
			"/tmp/core/test-cases-3 - " +
			"Expect [\"rwxrwx-**\"] != [\"rwxr-xr--\"] Actual\n" +
			"Path missing or having other access issues! Ref(s) { \"" +
			"[/tmp/core/test-cases-3s " +
			"/tmp/core/test-cases-3x]\" }",
	},
}
