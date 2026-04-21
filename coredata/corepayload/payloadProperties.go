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

package corepayload

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/core-v8/coreinterface/payloadinf"
)

type payloadProperties struct {
	payloadWrapper *PayloadWrapper
}

func (it *payloadProperties) SetBasicError(basicError errcoreinf.BasicErrWrapper) {
	it.payloadWrapper.InitializeAttributesOnNull()
	it.payloadWrapper.Attributes.SetBasicErr(basicError)
}

func (it payloadProperties) BasicError() errcoreinf.BasicErrWrapper {
	return it.payloadWrapper.BasicError()
}

func (it *payloadProperties) ReflectSetTo(toPointer any) error {
	return coredynamic.ReflectSetFromTo(it.payloadWrapper, toPointer)
}

func (it payloadProperties) AllSafe() (id, name, entity, category string, dynamicPayloads []byte) {
	return it.payloadWrapper.AllSafe()
}

func (it payloadProperties) All() (id, name, entity, category string, dynamicPayloads []byte) {
	return it.payloadWrapper.All()
}

func (it payloadProperties) Name() string {
	return it.payloadWrapper.Name
}

func (it *payloadProperties) SetName(name string) error {
	it.payloadWrapper.Name = name

	return nil
}

func (it *payloadProperties) SetNameMust(name string) {
	it.payloadWrapper.Name = name
}

func (it payloadProperties) IdInteger() int {
	return it.payloadWrapper.IdInteger()
}

func (it payloadProperties) IdUnsignedInteger() uint {
	return it.payloadWrapper.IdentifierUnsignedInteger()
}

func (it payloadProperties) IdString() string {
	return it.payloadWrapper.Identifier
}

func (it *payloadProperties) SetIdString(id string) error {
	it.payloadWrapper.Identifier = id

	return nil
}

func (it *payloadProperties) SetIdStringMust(id string) {
	it.payloadWrapper.Identifier = id
}

func (it payloadProperties) Category() string {
	return it.payloadWrapper.CategoryName
}

func (it *payloadProperties) SetCategory(category string) error {
	it.payloadWrapper.CategoryName = category

	return nil
}

func (it *payloadProperties) SetCategoryMust(category string) {
	it.payloadWrapper.CategoryName = category
}

func (it payloadProperties) EntityType() string {
	return it.payloadWrapper.EntityType
}

func (it *payloadProperties) SetEntityType(entityName string) error {
	it.payloadWrapper.EntityType = entityName

	return nil
}

func (it *payloadProperties) SetEntityTypeMust(entityName string) {
	it.payloadWrapper.EntityType = entityName
}

func (it payloadProperties) HasManyRecord() bool {
	return it.payloadWrapper.HasManyRecords
}

func (it payloadProperties) HasSingleRecordOnly() bool {
	return !it.HasManyRecord()
}

func (it *payloadProperties) SetSingleRecordFlag() {
	it.payloadWrapper.HasManyRecords = false
}

func (it *payloadProperties) SetManyRecordFlag() {
	it.payloadWrapper.HasManyRecords = true
}

func (it payloadProperties) DynamicPayloads() []byte {
	return it.payloadWrapper.Payloads
}

func (it *payloadProperties) SetDynamicPayloads(dynamicPayloads []byte) error {
	it.payloadWrapper.Payloads = dynamicPayloads

	return nil
}

func (it *payloadProperties) DynamicPayloadsDeserializedTo(toPtr any) error {
	return it.payloadWrapper.Deserialize(toPtr)
}

func (it *payloadProperties) SetDynamicPayloadsMust(dynamicPayloads []byte) {
	it.payloadWrapper.Payloads = dynamicPayloads
}

func (it payloadProperties) Json() corejson.Result {
	return it.payloadWrapper.Json()
}

func (it payloadProperties) JsonPtr() *corejson.Result {
	return it.payloadWrapper.JsonPtr()
}

func (it payloadProperties) AsPayloadPropertiesDefiner() payloadinf.PayloadPropertiesDefiner {
	return &it
}
