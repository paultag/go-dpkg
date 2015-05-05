go-dpkg
=======

Native Go bindings to abstract a cgo interface to `libdpkg`. Keep in mind
that this is a very voiltile interface, so any breakage against unstable's
`libdpkg-dev` should be reported.

Licensing
---------

This library statically links against `libdpkg.a`, which is licensed under
the terms of the GPLv2+. No matter the terms that are picked for this code,
the combined work is a derived work, so any distributions of the binary must
comply with the GPLv2+.
