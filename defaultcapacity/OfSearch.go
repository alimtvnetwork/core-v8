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

package defaultcapacity

import "github.com/alimtvnetwork/core-v8/constants"

func OfSearch(length int) int {
	if length <= constants.Capacity3 {
		return length
	}

	if length > constants.Capacity3 && length <= constants.N20 {
		return length / constants.Capacity3
	}

	defaultCapacity := length

	if length >= constants.ArbitraryCapacity1000 {
		defaultCapacity = constants.ArbitraryCapacity100
	} else if length > constants.ArbitraryCapacity250 {
		defaultCapacity = length / constants.N20
	} else if length >= constants.ArbitraryCapacity100 {
		defaultCapacity = length / constants.N10
	} else {
		defaultCapacity = length / constants.N5
	}

	return defaultCapacity
}
