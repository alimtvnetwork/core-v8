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

package loggerinf

import (
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coreinterface/enuminf"
	"github.com/alimtvnetwork/core-v8/coreinterface/errcoreinf"
)

type StandardLogger interface {
	StandardLoggerChecker
	ConditionalStandardLogger

	FullLogger() FullLogger
	EnvOptioner() enuminf.EnvironmentOptioner

	TaskNamedLogger(
		taskName string,
	) StandardLogger

	TaskWithPayloadLogger(
		taskName string,
		payloadAny any, // can be bytes, payloadWrapper, can be any
	) StandardLogger

	GetLoggerByTaskName(taskName string) StandardLogger
	GetLoggerByTaskNamer(taskNamer enuminf.Namer) StandardLogger

	Success(args ...any) StandardLogger
	Info(args ...any) StandardLogger
	Trace(args ...any) StandardLogger
	Debug(args ...any) StandardLogger
	Warn(args ...any) StandardLogger
	Error(args ...any) StandardLogger
	Fatal(args ...any) StandardLogger
	Panic(args ...any) StandardLogger

	SuccessFmt(format string, args ...any) StandardLogger
	InfoFmt(format string, args ...any) StandardLogger
	TraceFmt(format string, args ...any) StandardLogger
	DebugFmt(format string, args ...any) StandardLogger
	WarnFmt(format string, args ...any) StandardLogger
	ErrorFmt(format string, args ...any) StandardLogger
	FatalFmt(format string, args ...any) StandardLogger
	PanicFmt(format string, args ...any) StandardLogger

	SuccessExtend() SingleLogger
	InfoExtend() SingleLogger
	TraceExtend() SingleLogger
	DebugExtend() SingleLogger
	WarnExtend() SingleLogger
	FatalExtend() SingleLogger
	PanicExtend() SingleLogger

	ErrorDirect(err error) StandardLogger
	OnErrStackTrace(err error) StandardLogger
	ErrInterface(errInf errcoreinf.BaseErrorOrCollectionWrapper) StandardLogger
	ErrInterfaceStackTraces(errInfWithStackTraces errcoreinf.BaseErrorOrCollectionWrapper) StandardLogger

	ReflectSetter

	InfoOrError(isError bool) SingleLogger
	Log(loggerType enuminf.LoggerTyper) StandardLogger

	ErrorJsoner(jsoner corejson.Jsoner) StandardLogger
	DebugJsoner(jsoner corejson.Jsoner) StandardLogger
	ErrorJsonerTitle(title string, jsoner corejson.Jsoner) StandardLogger
	DebugJsonerTitle(title string, jsoner corejson.Jsoner) StandardLogger
}
