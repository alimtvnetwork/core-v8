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

package coretestcases

import (
	"github.com/alimtvnetwork/core/coretests/args"
)

// AnyGherkins is a GenericGherkins with all-any typed fields.
//
// Use when input and expected types are heterogeneous or unknown
// at compile time.
type AnyGherkins = GenericGherkins[any, any]

// StringGherkins is a GenericGherkins with string input and string expected.
//
// Use for text-based validation tests where both input and expected
// are plain strings.
type StringGherkins = GenericGherkins[string, string]

// StringBoolGherkins is a GenericGherkins with string input and bool expected.
//
// Use for matching/validation tests (e.g., regex, search) where input
// is a string and the expected outcome is a boolean.
type StringBoolGherkins = GenericGherkins[string, bool]

// MapGherkins is a GenericGherkins with args.Map for both input and expected.
//
// Use when test inputs and expectations are multi-field key-value pairs.
// Input holds arrange data (e.g., pattern, compareInput).
// Expected holds assertion data (e.g., isDefined, isApplicable, isMatch).
//
// This replaces opaque ExpectedLines with self-documenting semantic keys,
// making test cases readable without consulting the test runner.
//
// Example:
//
//	var testCases = []coretestcases.MapGherkins{
//	    {
//	        Title: "Lazy regex matches word pattern",
//	        When:  "given a simple word pattern",
//	        Input: args.Map{
//	            "pattern":      "hello",
//	            "compareInput": "hello world",
//	        },
//	        Expected: args.Map{
//	            "isDefined":    true,
//	            "isApplicable": true,
//	            "isMatch":      true,
//	            "isFailedMatch": false,
//	        },
//	    },
//	}
type MapGherkins = GenericGherkins[args.Map, args.Map]
