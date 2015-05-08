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
// #include <dpkg/pkg-show.h>
// #include <dpkg/parsedump.h>
// #include <dpkg/ehandle.h>
// #include <malloc.h>
// #include <stdlib.h>
//
// void go_dpkg_error_hider(const char *emsg, const void *data) {
//     return;
// }
//
// void parse_dependency(
//     bool *error,
//     struct pkginfo *pkg,
//     struct pkgbin *pkgbin,
//     struct parsedb_state *ps,
//     const char *value,
//     const struct fieldinfo *fip
// ) {
//     jmp_buf jbuf;
//     if (setjmp(jbuf)) {
//         pop_error_context(ehflag_normaltidy); /* EVERYTHING IS NORMAL */
//         *error = true;
//         return;
//     }
//     *error = false;
//     push_error_context_jump(&jbuf, go_dpkg_error_hider, NULL);
//     f_dependency(pkg, pkgbin, ps, value, fip);
//     pop_error_context(ehflag_normaltidy);
// }
//
import "C"

import (
	"errors"
	"unsafe"
)

/*
 */
type Dependency struct {
	Name         string
	Arch         *Arch
	Version      *Version
	ImplicitArch bool
}

/*
 */
type Relation struct {
	Possibilities []*Dependency
	Type          string
}

/*
 */
func (dependency *C.struct_dependency) toRelation() *Relation {
	relation := Relation{
		Possibilities: []*Dependency{},
		Type:          "",
	}

	dep := dependency.list
	for {
		relation.Possibilities = append(
			relation.Possibilities,
			dep.toDependency(),
		)
		dep = dep.next
		if dep == nil {
			break
		}
	}

	return &relation
}

/*
 */
func (dependency *C.struct_dependency) toRelations() []*Relation {
	relations := []*Relation{}

	for {
		if dependency == nil {
			break
		}

		relations = append(relations, dependency.toRelation())
		dependency = dependency.next
	}
	return relations
}

/*
 */
func (dependency *C.struct_deppossi) toDependency() *Dependency {
	arch := (*Arch)(nil)

	if dependency.arch != nil {
		arch = dependency.arch.toArch()
	}

	return &Dependency{
		Name:         C.GoString(dependency.ed.name),
		Version:      dependency.version.toVersion(),
		ImplicitArch: bool(dependency.arch_is_implicit),
		Arch:         arch,
	}
}

/*
 */
func ParseDepends(depends string) ([]*Relation, error) {
	pkg := C.struct_pkginfo{}
	pkgBin := C.struct_pkgbin{}

	fakeName := C.CString("<go-dpkg>")
	defer C.free(unsafe.Pointer(fakeName))

	name := C.CString("Depends") /* Don't change without changing namelen */
	defer C.free(unsafe.Pointer(name))

	cDepends := C.CString(depends + "\n")
	defer C.free(unsafe.Pointer(cDepends))

	ps := C.struct_parsedb_state{
		_type:    0,
		flags:    0,
		pkg:      nil, //&pkg, /* XXX: OUCH */
		pkgbin:   &pkgBin,
		data:     nil,
		dataptr:  nil,
		endptr:   nil,
		filename: fakeName,
		fd:       -1,
		lno:      -1,
	}

	fi := C.struct_fieldinfo{
		name:    name,
		namelen: 7,
		rcall:   &C.f_dependency,
		wcall:   nil, /* Hah. This'll backfire some day */
		integer: C.dep_depends,
	}

	err := C.bool(false)

	C.parse_dependency(&err, &pkg, &pkgBin, &ps, cDepends, &fi)

	if bool(err) == true {
		return nil, errors.New("ohshit")
	}

	return pkgBin.depends.toRelations(), nil
}
