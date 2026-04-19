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

package errcore

import "github.com/alimtvnetwork/core/constants"

const (
	ReferenceStart                           = "Reference(s) ("
	ReferenceEnd                             = ")"
	ReferenceFormat                          = " Ref(s) { \"%v\" }"
	rangeWithRangeFormat                     = "Range must be in between %v and %v. Ranges must be one of these {%v}"
	rangeWithoutRangeFormat                  = "Range must be in between %v and %v."
	CannotConvertStringToByteForLessThanZero = "Cannot convert string to byte. String cannot be less than 0 for byte."
	CannotConvertStringToByteForMoreThan255  = "Cannot convert string to byte. String is a number " +
		"but larger than byte size. At max it could be 255."
	CannotConvertStringToByte = "Cannot convert string to byte."
	// expectingMessageFormat "%s - expecting (type:[%T]) : [\"%v\"], but received or
	// actual (type:[%T]) : [\"%v\"]"
	expectingMessageFormat = "%s - expecting (type:[%T]) : [\"%v\"], but received " +
		"or actual (type:[%T]) : [\"%v\"]"
	expectingSimpleMessageFormat                  = "%s - Expect (type:\"%T\")[\"%v\"] != [\"%v\"](type:\"%T\") Actual"
	expectingSimpleNoTypeMessageFormat            = "%s - Expect [\"%v\"] != [\"%v\"] Actual"
	expectingNotMatchingSimpleNoTypeMessageFormat = "%s - Expect [\"%v\"] Not Matching [\"%v\"] Actual"
	var2Format                                    = "(%s, %s) = (%v, %v)"
	var2WithTypeFormat                            = "(%s [t:%T], %s[t:%T]) = (%v, %v)"
	var3Format                                    = "(%s, %s, %s) = (%v, %v, %v)"
	keyValFormat                                  = constants.KeyValShortFormat
	var3WithTypeFormat                            = "(%s [t:%T], %s[t:%T], %s[t:%T]) = (%v, %v, %v)"
	messageVar2Format                             = "%s (%s, %s) = (%v, %v)"
	messageVar3Format                             = "%s (%s, %s, %s) = (%v, %v, %v)"
	messageMapFormat                              = constants.MessageReferenceWrapFormat
	messageWithTracesWithoutRefFormat             = "%s \n%s"
	refsWithoutQuotation                          = " Ref (s) { %v }"
	messageWithRefWithoutQuoteFormat              = "%s" + refsWithoutQuotation
	messageWithOtherMsgWithRefWithoutQuoteFormat  = "%s " + messageWithRefWithoutQuoteFormat
	messageWithTracesRefFormat                    = messageWithTracesWithoutRefFormat + refsWithoutQuotation
	PrefixStackTrace                              = constants.Hyphen + constants.Space
	PrefixStackTraceNewLine                       = constants.DefaultLine + PrefixStackTrace
	NewLineCodeStacksHeader                       = "\nCode Stacks :\n"
	CodeStacksHeaderNewLine                       = "Code Stacks :\n"
	ShouldBeMessageFormat                         = "\"%v\" {actual} should be \"%v\" {expecting}" // actual, expecting
	stackEnhanceFormat                            = "%s - \n%s\n %s: \n  - %s"                     // "%s - \n%s\n %s: \n  - %s"
	combineMsgWithErrorFormat                     = "%s - %s"                                      // "%s - %s"
)
