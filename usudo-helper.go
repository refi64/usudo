/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func hang() {
	fmt.Print("Press Enter to exit...")
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
}

func errcode(e error) (int, error) {
	if e == nil {
		return 0, nil
	}

	e2, ok := e.(*exec.ExitError)
	if !ok {
		return 1, e
	}

	if w, ok := e2.Sys().(syscall.WaitStatus); ok {
		return int(w.ExitCode), nil
	} else {
		return 1, e
	}
}

func run(exe string, args ...string) int {
	cmd := exec.Command(exe, args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	ret, err := errcode(cmd.Run())
	if err != nil {
		fmt.Fprintf(os.Stderr, "usudo-helper: error running command: %v\n", err)
		hang()
		os.Exit(1)
	}

	return ret
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "usudo-helper needs a command to run!!\n")
		hang()
		os.Exit(1)
	}

	ret := run(os.Args[1], os.Args[2:]...)
	run("cmd.exe", "/C", "pause")
	os.Exit(ret)
}
