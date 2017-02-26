/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// https://blogs.msdn.microsoft.com/twistylittlepassagesallalike/2011/04/23/everyone-quotes-command-line-arguments-the-wrong-way/

func writeslashes(buf *bytes.Buffer, n int) {
	buf.WriteString(strings.Repeat(`\`, n))
}

func quoteargs(args []string) string {
	var res bytes.Buffer

	if len(args) == 0 {
		return ""
	}

	for _, arg := range args {
		if arg != "" && !strings.ContainsAny(arg, " \t\n\v\"") {
			res.WriteString(arg + " ")
			continue
		}

		res.WriteString(`"`)

		for i := 0; ; i++ {
			nslashes := 0

			for ; i < len(arg) && arg[i] == '\\'; i++ {
				nslashes++
			}

			if i == len(arg) {
				writeslashes(&res, nslashes*2)
				break
			} else {
				if arg[i] == '"' {
					writeslashes(&res, nslashes*2+1)
				} else {
					writeslashes(&res, nslashes)
				}
				res.WriteString(string(arg[i]))
			}
		}

		res.WriteString(`" `)
	}

	// Cut out the trailing space.
	res.Truncate(res.Len() - 1)
	return res.String()
}

func check() {
	if len(os.Args) == 0 {
		fmt.Fprintf(os.Stderr, "usudo needs a program to run!\n")
		os.Exit(1)
	}
}

func run(exe string, unquoted_args []string) {
	args := quoteargs(unquoted_args)

	err := ShellExecute("runas", exe, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error running program: %v\n", err)
		os.Exit(1)
	}
}
