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
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corejson"
)

func (it Info) Json() corejson.Result {
	return corejson.New(it)
}

func (it Info) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *Info) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	return jsonResult.Deserialize(it)
}

func (it *Info) JsonString() string {
	if it.IsNull() {
		return ""
	}

	return it.JsonPtr().JsonString()
}

func (it *Info) PrettyJsonStringWithPayloads(
	payloads []byte,
) string {
	return corejson.
		NewPtr(it.MapWithPayload(payloads)).
		PrettyJsonString()
}

func (it *Info) PrettyJsonString() string {
	if it.IsNull() {
		return ""
	}

	return corejson.
		NewPtr(it).
		PrettyJsonString()
}

func (it *Info) LazyMapPrettyJsonString() string {
	lazyMap := it.LazyMap()

	return corejson.
		NewPtr(lazyMap).
		PrettyJsonString()
}

func (it Info) JsonStringMust() string {
	jsonResult := it.Json()
	jsonResult.MustBeSafe()

	return jsonResult.JsonString()
}

func (it *Info) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it *Info) Deserialize(toPtr any) (parsingErr error) {
	return it.JsonPtr().Deserialize(toPtr)
}

func (it *Info) ExamplesAsString() (compiledString string) {
	if it.IsNull() {
		return ""
	}

	return strings.Join(
		it.Examples,
		constants.CommaSpace)
}

func (it *Info) String() string {
	if it == nil {
		return ""
	}

	return it.PrettyJsonString()
}

func (it Info) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}
