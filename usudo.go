/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	check()
	this, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting usudo path: %v\n", err)
		os.Exit(1)
	}
	run(path.Join(path.Dir(this), "usudo-helper.exe"), os.Args[1:])
}
