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

	return &Version{
		Epoch:    int(version.epoch),
		Version:  C.GoString(version.version),
		Revision: C.GoString(version.revision),
	}, nil
}
