// From https://github.com/AllenDang/w32/blob/master/shell32.go
// Copyright (c) 2010-2012 The w32 Authors. All rights reserved.
// Modifications (c) 2017 Ryan Gonzalez.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
// 3. The names of the authors may not be used to endorse or promote products
//    derived from this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE AUTHORS ``AS IS'' AND ANY EXPRESS OR
// IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
// OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY DIRECT, INDIRECT,
// INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
// THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package main

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

var (
	modshell32       = syscall.NewLazyDLL("shell32.dll")
	procShellExecute = modshell32.NewProc("ShellExecuteW")
)

const (
	ERROR_FILE_NOT_FOUND   = 2
	ERROR_PATH_NOT_FOUND   = 3
	ERROR_BAD_FORMAT       = 11
	SE_ERR_ACCESSDENIED    = 5
	SE_ERR_OOM             = 8
	SE_ERR_DLLNOTFOUND     = 32
	SE_ERR_SHARE           = 26
	SE_ERR_ASSOCINCOMPLETE = 27
	SE_ERR_DDETIMEOUT      = 28
	SE_ERR_DDEFAIL         = 29
	SE_ERR_DDEBUSY         = 30
	SE_ERR_NOASSOC         = 31
)

func ShellExecute(lpOperation, lpFile, lpParameters string) error {
	var op, param uintptr
	if len(lpOperation) != 0 {
		op = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpOperation)))
	}
	if len(lpParameters) != 0 {
		param = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpParameters)))
	}

	ret, _, _ := procShellExecute.Call(
		uintptr(0),
		op,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpFile))),
		param,
		uintptr(0),
		uintptr(1))

	errorMsg := ""
	if ret != 0 && ret <= 32 {
		switch int(ret) {
		case ERROR_FILE_NOT_FOUND:
			errorMsg = "The specified file was not found."
		case ERROR_PATH_NOT_FOUND:
			errorMsg = "The specified path was not found."
		case ERROR_BAD_FORMAT:
			errorMsg = "The .exe file is invalid (non-Win32 .exe or error in .exe image)."
		case SE_ERR_ACCESSDENIED:
			errorMsg = "The operating system denied access to the specified file."
		case SE_ERR_ASSOCINCOMPLETE:
			errorMsg = "The file name association is incomplete or invalid."
		case SE_ERR_DDEBUSY:
			errorMsg = "The DDE transaction could not be completed because other DDE transactions were being processed."
		case SE_ERR_DDEFAIL:
			errorMsg = "The DDE transaction failed."
		case SE_ERR_DDETIMEOUT:
			errorMsg = "The DDE transaction could not be completed because the request timed out."
		case SE_ERR_DLLNOTFOUND:
			errorMsg = "The specified DLL was not found."
		case SE_ERR_NOASSOC:
			errorMsg = "There is no application associated with the given file name extension. This error will also be returned if you attempt to print a file that is not printable."
		case SE_ERR_OOM:
			errorMsg = "There was not enough memory to complete the operation."
		case SE_ERR_SHARE:
			errorMsg = "A sharing violation occurred."
		default:
			errorMsg = fmt.Sprintf("Unknown error occurred with error code %v", ret)
		}
	} else {
		return nil
	}

	return errors.New(errorMsg)
}
