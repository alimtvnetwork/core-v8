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

package coregenerictests

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata/coregeneric"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ==========================================
// Test: AddLock — concurrent goroutines
// ==========================================

func Test_GenericCollection_AddLock_ConcurrentSafety(t *testing.T) {
	// Arrange
	tc := collectionAddLockConcurrencyTestCase
	const goroutines = 500
	col := coregeneric.EmptyCollection[int]()

	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			col.AddLock(idx)
			wg.Done()
		}(i)
	}

	wg.Wait()

	// Act
	actual := args.Map{"length": col.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// Test: AddsLock — concurrent batch appends
// ==========================================

func Test_GenericCollection_AddsLock_ConcurrentSafety(t *testing.T) {
	// Arrange
	tc := collectionAddsLockConcurrencyTestCase
	const goroutines = 200
	const batchSize = 5
	col := coregeneric.EmptyCollection[string]()

	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			batch := make([]string, batchSize)
			for b := 0; b < batchSize; b++ {
				batch[b] = fmt.Sprintf("g%d-b%d", idx, b)
			}
			col.AddsLock(batch...)
			wg.Done()
		}(i)
	}

	wg.Wait()

	// Act
	actual := args.Map{"length": col.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// Test: LengthLock — concurrent reads during writes
// ==========================================

func Test_GenericCollection_LengthLock_ConcurrentReadsWrites(t *testing.T) {
	// Arrange
	tc := collectionLengthLockConcurrencyTestCase
	const writers = 100
	const readers = 100
	col := coregeneric.EmptyCollection[int]()

	wg := sync.WaitGroup{}
	wg.Add(writers + readers)

	var noNegativeLen atomic.Bool
	noNegativeLen.Store(true)

	// concurrent writers
	for i := 0; i < writers; i++ {
		go func(idx int) {
			col.AddLock(idx)
			wg.Done()
		}(i)
	}

	// concurrent readers
	for i := 0; i < readers; i++ {
		go func() {
			length := col.LengthLock()
			if length < 0 {
				noNegativeLen.Store(false)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	// Act
	actual := args.Map{
		"finalLength":   col.Length(),
		"noNegativeLen": noNegativeLen.Load(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// Test: IsEmptyLock — concurrent check with writes
// ==========================================

func Test_GenericCollection_IsEmptyLock_ConcurrentSafety(t *testing.T) {
	// Arrange
	tc := collectionIsEmptyLockConcurrencyTestCase
	const goroutines = 100
	col := coregeneric.EmptyCollection[int]()

	wg := sync.WaitGroup{}
	wg.Add(goroutines * 2)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			col.AddLock(idx)
			wg.Done()
		}(i)
		go func() {
			_ = col.IsEmptyLock() // must not panic
			wg.Done()
		}()
	}

	wg.Wait()

	// Act
	actual := args.Map{"length": col.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
