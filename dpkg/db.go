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
// #include <dpkg/dpkg-db.h>
import "C"

import (
	"fmt"
)

type Version struct {
	Epoch    int
	Version  string
	Revision string
}

func ParseVersion(str string) (ret *Version, err error) {
	version := C.struct_dpkg_version{}
	derr := C.struct_dpkg_error{}
	failed := C.parseversion(&version, C.CString(str), &derr)

	if failed < 0 {
		return nil, fmt.Errorf("Bad version syntax: %s", C.GoString(derr.str))
	}

	return version.toVersion(), nil
}

func (version *C.struct_dpkg_version) toVersion() *Version {
	return &Version{
		Epoch:    int(version.epoch),
		Version:  C.GoString(version.version),
		Revision: C.GoString(version.revision),
	}
}

func (this *Version) toCStruct() *C.struct_dpkg_version {
	return &C.struct_dpkg_version{
		epoch:    C.uint(this.Epoch),
		version:  C.CString(this.Version),
		revision: C.CString(this.Revision),
	}
}

func (this *Version) Describe() string {
	version := this.toCStruct()
	explain := C.versiondescribe(
		version,
		C.vdew_nonambig,
	)
	return C.GoString(explain)
}
