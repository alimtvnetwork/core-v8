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

// ══════════════════════════════════════════════════════════════════════════════
//  — coredata/coredynamic accepted gaps
//
// Categories:
//
// 1. Nil receiver guards after Lock() — dead code
//    - CollectionLock.LengthLock:15 — Lock() panics on nil before nil check
//    - CollectionLock.ItemsLock:125-127 — nil items after Lock()
//
// 2. json.Marshal error branches — defensive dead code
//    - AnyCollection.JsonString:485-487, JsonStringMust:495-499
//    - DynamicCollection.JsonString:416-418, JsonStringMust:426-430
//    - DynamicJson.go:54 (MarshalJSON), :123 (ParseInjectMust panic)
//    - DynamicJson.go:139-141, 149-151, 159-163 (cascading errors)
//    - Collection.JsonString:355-357, JsonStringMust:364-365
//    - TypedDynamic.JsonString:117-119
//
// 3. ReflectSetFromTo byte conversion — requires specific reflect types
//    - ReflectSetFromTo.go:159-167 (marshal error in byte conversion)
//    - ReflectSetFromTo.go:174-180 (unexpected state fallback)
//
// 4. Unreachable code
//    - ReflectInterfaceVal.go:20 — exhaustive if/else makes line 20 dead
//    - MapAnyItems.GetItemRef:362-373 — exhaustive Ptr check (lines 350+356)
//
// 5. KeyVal/KeyValCollection defensive errors
//    - KeyVal.ReflectSetKeyValue:134-136
//    - KeyValCollection lines 139-141, 342-344, 365-366, 385-387, 395-397
//
// 6. MapAnyItems.ToKeyValCollection AddAny error:903-904
//
// ══════════════════════════════════════════════════════════════════════════════
