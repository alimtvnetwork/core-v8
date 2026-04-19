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

package converters

import (
	"github.com/alimtvnetwork/core/internal/strutilinternal"
)

type StringsToMapConverter []string

func (it StringsToMapConverter) SafeStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	return it
}

func (it StringsToMapConverter) Strings() []string {
	return it
}

func (it *StringsToMapConverter) Length() int {
	if it == nil || *it == nil {
		return 0
	}

	return len(*it)
}

func (it *StringsToMapConverter) IsEmpty() bool {
	return it.Length() == 0
}

func (it *StringsToMapConverter) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *StringsToMapConverter) LastIndex() int {
	return it.Length() - 1
}

func (it StringsToMapConverter) LineSplitMapOptions(
	isTrim bool,
	splitter string,
) map[string]string {
	if isTrim {
		return it.LineSplitMapTrim(splitter)
	}

	return it.LineSplitMap(splitter)
}

func (it StringsToMapConverter) LineProcessorMapOptions(
	isTrimBefore bool,
	processorFunc func(line string) (key, val string),
) map[string]string {
	return strutilinternal.
		SliceToMapConverter(it.Strings()).
		LineProcessorMapOptions(
			isTrimBefore,
			processorFunc)
}

func (it StringsToMapConverter) LineProcessorMapStringIntegerTrim(
	processorFunc func(line string) (key string, val int),
) map[string]int {
	return it.LineProcessorMapStringIntegerOptions(
		true,
		processorFunc)
}

func (it StringsToMapConverter) LineProcessorMapStringIntegerOptions(
	isTrimBefore bool,
	processorFunc func(line string) (key string, val int),
) map[string]int {
	return strutilinternal.
		SliceToMapConverter(it.Strings()).
		LineProcessorMapStringIntegerOptions(
			isTrimBefore,
			processorFunc)
}

func (it StringsToMapConverter) LineProcessorMapStringAnyTrim(
	processorFunc func(line string) (key string, val any),
) map[string]any {
	return it.LineProcessorMapStringAnyOptions(
		true,
		processorFunc)
}

func (it StringsToMapConverter) LineProcessorMapStringAnyOptions(
	isTrimBefore bool,
	processorFunc func(line string) (key string, val any),
) map[string]any {
	return strutilinternal.
		SliceToMapConverter(it.Strings()).
		LineProcessorMapStringAnyOptions(
			isTrimBefore,
			processorFunc)
}

func (it StringsToMapConverter) LineSplitMapTrim(
	splitter string,
) map[string]string {
	return strutilinternal.
		SliceToMapConverter(it.Strings()).
		LineSplitMapTrim(
			splitter,
		)
}

func (it StringsToMapConverter) LineSplitMap(
	splitter string,
) map[string]string {
	return strutilinternal.
		SliceToMapConverter(it.Strings()).
		LineSplitMap(
			splitter,
		)
}
