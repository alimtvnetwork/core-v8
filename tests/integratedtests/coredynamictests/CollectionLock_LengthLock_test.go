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

package coredynamictests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ── CollectionLock.LengthLock (line 15) ──

func Test_CollectionLock_LengthLock_I29(t *testing.T) {
	// Arrange
	coll := coredynamic.NewCollection[int](5)
	coll.Add(1)
	coll.Add(2)

	// Act
	length := coll.LengthLock()

	// Assert
	actual := args.Map{"length": length}
	expected := args.Map{"length": 2}
	actual.ShouldBeEqual(t, 1, "CollectionLock LengthLock", expected)
}

// ── CollectionLock.RemoveAtLock invalid index (line 125) ──

func Test_CollectionLock_RemoveAtLock_InvalidIndex_I29(t *testing.T) {
	// Arrange
	coll := coredynamic.NewCollection[string](2)
	coll.Add("a")

	// Act
	removed := coll.RemoveAtLock(99)

	// Assert
	actual := args.Map{"removed": removed}
	expected := args.Map{"removed": false}
	actual.ShouldBeEqual(t, 1, "CollectionLock RemoveAtLock invalid index", expected)
}
