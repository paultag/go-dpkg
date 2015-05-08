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

package dpkg_test

import (
	"testing"

	"github.com/paultag/go-dpkg/dpkg"
)

/* */
func TestSimpleDependency(t *testing.T) {
	dependency, err := dpkg.ParseDepends("foo")
	ok(t, err)
	equals(t, 1, len(dependency))
}

/* */
func TestDeppossiDependency(t *testing.T) {
	dependency, err := dpkg.ParseDepends("foo, bar | baz | quix")
	ok(t, err)
	equals(t, 2, len(dependency))
	equals(t, 3, len(dependency[1].Possibilities))
}

/* */
func TestInvalidDependency(t *testing.T) {
	dependency, err := dpkg.ParseDepends("etc foo")

	if err == nil {
		t.FailNow()
	}

	assert(t, dependency == nil, "Dependency isn't nil.")
}

/* */
func TestVersionDependency(t *testing.T) {
	_, err := dpkg.ParseDepends("libc6 (>= 2.2.1), exim | mail-transport-agent")
	ok(t, err)
}

/* */
func TestArchDependency(t *testing.T) {
	t.Skip()
	/* This is broken */
	_, err := dpkg.ParseDepends("foo [i386]")
	ok(t, err)
}
