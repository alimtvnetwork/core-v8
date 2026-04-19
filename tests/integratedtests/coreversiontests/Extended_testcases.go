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

package coreversiontests

type versionCreationCase struct {
	name          string
	input         string
	expectedMajor int
	expectedMinor int
	expectedPatch int
}

var versionCreationCases = []versionCreationCase{
	{"full version", "v1.2.3", 1, 2, 3},
	{"without v prefix", "1.2.3", 1, 2, 3},
	{"major only", "v5", 5, 0, 0},
	{"major.minor", "v1.2", 1, 2, 0},
	{"with build", "v1.2.3.4", 1, 2, 3},
	{"empty", "", -1, -1, -1},
	{"just v", "v", -1, -1, -1},
	{"spaces", "  v1.2.3  ", 1, 2, 3},
}
