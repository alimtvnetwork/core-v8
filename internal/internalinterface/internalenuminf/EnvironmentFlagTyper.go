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

package internalenuminf

type EnvironmentFlagTyper interface {
	BasicEnumer
	ByteValuePlusEqualer

	ToNumberString() string
	IsValidInvalidChecker

	IsVerbose() bool
	IsDebug() bool
	IsProduction() bool
	IsStackTrace() bool
	IsJson() bool
	IsLog() bool
	IsLoggedIn() bool

	SessionDetailsDynamic() []byte
	KeyValues() map[string]string
	AnyKeyValues() map[string]any
	NameBoolMap() map[string]bool

	FlagValue(name string) string
	FlagAnyValue(name string) any
	FlagAnyValueReflectSet(name string, toPointer any) error

	HasFlag(name string) bool
	HasAnyFlag(names ...string) bool
	HasAllFlags(names ...string) bool

	IsFlagEnabled(name string) bool
	IsFlagDisabled(name string) bool
}
