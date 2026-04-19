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

package corecmptests

// NOTE: The 9 uncovered statements in corecmp are all unreachable dead code.
//
// Functions: Byte, Integer, Integer8, Integer16, Integer32, Integer64, Time,
//            VersionSliceByte, VersionSliceInteger
//
// Each has a final `return corecomparator.NotEqual` after exhaustive
// if/else-if chains covering equal, less-than, and greater-than.
// Since these three conditions are logically exhaustive for comparable types,
// the fallthrough return is dead code — it exists only as a safety net.
//
// These 9 statements CANNOT be covered by any test input.
// Filed as known dead code in spec/13-app-issues/golang/41-unreachable-branches-corejson-coreonce.md.
