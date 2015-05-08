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
func TestVersionParse(t *testing.T) {
	version, err := dpkg.ParseVersion("1.0-1")
	ok(t, err)

	equals(t, 0, version.Epoch)
	equals(t, "1.0", version.Version)
	equals(t, "1", version.Revision)
}
