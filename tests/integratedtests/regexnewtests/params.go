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

package regexnewtests

var params = struct {
	// Input keys
	pattern      string
	compareInput string
	content      string
	comparing    string
	input        string
	isLock       string
	customizer   string

	// Expected keys — state
	patternResult string
	stringResult  string
	isDefined     string
	isUndefined   string
	isApplicable  string
	isCompiled    string
	regexNotNil   string
	regexIsNil    string

	// Expected keys — matching
	isMatch            string
	isFailedMatch      string
	isMatchBytes       string
	isFailedMatchBytes string
	isFailed           string
	firstMatch         string
	isInvalidMatch     string

	// Expected keys — error / validation
	hasError      string
	isInvalid     string
	isNoError     string
	isCustomError string
	errorContains string

	// Expected keys — misc
	isNotEmpty  string
	panicked    string
	samePointer string
	mapLength   string
}{
	// Input keys
	pattern:      "pattern",
	compareInput: "compareInput",
	content:      "content",
	comparing:    "comparing",
	input:        "input",
	isLock:       "isLock",
	customizer:   "customizer",

	// Expected keys — state
	patternResult: "patternResult",
	stringResult:  "stringResult",
	isDefined:     "isDefined",
	isUndefined:   "isUndefined",
	isApplicable:  "isApplicable",
	isCompiled:    "isCompiled",
	regexNotNil:   "regexNotNil",
	regexIsNil:    "regexIsNil",

	// Expected keys — matching
	isMatch:            "isMatch",
	isFailedMatch:      "isFailedMatch",
	isMatchBytes:       "isMatchBytes",
	isFailedMatchBytes: "isFailedMatchBytes",
	isFailed:           "isFailed",
	firstMatch:         "firstMatch",
	isInvalidMatch:     "isInvalidMatch",

	// Expected keys — error / validation
	hasError:      "hasError",
	isInvalid:     "isInvalid",
	isNoError:     "isNoError",
	isCustomError: "isCustomError",
	errorContains: "errorContains",

	// Expected keys — misc
	isNotEmpty:  "isNotEmpty",
	panicked:    "panicked",
	samePointer: "samePointer",
	mapLength:   "mapLength",
}
