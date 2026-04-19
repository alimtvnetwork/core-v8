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

package coretestcases

// IsFailedToMatch returns the inverse of IsMatching.
//
// Use when validating that a mismatch is expected.
func (it *GenericGherkins[TInput, TExpect]) IsFailedToMatch() bool {
	return !it.IsMatching
}

// HasExtraArgs returns true if ExtraArgs is defined and non-empty.
func (it *GenericGherkins[TInput, TExpect]) HasExtraArgs() bool {
	return it != nil && len(it.ExtraArgs) > 0
}

// GetExtra returns the value for a key in ExtraArgs, or nil if not found.
func (it *GenericGherkins[TInput, TExpect]) GetExtra(key string) any {
	if it == nil || it.ExtraArgs == nil {
		return nil
	}

	return it.ExtraArgs[key]
}

// GetExtraAsString returns the string value for a key in ExtraArgs.
func (it *GenericGherkins[TInput, TExpect]) GetExtraAsString(key string) (string, bool) {
	if it == nil || it.ExtraArgs == nil {
		return "", false
	}

	return it.ExtraArgs.GetAsString(key)
}

// GetExtraAsBool returns the bool value for a key in ExtraArgs.
func (it *GenericGherkins[TInput, TExpect]) GetExtraAsBool(key string) (value bool, isValid bool) {
	if it == nil || it.ExtraArgs == nil {
		return false, false
	}

	return it.ExtraArgs.GetAsBool(key)
}

// GetExtraAsBoolDefault returns the bool value for a key in ExtraArgs,
// or the default value if not found or not a bool.
func (it *GenericGherkins[TInput, TExpect]) GetExtraAsBoolDefault(key string, defaultVal bool) bool {
	if it == nil || it.ExtraArgs == nil {
		return defaultVal
	}

	return it.ExtraArgs.GetAsBoolDefault(key, defaultVal)
}
