/* Copyright (c) 2015  Paul R. Tagliamonte <paultag@debian.org>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along
 * with this program; if not, write to the Free Software Foundation, Inc.,
 * 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA. */

package dpkg

// #cgo LDFLAGS: /usr/lib/libdpkg.a
// #cgo CFLAGS: -DLIBDPKG_VOLATILE_API
// #include <dpkg/version.h>
import "C"

type Arch struct {
	Name string
}

func (arch *C.struct_dpkg_arch) toArch() *Arch {
	return &Arch{
		Name: C.GoString(arch.name),
	}
}
