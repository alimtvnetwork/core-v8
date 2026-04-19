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

package enumimpl

const (
	errUnmappedMessage = "typename:[%s], value given : [\"%s\"], cannot find in the enum map. " +
		"reference values can be [%s], " +
		"brackets values can be used to unmarshal as well."
	typeNameTemplateKey                 = "type-name"
	nameKey                             = "name"
	valueKey                            = "value"
	diffBetweenMapShouldBeMessageFormat = "%s\n\nDifference Between Map:\n\n{%s}"                // title, diff string
	actualVsExpectingMessageFormat      = "%s :\n\nActual:\n%s\n\nExpecting:\n%s\n\n"            // title, actual, expecting
	curlyWrapFormat                     = "{\n\n%s\n\n}"                                         // jsonValueString
	currentValueNotFoundInJsonMapFormat = "current given value (%v) is not found in the map. %s" // anyValue, RangesInvalidMessage
	defaultStackSkipForSpecificMethod   = 4
)
