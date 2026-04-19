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

package dtformats

type Layout string

const (
	ANSIC       Layout = "Mon Jan _2 15:04:05 2006"
	UnixDate    Layout = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    Layout = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      Layout = "02 Jan 06 15:04 MST"
	RFC822Z     Layout = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      Layout = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     Layout = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    Layout = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     Layout = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano Layout = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     Layout = "3:04PM"
	Stamp       Layout = "Jan _2 15:04:05"
	StampMilli  Layout = "Jan _2 15:04:05.000"
	StampMicro  Layout = "Jan _2 15:04:05.000000"
	StampNano   Layout = "Jan _2 15:04:05.000000000"
)

func (layout Layout) Value() string {
	return string(layout)
}

func (layout Layout) Is(format string) bool {
	return string(layout) == format
}

func (layout Layout) IsTimeOnly() bool {
	return layout == Kitchen
}

func (layout Layout) IsTimeFocused() bool {
	return layout == Stamp ||
		layout == StampMilli ||
		layout == StampMicro ||
		layout == Kitchen ||
		layout == StampNano
}

func (layout Layout) IsDateTime() bool {
	return layout == ANSIC ||
		layout == UnixDate ||
		layout == RubyDate ||
		layout == RFC822 ||
		layout == RFC822Z ||
		layout == RFC850 ||
		layout == RFC1123 ||
		layout == RFC1123Z ||
		layout == RFC3339 ||
		layout == RFC3339Nano
}

func (layout Layout) HasTimeZone() bool {
	return layout == UnixDate ||
		layout == RubyDate ||
		layout == RFC822 ||
		layout == RFC822Z ||
		layout == RFC850 ||
		layout == RFC1123 ||
		layout == RFC1123Z ||
		layout == RFC3339 ||
		layout == RFC3339Nano
}
