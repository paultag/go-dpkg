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
// #include <dpkg/dpkg-db.h>
// #include <dpkg/pkg-array.h>
// #include <dpkg/pkg-spec.h>
// #include <dpkg/pkg-show.h>
import "C"

import (
	"fmt"
	"unsafe"
)

/*
 */
func (array *C.struct_pkg_array) toSlice() (ret []*C.struct_pkginfo) {
	width := unsafe.Sizeof(&C.struct_pkginfo{})
	base_pointer := unsafe.Pointer(array.pkgs)
	current_address := uintptr(base_pointer)
	for i := 0; i < int(array.n_pkgs); i++ {
		pkg_pointer := (**C.struct_pkginfo)(unsafe.Pointer(current_address))
		current_address += width
		pkg := *pkg_pointer
		ret = append(ret, pkg)
	}

	return
}

type Package struct {
	Want     rune
	EFlag    rune
	Status   rune
	Priority string
	Section  string
	Name     string

	/* C internals */
	cPackage *C.struct_pkginfo
}

func (pkg *C.struct_pkginfo) toPackage() *Package {
	return &Package{
		Name:     C.GoString(C.pkg_name(pkg, C.pnaw_nonambig)),
		Want:     rune(C.pkg_abbrev_want(pkg)),
		Status:   rune(C.pkg_abbrev_status(pkg)),
		EFlag:    rune(C.pkg_abbrev_eflag(pkg)),
		Priority: C.GoString(C.pkg_priority_name(pkg)),
		Section:  C.GoString(pkg.section),

		/* */
		cPackage: pkg,
	}
}

/*
 */
func Foo() {
	C.modstatdb_open(C.msdbrw_readonly)

	array := C.struct_pkg_array{}
	C.pkg_array_init_from_db(&array)
	C.pkg_array_sort(
		&array,
		(*C.pkg_sorter_func)(C.pkg_sorter_by_nonambig_name_arch),
	)

	for _, pkg := range array.toSlice() {
		if pkg.status == C.PKG_STAT_NOTINSTALLED {
			continue
		}

		goPkg := pkg.toPackage()
		fmt.Printf("%c%c %s %s\n",
			goPkg.Want,
			goPkg.Status,
			goPkg.Name,
			goPkg.Section,
		)

		// fmt.Printf("%c%c %c %s %s %s\n",
		// 	C.pkg_abbrev_want(pkg),
		// 	C.pkg_abbrev_status(pkg),
		// 	C.pkg_abbrev_eflag(pkg),
		// 	C.GoString(C.pkg_name(pkg, C.pnaw_nonambig)),
		// 	C.GoString(C.versiondescribe(
		// 		&pkg.installed.version,
		// 		C.vdew_nonambig,
		// 	)),
		// 	C.GoString(C.dpkg_arch_describe(pkg.installed.arch)))
	}
}
