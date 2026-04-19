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

package defaulterr

import "github.com/alimtvnetwork/core/errcore"

var (
	Marshalling = errcore.
			MarshallingFailedType.
			ErrorNoRefs("Cannot marshal object to serialize form.")

	UnMarshalling = errcore.
			UnMarshallingFailedType.
			ErrorNoRefs("Cannot unmarshal data to object form.")

	UnMarshallingPlusCannotFindingEnumMap = errcore.
						UnMarshallingFailedType.
						ErrorNoRefs(
			"Cannot find in the enum map. " +
				"Reference data given as : ")

	MarshallingFailedDueToNilOrEmpty = errcore.
						MarshallingFailedType.
						ErrorNoRefs("Cannot marshal to serialize data because of nil or empty object.")

	UnmarshallingFailedDueToNilOrEmpty = errcore.
						UnMarshallingFailedType.
						ErrorNoRefs("Cannot unmarshal to object because of nil or empty serialized data.")

	CannotProcessNilOrEmpty = errcore.
				CannotBeNilOrEmptyType.
				ErrorNoRefs("Cannot process nil or empty.")

	OutOfRange = errcore.
			OutOfRangeType.
			ErrorNoRefs("Cannot process out of range data.")

	NegativeDataCannotProcess = errcore.
					CannotBeNegativeType.
					ErrorNoRefs("Cannot process negative values.")

	NilResult = errcore.
			NullResultType.
			ErrorNoRefs("Cannot process nil result.")

	UnexpectedValue = errcore.
			UnexpectedValueType.
			ErrorNoRefs("Cannot process unexpected value or values.")

	CannotRemoveFromEmptyCollection = errcore.
					CannotRemoveIndexesFromEmptyCollectionType.
					ErrorNoRefs("Cannot process request: cannot remove from empty collection.")

	CannotConvertStringToByte = errcore.
					FailedToConvertType.
					ErrorNoRefs("Cannot convert string to byte.")

	AttributeNull = errcore.
			NullResultType.
			ErrorNoRefs("attribute is nil!")
	JsonResultNull = errcore.
			CannotBeNilOrEmptyType.
			ErrorNoRefs("JsonResult is given as nil")

	KeyNotExistInMap = errcore.
				KeyNotExistInMapType.
				ErrorNoRefs("key doesn't exist in the map.")
)
