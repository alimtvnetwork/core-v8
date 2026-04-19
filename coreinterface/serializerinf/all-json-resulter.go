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

package serializerinf

type JsonResulter interface {
	BasicJsonResulter

	SafeBytesTypeNameGetter
	BytesTypeNameGetter

	SafeStringGetter
	JsonStringGetter
	JsonStringPtrGetter
	PrettyJsonBufferGetter
	PrettyJsonStringGetter
	PrettyJsonStringOrErrStringGetter

	LengthGetter
	HasErrorChecker
	ErrorStringGetter
	IsErrorEqualChecker
	SafeBytesGetter
	BytesValuesGetter
	SafeValuesGetter
	SafeValuesPtrGetter
	RawSerializeGetter
	MustRawSerializeGetter
	RawStringSerializeGetter
	MustRawStringSerializeGetter
	RawErrStringGetter
	RawPrettyStringGetter
	MeaningfulErrorMessageGetter
	MeaningfulErrorGetter
	IsEmptyErrorChecker
	HasSafeItemsChecker
	IsAnyNullChecker
	HasIssuesOrEmptyChecker
	ErrorHandler
	MustBeSafer
	ErrorHandlerWithMessager
	HasBytesChecker
	HasJsonBytesChecker
	IsEmptyJsonBytesChecker
	IsEmptyChecker
	IsEmptyJsonChecker

	Deserializer
	MustDeserializer

	MustUnmarshaler
	Unmarshaler

	SkipExistingIssuesSerializer

	SelfSerializer
	MustSelfSerializer

	CombineErrorWithRefString(references ...string) string
	CombineErrorWithRefError(references ...string) error
	Dispose()
}

type SimpleBytesResulter interface {
	LengthGetter

	BytesValuesGetter
	SafeValuesGetter
	MeaningfulErrorGetter

	RawSerializeGetter

	IsEmptyChecker
	HasAnyItemChecker

	HasErrorChecker
}

type BaseJsonResulter interface {
	SimpleBytesResulter

	BytesTypeNameGetter

	SafeStringGetter

	LengthGetter
	HasErrorChecker

	MeaningfulErrorGetter

	BytesValuesGetter
	SafeValuesGetter

	RawSerializeGetter

	IsEmptyChecker

	SelfSerializer
}

type BasicJsonResulter interface {
	BaseJsonResulter

	SafeBytesTypeNameGetter
	BytesTypeNameGetter

	SafeStringGetter
	JsonStringGetter

	LengthGetter
	HasErrorChecker
	ErrorStringGetter
	SafeBytesGetter
	BytesValuesGetter
	SafeValuesGetter

	RawSerializeGetter
	MustRawSerializeGetter

	IsEmptyErrorChecker
	HasSafeItemsChecker
	HasIssuesOrEmptyChecker
	ErrorHandler

	Deserializer
	MustDeserializer

	Unmarshaler

	SkipExistingIssuesSerializer

	SelfSerializer
	MustSelfSerializer
}
