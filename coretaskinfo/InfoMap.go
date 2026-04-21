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

package coretaskinfo

import (
	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
)

func (it *Info) Map() map[string]string {
	if it.IsNull() || it.IsExcludeAdditionalErrorWrap() {
		return map[string]string{}
	}

	newMap := make(
		map[string]string,
		constants.Capacity8)

	if it.IsIncludeRootName() {
		newMap[infoFieldName] = it.RootName
	}

	if it.IsIncludeDescription() {
		newMap[infoFieldDescription] = it.Description
	}

	if it.IsIncludeUrl() {
		newMap[infoFieldUrl] = it.Url
	}

	if it.IsIncludeHintUrl() {
		newMap[infoFieldHintUrl] = it.HintUrl
	}

	if it.IsIncludeErrorUrl() {
		newMap[infoFieldErrorUrl] = it.ErrorUrl
	}

	if it.IsIncludeExampleUrl() {
		newMap[infoFieldExampleUrl] = it.ExampleUrl
	}

	if it.IsIncludeSingleExample() {
		newMap[infoFieldSingleExample] = it.SingleExample
	}

	if it.IsIncludeExamples() {
		newMap[infoFieldExamples] = it.ExamplesAsString()
	}

	return newMap
}

func (it *Info) MapWithPayload(
	payloads []byte,
) map[string]string {
	compiledMap := it.Map()

	if it.IsIncludePayloads() {
		compiledMap[payloadsField] = corejson.BytesToString(payloads)
	}

	return compiledMap
}

func (it *Info) LazyMapWithPayload(
	payloads []byte,
) map[string]string {
	compiledMap := it.LazyMap()

	if it.IsIncludePayloads() {
		compiledMap[payloadsField] = corejson.BytesToString(payloads)
	}

	return compiledMap
}

func (it *Info) MapWithPayloadAsAny(
	payloadsAny any,
) map[string]string {
	compiledMap := it.Map()

	if it.IsExcludePayload() {
		return compiledMap
	}

	jsonResult := corejson.
		AnyTo.
		SerializedJsonResult(payloadsAny)

	if jsonResult.HasError() {
		compiledMap[payloadsErrField] = jsonResult.MeaningfulErrorMessage()
	}

	compiledMap[payloadsField] = jsonResult.JsonString()

	return compiledMap
}

func (it *Info) LazyMapWithPayloadAsAny(
	payloadsAny any,
) map[string]string {
	compiledMap := it.LazyMap()

	if it.IsExcludePayload() {
		return compiledMap
	}

	jsonResult := corejson.
		AnyTo.
		SerializedJsonResult(payloadsAny)

	if jsonResult.HasError() {
		compiledMap[payloadsErrField] = jsonResult.MeaningfulErrorMessage()
	}

	compiledMap[payloadsField] = jsonResult.JsonString()

	return compiledMap
}

func (it *Info) LazyMap() map[string]string {
	if it.IsNull() || it.IsExcludeAdditionalErrorWrap() {
		return map[string]string{}
	}

	if it.lazyMap != nil {
		return it.lazyMap
	}

	it.lazyMap = it.Map()

	return it.lazyMap
}
