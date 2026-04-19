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

package codestack

import "runtime"

type newCreator struct {
	traces     newTraceCollection
	StackTrace newStacksCreator
}

func (it newCreator) Default() Trace {
	return it.Create(defaultInternalSkip)
}

func (it newCreator) SkipOne() Trace {
	return it.Create(Skip2)
}

func (it newCreator) Ptr(skipIndex int) *Trace {
	trace := it.Create(skipIndex + defaultInternalSkip)

	return &trace
}

func (it newCreator) Create(skipIndex int) Trace {
	pc, file, line, isOkay := runtime.Caller(skipIndex + defaultInternalSkip)

	if !isOkay {
		return Trace{
			SkipIndex: skipIndex,
			IsOkay:    false,
		}
	}

	funcInfo := runtime.FuncForPC(pc)
	fullFuncName := ""
	if funcInfo != nil {
		fullFuncName = funcInfo.Name()
	}

	fullMethodSignature, packageName, methodName := NameOf.All(fullFuncName)

	return Trace{
		SkipIndex:         skipIndex,
		PackageName:       packageName,
		MethodName:        methodName,
		PackageMethodName: fullMethodSignature,
		FilePath:          file,
		Line:              line,
		IsOkay:            isOkay,
		IsSkippable:       isSkippablePackage(packageName),
	}
}
