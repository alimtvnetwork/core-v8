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

package mutexbykey

import (
	"sync"

	"github.com/alimtvnetwork/core-v8/constants"
)

type mutexMap struct {
	items map[string]*sync.Mutex
}

var globalMutex = sync.Mutex{}

var items = make(
	map[string]*sync.Mutex,
	constants.ArbitraryCapacity10)

var internalMap = mutexMap{
	items: items,
}

func Get(key string) *sync.Mutex {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	mutex, has := internalMap.items[key]

	if has {
		return mutex
	}

	// not there
	newMutex := &sync.Mutex{}
	internalMap.items[key] = newMutex

	return newMutex
}

// Delete removes the mutex for the given key from the registry.
//
// WARNING: The caller MUST ensure no other goroutine currently holds or is
// waiting on the mutex for this key. Deleting a mutex while it is in use
// will cause a race condition — a new mutex may be created for the same key
// via Get(), allowing two goroutines to operate concurrently on the same
// logical key.
func Delete(key string) bool {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	_, has := internalMap.items[key]

	if has {
		delete(internalMap.items, key)
	}

	return has
}
