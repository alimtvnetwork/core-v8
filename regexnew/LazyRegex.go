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

package regexnew

import (
	"errors"
	"regexp"
	"sync"
)

// LazyRegex
//
//	lazy regex for future unwrapping or compiled but only once.
type LazyRegex struct {
	mu           sync.Mutex
	isCompiled   bool
	isApplicable bool // no err, pattern defined, not null
	pattern      string
	regex        *regexp.Regexp
	compiledErr  error
	compiler     func(pattern string) (
		*regexp.Regexp, error,
	) // must be a locked function, cannot use non lock version
}

func (it *LazyRegex) IsNull() bool {
	return it == nil
}

func (it *LazyRegex) IsDefined() bool {
	return it != nil && it.pattern != "" && it.compiler != nil
}

func (it *LazyRegex) IsUndefined() bool {
	return it == nil || it.pattern == "" || it.compiler == nil
}

// IsApplicable
//
//	it unwraps the regex and compiles so take memory for once.
func (it *LazyRegex) IsApplicable() bool {
	if it == nil {
		return false
	}

	it.mu.Lock()
	if it.isApplicable {
		it.mu.Unlock()
		return true
	}
	it.mu.Unlock()

	if it.IsUndefined() {
		return false
	}

	// compile required
	// fine to swallow err
	// updates isApplicable
	it.Compile()

	it.mu.Lock()
	defer it.mu.Unlock()
	return it.isApplicable
}

// Compile
//
//	it is done through the locking mechanism
func (it *LazyRegex) Compile() (regex *regexp.Regexp, err error) {
	it.mu.Lock()
	defer it.mu.Unlock()

	if it.isCompiled {
		return it.regex, it.compiledErr
	}

	if it.IsUndefined() {
		return nil, errors.New("lazy regex is undefined or nil")
	}

	// defined
	compiledRegex, regExErr := it.compiler(it.pattern) // compiler should be locked func, must be
	it.isApplicable = compiledRegex != nil && regExErr == nil
	it.regex = compiledRegex
	it.compiledErr = regExErr
	it.isCompiled = true

	return compiledRegex, regExErr
}

func (it *LazyRegex) CompileMust() (regex *regexp.Regexp) {
	regexCompiled, err := it.Compile()

	if err != nil {
		panic(err)
	}

	return regexCompiled
}

func (it *LazyRegex) IsCompiled() bool {
	if it == nil {
		return false
	}

	it.mu.Lock()
	defer it.mu.Unlock()

	return it.isCompiled
}

func (it *LazyRegex) OnRequiredCompiled() error {
	if it == nil {
		return errors.New("nil LazyRegex cannot compile")
	}

	if it.IsCompiled() {
		return it.compiledErr
	}

	_, err := it.Compile()

	return err
}

func (it *LazyRegex) OnRequiredCompiledMust() {
	err := it.OnRequiredCompiled()

	if err != nil {
		panic(err)
	}
}

func (it *LazyRegex) HasError() bool {
	// fine to swallow
	it.OnRequiredCompiled()

	return it != nil && it.compiledErr != nil
}

func (it *LazyRegex) HasAnyIssues() bool {
	if it == nil {
		return true
	}

	return !it.IsApplicable()
}

func (it *LazyRegex) IsInvalid() bool {
	if it == nil {
		return true
	}

	return !it.IsApplicable()
}

func (it *LazyRegex) CompiledError() error {
	return it.OnRequiredCompiled()
}

func (it *LazyRegex) Error() error {
	return it.OnRequiredCompiled()
}

func (it *LazyRegex) MustBeSafe() {
	compiledErr := it.CompiledError()

	if compiledErr != nil {
		panic(compiledErr)
	}
}

func (it *LazyRegex) String() (pattern string) {
	if it == nil {
		return ""
	}

	return it.pattern
}

func (it *LazyRegex) FullString() (detail string) {
	if it == nil {
		return ""
	}

	isApplicable := it.IsApplicable()
	isCompiled := it.IsCompiled()
	compiledErr := it.CompiledError()

	newMap := map[string]any{
		"pattern":      it.Pattern(),
		"isCompiled":   isCompiled,
		"isApplicable": isApplicable,
		"error":        compiledErr,
	}

	return prettyJson(newMap)
}

func (it *LazyRegex) Pattern() (pattern string) {
	if it == nil {
		return ""
	}

	return it.pattern
}

func (it *LazyRegex) MatchError(matchingPattern string) error {
	regEx, compiledErr := it.Compile()

	if regEx != nil && regEx.MatchString(matchingPattern) {
		return nil
	}

	return regExMatchValidationError(
		it.pattern,
		matchingPattern,
		compiledErr,
		regEx)
}

// MatchUsingFuncError
//
//	creates new regex using lock
//	and then calls match.
//	On condition mismatch returns error
//	or else nil
func (it *LazyRegex) MatchUsingFuncError(
	comparing string,
	matchFunc RegexValidationFunc,
) error {
	regEx, compiledErr := it.Compile()

	if regEx != nil && matchFunc(regEx, comparing) {
		return nil
	}

	return regExMatchValidationError(
		it.pattern,
		comparing,
		compiledErr,
		regEx)
}

func (it *LazyRegex) IsMatch(
	comparing string,
) bool {
	regEx, compiledErr := it.Compile()

	if regEx == nil || compiledErr != nil {
		return false
	}

	return regEx.MatchString(comparing)
}

func (it *LazyRegex) IsMatchBytes(
	comparingBytes []byte,
) bool {
	regEx, compiledErr := it.Compile()

	if regEx == nil || compiledErr != nil {
		return false
	}

	return regEx.Match(comparingBytes)
}

func (it *LazyRegex) IsFailedMatch(
	comparing string,
) bool {
	regEx, compiledErr := it.Compile()

	if regEx == nil || compiledErr != nil {
		return true
	}

	return !regEx.MatchString(comparing)
}

func (it *LazyRegex) IsFailedMatchBytes(
	comparingBytes []byte,
) bool {
	regEx, compiledErr := it.Compile()

	if regEx == nil || compiledErr != nil {
		return true
	}

	return !regEx.Match(comparingBytes)
}

func (it *LazyRegex) FirstMatchLine(
	content string,
) (firstMatch string, isInvalidMatch bool) {
	regEx, compiledErr := it.Compile()

	if regEx == nil || compiledErr != nil {
		return "", true
	}

	lines := regEx.FindStringSubmatch(content)

	if len(lines) > 0 {
		// valid
		return lines[0], false
	}

	// invalid
	return "", true
}
