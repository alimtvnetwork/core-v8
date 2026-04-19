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

package corestrtests

// ══════════════════════════════════════════════════════════════════════════════
// Coverage — coredata/corestr remaining gaps (90 uncovered lines)
//
// Most gaps are nil-receiver guards and defensive dead code across:
// - CharCollectionMap: nil map guards, json marshal errors
// - CharHashsetMap: nil hashset guards, defensive branches
// - Collection: nil receiver, json errors
// - CollectionsOfCollection: empty return guards
// - LinkedCollections/LinkedList/LinkedListNode: nil node guards
// - SimpleSlice: nil receiver, boundary guards
// - SimpleStringOnce: lazy init fallback
// - ValidValue: dead fallback return
//
// These require either internal tests or represent defensive dead code.
// Documented for future internal test coverage.
// ══════════════════════════════════════════════════════════════════════════════
