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

package constants

const (
	SprintValueFormat                            = "%v"
	SprintValueDoubleQuotationFormat             = "\"%s\""
	SprintNumberFormat                           = "%d"
	SprintFullPropertyNameValueFormat            = "%#v"
	SprintPropertyNameValueFormat                = "%+v"
	SprintPropertyValueWithTypeFormat            = "%+v (%T)"
	SprintTypeFormat                             = "%T"
	SprintTypeInParenthesisFormat                = "(type : %T)"
	SprintNilValueTypeInParenthesisFormat        = "<nil> (type : %T)"
	SprintValueWithTypeFormat                    = "%v " + SprintTypeInParenthesisFormat
	SprintDoubleQuoteFormat                      = "%q"
	SprintStartStringEndCharFormat               = "%c%s%c"
	SprintSingleQuoteFormat                      = "'%s'"
	SprintStringFormat                           = "%s"
	SprintThirdBracketQuoteFormat                = "[\"%v\"]"
	KeyValuePariSimpleFormat                     = "{ Key (Type - %T): %v} - { Value (Type - %T) : %v  }"
	SprintFormatNumberWithColon                  = "%d:%d"
	SprintFormatAnyValueWithColon                = "%v:%v"
	TitleValueFormat                             = "%s : %v"
	CurlyTitleWrapFormat                         = "%s: {%s}"        // Title, Value
	QuotationTitleWrapFormat                     = "%v: \"%v\""      // Title, Value
	QuotationTitleMetaWrapFormat                 = "%s: \"%s\" (%v)" // Title, Value, Meta
	CurlyTitleMetaWrapFormat                     = "%s: {%s} (%v)"   // Title, Value, Meta
	SquareTitleWrapFormat                        = "%s: [%s]"        // Title, Value
	SquareTitleMetaWrapFormat                    = "%s: [%s] (%v)"   // Title, Value, Meta
	SprintFormatAnyValueWithComma                = "%v,%v"
	SprintFormatWithNewLine                      = "%v\n%v"
	SprintFormatAnyValueWithPipe                 = "%v|%v"
	SprintFormatAnyNameValueWithColon            = "%#v:%#v"
	SprintFormatAnyNameValueWithPipe             = "%#v|%#v"
	SprintFormatNumberWithHyphen                 = "%d-%d"
	SprintFormatNumberWithPipe                   = "%d|%d"
	ThreeValueNewLineJoin                        = "%v\n%v\n%v"
	ThreeValueNewLineSpaceJoin                   = " %v\n %v\n %v"
	BracketWrapFormat                            = "[%s]"
	BracketQuotationWrapFormat                   = "[\"%s\"]"
	CurlyWrapFormat                              = "{%s}"
	SquareWrapFormat                             = "[%s]"
	ParenthesisWrapFormat                        = "(%s)"
	CurlyQuotationWrapFormat                     = "{\"%s\"}"
	ParenthesisQuotationWrap                     = "(\"%s\")"
	ReferenceWrapFormat                          = "Ref (s) { %v }"
	MessageReferenceWrapFormat                   = "%s Ref (s) { %v }"
	StringWithBracketWrapNumberFormat            = "%s[%d]"
	DoubleQuoteStringWithBracketWrapNumberFormat = "\"%s\"[%d]"
	SpaceHyphenAngelBracketSpaceRefWrapFormat    = " -> Ref(%v)"
	ValueWithDoubleQuoteFormat                   = "\"%v\""
	ValueWithSingleQuoteFormat                   = "'%v'"
	StringWithDoubleQuoteFormat                  = "\"%s\""
	StringWithSingleQuoteFormat                  = "'%s'"
	StringTypeFormat                             = "'%s - %T'"
	TypeWithSingleQuoteFormat                    = "'%T'"
	TypeWithDoubleQuoteFormat                    = "\"%T\""
	MessageWrapMessageFormat                     = "%s (%s)"
	FromToFormat                                 = "{From : %q, To: %q}"            // From, To name
	SourceDestinationFormat                      = "{Source : %q, Destination: %q}" // source, destination
	RenameFormat                                 = "{Existing : %q, New: %q}"       // existing, new
	ValueWrapValueFormat                         = "%v (%v)"
	StringWrapValueFormat                        = "%s (%s)"
	FilePathEmpty                                = "File path was empty(\"\")."
	EnumOnlySupportedFormat                      = "enum: %T, " +
		"not supported (\"%s\") | only supported { %s }" // enumSelf, enumSelf, csv-support
	EnumOnlySupportedWithMessageFormat = "enum: %T, " +
		"not supported (\"%s\") | %s | only supported { %s }" // enumSelf, enumSelf, message, csv-support
)
