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

package namevaluetests

import (
	"testing"

	"github.com/alimtvnetwork/core-v8/namevalue"
	"github.com/smartystreets/goconvey/convey"
)

// ══════════════════════════════════════════════════════════════════════════════
// Coverage9 — Instance.String nil receiver + Collection.JsonString error branch
//
// Target 1: namevalue/Instance.go:20-22
//   it.IsNull() → return EmptyString (nil receiver)
//
// Target 2: namevalue/Collection.go:385-387
//   err != nil || len(jsonBytes) == 0 → return EmptyString
//   json.Marshal of a valid Collection won't fail; this is defensive dead code.
//   Documented as accepted gap.
// ══════════════════════════════════════════════════════════════════════════════

func Test_Instance_String_NilReceiver(t *testing.T) {
	// Arrange
	var inst *namevalue.Instance[string, string]

	// Act
	result := inst.String()

	// Assert
	convey.Convey("Instance.String returns empty for nil receiver", t, func() {
		convey.So(result, convey.ShouldBeEmpty)
	})
}

// Coverage note: Collection.JsonString line 385-387 (err != nil || len==0)
// is defensive dead code — json.Marshal on Collection[K,V] with basic types
// cannot fail. Documented as accepted dead-code gap.
