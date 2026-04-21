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

package coreoncetests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/coreonce"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

func Test_ByteOnce_Core(t *testing.T) {
	for caseIndex, tc := range byteOnceCoreTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewByteOncePtr(func() byte { return initVal })

		// Act
		actual := args.Map{
			"value":      int(once.Value()),
			"int":        once.Int(),
			"string":     once.String(),
			"isEmpty":    once.IsEmpty(),
			"isZero":     once.IsZero(),
			"isNegative": once.IsNegative(),
			"isPositive": once.IsPositive(),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ByteOnce_Caching(t *testing.T) {
	for caseIndex, tc := range byteOnceCachingTestCases {
		// Arrange
		callCount := 0
		initVal := tc.InitValue
		once := coreonce.NewByteOncePtr(func() byte {
			callCount++

			return initVal
		})

		// Act
		r1 := once.Value()
		r2 := once.Value()

		// Assert
		actual := args.Map{
			"r1":        int(r1),
			"r2":        int(r2),
			"callCount": callCount,
		}
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ByteOnce_Json(t *testing.T) {
	for caseIndex, tc := range byteOnceJsonTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewByteOncePtr(func() byte { return initVal })

		// Act
		data, err := once.MarshalJSON()

		// Assert
		actual := args.Map{
			"noError":        err == nil,
			"marshaledValue": string(data),
		}
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ByteOnce_Serialize(t *testing.T) {
	for caseIndex, tc := range byteOnceSerializeTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewByteOncePtr(func() byte { return initVal })

		// Act
		data, err := once.Serialize()

		// Assert
		actual := args.Map{
			"noError":         err == nil,
			"serializedValue": string(data),
		}
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ByteOnce_Constructor(t *testing.T) {
	for caseIndex, tc := range byteOnceConstructorTestCases {
		// Arrange
		initVal := tc.InitValue
		once := coreonce.NewByteOnce(func() byte { return initVal })

		// Act
		actual := args.Map{
			"constructedValue": int(once.Value()),
		}

		// Assert
		tc.Case.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
