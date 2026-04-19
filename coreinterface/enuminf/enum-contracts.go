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

package enuminf

type BasicContractsEnumer interface {
	BasicEnumer
	TypeNameWithRangeNamesCsvGetter
}

type BasicByteContractsEnumer interface {
	BasicContractsEnumer
	IsValueByteEqualer
	IsAnyValueByteEqualer
	BasicByteEnumer
}

type BasicByteEnumContractsBinder interface {
	BasicByteContractsEnumer
	AsBasicByteEnumContractsBinder() BasicByteEnumContractsBinder
}

type BasicByteEnumContractsDelegateBinder interface {
	AsBasicByteEnumContractsDelegateBinder() BasicByteEnumContractsDelegateBinder
}

type BasicInt8ContractsEnumer interface {
	BasicContractsEnumer
	BasicInt8Enumer
	Int8ToEnumStringer
	IsValueInteger8Equaler
	IsAnyValueInteger8Equaler
}

type BasicInt8EnumContractsBinder interface {
	BasicInt8ContractsEnumer
	AsBasicInt8EnumContractsBinder() BasicInt8EnumContractsBinder
}

type BasicInt16ContractsEnumer interface {
	BasicContractsEnumer
	BasicInt16Enumer
	Int16ToEnumStringer
	IsValueInteger16Equaler
	IsAnyValueInteger16Equaler
}

type BasicInt16EnumContractsBinder interface {
	BasicInt16ContractsEnumer
	AsBasicIn16EnumContractsBinder() BasicInt16ContractsEnumer
}

type BasicInt32ContractsEnumer interface {
	BasicContractsEnumer
	BasicInt32Enumer
	Int32ToEnumStringer
	IsValueInteger32Equaler
	IsAnyValueInteger32Equaler
}

type BasicInt32EnumContractsBinder interface {
	BasicInt32ContractsEnumer
	AsBasicInt32EnumContractsBinder() BasicInt32ContractsEnumer
}

type BasicEnumContractsBinder interface {
	BasicContractsEnumer
	AsBasicEnumContractsBinder() BasicEnumContractsBinder
}

type StandardEnumerContractsBinder interface {
	StandardEnumer
	AsStandardEnumerContractsBinder() StandardEnumerContractsBinder
}
