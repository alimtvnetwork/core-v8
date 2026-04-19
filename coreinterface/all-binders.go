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

package coreinterface

type BasicSlicerContractsBinder interface {
	BasicSlicer
	AsBasicSliceContractsBinder() BasicSlicerContractsBinder
}

type StandardSlicerContractsBinder interface {
	StandardSlicer
	AsStandardSlicerContractsBinder() StandardSlicerContractsBinder
}

type DynamicLinqWithPagingContractsBinder interface {
	DynamicLinqWithPaging
	AsDynamicLinqWithPagingContractsBinder() DynamicLinqWithPagingContractsBinder
}

type DynamicLinqContractsBinder interface {
	DynamicLinq
	AsDynamicLinqContractsBinder() DynamicLinqContractsBinder
}

type SimpleValidInvalidCheckerContractsBinder interface {
	SimpleValidInvalidChecker
	AsSimpleValidInvalidChecker() SimpleValidInvalidChecker
}

type JsonBytesStringerContractsBinder interface {
	JsonByter
	JsonCombineStringer
	AsJsonBytesStringerContractsBinder() JsonBytesStringerContractsBinder
}

type CountStateTrackerBinder interface {
	CountStateTracker
	AsCountStateTrackerBinder() CountStateTrackerBinder
}

type KeyValueStringDefinerBinder interface {
	KeyValueStringDefiner
	AsKeyValueStringDefinerBinder() KeyValueStringDefinerBinder
}

type KeyAnyValueDefinerBinder interface {
	KeyAnyValueDefiner
	AsKeyAnyValueDefinerBinder() KeyAnyValueDefinerBinder
}

type KeyStringValuesCollectionDefinerBinder interface {
	KeyStringValuesCollectionDefiner
	AsKeyStringValuesCollectionDefinerBinder() KeyStringValuesCollectionDefinerBinder
}

type KeyAnyValuesCollectionDefinerBinder interface {
	KeyAnyValuesCollectionDefiner
	AsKeyAnyValuesCollectionDefinerBinder() KeyAnyValuesCollectionDefinerBinder
}
