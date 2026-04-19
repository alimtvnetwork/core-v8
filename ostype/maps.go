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

package ostype

import "github.com/alimtvnetwork/core/osconsts"

var (
	osTypesStrings = [...]string{
		Any:           osconsts.Any,
		Unknown:       osconsts.Unknown,
		Windows:       osconsts.Windows,
		Linux:         osconsts.Linux,
		DarwinOrMacOs: osconsts.DarwinOrMacOs,
		FreeBsd:       osconsts.FreeBsd,
		NetBsd:        osconsts.NetBsd,
		OpenBsd:       osconsts.OpenBsd,
		DragonFly:     osconsts.DragonFly,
		Android:       osconsts.Android,
		Plan9:         osconsts.Plan9,
		Solaris:       osconsts.Solaris,
		Nacl:          osconsts.Nacl,
		Illumos:       osconsts.Illumos,
		IOs:           osconsts.IOs,
		Aix:           osconsts.Aix,
	}

	OsVariantToStringMap = map[Variation]string{
		Any:           osconsts.Any,
		Unknown:       osconsts.Unknown,
		Windows:       osconsts.Windows,
		Linux:         osconsts.Linux,
		DarwinOrMacOs: osconsts.DarwinOrMacOs,
		FreeBsd:       osconsts.FreeBsd,
		NetBsd:        osconsts.NetBsd,
		OpenBsd:       osconsts.OpenBsd,
		DragonFly:     osconsts.DragonFly,
		Android:       osconsts.Android,
		Plan9:         osconsts.Plan9,
		Solaris:       osconsts.Solaris,
		Nacl:          osconsts.Nacl,
		Illumos:       osconsts.Illumos,
		IOs:           osconsts.IOs,
		Aix:           osconsts.Aix,
	}

	OsStringToVariantMap = map[string]Variation{
		osconsts.Any:           Any,
		osconsts.Unknown:       Unknown,
		osconsts.Windows:       Windows,
		osconsts.Linux:         Linux,
		osconsts.DarwinOrMacOs: DarwinOrMacOs,
		osconsts.FreeBsd:       FreeBsd,
		osconsts.NetBsd:        NetBsd,
		osconsts.OpenBsd:       OpenBsd,
		osconsts.DragonFly:     DragonFly,
		osconsts.Android:       Android,
		osconsts.Plan9:         Plan9,
		osconsts.Solaris:       Solaris,
		osconsts.Nacl:          Nacl,
		osconsts.Illumos:       Illumos,
		osconsts.IOs:           IOs,
		osconsts.Aix:           Aix,
	}
)
