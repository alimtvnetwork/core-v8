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

package corestr

import (
	"fmt"
	"strings"

	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/internal/reflectinternal"
)

type KeyAnyValuePair struct {
	Key         string
	valueString SimpleStringOnce
	Value       any
}

func (it *KeyAnyValuePair) KeyName() string {
	return it.Key
}

func (it *KeyAnyValuePair) VariableName() string {
	return it.Key
}

func (it *KeyAnyValuePair) ValueAny() any {
	return it.Value
}

func (it *KeyAnyValuePair) IsVariableNameEqual(name string) bool {
	return it.Key == name
}

func (it KeyAnyValuePair) SerializeMust() (jsonBytes []byte) {
	return corejson.NewPtr(it).RawMust()
}

func (it *KeyAnyValuePair) Compile() string {
	return it.String()
}

func (it *KeyAnyValuePair) IsValueNull() bool {
	return it == nil || reflectinternal.Is.Null(it.Value)
}

func (it *KeyAnyValuePair) HasNonNull() bool {
	return it != nil && reflectinternal.Is.Defined(it.Value)
}

func (it *KeyAnyValuePair) HasValue() bool {
	return it != nil && reflectinternal.Is.Defined(it.Value)
}

func (it *KeyAnyValuePair) IsValueEmptyString() bool {
	return it == nil || it.ValueString() == ""
}

func (it *KeyAnyValuePair) IsValueWhitespace() bool {
	return it == nil || strings.TrimSpace(it.ValueString()) == ""
}

func (it *KeyAnyValuePair) ValueString() string {
	if it.valueString.IsInitialized() {
		return it.valueString.String()
	}

	if it.HasNonNull() {
		valueString := fmt.Sprintf(constants.SprintValueFormat, it.Value)

		return it.
			valueString.
			GetSetOnce(valueString)
	}

	return it.
		valueString.
		GetOnce()
}

func (it *KeyAnyValuePair) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it *KeyAnyValuePair) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*KeyAnyValuePair, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *KeyAnyValuePair) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *KeyAnyValuePair {
	parsed, err := it.
		ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return parsed
}

func (it KeyAnyValuePair) Json() corejson.Result {
	return corejson.New(it)
}

func (it *KeyAnyValuePair) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *KeyAnyValuePair) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *KeyAnyValuePair) AsJsoner() corejson.Jsoner {
	return it
}

func (it *KeyAnyValuePair) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *KeyAnyValuePair) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it KeyAnyValuePair) String() string {
	return fmt.Sprintf(
		keyValuePrintFormat,
		it.Key,
		it.Value,
	)
}

func (it *KeyAnyValuePair) Clear() {
	if it == nil {
		return
	}

	it.Key = ""
	it.Value = nil
	it.valueString.Dispose()
}

func (it *KeyAnyValuePair) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
}
