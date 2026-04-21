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

package keymk

import (
	"encoding/json"

	"github.com/alimtvnetwork/core-v8/coredata/corejson"
)

func (it *Key) TemplateReplacer() templateReplacer {
	return templateReplacer{
		it,
	}
}

func (it *Key) JsonModel() keyModel {
	return keyModel{
		Option:        it.option,
		MainName:      it.mainName,
		KeyChains:     it.keyChains,
		CompiledChain: it.compiledChain,
	}
}

func (it *Key) JsonModelAny() any {
	return it.JsonModel()
}

func (it Key) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it *Key) MarshalJSON() ([]byte, error) {
	return json.Marshal(it.JsonModel())
}

func (it *Key) UnmarshalJSON(data []byte) error {
	var deserializedModel keyModel
	err := json.Unmarshal(data, &deserializedModel)

	if err == nil {
		it.option = deserializedModel.Option
		it.mainName = deserializedModel.MainName
		it.keyChains = deserializedModel.KeyChains
		it.compiledChain = deserializedModel.CompiledChain
	}

	return err
}

func (it Key) Json() corejson.Result {
	return corejson.New(it)
}

func (it Key) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it Key) JsonString() string {
	return corejson.NewPtr(it).JsonString()
}

// ParseInjectUsingJson It will not update the self but creates a new one.
func (it *Key) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Key, error) {
	err := jsonResult.Unmarshal(it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

// ParseInjectUsingJsonMust Panic if error
func (it *Key) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Key {
	deserialized, err := it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return deserialized
}

func (it *Key) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Key) AsJsoner() corejson.Jsoner {
	return it
}

func (it *Key) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it Key) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return &it
}

func (it Key) AsJsonMarshaller() corejson.JsonMarshaller {
	return &it
}
