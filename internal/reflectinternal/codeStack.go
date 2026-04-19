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

package reflectinternal

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/alimtvnetwork/core/constants"
)

type codeStack struct{}

func (it codeStack) New(skipStack int) StackTrace {
	pc, file, line, isOkay := runtime.Caller(skipStack + defaultInternalSkip)
	funcInfo := runtime.FuncForPC(pc)
	fullFuncName := funcInfo.Name()

	fullMethodSignature, packageName, methodName := GetFunc.All(fullFuncName)

	return StackTrace{
		SkipIndex:         skipStack,
		PackageName:       packageName,
		MethodName:        methodName,
		PackageMethodName: fullMethodSignature,
		FilePath:          file,
		Line:              line,
		IsOkay:            isOkay,
	}
}

func (it codeStack) NewDefault() StackTrace {
	return it.New(defaultInternalSkip)
}

func (it codeStack) MethodName(skipStack int) string {
	pc, _, _, _ := runtime.Caller(skipStack + defaultInternalSkip)
	funcInfo := runtime.FuncForPC(pc)
	fullFuncName := funcInfo.Name()

	_, _, methodName := GetFunc.All(fullFuncName)

	return methodName
}

func (it codeStack) MethodNameWithLine(skipStack int) string {
	pc, _, line, _ := runtime.Caller(skipStack + defaultInternalSkip)
	funcInfo := runtime.FuncForPC(pc)
	fullFuncName := funcInfo.Name()

	_, _, methodName := GetFunc.All(fullFuncName)

	return fmt.Sprintf(
		"%s:%d",
		methodName,
		line,
	)
}

func (it codeStack) FileWithLine(skipStack int) string {
	pc, file, line, _ := runtime.Caller(skipStack + defaultInternalSkip)
	funcInfo := runtime.FuncForPC(pc)
	fullFuncName := funcInfo.Name()

	_, _, methodName := GetFunc.All(fullFuncName)

	return fmt.Sprintf(
		shortStringFormat,
		methodName,
		line,
		file,
		line,
	)
}

func (it codeStack) LastFileWithLines(skipStack, count int) []string {
	lines := make([]string, 0, count)

	for i := 0; i < count; i++ {
		lines = append(lines, it.FileWithLine(skipStack+i))
	}

	return lines
}

func (it codeStack) LastFileWithLine(skipStack, count int) string {
	lines := it.LastFileWithLines(defaultInternalSkip+skipStack, count)

	return strings.Join(lines, constants.NewLineUnix)
}

func (it codeStack) FilePath(skipStack int) string {
	_, file, _, isOkay := runtime.Caller(skipStack + defaultInternalSkip)

	if isOkay {
		return file
	}

	return constants.EmptyString
}

func (it codeStack) NewFileWithLines(skipStack, count int) []FileWithLine {
	lines := make([]FileWithLine, 0, count)

	for i := 0; i < count; i++ {
		_, file, line, isOkay := runtime.Caller(skipStack + defaultInternalSkip + i)

		isCallerFailed := !isOkay

		if isCallerFailed {
			return lines
		}

		f := FileWithLine{
			FilePath: file,
			Line:     line,
		}

		lines = append(
			lines,
			f,
		)
	}

	return lines
}

func (it codeStack) NewFileWithLine(skipStack int) FileWithLine {
	_, file, line, _ := runtime.Caller(skipStack + defaultInternalSkip)

	return FileWithLine{
		FilePath: file,
		Line:     line,
	}
}

func (it codeStack) NewStacks(skipStack, count int) []StackTrace {
	slice := make([]StackTrace, 0, count)

	for i := 0; i < count; i++ {
		f := it.New(defaultInternalSkip + skipStack)

		slice = append(
			slice,
			f,
		)
	}

	return slice
}

func (it codeStack) StacksStrings(skipStack int) []string {
	return it.StacksStringsCount(
		skipStack+defaultInternalSkip,
		defaultStackCount,
	)
}

func (it codeStack) StacksStringsCount(skipStack, count int) []string {
	fileWithLines := it.NewFileWithLines(
		skipStack,
		count,
	)

	lines := make([]string, 0, len(fileWithLines))

	for _, fileWithLine := range fileWithLines {
		newLine := fmt.Sprintf(
			fileWithLineFormat,
			fileWithLine.FilePath,
			fileWithLine.Line,
		)

		lines = append(lines, newLine)
	}

	return lines
}

// StacksStringsFiltered returns stack trace lines filtered to exclude
// Go standard library frames (runtime/, testing/, etc.).
func (it codeStack) StacksStringsFiltered(skipStack, count int) []string {
	goRoot := runtime.GOROOT()
	fileWithLines := it.NewFileWithLines(
		skipStack+defaultInternalSkip,
		count+4, // fetch extra to compensate for filtered frames
	)

	lines := make([]string, 0, count)

	for _, fileWithLine := range fileWithLines {
		if isSystemLibraryPath(goRoot, fileWithLine.FilePath) {
			continue
		}

		newLine := fmt.Sprintf(
			fileWithLineFormat,
			fileWithLine.FilePath,
			fileWithLine.Line,
		)

		lines = append(lines, newLine)

		if len(lines) >= count {
			break
		}
	}

	return lines
}

// isSystemLibraryPath returns true if the file path belongs to
// Go standard library or runtime (under GOROOT).
func isSystemLibraryPath(goRoot, filePath string) bool {
	if goRoot != "" && strings.HasPrefix(filePath, goRoot) {
		return true
	}

	return false
}

func (it codeStack) StacksString(skipStack int) string {
	lines := it.StacksStrings(skipStack + defaultInternalSkip)

	return strings.Join(lines, constants.NewLineUnix)
}

func (it codeStack) StacksStringDefault(skipStack int) string {
	return it.StacksStringCount(
		skipStack+defaultInternalSkip,
		defaultStackCount,
	)
}

func (it codeStack) StacksStringCount(skipStack, count int) string {
	lines := it.StacksStringsCount(skipStack+defaultInternalSkip, count)

	joinedLines := strings.Join(lines, "\n  - ")

	return fmt.Sprintf(
		"%s :\n  - %s",
		constants.StackTrace,
		joinedLines,
	)
}

func (it codeStack) SingleStack(skipStack int) string {
	lines := it.StacksStringsCount(skipStack+defaultInternalSkip, 1)

	if len(lines) > 0 {
		return lines[0]
	}

	return ""
}
