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

package internalinterface

type BaseErrorTyper interface {
	NameWithNameEqualer
	NameValue() string
	IsValid() bool
	IsInvalid() bool
	IsRawValue(rawValue uint16) bool
	IsNoError() bool
	IsEmptyError() bool
	HasError() bool
	Combine(
		additionalMessage,
		varName string,
		val any,
	) string
	CombineNoRefs(
		additionalMessage string,
	) string
	Error(
		additionalMessage,
		varName string,
		val any,
	) error
	ErrorReferences(
		additionalMessage string,
		references ...any,
	) error
	ErrorNoRefs(
		additionalMessage string,
	) error
	Panic(
		additionalMessage,
		varName string,
		val any,
	)
	PanicNoRefs(
		additionalMessage string,
	)
	// CodeWithTypeName
	//
	// 	errconsts.ErrorCodeHyphenTypeNameFormat  = "(#%d - %s)"
	CodeWithTypeName() string
	TypeName() string
	CodeTypeNameWithCustomMessage(
		customMessage string,
	) string
	ReferencesCsv(
		additionalMessage string,
		references ...any,
	) string
	ReferencesLines(
		additionalMessage string,
		referencesLines ...any,
	) string
	ReferencesLinesError(
		additionalMessage string,
		referencesLines ...any,
	) error
	ReferencesCsvError(
		additionalMessage string,
		references ...any,
	) error
	ShortReferencesCsv(
		references ...any,
	) string
	ShortReferencesCsvError(
		references ...any,
	) error
	RawValue() uint16
	Value() uint16
	ValueInt16() int16
	ValueInt() int
	ValueUInt() uint
}
