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

package payloadinf

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coreinterface"
	"github.com/alimtvnetwork/core/coreinterface/entityinf"
	"github.com/alimtvnetwork/core/coreinterface/errcoreinf"
)

type AttributesBinder interface {
	Length() int
	IsEmpty() bool
	HasItems() bool
	HasAnyItem() bool
	HasSafeItems() bool

	Payloads() []byte
	Capacity() int
	AnyKeyValMap() map[string]any
	Hashmap() map[string]string
	CompiledError() error
	HasError() bool
	IsSafeValid() bool
	IsInvalid() bool
	IsValid() bool
	MustBeEmptyError()
	HandleErr()

	DeserializeDynamicPayloads(
		toPtr any,
	) error
	IsEmptyError() bool
	DynamicBytesLength() int
	StringKeyValuePairsLength() int
	AnyKeyValuePairsLength() int

	AuthType() string
	ResourceName() string
	HasStringKeyValuePairs() bool
	HasAnyKeyValuePairs() bool
	HasDynamicPayloads() bool
	DynamicPayloadsDeserialize(
		unmarshallingPointer any,
	) error
	DynamicPayloadsDeserializeMust(
		unmarshallingPointer any,
	)
	AddOrUpdateString(
		key, value string,
	) (isNewlyAdded bool)
	AddOrUpdateAnyItem(
		key string,
		anyItem any,
	) (isNewlyAdded bool)
	String() string
	JsonModelAny() any
	SetBasicErr(
		basicErr errcoreinf.BasicErrWrapper,
	) AttributesBinder

	HasIssuesOrEmpty() bool
	IsNull() bool

	IsErrorEqual(basicErr errcoreinf.BasicErrWrapper) bool

	HasStringKey(key string) bool
	HasAnyKey(key string) bool

	AddNewStringKeyValueOnly(key, value string) (isAdded bool)
	AddNewAnyKeyValueOnly(key string, value any) (isAdded bool)

	GetStringKeyValue(key string) (value string, isFound bool)
	GetAnyKeyValue(key string) (valueAny any, isFound bool)
	AnyKeyReflectSetTo(key string, toPtr any) error

	corejson.Jsoner
	coreinterface.ErrorHandler
	coreinterface.ReflectSetter

	Clear()
	Dispose()
}

type PayloadsBinder interface {
	HasSafeItems() bool

	DynamicPayloads() []byte
	SetDynamicPayloads(payloads []byte) error

	AnyAttributes() any
	ReflectSetAttributes(
		toPointer any,
	) error
	AttrAsBinder() AttributesBinder

	IdString() string
	IdInteger() int

	IsStandardTaskEntityEqual(
		entity entityinf.StandardTaskEntityDefiner,
	) bool
	ValueReflectSet(
		setterPtr any,
	) error
	Serialize() ([]byte, error)
	SerializeMust() []byte

	InitializeAttributesOnNull() AttributesBinder

	Username() string
	Value() any
	Error() error
	BasicError() errcoreinf.BasicErrWrapper

	IsPayloadsEqual(nextPayloads []byte) bool
	IsName(name string) bool
	IsIdentifier(id string) bool
	IsTaskTypeName(taskType string) bool
	IsEntityType(entityType string) bool
	IsCategory(category string) bool

	String() string
	PrettyJsonString() string
	JsonString() string
	JsonStringMust() string

	HasAnyItem() bool
	HasIssuesOrEmpty() bool
	HasError() bool
	IsEmptyError() bool
	HasAttributes() bool
	IsEmptyAttributes() bool
	HasSingleRecord() bool

	IsNull() bool
	HasAnyNil() bool

	coreinterface.LengthGetter
	coreinterface.CountGetter

	IsEmpty() bool
	HasItems() bool

	IdentifierInteger() int
	IdentifierUnsignedInteger() uint
	Deserialize(
		unmarshallingPointer any,
	) error
	DeserializeMust(
		unmarshallingPointer any,
	)
	PayloadDeserialize(
		unmarshallingPointer any,
	) error
	PayloadDeserializeMust(
		unmarshallingPointer any,
	)

	PayloadDeserializeToPayloadBinder() (PayloadsBinder, error)
	JsonModelAny() any

	All() (id, name, entity, category string, dynamicPayloads []byte)
	AllSafe() (id, name, entity, category string, dynamicPayloads []byte)
	PayloadName() string
	PayloadCategory() string
	PayloadTaskType() string
	PayloadEntityType() string
	PayloadDynamic() []byte
	PayloadProperties() PayloadPropertiesDefiner

	corejson.Jsoner
	coreinterface.ErrorHandler
	coreinterface.ReflectSetter

	Clear()
	Dispose()
}

type PayloadPropertiesDefiner interface {
	All() (id, name, entity, category string, dynamicPayloads []byte)
	AllSafe() (id, name, entity, category string, dynamicPayloads []byte)

	BasicError() errcoreinf.BasicErrWrapper
	SetBasicError(basicError errcoreinf.BasicErrWrapper)

	Name() string
	SetName(name string) error
	SetNameMust(name string)

	IdInteger() int
	IdUnsignedInteger() uint
	IdString() string
	SetIdString(id string) error
	SetIdStringMust(id string)

	Category() string
	SetCategory(category string) error
	SetCategoryMust(category string)

	EntityType() string
	SetEntityType(entityName string) error
	SetEntityTypeMust(entityName string)

	HasManyRecord() bool
	HasSingleRecordOnly() bool

	SetSingleRecordFlag()
	SetManyRecordFlag()

	DynamicPayloads() []byte
	DynamicPayloadsDeserializedTo(toPtr any) error
	SetDynamicPayloads(dynamicPayloads []byte) error
	SetDynamicPayloadsMust(dynamicPayloads []byte)

	coreinterface.ReflectSetter

	corejson.Jsoner
}
