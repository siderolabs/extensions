// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"syscall"
)

func utsFieldToString(field [65]int8) string {
	b := make([]byte, 0, len(field))
	for _, c := range field {
		if c == 0 {
			break
		}
		b = append(b, byte(c))
	}
	return string(b)
}

func main() {
	var uts syscall.Utsname
	if err := syscall.Uname(&uts); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("%s %s %s %s\n",
		utsFieldToString(uts.Sysname),
		utsFieldToString(uts.Release),
		utsFieldToString(uts.Version),
		utsFieldToString(uts.Machine),
	)
}
